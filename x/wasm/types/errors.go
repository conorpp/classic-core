package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Wasm Errors (increased to not conflict)
var (
	ErrAccountExists             = sdkerrors.Register(ModuleName, 100002, "contract account already exists")
	ErrInstantiateFailed         = sdkerrors.Register(ModuleName, 100003, "instantiate wasm contract failed")
	ErrExecuteFailed             = sdkerrors.Register(ModuleName, 100004, "execute wasm contract failed")
	ErrGasLimit                  = sdkerrors.Register(ModuleName, 100005, "insufficient gas")
	ErrInvalidGenesis            = sdkerrors.Register(ModuleName, 100006, "invalid genesis")
	ErrNotFound                  = sdkerrors.Register(ModuleName, 100007, "not found")
	ErrInvalidMsg                = sdkerrors.Register(ModuleName, 100008, "invalid Msg from the contract")
	ErrNoRegisteredQuerier       = sdkerrors.Register(ModuleName, 100009, "failed to find querier for route")
	ErrNoRegisteredParser        = sdkerrors.Register(ModuleName, 1000010, "failed to find parser for route")
	ErrMigrationFailed           = sdkerrors.Register(ModuleName, 1000011, "migrate wasm contract failed")
	ErrNotMigratable             = sdkerrors.Register(ModuleName, 1000012, "the contract is not migratable ")
	ErrStoreCodeFailed           = sdkerrors.Register(ModuleName, 1000013, "store wasm contract failed")
	ErrContractQueryFailed       = sdkerrors.Register(ModuleName, 1000014, "contract query failed")
	ErrExceedMaxContractSize     = sdkerrors.Register(ModuleName, 1000015, "exceeds max contract size limit")
	ErrExceedMaxContractMsgSize  = sdkerrors.Register(ModuleName, 1000016, "exceeds max contract msg size limit")
	ErrExceedMaxContractDataSize = sdkerrors.Register(ModuleName, 1000017, "exceeds max contract data size limit")
	ErrReplyFailed               = sdkerrors.Register(ModuleName, 1000018, "reply wasm contract failed")
	ErrExceedMaxQueryDepth       = sdkerrors.Register(ModuleName, 1000019, "exceed max query depth")
	ErrPinContractFailed         = sdkerrors.Register(ModuleName, 1000020, "pinning contract failed")
)
