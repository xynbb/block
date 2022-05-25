package main

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/builtin/erc20"
)

var (
	zeroX       = ethgo.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	owerAddr    = ethgo.HexToAddress("1000")
	spenderAddr = ethgo.HexToAddress("2000")
)

func calculator(a string) {
	fmt.Printf("%s\n", a)
}

func main() {
	erc20 := erc20.NewERC20(zeroX)

	// Name calls the name method in the solidity contract
	name, _ := erc20.Name()
	fmt.Println(name)

	//Symbol calls the symbol method in the solidity contract
	symbol, _ := erc20.Symbol()
	fmt.Println(symbol)

	//Decimals calls the decimals method in the solidity contract
	decimals, _ := erc20.Decimals()
	fmt.Println(decimals)

	// TotalSupply calls the totalSupply method in the solidity contract
	supply, _ := erc20.TotalSupply()
	fmt.Println(supply)

	//BalanceOf calls the balanceOf method in the solidity contract
	erc20.BalanceOf(owerAddr)

	//Transfer sends a transfer transaction in the solidity contract
	erc20.Approve(spenderAddr, supply)
	erc20.Allowance(owerAddr, spenderAddr)

	//transfer event
	approveHash := erc20.ApprovalEventSig()
	transferHash := erc20.TransferEventSig()

	bus := EventBus.New()
	bus.Subscribe("main:calculator", calculator)
	bus.Publish("main:calculator", approveHash)
	bus.Publish("main:calculator", transferHash)
	bus.Unsubscribe("main:calculator", calculator)
}
