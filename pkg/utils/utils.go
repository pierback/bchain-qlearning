package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	_, filename, _, _ = runtime.Caller(0)
)

func GetLatestContractAddress() common.Address {
	dir := path.Join(path.Dir(filename), "../../..", "smart-contracts", "BeverageList", "contractAddress")
	dat, err := ioutil.ReadFile(dir)
	fmt.Printf("Contract Address 0x%s\n", common.Bytes2Hex(dat))
	if err != nil {
		fmt.Println("Error ReadFile:", err)
	}
	return common.BytesToAddress(dat)
}

//GetLocalIP get local ip
func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

//GetEthWsAddr get websocket eth address
func GetEthWsAddr() string {
	ip := GetLocalIP()
	ipStr := fmt.Sprintf("%s", ip)
	return "ws://" + ipStr + ":8546"
}

func GetUploadIP() string {
	return "http://" + os.Getenv("SIP") + ":9090/upload"
}

func GetBchainIP() string {
	return "ws://" + os.Getenv("SIP") + ":8546"
}

//GetClientConnection returns eth client
func GetClientConnection() *ethclient.Client {
	address := os.Getenv("BCIP")
	if address == "" {
		dir := path.Join(path.Dir(filename), "../../..", "private-net-docker", "bchainIp")
		dat, _ := ioutil.ReadFile(dir)
		address = strings.TrimSuffix(fmt.Sprintf("%s", dat), "\n")
	}

	client, err := ethclient.Dial(address)
	PrintError(err)
	return client
}

//PrintError printing error
func PrintError(err error) {
	if err != nil {
		log.Printf("%s", err)
	}
}

func PostFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", path.Base(filename))
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}
	fmt.Println("filename: ", filename)

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}
