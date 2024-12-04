package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/charles08200/cosmos/testutil/keeper"
	"github.com/charles08200/cosmos/x/cosmos/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.cosmosKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
