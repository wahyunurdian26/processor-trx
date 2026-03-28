package micro

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	pbaccount "github.com/wahyunurdian26/client/account"
	"github.com/wahyunurdian26/processor-trx/repository"
)

type accountClient struct {
	deductBudgetEndpoint endpoint.Endpoint
}

func NewAccountClient(conn *grpc.ClientConn) repository.AccountClient {
	deductBudgetEndpoint := grpctransport.NewClient(
		conn,
		"account.AccountService",
		"DeductBudget",
		encodeDeductBudgetRequest,
		decodeDeductBudgetResponse,
		pbaccount.DeductResponse{},
	).Endpoint()

	return &accountClient{
		deductBudgetEndpoint: deductBudgetEndpoint,
	}
}

type deductBudgetRequest struct {
	AccountID string
	Amount    float64
}

type deductBudgetResponse struct {
	Success bool
	Message string
	Err     error
}

func (r deductBudgetResponse) Error() error { return r.Err }

func encodeDeductBudgetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(deductBudgetRequest)
	return &pbaccount.DeductRequest{
		AccountId: req.AccountID,
		Amount:    req.Amount,
	}, nil
}

func decodeDeductBudgetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pbaccount.DeductResponse)
	if !resp.Success {
		return deductBudgetResponse{Success: false, Message: resp.Message, Err: errors.New(resp.Message)}, nil
	}
	return deductBudgetResponse{
		Success: resp.Success,
		Message: resp.Message,
	}, nil
}

func (c *accountClient) DeductBudget(ctx context.Context, accountID string, amount float64) error {
	req := deductBudgetRequest{
		AccountID: accountID,
		Amount:    amount,
	}
	resp, err := c.deductBudgetEndpoint(ctx, req)
	if err != nil {
		return err
	}

	response := resp.(deductBudgetResponse)
	return response.Err
}
