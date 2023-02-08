package challenge

import (
	"github.com/bnb-chain/greenfield/x/challenge/keeper"
	"github.com/bnb-chain/greenfield/x/challenge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

// coolingOffMultiplier is used to purge slashes. The purpose of it is for: if the cooling-off period is increase, e.g.
// by gov, then we still keep some recent slashes for the new parameter. However, if it is increased too large, then
// there is the possibility that some recent slashes for the new parameter will be gone, and we think this is acceptable.
const coolingOffMultiplier = 3

func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	// reset count of challenge in current block to zero
	keeper.ResetChallengeCount(ctx)

	// delete expired challenges at this height
	events := make([]proto.Message, 0)
	height := uint64(ctx.BlockHeight())
	minHeight := height - keeper.ChallengeExpirePeriod(ctx)
	challenges := keeper.GetAllOngoingChallenge(ctx)
	for _, elem := range challenges {
		if elem.Height < minHeight {
			events = append(events, &types.EventExpireChallenge{
				ChallengeId: elem.Id,
			})
			keeper.RemoveOngoingChallenge(ctx, elem.Id)
		}
	}

	_ = ctx.EventManager().EmitTypedEvents(events...)

	// delete too old slashes at this height
	coolingOff := keeper.SlashCoolingOffPeriod(ctx)
	minHeight = height - coolingOff*coolingOffMultiplier
	slashes := keeper.GetAllRecentSlash(ctx)
	for _, elem := range slashes {
		if elem.Height < minHeight {
			keeper.RemoveRecentSlash(ctx, elem.Id)
		}
	}
}

func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	// emit new challenge events if more challenges are needed
	count := keeper.GetChallengeCount(ctx)
	needed := keeper.EventCountPerBlock(ctx)
	events := make([]proto.Message, 0)
	for count < needed {
		// TODO: random object to challenge
		count++
	}
	_ = ctx.EventManager().EmitTypedEvents(events...)
}
