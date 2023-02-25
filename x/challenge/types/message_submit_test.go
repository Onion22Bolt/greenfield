package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/bnb-chain/greenfield/testutil/sample"
	storagetypes "github.com/bnb-chain/greenfield/x/storage/types"
)

func TestMsgSubmit_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSubmit
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSubmit{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid bucket name",
			msg: MsgSubmit{
				Creator:           sample.AccAddress(),
				SpOperatorAddress: sample.AccAddress(),
				BucketName:        "1",
			},
			err: storagetypes.ErrInvalidBucketName,
		}, {
			name: "invalid object name",
			msg: MsgSubmit{
				Creator:           sample.AccAddress(),
				SpOperatorAddress: sample.AccAddress(),
				BucketName:        "bucket",
				ObjectName:        "",
			},
			err: storagetypes.ErrInvalidObjectName,
		}, {
			name: "valid message with random index",
			msg: MsgSubmit{
				Creator:           sample.AccAddress(),
				SpOperatorAddress: sample.AccAddress(),
				BucketName:        "bucket",
				ObjectName:        "object",
				RandomIndex:       true,
				SegmentIndex:      10,
			},
		}, {
			name: "valid message with specific index",
			msg: MsgSubmit{
				Creator:           sample.AccAddress(),
				SpOperatorAddress: sample.AccAddress(),
				BucketName:        "bucket",
				ObjectName:        "object",
				RandomIndex:       false,
				SegmentIndex:      2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
