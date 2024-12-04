package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/charles08200/cosmos/testutil/keeper"
	"github.com/charles08200/cosmos/x/cosmos/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.cosmosKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
