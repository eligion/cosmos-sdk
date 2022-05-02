package store

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/eligion/cosmos-sdk/store/cache"
	"github.com/eligion/cosmos-sdk/store/rootmulti"
	"github.com/eligion/cosmos-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
