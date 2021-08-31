package eth

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/glog"
)

// func (b *Block) Number() *big.Int     { return new(big.Int).Set(b.header.Number) }
// func (b *Block) GasLimit() uint64     { return b.header.GasLimit }
// func (b *Block) GasUsed() uint64      { return b.header.GasUsed }
// func (b *Block) Difficulty() *big.Int { return new(big.Int).Set(b.header.Difficulty) }
// func (b *Block) Time() uint64         { return b.header.Time }

// func (b *Block) NumberU64() uint64        { return b.header.Number.Uint64() }
// func (b *Block) MixDigest() common.Hash   { return b.header.MixDigest }
// func (b *Block) Nonce() uint64            { return binary.BigEndian.Uint64(b.header.Nonce[:]) }
// func (b *Block) Bloom() Bloom             { return b.header.Bloom }
// func (b *Block) Coinbase() common.Address { return b.header.Coinbase }
// func (b *Block) Root() common.Hash        { return b.header.Root }
// func (b *Block) ParentHash() common.Hash  { return b.header.ParentHash }
// func (b *Block) TxHash() common.Hash      { return b.header.TxHash }
// func (b *Block) ReceiptHash() common.Hash { return b.header.ReceiptHash }
// func (b *Block) UncleHash() common.Hash   { return b.header.UncleHash }
// func (b *Block) Extra() []byte            { return common.CopyBytes(b.header.Extra) }
func printBlockInfo(block *types.Block) {
	if block == nil {
		glog.Warning("block is nil")
		return
	}
	glog.Infoln("====Block====")
	glog.Infoln("Number:", block.Number().Uint64())
	glog.Infoln("GasLimit:", block.GasLimit())
	glog.Infoln("GasUsed:", block.GasUsed())
	glog.Infoln("Difficulty:", block.Difficulty().Uint64())
	glog.Infoln("Time:", block.Time())
	glog.Infoln("Hash:", block.Hash().Hex())
	glog.Infoln("ParentHash:", block.ParentHash().Hex())
	glog.Infoln("UncleHash:", block.UncleHash().Hex())

}
