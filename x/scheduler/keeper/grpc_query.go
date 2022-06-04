package keeper

import (
	"gert/x/scheduler/types"
)

var _ types.QueryServer = Keeper{}
