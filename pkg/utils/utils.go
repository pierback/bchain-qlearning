package utils

import (
	"bytes"
	"encoding/json"
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
	var result map[string]interface{} = DownloadFile("bvgl.json")
	log.Println("result: ", result["address"])

	fmt.Printf("\nContract Address 0x%s\n", result["address"])
	return common.HexToAddress(result["address"].(string))
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

func DownloadIP() string {
	return "http://" + os.Getenv("SIP") + ":9090/files/"
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

func DownloadFile(filename string) map[string]interface{} {
	url := DownloadIP() + filename
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Println(readErr)
	}

	var result map[string]interface{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Println(jsonErr)
	}

	return result
}

func PostFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", path.Base(filename))
	if err != nil {
		log.Println("error writing to buffer")
		return err
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		log.Println("error opening file")
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
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}
