package cosmos_test

import (
	"testing"

	keepertest "github.com/charles08200/cosmos/testutil/keeper"
	"github.com/charles08200/cosmos/testutil/nullify"
	cosmos "github.com/charles08200/cosmos/x/cosmos/module"
	"github.com/charles08200/cosmos/x/cosmos/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.cosmosKeeper(t)
	cosmos.InitGenesis(ctx, k, genesisState)
	got := cosmos.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
