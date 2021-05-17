package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//共识算法管理文件

//实现POW实力以及相关功能
// 目标难度值
const targetBit = 16
//工作量证明的结构
type ProofOfWork struct {
	//需要共识验证的区块
	Block *Block
	//目标难度的哈希（大数据存储）
	target  *big.Int

}

func NewProofOfWork(block *Block) *ProofOfWork{
		target :=big.NewInt(1)
		fmt.Println("tar:",target)
		//数据长度为8位
		//需求：满足前两位为0
		target = target.Lsh(target,256-targetBit)
		return &ProofOfWork{Block: block,target: target}
}
func (pw * ProofOfWork)Run()([]byte,int){
	var nonce = 0
	var hashInt big.Int
	var hash [32]byte
	//无限循环
	for  {
		//生成准备数据
		dataBytes := pw.prepareData(int64(nonce))
		hash = sha256.Sum256(dataBytes)
		hashInt.SetBytes(hash[:])
		//检测生成的哈希是否符合条件
		 if pw.target.Cmp(&hashInt) ==1{
		 	//找到符合条件的哈希
		 	break
		 }
		 nonce++
	}
	fmt.Printf("\n碰撞次数：%d\n",nonce)
	return hash[:],nonce
}
//生成准备数据
func (pw *ProofOfWork)prepareData(nonce int64) []byte{
	var data []byte
	timeStapBytes := IntToHex(pw.Block.TimeStamp)
	heightBytes:= IntToHex(pw.Block.Height)
	data = bytes.Join([][]byte{
		heightBytes,
		timeStapBytes,
		pw.Block.PrevBlockHash,
		pw.Block.Data,
		IntToHex(nonce),
		IntToHex(targetBit),
	},[]byte{})
	return data
}

