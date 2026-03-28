package service

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/metadata"

	uLog "github.com/wahyunurdian26/util/logger"
	"github.com/wahyunurdian26/util/requestid"
)

func (s *trxProcessorService) SubscribeTransaction(ctx context.Context) error {
	if s.broker == nil {
		return fmt.Errorf("broker is not initialized")
	}

	return s.broker.Subscribe("transaction_initiated_queue", s.handleTransactionMessage)
}

func (s *trxProcessorService) handleTransactionMessage(exchange string, queue string, headers map[string]interface{}, body []byte) error {
	ctx := requestid.MiddlewareRequestId(context.Background(), metadata.Pairs(requestid.RequestIdAttr, fmt.Sprintf("%v", headers[requestid.RequestIdAttr])))

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
