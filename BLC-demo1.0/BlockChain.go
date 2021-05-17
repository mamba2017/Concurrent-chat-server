package BLC

type BlockChain struct {
	Blocks []*Block //区块链

}
//初始化区块链
func CreateBlockChainWithGenesisBlock() *BlockChain {
	//生成创世区块
	block := CreateGenesisBlock([]byte("init blockchain"))
	return &BlockChain{[]*Block{block}}
}
//添加区块到区块链中
func (b *BlockChain)AddBlock(height int64,preBlockHash, data []byte)  {
	//var newBlock *Block
	newBlock := NewBlock(height,preBlockHash,data)
	b.Blocks = append(b.Blocks,newBlock)
}
