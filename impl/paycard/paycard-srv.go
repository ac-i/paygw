package paycard

import (
	"context"
	"paygw/proto/card"
	"paygw/proto/currency"
	"paygw/proto/paycard"
	"strconv"
	"strings"
	"time"
)

func (s *PayCardService) authorize(ctx context.Context, in *paycard.Authorization) (*paycard.AuthorizationState, error) {
	t := time.Now()
	s.app.Log.Debug().
		Str("rpc", "authorize").
		Interface("in", in).
		Send()

	// in = &paycard.Authorization{
	// 	Uid: in.GetUid(),
	// 	Party: &party.Party{
	// 		Uid:  "uid-party",
	// 		Name: "name-party",
	// 		Role: "role-party-merchant",
	// 	},
	// 	BasicAuth: &app.BasicAuth{
	// 		Username: "merchant-user",
	// 		Password: "merchant-pasw",
	// 	},
	// 	Card: &card.Card{
	// 		Brand:    "card-brand",
	// 		Number:   "card-number",
	// 		ExpMonth: "card-mm",
	// 		ExpYear:  "card-yy",
	// 		Names:    []string{"card-user", "card-user-com"},
	// 		Cvc:      "card-cvc",
	// 	},
	// 	Currency: &currency.Currency{
	// 		Code3:     "EUR",
	// 		Amount100: 100_000,
	// 	},
	// }
	out := &paycard.AuthorizationState{
		Ok: true,
		State: &paycard.State{
			Ok:         true,
			Uid:        in.GetUid(),
			Action:     "authorize",
			Code:       0,
			Message:    "",
			DataFields: []string{"ok", "state", "authorization", "currency"},
			Log: &paycard.TransactionLog{
				Authorization: &card.Authorization{
					Brand: "card-brand",
					Uid:   strings.Join([]string{"atz", strconv.FormatInt(t.UnixNano(), 32)}, "-"),
					Ns:    t.UnixNano(),
				},
				Authorized: &currency.Currency{
					Code3:     "EUR",
					Amount100: 100_000,
				},
				Available: &currency.Currency{
					Code3:     "EUR",
					Amount100: 100_000,
				},
			},
		},
		Authorization: &card.Authorization{
			Brand: "card-brand",
			Uid:   strings.Join([]string{"atz", strconv.FormatInt(t.UnixNano(), 32)}, "-"),
			Ns:    t.UnixNano(),
		},
		Currency: &currency.Currency{
			Code3:     "EUR",
			Amount100: 100_000,
		},
	}
	s.storeTransactionLog(s.app.Ctx, out.State.Log)
	s.app.Log.Debug().
		Str("rpc", "authorize").
		Interface("in", in).
		Interface("out", out).
		Send()
	// return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
	return out, nil
}
func (s *PayCardService) process(ctx context.Context, in *paycard.Transaction) (*paycard.TransactionState, error) {
	t := time.Now()
	s.app.Log.Debug().
		Str("rpc", "process").
		Str("action", in.GetAction()).
		Interface("in", in).
		Send()
	tr, _ := s.getTransactionLog(ctx, in.Authorization)
	s.app.Log.Debug().
		Str("rpc", "process").
		Str("action", in.GetAction()).
		Interface("tr", tr).
		Send()

	// in = &paycard.Transaction{
	// 	Uid:    "uid-transaction-indepotency",
	// 	Action: in.GetAction(),
	// 	Authorization: &card.Authorization{
	// 		Brand: "card-brand",
	// 		Uid:   "card-authorization-uid",
	// 		Ns:    t.UnixNano(), // TODO
	// 	},
	// 	Currency: &currency.Currency{
	// 		Code3:     "EUR",
	// 		Amount100: 10_000,
	// 	},
	// }
	out := &paycard.TransactionState{
		Ok: true,
		State: &paycard.State{
			Ok:         true,
			Uid:        in.GetUid(),
			Action:     in.GetAction(),
			Code:       0,
			Message:    "",
			DataFields: []string{"ok", "state", "currency"},
			Log: &paycard.TransactionLog{
				Authorization: in.Authorization,
				// Currency: , // TODO
				Authorized: &currency.Currency{
					Code3:     "EUR",
					Amount100: 100_000,
				},
				Available: &currency.Currency{
					Code3:     "EUR",
					Amount100: 90_000,
				},
				Capture: &currency.Log100{
					Sum:     in.Currency.Amount100,
					Count:   1,
					Amounts: []int64{in.Currency.Amount100},
					Ns:      []int64{t.UnixNano()},
				},
			},
		},
		Available: &currency.Currency{
			Code3:     "EUR",
			Amount100: 90_000,
		},
	}
	s.app.Log.Debug().
		Str("rpc", "process").
		Str("action", in.GetAction()).
		Interface("in", in).
		Interface("out", out).
		Send()
	// return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
	return out, nil
}

/*
example fixed values

func (s *PayCardService) authorize(ctx context.Context, in *paycard.Authorization) (*paycard.AuthorizationState, error) {
	t := time.Now()
	s.app.Log.Debug().Interface("authorize", in).Send()
	s.app.Log.Debug().
		Str("rpc", "authorize").
		Interface("in", in).
		Send()
	in = &paycard.Authorization{
		Uid: in.GetUid(),
		Party: &party.Party{
			Uid:  "uid-party",
			Name: "name-party",
			Role: "role-party-merchant",
		},
		BasicAuth: &app.BasicAuth{
			Username: "merchant-user",
			Password: "merchant-pasw",
		},
		Card: &card.Card{
			Brand:    "card-brand",
			Number:   "card-number",
			ExpMonth: "card-mm",
			ExpYear:  "card-yy",
			Names:    []string{"card-user", "card-user-com"},
			Cvc:      "card-cvc",
		},
		Currency: &currency.Currency{
			Code3:     "EUR",
			Amount100: 100_000,
		},
	}
	out := &paycard.AuthorizationState{
		Ok: true,
		State: &paycard.State{
			Ok:         true,
			Uid:        in.GetUid(),
			Action:     "authorize",
			Code:       0,
			Message:    "",
			DataFields: []string{"ok", "state", "authorization", "currency"},
			Log: &paycard.TransactionLog{
				Authorization: &card.Authorization{
					Brand: "card-brand",
					Uid:   "authorization-uid",
					Ns:    t.UnixNano(),
				},
				Authorized: &currency.Currency{
					Code3:     "EUR",
					Amount100: 100_000,
				},
				Available: &currency.Currency{
					Code3:     "EUR",
					Amount100: 100_000,
				},
			},
		},
		Authorization: &card.Authorization{
			Brand: "card-brand",
			Uid:   "authorization-uid",
			Ns:    t.UnixNano(),
		},
		Currency: &currency.Currency{
			Code3:     "EUR",
			Amount100: 100_000,
		},
	}
	s.app.Log.Debug().
		Str("rpc", "authorize").
		Interface("in", in).
		Interface("out", out).
		Send()
	// return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
	return out, nil
}
func (s *PayCardService) process(ctx context.Context, in *paycard.Transaction) (*paycard.TransactionState, error) {
	t := time.Now()
	s.app.Log.Debug().
		Str("rpc", "process").
		Str("action", in.GetAction()).
		Interface("in", in).
		Send()
	in = &paycard.Transaction{
		Uid:    "uid-transaction-indepotency",
		Action: in.GetAction(),
		Authorization: &card.Authorization{
			Brand: "card-brand",
			Uid:   "card-authorization-uid",
			Ns:    t.UnixNano(), // TODO
		},
		Currency: &currency.Currency{
			Code3:     "EUR",
			Amount100: 10_000,
		},
	}
	out := &paycard.TransactionState{
		Ok: true,
		State: &paycard.State{
			Ok:         true,
			Uid:        in.GetUid(),
			Action:     in.GetAction(),
			Code:       0,
			Message:    "",
			DataFields: []string{"ok", "state", "currency"},
			Log: &paycard.TransactionLog{
				Authorization: in.Authorization,
				// Currency: , // TODO
				Authorized: &currency.Currency{
					Code3:     "EUR",
					Amount100: 100_000,
				},
				Available: &currency.Currency{
					Code3:     "EUR",
					Amount100: 90_000,
				},
				Capture: &currency.Log100{
					Sum:     in.Currency.Amount100,
					Count:   1,
					Amounts: []int64{in.Currency.Amount100},
					Ns:      []int64{t.UnixNano()},
				},
			},
		},
		Available: &currency.Currency{
			Code3:     "EUR",
			Amount100: 90_000,
		},
	}
	s.app.Log.Debug().
		Str("rpc", "process").
		Str("action", in.GetAction()).
		Interface("in", in).
		Interface("out", out).
		Send()
	// return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
	return out, nil
}


*/
