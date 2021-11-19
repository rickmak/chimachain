package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/rickmak/chimachain/testutil/keeper"
	"github.com/rickmak/chimachain/x/chimachain/keeper"
	"github.com/rickmak/chimachain/x/chimachain/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ChimachainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
