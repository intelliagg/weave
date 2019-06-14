package username

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/x"
	"github.com/iov-one/weave/x/cash"
)

const (
	registerTokenCost     = 0
	changeTokenOwnerCost  = 0
	changeTokenTargetCost = 0
)

func RegisterRoutes(r weave.Registry, auth x.Authenticator, cash cash.Controller) {
	r = migration.SchemaMigratingRegistry("username", r)

	b := NewTokenBucket()
	r.Handle(RegisterTokenMsg{}.Path(), &registerTokenHandler{auth: auth, bucket: b})
	r.Handle(ChangeTokenOwnerMsg{}.Path(), &changeTokenOwnerHandler{auth: auth, bucket: b})
	r.Handle(ChangeTokenTargetMsg{}.Path(), &changeTokenTargetHandler{auth: auth, bucket: b})
}

type registerTokenHandler struct {
	auth   x.Authenticator
	bucket orm.ModelBucket
}

func (h *registerTokenHandler) Check(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	if _, err := h.validate(ctx, db, tx); err != nil {
		return nil, err
	}
	return &weave.CheckResult{GasAllocated: registerTokenCost}, nil
}

func (h *registerTokenHandler) Deliver(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	msg, err := h.validate(ctx, db, tx)
	if err != nil {
		return nil, err
	}

	owner := x.MainSigner(ctx, h.auth).Address()
	if len(owner) == 0 {
		return nil, errors.Wrap(errors.ErrUnauthorized, "message must be signed")
	}

	token := Token{
		Metadata: &weave.Metadata{Schema: 1},
		Target:   msg.Target,
		Owner:    owner,
	}
	if _, err := h.bucket.Put(db, msg.Username.Bytes(), &token); err != nil {
		return nil, errors.Wrap(err, "cannot store token")
	}
	return &weave.DeliverResult{Data: msg.Username.Bytes()}, nil
}

func (h *registerTokenHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*RegisterTokenMsg, error) {
	var msg RegisterTokenMsg
	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, errors.Wrap(err, "load msg")
	}

	switch err := h.bucket.Has(db, msg.Username.Bytes()); {
	case err == nil:
		// All good.
	case errors.ErrNotFound.Is(err):
		return nil, errors.Wrapf(errors.ErrDuplicate, "username %q already registered", msg.Username)
	default:
		return nil, errors.Wrap(err, "cannot check if username is unique")
	}
	return &msg, nil
}

type changeTokenOwnerHandler struct {
	auth   x.Authenticator
	bucket orm.ModelBucket
}

func (h *changeTokenOwnerHandler) Check(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	if _, _, err := h.validate(ctx, db, tx); err != nil {
		return nil, err
	}
	return &weave.CheckResult{GasAllocated: changeTokenOwnerCost}, nil
}

func (h *changeTokenOwnerHandler) Deliver(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	msg, token, err := h.validate(ctx, db, tx)
	if err != nil {
		return nil, err
	}

	token.Owner = msg.NewOwner
	if _, err := h.bucket.Put(db, msg.Username.Bytes(), token); err != nil {
		return nil, errors.Wrap(err, "cannot store token")
	}
	return &weave.DeliverResult{Data: msg.Username.Bytes()}, nil
}

func (h *changeTokenOwnerHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*ChangeTokenOwnerMsg, *Token, error) {
	var msg ChangeTokenOwnerMsg
	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, nil, errors.Wrap(err, "load msg")
	}

	var token Token
	if err := h.bucket.One(db, msg.Username.Bytes(), &token); err != nil {
		return nil, nil, errors.Wrap(err, "cannot get token from database")
	}

	if !h.auth.HasAddress(ctx, token.Owner) {
		return nil, nil, errors.Wrap(errors.ErrUnauthorized, "only the token owner can execute this operation")
	}

	return &msg, &token, nil
}

type changeTokenTargetHandler struct {
	auth   x.Authenticator
	bucket orm.ModelBucket
}

func (h *changeTokenTargetHandler) Check(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	if _, _, err := h.validate(ctx, db, tx); err != nil {
		return nil, err
	}
	return &weave.CheckResult{GasAllocated: changeTokenTargetCost}, nil
}

func (h *changeTokenTargetHandler) Deliver(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	msg, token, err := h.validate(ctx, db, tx)
	if err != nil {
		return nil, err
	}

	token.Target = msg.NewTarget
	if _, err := h.bucket.Put(db, msg.Username.Bytes(), token); err != nil {
		return nil, errors.Wrap(err, "cannot store token")
	}
	return &weave.DeliverResult{Data: msg.Username.Bytes()}, nil
}

func (h *changeTokenTargetHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*ChangeTokenTargetMsg, *Token, error) {
	var msg ChangeTokenTargetMsg
	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, nil, errors.Wrap(err, "load msg")
	}

	var token Token
	if err := h.bucket.One(db, msg.Username.Bytes(), &token); err != nil {
		return nil, nil, errors.Wrap(err, "cannot get token from database")
	}

	if !h.auth.HasAddress(ctx, token.Owner) {
		return nil, nil, errors.Wrap(errors.ErrUnauthorized, "only the token owner can execute this operation")
	}

	return &msg, &token, nil
}
