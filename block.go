package main

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/jsonrpc"
	"log"
)

func calculator(a string) {
	fmt.Printf("%s\n", a)
}

func main() {
	client, err := jsonrpc.NewClient("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	bus := EventBus.New()
	bus.Subscribe("main:calculator", calculator)

	//订阅新区块
	data := make(chan []byte)
	cancel, err := client.Subscribe("newHeads", func(b []byte) {
		data <- b
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case buf := <-data: //读取区块
			var block ethgo.Block
			if err := block.UnmarshalJSON(buf); err != nil {
				log.Fatal(err)
			}
			//GetTransactionReceipt returns the receipt of a transaction by transaction hash.
			receipt, err := client.Eth().GetTransactionReceipt(block.Hash)
			if err != nil {
				log.Fatal(err)
			}
			if len(receipt.Logs) > 0 {
				for _, v := range receipt.Logs {
					address := v.Topics[0]
					//判断erc20 transfer
					evm := new EVM(address)
					funcs := evm.getFunctions()
					if funcs[0] == "transfer(address,uint256)" {
						bus.Publish("main:calculator", address)
					}
				}
			}
		}
	}
	bus.Unsubscribe("main:calculator", calculator)
	log.Fatal(cancel)

}
