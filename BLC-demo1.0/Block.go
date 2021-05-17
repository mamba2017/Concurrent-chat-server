package BLC

import (
	"bytes"
	"crypto/sha256"
	"time"
)

//区块基本结构与功能管理

type Block struct {
	TimeStamp     	int64       //时间戳
	Hash         	[]byte      //当前区块哈希
	PrevBlockHash 	[]byte   //前区块哈希
	Height        	int64    //区块高度
	Data   			[]byte   //交易数据
	Nonce           int64
}

func NewBlock(height int64,prevBlockHash ,data []byte) *Block{
	var block Block
	block = Block{
		TimeStamp: time.Now().Unix(),
		Hash: nil,
		PrevBlockHash: prevBlockHash,
		Height: height,
		Data: data,
	}
	//block.SetHash()
	//生成hash
	pow := NewProofOfWork(&block)
	hash,nonce := pow.Run()
	block.Hash = hash
	block.Nonce = int64(nonce)
	return &block
}

//计算区块哈希
func (b *Block)SetHash()  {
	//调用sha256
	//实现int->hash
	timeStapBytes := IntToHex(b.TimeStamp)
	heightBytes:= IntToHex(b.Height)
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		timeStapBytes,
		b.PrevBlockHash,
		b.Data,
	},[]byte{})
	hash := sha256.Sum256(blockBytes)
	b.Hash = hash[:]
}
//生成创世区块
func CreateGenesisBlock(data []byte) *Block {
	return NewBlock(1,nil,data)
}
