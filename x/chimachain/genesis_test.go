package chimachain_test

import (
	"testing"

	keepertest "github.com/rickmak/chimachain/testutil/keeper"
	"github.com/rickmak/chimachain/x/chimachain"
	"github.com/rickmak/chimachain/x/chimachain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChimachainKeeper(t)
	chimachain.InitGenesis(ctx, *k, genesisState)
	got := chimachain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
