package nft

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	"gitlab.cschain.tech/csmod/modules/nft"
)

func init() {
	codec.RegisterAppModules(nft.AppModuleBasic{})
}