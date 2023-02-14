// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package kvcoord

import (
	"context"
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/util/log"

	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/util/hlc"
	"github.com/cockroachdb/cockroach/pkg/util/timeutil"
)

// txnMetricRecorder is a txnInterceptor in charge of updating metrics about
// the behavior and outcome of a transaction. It records information about the
// requests that a transaction sends and updates counters and histograms when
// the transaction completes.
//
// TODO(nvanbenschoten): Unit test this file.
type txnMetricRecorder struct {
	wrapped lockedSender
	metrics *TxnMetrics
	clock   *hlc.Clock

	txn            *roachpb.Transaction
	txnStartNanos  int64
	onePCCommit    bool
	parallelCommit bool
}

// SendLocked is part of the txnInterceptor interface.
func (m *txnMetricRecorder) SendLocked(
	ctx context.Context, ba roachpb.BatchRequest,
) (*roachpb.BatchResponse, *roachpb.Error) {
	if m.txnStartNanos == 0 {
		m.txnStartNanos = timeutil.Now().UnixNano()
	}

	br, pErr := m.wrapped.SendLocked(ctx, ba)
	if pErr != nil {
		return br, pErr
	}

	if length := len(br.Responses); length > 0 {
		if et := br.Responses[length-1].GetEndTxn(); et != nil {
			// Check for 1-phase commit.
			m.onePCCommit = et.OnePhaseCommit

			// Check for parallel commit.
			m.parallelCommit = !et.StagingTimestamp.IsEmpty()
		}
	}
	return br, nil
}

// setWrapped is part of the txnInterceptor interface.
func (m *txnMetricRecorder) setWrapped(wrapped lockedSender) { m.wrapped = wrapped }

// populateLeafInputState is part of the txnInterceptor interface.
func (*txnMetricRecorder) populateLeafInputState(*roachpb.LeafTxnInputState) {}

// populateLeafFinalState is part of the txnInterceptor interface.
func (*txnMetricRecorder) populateLeafFinalState(*roachpb.LeafTxnFinalState) {}

// importLeafFinalState is part of the txnInterceptor interface.
func (*txnMetricRecorder) importLeafFinalState(context.Context, *roachpb.LeafTxnFinalState) {}

// epochBumpedLocked is part of the txnInterceptor interface.
func (*txnMetricRecorder) epochBumpedLocked() {}

// createSavepointLocked is part of the txnInterceptor interface.
func (*txnMetricRecorder) createSavepointLocked(context.Context, *savepoint) {}

// rollbackToSavepointLocked is part of the txnInterceptor interface.
func (*txnMetricRecorder) rollbackToSavepointLocked(context.Context, savepoint) {}

// closeLocked is part of the txnInterceptor interface.
func (m *txnMetricRecorder) closeLocked() {
	ctx := context.Background()
	if m.onePCCommit {
		m.metrics.Commits1PC.Inc(1)
	}
	if m.parallelCommit {
		m.metrics.ParallelCommits.Inc(1)
	}

	if m.txnStartNanos != 0 {
		duration := timeutil.Now().UnixNano() - m.txnStartNanos
		if duration >= 0 {
			m.metrics.Durations.RecordValue(duration)
		}
	}
	restarts := int64(m.txn.Epoch)
	status := m.txn.Status

	// TODO(andrei): We only record txn that had any restarts, otherwise the
	// serialization induced by the histogram shows on profiles. Figure something
	// out to make it cheaper - maybe augment the histogram library with an
	// "expected value" that is a cheap counter for the common case. See #30644.
	// Also, the epoch is not currently an accurate count since we sometimes bump
	// it artificially (in the parallel execution queue).
	if restarts > 0 {
		m.metrics.Restarts.RecordValue(restarts)
	}
	switch status {
	case roachpb.ABORTED:
		m.metrics.Aborts.Inc(1)
		// m.metrics.GetEva(ctx)
		// log.Info(ctx,"roachpb.ABORTED")
	case roachpb.PENDING:
		// NOTE(andrei): Getting a PENDING status here is possible when this
		// interceptor is closed without a rollback ever succeeding.
		// We increment the Aborts metric nevertheless; not sure how these
		// transactions should be accounted.
		m.metrics.Aborts.Inc(1)
		// Record failed aborts separately as in this case EndTxn never succeeded
		// which means intents are left for subsequent cleanup by reader.
		m.metrics.RollbacksFailed.Inc(1)
		// m.metrics.GetEva(ctx)
		// log.Info(ctx,"roachpb.PENDING")
	case roachpb.COMMITTED:
		// Note that successful read-only txn are also counted as committed, even
		// though they never had a txn record.
		m.metrics.Commits.Inc(1)
		// m.metrics.GetEva(ctx)
		// log.Info(ctx,"roachpb.COMMITTED")
	}
	log.Info(ctx, m.stringifyMetrics())
}

func (m *txnMetricRecorder) stringifyMetrics() string {
	return fmt.Sprintf(`{transaction: "%s", Aborts: "%d", RestartsTxnAborted: "%d", Commits: "%d", Commits1PC: "%d", ParallelCommits: "%d", CommitWaits: "%d", RefreshSuccess: "%d", RefreshFail: "%d", RefreshFailWithCondensedSpans: "%d", RefreshMemoryLimitExceeded: "%d", RefreshAutoRetries: "%d", Durations: "%d", TxnsWithCondensedIntents: "%d"}, TxnsRejectedByLockSpanBudget: "%d", Restarts: "%d", RestartsWriteTooOld: "%d", RestartsWriteTooOldMulti: "%d", RestartsSerializable: "%d", RestartsAsyncWriteFailure: "%d", RestartsCommitDeadlineExceeded: "%d", RestartsReadWithinUncertainty: "%d", RestartsTxnPush: "%d", RestartsUnknown: "%d", RollbacksFailed: "%d", AsyncRollbacksFailed: "%d" }`,
		m.txn.ID, m.metrics.Aborts.Count(), m.metrics.RestartsTxnAborted.Count(), m.metrics.Commits.Count(), m.metrics.Commits1PC.Count(), m.metrics.ParallelCommits.Count(), m.metrics.CommitWaits.Count(), m.metrics.RefreshSuccess.Count(), m.metrics.RefreshFail.Count(), m.metrics.RefreshFailWithCondensedSpans.Count(), m.metrics.RefreshMemoryLimitExceeded.Count(), m.metrics.RefreshAutoRetries.Count(), m.metrics.Durations.TotalCount(), m.metrics.TxnsWithCondensedIntents.Count(), m.metrics.TxnsRejectedByLockSpanBudget.Count(), m.metrics.Restarts.TotalCount(), m.metrics.RestartsWriteTooOld.Count(), m.metrics.RestartsWriteTooOldMulti.Count(), m.metrics.RestartsSerializable.Count(), m.metrics.RestartsAsyncWriteFailure.Count(), m.metrics.RestartsCommitDeadlineExceeded.Count(), m.metrics.RestartsReadWithinUncertainty.Count(), m.metrics.RestartsTxnPush.Count(), m.metrics.RestartsUnknown.Count(), m.metrics.RollbacksFailed.Count(), m.metrics.AsyncRollbacksFailed.Count())
}
