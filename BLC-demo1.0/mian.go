package main

import (
	"fmt"
	"pro510.com/BLC"
)

func main() {
	//block := BLC.NewBlock(1,nil,[]byte("fisrt block chain"))
	//fmt.Println(block)
	bc := BLC.CreateBlockChainWithGenesisBlock() //创建区块链并初始化
	fmt.Println(len(bc.Blocks))
	fmt.Println("sss",bc.Blocks[0].Data)
	//测试左移
	//target :=big.NewInt(1)
	//target = target.Lsh(target,8-3)
	//fmt.Println(target)
	//AddBlock(height int64,preBlockHash, data []byte)
	bc.AddBlock(bc.Blocks[len(bc.Blocks) - 1].Height+1,   //添加区块
		bc.Blocks[len(bc.Blocks) - 1].Hash,[]byte("alice send bob 100"))

	bc.AddBlock(bc.Blocks[len(bc.Blocks) - 1].Height+1,   //添加区块
		bc.Blocks[len(bc.Blocks) - 1].Hash,[]byte("alice send bob 200"))
	for i,block := range bc.Blocks{
		fmt.Printf("区块%d,PrevBlockHash:%x,current:%x\n",i,block.PrevBlockHash,block.Hash)
	}
}
