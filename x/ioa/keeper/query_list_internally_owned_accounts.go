package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jinpk/gesia/x/ioa/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListInternallyOwnedAccounts(goCtx context.Context, req *types.QueryListInternallyOwnedAccountsRequest) (*types.QueryListInternallyOwnedAccountsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InternallyOwnedAccountPrefix))

	var accounts []types.InternallyOwnedAccount

	pageRes, err := query.Paginate(store, req.Pagination, func(key, value []byte) error {
		var account types.InternallyOwnedAccount
		if err := k.cdc.Unmarshal(value, &account); err != nil {
			return err
		}

		accounts = append(accounts, account)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListInternallyOwnedAccountsResponse{
		InternallyOwnedAccounts: accounts,
		Pagination:              pageRes,
	}, nil
}
