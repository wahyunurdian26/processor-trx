package service

import (
	"context"
	"fmt"

	uLog "github.com/wahyunurdian26/util/logger"
)

func (s *trxProcessorService) ProcessTransaction(ctx context.Context, trxID string) error {
	uLog.LogInfo(ctx, "ProcessTransaction", fmt.Sprintf("Processing transaction: %s", trxID))

	trx, err := s.repo.FindByID(ctx, trxID)
	if err != nil {
		uLog.LogError(ctx, "s.repo.FindByID", err)
		return err
	}
	if trx == nil {
		return fmt.Errorf("transaction not found: %s", trxID)
	}

	if trx.Status != "PENDING" {
		uLog.LogInfo(ctx, "ProcessTransaction", fmt.Sprintf("Transaction %s is already processed with status %s", trxID, trx.Status))
		return nil
	}

	// Call Account Service to deduct budget
	err = s.accountClient.DeductBudget(ctx, trx.AccountID, trx.Amount)
	if err != nil {
		uLog.LogError(ctx, "s.accountClient.DeductBudget", err)
		s.repo.UpdateStatus(ctx, trxID, "FAILED")
		return err
	}

	// Update status to SUCCESS
	err = s.repo.UpdateStatus(ctx, trxID, "SUCCESS")
	if err != nil {
		uLog.LogError(ctx, "s.repo.UpdateStatus", err)
		return err
	}

	// Publish PaymentCompleted event for Audit and Email
	if s.broker != nil {
		event := map[string]interface{}{
			"transaction_id": trx.ID,
			"account_id":     trx.AccountID,
			"amount":         trx.Amount,
			"merchant_name":  trx.MerchantName,
			"status":         "SUCCESS",
			"timestamp":      trx.CreatedAt,
		}
		_ = s.broker.PublishPaymentCompleted(ctx, event)
	}

	uLog.LogInfo(ctx, "ProcessTransaction", fmt.Sprintf("Transaction %s successfully processed", trxID))
	return nil
}
