package paycard

import (
	"context"

	"paygw/impl/app/instance"
	"paygw/impl/app/storage/mem"
	"paygw/proto/paycard"
	"paygw/proto/paycard_srv"
)

// PayCardService PayCardServiceServer (PayCard Domain Service Server)
type PayCardService struct {
	paycard_srv.UnimplementedPayCardServiceServer
	app      *instance.App
	storeMEM *mem.Storage
}

func NewPayCardService(app *instance.App) *PayCardService {
	return &PayCardService{
		app:      app,
		storeMEM: mem.NewStorageMEM(),
	}
}

func (s *PayCardService) Authorize(ctx context.Context, in *paycard.Authorization) (*paycard.AuthorizationState, error) {
	return s.authorize(ctx, in)
}
func (s *PayCardService) Process(ctx context.Context, in *paycard.Transaction) (*paycard.TransactionState, error) {
	return s.process(ctx, in)
}
