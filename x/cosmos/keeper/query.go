package keeper

import (
	"github.com/charles08200/cosmos/x/cosmos/types"
)

var _ types.QueryServer = Keeper{}
