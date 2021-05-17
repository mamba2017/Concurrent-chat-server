package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
)

//实现int64转[]byte
func IntToHex(data int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer,binary.BigEndian,data)
	if err != nil{
		log.Panic("transact failed",err)
	}
	return buffer.Bytes()
}
