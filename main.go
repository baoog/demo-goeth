package main

import (
	"context"
	"encoding/hex"
	"fmt"
	store "github.com/baoog/demo-goeth/abi"
	client2 "github.com/baoog/demo-goeth/client"
	_ "github.com/baoog/demo-goeth/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func main() {
	client := client2.NewClient()
	defer client.Close()

	contractAddress := common.HexToAddress("0x9d4032ba814cd49085c991e787585d4b4581417a")
	bytes, err := client.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

	//print bytes code
	fmt.Println("Bytecode", fmt.Sprintf("0x%v", hex.EncodeToString(bytes)))

	instance, err := store.NewToken(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xb4b9e5a6e83e42ffea43462d8dd10a51e965c871")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	domains, err := instance.DOMAINSEPARATOR(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance", bal) // Balanace
	fmt.Println("DOMAIN_SEPARATOR", fmt.Sprintf("0x%v", hex.EncodeToString(domains[:])))

}
