package keeper

import (
	"github.com/rickmak/chimachain/x/chimachain/types"
)

var _ types.QueryServer = Keeper{}
