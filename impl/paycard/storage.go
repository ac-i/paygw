package paycard

import (
	"context"
	"fmt"
	"paygw/proto/card"
	"paygw/proto/paycard"

	"google.golang.org/protobuf/proto"
)

func (s *PayCardService) storeTransactionLog(ctx context.Context, in *paycard.TransactionLog) error {
	if in.Authorization.Uid == "" {
		return fmt.Errorf("empty Authorization.Uid")
	}
	mb, err := proto.Marshal(in)
	if err != nil {
		return err
	}
	s.storeMEM.Conn.Set([]byte(in.Authorization.Uid), mb)
	return nil
}

func (s *PayCardService) getTransactionLog(ctx context.Context, in *card.Authorization) (*paycard.TransactionLog, error) {
	if in.Uid == "" {
		return nil, fmt.Errorf("empty Authorization.Uid")
	}
	out := new(paycard.TransactionLog)
	err := proto.Unmarshal(s.storeMEM.Conn.Get([]byte{}, []byte(in.Uid)), out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
