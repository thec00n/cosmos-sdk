package appmanager

import (
	"context"

	"cosmossdk.io/server/v2/core/transaction"
)

type Identity = []byte

type Tx = transaction.Tx

type MsgRouterBuilder interface {
	RegisterHandler(msg Type, handlerFunc func(ctx context.Context, msg Type) (resp Type, err error))
}

type QueryRouterBuilder = MsgRouterBuilder

type PreMsgRouterBuilder interface {
	RegisterPreHandler(msg Type, preHandler func(ctx context.Context, msg Type) error)
}

type PostMsgRouterBuilder interface {
	RegisterPostHandler(msg Type, postHandler func(ctx context.Context, msg, msgResp Type) error)
}

type STFModule[T transaction.Tx] interface {
	Name() string
	RegisterMsgHandlers(router MsgRouterBuilder)
	RegisterQueryHandler(router QueryRouterBuilder)
}

type HasPreHandlers interface {
	RegisterPreMsgHandler(router PreMsgRouterBuilder)
	RegisterPostMsgHandler(router PostMsgRouterBuilder)
}

type HasTxValidator[T transaction.Tx] interface {
	TxValidator() func(ctx context.Context, tx T)
}

// TODO move that to core
// Deprecate HasABCIEndblock and have module manager handle it properly
type HasUpdateValidators interface {
	UpdateValidators() func(ctx context.Context) ([]ValidatorUpdate, error)
}

// ValidatorUpdate defines what is expected to be returned
// TODO move that to core
type ValidatorUpdate struct {
	PubKey []byte
	Power  int64 // updated power of the validtor
}
