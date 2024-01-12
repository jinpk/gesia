package keeper

import (
	"time"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	types "github.com/jinpk/gesia/x/ioa/types"
)

/** IOA 생성 */
func (k Keeper) CreateInternallyOwnedAccount(ctx sdk.Context, withdrawal sdk.AccAddress, contractAddress string) (error, types.InternallyOwnedAccount) {
	key := append([]byte("ioa"), sdk.Uint64ToBigEndian(uint64(time.Now().Unix()))...)
	var ioaAddress sdk.AccAddress = address.Module(types.ModuleName, key)

	internallyOwnedAccount := &types.InternallyOwnedAccount{
		Address:           ioaAddress.String(),
		WithdrawalAddress: withdrawal.String(),
		ContractAddress:   contractAddress,
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InternallyOwnedAccountPrefix))

	bz := k.cdc.MustMarshal(internallyOwnedAccount)

	store.Set(ioaAddress.Bytes(), bz)

	return nil, *internallyOwnedAccount

}
