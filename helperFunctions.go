package main

import (
	"github.com/FactomProject/factoid"
	"github.com/FactomProject/fctwallet/Wallet"

	"fmt"
)

var _ = fmt.Sprintf("")

type Address struct {
	Name    string
	Address string
	Balance float64
	Type    string
}

type Addresses struct {
	FactoidAddresses []Address
	ECAddresses      []Address
}

func GetAddresses() (*Addresses, error) {
	addresses := Wallet.GetAddresses()

	answer := new(Addresses)
	for _, we := range addresses {
		tmp := Address{}
		tmp.Type = we.GetType()
		if we.GetType() == "ec" {
			address, err := we.GetAddress()
			if err != nil {
				return nil, err
			}

			tmp.Address = factoid.ConvertECAddressToUserStr(address)
			tmp.Name = string(we.GetName())
			bal, err := Wallet.ECBalance(tmp.Address)
			if err != nil {
				return nil, err
			}
			tmp.Balance = float64(bal)
			answer.ECAddresses = append(answer.ECAddresses, tmp)
		} else {
			address, err := we.GetAddress()
			if err != nil {
				return nil, err
			}

			tmp.Address = factoid.ConvertFctAddressToUserStr(address)
			tmp.Name = string(we.GetName())
			bal, err := Wallet.FactoidBalance(tmp.Address)
			if err != nil {
				return nil, err
			}
			tmp.Balance = factoid.ConvertDecimalToFloat(uint64(bal))
			answer.FactoidAddresses = append(answer.FactoidAddresses, tmp)
		}
	}
	return answer, nil
}

type Balances struct {
	FactoidBalances float64
	ECBalances      float64
}

func GetBalances() (*Balances, error) {
	addresses, err := GetAddresses()
	if err != nil {
		return nil, err
	}
	answer := new(Balances)
	for _, v := range addresses.FactoidAddresses {
		answer.FactoidBalances += v.Balance
	}
	for _, v := range addresses.ECAddresses {
		answer.ECBalances += v.Balance
	}
	return answer, nil
}

func GetAddressMapByAddress() (map[string]Address, error) {
	addresses, err := GetAddresses()
	if err != nil {
		return nil, err
	}
	addressMap := map[string]Address{}
	for _, v := range addresses.FactoidAddresses {
		addressMap[v.Address] = v
	}
	for _, v := range addresses.ECAddresses {
		addressMap[v.Address] = v
	}
	return addressMap, nil
}

func GetAddressMapByName() (map[string]Address, error) {
	addresses, err := GetAddresses()
	if err != nil {
		return nil, err
	}
	addressMap := map[string]Address{}
	for _, v := range addresses.FactoidAddresses {
		addressMap[v.Name] = v
	}
	for _, v := range addresses.ECAddresses {
		addressMap[v.Name] = v
	}
	return addressMap, nil
}

type Transaction struct {
	Key                  string
	Timestamp            string
	FactoidBalanceChange float64
	ECBalanceChange      float64
}

func GetTransactions() ([]*Transaction, error) {
	fmt.Println("GetTransactions")
	keys, transactions, err := Wallet.GetTransactions()
	if err != nil {
		fmt.Printf("Error - %v\n", err)
		return nil, err
	}
	fmt.Printf("%v, %v\n", len(keys), len(transactions))
	/*
		addressMap, err := GetAddressMapByAddress()
		if err != nil {
			return nil, err
		}*/

	for _, v := range transactions {
		//var FactoidBalanceDelta float64 = 0
		//	var ECBalanceDelta float64 = 0

		ins := v.GetInputs()
		for _, i := range ins {
			fmt.Printf("%v\n", i.String())
		}
	}
	return nil, nil
}
