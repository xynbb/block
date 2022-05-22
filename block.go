package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/asaskevich/EventBus"
	"strconv"
	"time"
)

type Block struct {
	Index         int64
	TimeStamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}
type Blockchain struct {
	blocks []*Block
}

func NewBlock(index int64, data, prevBlockHash []byte) *Block {
	block := &Block{index, time.Now().Unix(), data, prevBlockHash, []byte{}}
	block.setHash()
	return block
}
func NewGenesiBlock() *Block {
	return NewBlock(0, []byte("first block"), []byte{})
}
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesiBlock()}}
}
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, []byte(data), prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
func (b *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.TimeStamp, 10))
	index := []byte(strconv.FormatInt(b.TimeStamp, 10))
	headers := bytes.Join([][]byte{timestamp, index, b.PrevBlockHash}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
func calculator(a int, b int) {
	fmt.Printf("%d\n", a+b)
}

func main() {
	bc := NewBlockchain()
	bc.AddBlock("send 1 BTC")
	bc.AddBlock("send 2 BTC")

	for _, block := range bc.blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("TimeStamp: %d\n", block.TimeStamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("==========================================")
	}

	bus := EventBus.New()
	bus.Subscribe("main:calculator", calculator)
	bus.Publish("main:calculator", 20, 40)
	bus.Unsubscribe("main:calculator", calculator)
}
