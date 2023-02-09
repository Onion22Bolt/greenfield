package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/challenge module sentinel errors
var (
	ErrUnknownSp           = sdkerrors.Register(ModuleName, 1, "unknown storage provider")
	ErrUnknownObject       = sdkerrors.Register(ModuleName, 1, "unknown object info")
	ErrInvalidObjectStatus = sdkerrors.Register(ModuleName, 1, "invalid object status to challenge")
	ErrInvalidSpStatus     = sdkerrors.Register(ModuleName, 2, "invalid storage provider status")
	ErrUnknownChallenge    = sdkerrors.Register(ModuleName, 3, "unknown challenge")
	ErrInvalidValSet       = sdkerrors.Register(ModuleName, 4, "invalid validator set")
	ErrInvalidBlsPubKey    = sdkerrors.Register(ModuleName, 5, "invalid bls public key")
	ErrVotesNotEnough      = sdkerrors.Register(ModuleName, 6, "attest votes are not enough")
	ErrInvalidBlsSignature = sdkerrors.Register(ModuleName, 7, "invalid bls signature")
	ErrInvalidGenesis      = sdkerrors.Register(ModuleName, 8, "invalid genesis state")
)
