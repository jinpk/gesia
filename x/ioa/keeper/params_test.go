package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/jinpk/gesia/testutil/keeper"
	"github.com/jinpk/gesia/x/ioa/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.IoaKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
