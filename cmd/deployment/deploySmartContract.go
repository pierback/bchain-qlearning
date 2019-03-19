package deployment

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	en "github.com/pierback/bchain-qlearning/cmd/environment"
	bl "github.com/pierback/bchain-qlearning/internal/contracts/BeverageList"
	cc "github.com/pierback/bchain-qlearning/internal/contracts/CoffeeCoin"
	ccp "github.com/pierback/bchain-qlearning/internal/contracts/CoffeeCoinParent"
	pt "github.com/pierback/bchain-qlearning/internal/contracts/Parent"

	ut "github.com/pierback/bchain-qlearning/pkg/utils"
)

const key = `{"address":"e8816898d851d5b61b7f950627d04d794c07ca37","crypto":{"cipher":"aes-128-ctr","ciphertext":"1ff4add6955cba7ddaf29f66d7d21c5e1d714ef6191fbc651ae60f2ea3c95e8f","cipherparams":{"iv":"3ff869fbdbe1a523cdb327780365976e"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"7372dbae5fb318f8684902e099c311d4188721d677974d729711762c7ef6030c"},"mac":"485fa5dc701067782baa1589716a53110c7f917eb259e35ebca7265bbb7150b1"},"id":"89edb004-5b00-4607-a3af-a0d9ab9b1c34","version":3}`

func DeploySC() {
	client := ut.GetClientConnection()
	auth, err := bind.NewTransactor(strings.NewReader(key), "password")

	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	if *en.DplFlag == "cffcn" {
		// ccDeploy(auth, client)
		TestCC(auth)
	} else {
		bvglDeploy(auth, client)
	}
	// bvglDeploy(auth, client)
	// parentDeploy(auth, client)
}

func ccParentDeploy(auth *bind.TransactOpts, client *ethclient.Client) {
	/* address, _, _, err1 := pt.DeployCoffeCoinParent(auth, client)
	if err1 != nil {
		fmt.Println("err1: ", err1)
		log.Fatal(err1)
	}

	fmt.Printf("Contract CC-Parent pending deploy: 0x%x\n", address)

	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../..", "smart-contracts", "CoffeeCoin", "contractAddress")

	bvglJSON := createSCJson(address.Hex(), string(cc.CoffeecoinParentABI))
	bvglJSONDir := path.Join(path.Dir(filename), "ccParent.json")

	err := ioutil.WriteFile(bvglJSONDir, bvglJSON, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
	}

	err = ut.PostFile(bvglJSONDir, os.Getenv("UPID"))
	ut.PrintError(err)

	err = os.Remove(bvglJSONDir)

	if err != nil {
		fmt.Println(err)
	} */
}

func bvglParentDeploy(auth *bind.TransactOpts, client *ethclient.Client) {
	/* address, _, _, err1 := pt.DeployBeverageListParent(auth, client)
	if err1 != nil {
		fmt.Println("err1: ", err1)
		log.Fatal(err1)
	}

	fmt.Printf("Contract BeverageList Parent pending deploy: 0x%x\n", address)

	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../..", "smart-contracts", "BeverageList", "contractAddress")

	// bvglJSON := createBvglJson(address)

	err12 := ioutil.WriteFile(dir, address.Bytes(), 0644)
	check(err12)

	bvglJSON := createSCJson(address.Hex(), string(bl.BeveragelistParentABI))
	bvglJSONDir := path.Join(path.Dir(filename), "bvglParent.json")

	err := ioutil.WriteFile(bvglJSONDir, bvglJSON, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
	}

	err = ut.PostFile(bvglJSONDir, os.Getenv("UPID"))
	ut.PrintError(err)

	err = os.Remove(bvglJSONDir)

	if err != nil {
		fmt.Println(err)
	} */
}

func upgradeContract(auth *bind.TransactOpts, client *ethclient.Client) {
	address := common.HexToAddress("0xa5073710ee54574b77bfe8faebebe07823dc7dac")
	instance, err := pt.NewParent(address, client)

	bvglAddress := common.HexToAddress("0xf189078a9969cc0247b01437d2a6deb7be83e7bf")

	key := [32]byte{}
	copy(key[:], []byte("1"))

	if *en.DplFlag == "upgrade" {
		_, err = instance.RegisterBeverageList(auth, key, bvglAddress)
	} else {
		_, err = instance.UpgradeBeverageList(auth, key, bvglAddress)
	}
	time.Sleep(2 * time.Second)
	check(err)

	drink := [32]byte{}
	weekday := [32]byte{}
	ti := [32]byte{}
	copy(drink[:], []byte("mate"))
	copy(weekday[:], []byte("dienstag"))
	copy(ti[:], "2019-03-15T21:28:05")
	// copy(weekday[:], []byte(time.Now().Weekday().String()))
	// copy(ti[:], []byte(time.Now().String()))

	/* usr := common.HexToAddress("e8816898d851d5b61b7f950627d04d794c07ca37")
	bvgInstance, err1 := bl.NewBeveragelist(bvglAddress, client)
	useror, err := bvgInstance.IsUser(nil, usr)
	fmt.Println("useror: ", useror)
	check(err1)
	check(err)

	dsa, _ := bvgInstance.DataStorage(nil)
	fmt.Println("DataStorage: ", dsa.Hex())

	_, err = bvgInstance.SetDrinkData(auth, ti, drink, weekday)
	check(err) */

	time.Sleep(10 * time.Second)

	// drink1, weekday1, err := bvgInstance.GetDrinkData(nil, ti)
	// fmt.Println("getFromTime drink", drink1)
	/* fmt.Println("getFromTime drink", string(drink1[:]))
	// fmt.Println("getFromTime weekday", weekday1)
	fmt.Println("getFromTime weekday", string(weekday1[:]))
	check(err) */

	/*

		time.Sleep(10 * time.Second) */
	_ = instance
}

func bvglDeploy(auth *bind.TransactOpts, client *ethclient.Client) {
	// parentDeploy(auth, client)

	address, _, _, err1 := bl.DeployBeveragelist(auth, client)
	if err1 != nil {
		fmt.Println("err1: ", err1)
		log.Fatal(err1)
	}

	fmt.Printf("Contract Beveragelist pending deploy: 0x%x\n", address)

	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../..", "smart-contracts", "CoffeeCoin", "contractAddress")

	// bvglJSON := createBvglJson(address)

	err12 := ioutil.WriteFile(dir, address.Bytes(), 0644)
	check(err12)

	bvglJSON := createSCJson(address.Hex(), string(bl.BeveragelistABI))
	bvglJSONDir := path.Join(path.Dir(filename), "bvgl.json")

	err := ioutil.WriteFile(bvglJSONDir, bvglJSON, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
	}

	err = ut.PostFile(bvglJSONDir, os.Getenv("UPID"))
	ut.PrintError(err)

	err = os.Remove(bvglJSONDir)

	if err != nil {
		fmt.Println(err)
	}
}

type SCJson struct {
	Address string `json:"address"`
	Abi     string `json:"abi"`
}

func createSCJson(address string, abi string) []byte {
	data := SCJson{
		Address: address,
		Abi:     abi,
	}

	output, err := json.MarshalIndent(&data, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
	}
	return output
}

func deployCCParent(auth *bind.TransactOpts, client *ethclient.Client, ccAddress common.Address) {

	// check if there is address in file if not deploy if so upgrade CoffeeCoin
	var result map[string]interface{} = ut.DownloadFile("ccParent.json")
	fmt.Println("result: ", result["address"])
	var parentAddress common.Address

	//deploy parent
	if result["address"] == nil {

		parentAddress, _, _, err1 := ccp.DeployCoffeecoinparent(auth, client)
		if err1 != nil {
			fmt.Println("err1: ", err1)
			log.Fatal(err1)
		}

		fmt.Printf("Contract CC-Parent pending deploy: 0x%x\n", parentAddress)

		_, filename, _, _ := runtime.Caller(0)
		// dir := path.Join(path.Dir(filename), "../../..", "smart-contracts", "CoffeeCoin", "contractparentAddress")

		ccpJSON := createSCJson(parentAddress.Hex(), string(ccp.CoffeecoinparentABI))
		ccpJSONDir := path.Join(path.Dir(filename), "ccParent.json")

		err := ioutil.WriteFile(ccpJSONDir, ccpJSON, 0644)
		if err != nil {
			fmt.Println("Error writing JSON to file:", err)
		}

		err = ut.PostFile(ccpJSONDir, os.Getenv("UPID"))
		ut.PrintError(err)

		err = os.Remove(ccpJSONDir)

		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(5 * time.Second)
	} else {
		parentAddress = common.HexToAddress(result["address"].(string))
	}

	instance, err := ccp.NewCoffeecoinparent(parentAddress, client)
	check(err)
	// bvglAddress := common.HexToAddress("0xf189078a9969cc0247b01437d2a6deb7be83e7bf")

	key := [32]byte{}
	copy(key[:], []byte("1"))

	_, err = instance.RegisterCoffeCoin(auth, key, ccAddress)
	check(err)
}

func TestCC(auth *bind.TransactOpts) {
	client := ut.GetClientConnection()
	contractAddress := common.HexToAddress("0x33a1bdf303d830d7264003dd9789a7f656c5c5e0")
	ccInstance, _ := cc.NewCoffeecoin(contractAddress, client)

	_, err := ccInstance.PayCoffee(auth)
	check(err)
	time.Sleep(5 * time.Second)

	ob, _ := ccInstance.GetChairBalance(nil)

	fmt.Println("own balance: ", ob)

	price, _ := ccInstance.GetCoffeePrice(nil)

	fmt.Println("price ", price)
}

func ccDeploy(auth *bind.TransactOpts, client *ethclient.Client) {
	chairAddress := common.HexToAddress("0x18ef96d887954472de5e9f47d60ba8dea371dbfe")
	coffeePrice := new(big.Int).SetUint64(3)
	matePrice := new(big.Int).SetUint64(5)
	waterPrice := new(big.Int).SetUint64(2)

	address, _, _, err1 := cc.DeployCoffeecoin(auth, client)
	if err1 != nil {
		fmt.Println("err1: ", err1)
		log.Fatal(err1)
	}
	fmt.Printf("Contract CoffeCoin pending deploy: 0x%x\n", address)
	time.Sleep(5 * time.Second)

	deployCCParent(auth, client, address)
	ccInstance, _ := cc.NewCoffeecoin(address, client)

	ccInstance.SeInitValues(auth, chairAddress, coffeePrice, matePrice, waterPrice)

	time.Sleep(5 * time.Second)

	cprice, _ := ccInstance.GetCoffeePrice(nil)
	fmt.Println("cprice: ", cprice)

	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../..", "smart-contracts", "CoffeeCoin", "contractAddress")

	err12 := ioutil.WriteFile(dir, address.Bytes(), 0644)

	check(err12)

	bvglJSON := createSCJson(address.Hex(), string(cc.CoffeecoinABI))
	bvglJSONDir := path.Join(path.Dir(filename), "cc.json")

	err := ioutil.WriteFile(bvglJSONDir, bvglJSON, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
	}

	err = ut.PostFile(bvglJSONDir, os.Getenv("UPID"))
	ut.PrintError(err)

	err = os.Remove(bvglJSONDir)

	if err != nil {
		fmt.Println(err)
	}
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

//old approach with privatekey did not work
func StartDeploy() {
	fmt.Println("StartDeploy")
	// connect to an ethereum node  hosted by infura
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasPrice: ", gasPrice)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)         // in wei
	auth.GasLimit = uint64(3141500000) // in units
	auth.GasPrice = gasPrice

	/* _st := "coffee"
		_lr := big.NewInt(700000)
		_gm := big.NewInt(700000)
		_ep := big.NewInt(700000)
		_epd := big.NewInt(700000)
		_sr := big.NewInt(0)
		_pd := big.NewInt(0)
	address, tx, instance, err1 := cfd.DeployCoffeedash(
			auth,
			client, _st,
			_lr,
			_gm,
			_ep,
			_epd,
			_sr,
			_pd,
		)
		if err1 != nil {
			fmt.Println("err1: ", err1)
			log.Fatal(err)
		}

		fmt.Printf("Contract pending deploy: 0x%x\n", address)
		fmt.Println(address.Hex()) // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
		fmt.Println(tx.Hash().Hex())

		_ = instance */
}
