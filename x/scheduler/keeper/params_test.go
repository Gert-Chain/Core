package keeper_test

import (
	"testing"

	"gert/x/scheduler/types"
	testkeeper "gert/testutil/keeper"

	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SchedulerKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
