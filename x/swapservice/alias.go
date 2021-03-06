package swapservice

import (
	"github.com/jpthor/cosmos-swap/x/swapservice/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewMsgSetPoolData  = types.NewMsgSetPoolData
	NewPoolStruct      = types.NewPoolStruct
	NewMsgSetAccData   = types.NewMsgSetAccData
	NewAccStruct       = types.NewAccStruct
	NewMsgSetStakeData = types.NewMsgSetStakeData
	NewStakeStruct     = types.NewStakeStruct
	NewMsgSwap         = types.NewMsgSwap
	ModuleCdc          = types.ModuleCdc
	RegisterCodec      = types.RegisterCodec
)

type (
	MsgSetPoolData     = types.MsgSetPoolData
	MsgSetAccData      = types.MsgSetAccData
	MsgSetStakeData    = types.MsgSetStakeData
	MsgSwap            = types.MsgSwap
	QueryResResolve    = types.QueryResResolve
	QueryResPoolDatas  = types.QueryResPoolDatas
	QueryResAccDatas   = types.QueryResAccDatas
	QueryResStakeDatas = types.QueryResStakeDatas
	PoolStruct         = types.PoolStruct
	AccStruct          = types.AccStruct
	StakeStruct        = types.StakeStruct
	AccStake           = types.AccStake
	Holding            = types.Holding
)
