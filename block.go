package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// 区块
type Block struct {
	Timestamp     int64   //当前时间戳(区块被创建的时间)
	Data          [] byte //区块中实际包含的有用信息
	PrevBlockHash [] byte //存储的前一个区块的哈希值
	Hash          [] byte //当前区块的哈希值
}

// 哈希加密
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// 一个新的块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// 一个创世块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
