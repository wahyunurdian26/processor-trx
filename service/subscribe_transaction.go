package service

import (
	"context"
	"encoding/json"
	"fmt"

	"microservice/util/requestid"
	uLog "microservice/util/logger"
)

func (s *trxProcessorService) SubscribeTransaction(ctx context.Context) error {
	if s.broker == nil {
		return fmt.Errorf("broker is not initialized")
	}

	return s.broker.Subscribe("transaction_initiated_queue", s.handleTransactionMessage)
}

func (s *trxProcessorService) handleTransactionMessage(exchange string, queue string, headers map[string]interface{}, body []byte) error {
	ctx := requestid.MiddlewareRequestIdAMQP(context.Background(), headers)
	
	uLog.LogInfo(ctx, "handleTransactionMessage", "Received transaction initiated event")

	var event map[string]interface{}
	if err := json.Unmarshal(body, &event); err != nil {
		uLog.LogError(ctx, "json.Unmarshal", err)
		return err
	}

	trxID, ok := event["transaction_id"].(string)
	if !ok {
		uLog.LogError(ctx, "handleTransactionMessage", fmt.Errorf("transaction_id not found in event"))
		return fmt.Errorf("transaction_id not found")
	}

	return s.ProcessTransaction(ctx, trxID)
}
