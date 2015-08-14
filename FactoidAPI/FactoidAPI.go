package FactoidAPI

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/factoid"
	"os"
	"strconv"
	"regexp"
	"bytes"
)


var serverfactoid = "localhost:8089"
var badChar,_ = regexp.Compile("[^A-Za-z0-9_-]")

type Response struct {
	Response string
	Success bool
}

func SendCommand(get bool, cmd string) (*Response, error) {
	var err error
	var resp *http.Response
	if get == true {
		resp, err = http.Get(cmd)
	} else {
		resp, err = http.PostForm(cmd, nil)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp.Body.Close()

	b := new(Response)
	if err := json.Unmarshal(body, b); err != nil {
		fmt.Printf("Failed to parse the response from factomd: %s\n",body)
		return nil, fmt.Errorf("blabla")
	}
	
	fmt.Printf("%v\n", b)
	
	if !b.Success {
		return nil, fmt.Errorf("blabla")
	}
	  
	return b, nil
}

func ValidateKey(key string) (msg string, valid bool) {
	if len(key) > factoid.ADDRESS_LENGTH	 { 
		return "Key is too long.  Keys must be less than 32 characters", false	 
	}
	if badChar.FindStringIndex(key)!=nil { 
		str := fmt.Sprintf("The key or name '%s' contains invalid characters.\n"+
		"Keys and names are restricted to alphanumeric characters,\n"+
		"minuses (dashes), and underscores", key)
		return str, false
	}
	return "", true
}


// Generates a new Address
func GenerateAddress(addressType, addressName string) error {
	var err error
	var Addr string
	switch addressType {
		case "ec": 
			Addr, err= factom.GenerateEntryCreditAddress(addressName)
		case "factoid":
			Addr, err= factom.GenerateFactoidAddress(addressName)
		default:
			panic("Expected ec|factoid name")
	}
	
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(addressType," = ",Addr)
	return nil
}

func GetAddresses() string {
	str := fmt.Sprintf("http://%s/v1/factoid-get-addresses/", serverfactoid)
	resp, _:=SendCommand(true, str)
	return resp.Response
}

func GetTransactions() {
	str := fmt.Sprintf("http://%s/v1/factoid-get-transactions/", serverfactoid)
	SendCommand(true, str)
	return 
}


func FactoidNewTransaction(key string) {
	msg, valid := ValidateKey(key) 
	if !valid {
		fmt.Println(msg)
		os.Exit(1)
	}
	str := fmt.Sprintf("http://%s/v1/factoid-new-transaction/%s", serverfactoid, key)
	SendCommand(false, str)
	return 
	
}

func FactoidDeleteTransaction(tx string) {
	msg, valid := ValidateKey(tx) 
	if !valid {
		fmt.Println(msg)
		os.Exit(1)
	}
	str := fmt.Sprintf("http://%s/v1/factoid-delete-transaction/%s", serverfactoid, tx)
	SendCommand(false, str)
	
	return
}

func FactoidAddFee(key, name string) {	
	msg, valid := ValidateKey(key) 
	if !valid {
		fmt.Println(msg)
		os.Exit(1)
	}
	
	msg, valid = ValidateKey(name) 
	if !valid {
		fmt.Println(msg)
		os.Exit(1)
	}
	
	str := fmt.Sprintf("http://%s/v1/factoid-add-fee/?key=%s&name=%s", 
					   serverfactoid, key, name)
	SendCommand(false, str)
	
	return 
}

func FactoidAddInput(key, name, amount string) {
	msg, valid := ValidateKey(key) 
	if !valid {
		fmt.Println(msg)
		os.Exit(1)
	}
	
	amt, err := factoid.ConvertFixedPoint(amount)
	if err != nil { 
		fmt.Println(err)
		os.Exit(1) 
	}
	
	ramt, err := strconv.ParseInt(amt,10,64)
	if err != nil { 
		fmt.Println(err)
		os.Exit(1) 
	}
 
	_, err = factoid.ValidateAmounts(uint64(ramt))
	if err != nil { 
		fmt.Println(err)
		os.Exit(1)
	}

	str := fmt.Sprintf("http://%s/v1/factoid-add-input/?key=%s&name=%s&amount=%s", 
					   serverfactoid, key, name, amt)
	SendCommand(false, str)
	
	return 
}


func FactoidAddOutput(key, name, amount string) {
	msg, valid := ValidateKey(key) 
	if !valid {
		fmt.Println(msg)
		os.Exit(1)
	}

	msg, valid = ValidateKey(name) 
	if !valid {
		fmt.Println(msg)
		os.Exit(1)
	}
	
	amt, err := factoid.ConvertFixedPoint(amount)
	if err != nil { 
		fmt.Println("Invalid format for a number: ",amount)
		os.Exit(1) 
	}
	
	ramt, err := strconv.ParseInt(amt,10,64)
	if err != nil { 
		fmt.Println(err)
		os.Exit(1) 
	}

	_, err = factoid.ValidateAmounts(uint64(ramt))
	if err != nil { 
		fmt.Println(err) 
		os.Exit(1)
	}
	
	
	str := fmt.Sprintf("http://%s/v1/factoid-add-output/?key=%s&name=%s&amount=%s", 
					   serverfactoid, key, name, amt)
	SendCommand(false, str)
	
	return 
}

func FactoidAddECOutput(key, name, amount string) {
	msg, valid := ValidateKey(key) 
	if !valid {
		fmt.Println(msg)
		os.Exit(1)
	}
	
	msg, valid = ValidateKey(name) 
	if !valid {
		fmt.Println(msg)
		os.Exit(1)
	}
	
	amt, err := factoid.ConvertFixedPoint(amount)
	if err != nil { 
		fmt.Println(err)
		os.Exit(1)  
	}
	
	ramt, err := strconv.ParseInt(amt,10,64)
	if err != nil { 
		fmt.Println(err)
		os.Exit(1) 
	}

	_, err = factoid.ValidateAmounts(uint64(ramt))
	if err!=nil { 
		fmt.Println(err)
		os.Exit(1)
	}
	
	str := fmt.Sprintf("http://%s/v1/factoid-add-ecoutput/?key=%s&name=%s&amount=%s", 
					   serverfactoid, key, name, amt)
	SendCommand(false, str)
	
	return 
}

func FactoidGetFee() {
	resp, err := http.Get(fmt.Sprintf("http://%s/v1/factoid-get-fee/",serverfactoid))
	if err != nil {
		fmt.Println("Command Failed Get")
		os.Exit(1) 
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Command Failed")
		os.Exit(1) 
	}
	resp.Body.Close()
	
	// We pull the fee.  If the fee isn't positive, or if we fail to marshal, then there is a failure
	type x struct { Fee int64 }
	b := new(x)
	b.Fee = -1
	if err := json.Unmarshal(body, b); err != nil || b.Fee == -1 {
		fmt.Println("Command Failed")
		os.Exit(1)
	}
	tv := b.Fee/100000000
	lv := b.Fee-(tv*100000000)
	r := fmt.Sprintf("Fee: %d.%08d",tv,lv)
	var i int; for i=len(r)-1; r[i]=='0'; i-- {}
	if string(r[i])=="." { i +=1 }
	fmt.Println(r[:i+1])
	 
}


func FactoidSign(key string) error {
	msg, valid := ValidateKey(key) 
	if !valid {
		fmt.Println(msg)
		return fmt.Errorf("%v", msg)
	}
	
	str := fmt.Sprintf("http://%s/v1/factoid-sign-transaction/%s", serverfactoid, key)
	SendCommand(false, str)
	return nil
}


func FactoidSubmit(key string) {
	msg, valid := ValidateKey(key) 
	if !valid {
		fmt.Println(msg)
		os.Exit(1)
	}
	
	s := struct{Transaction string}{key}
	
	jdata, err := json.Marshal(s)
	if err != nil {
		fmt.Println("Submitt failed")
		os.Exit(1) 
	}
	
	str:=fmt.Sprintf("http://%s/v1/factoid-submit/%s", serverfactoid,bytes.NewBuffer(jdata))
	SendCommand(false, str)
}

func FactoidSetup(transaction string) {
	s := struct{Transaction string}{transaction}
	
	jdata, err := json.Marshal(s)
	if err != nil {
		fmt.Println("Submitt failed")
		os.Exit(1) 
	}
	
	str:=fmt.Sprintf("http://%s/v1/factoid-setup/%s", serverfactoid,bytes.NewBuffer(jdata))
	SendCommand(false, str)
}
