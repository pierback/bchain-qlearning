package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/xujiajun/nutsdb"

	ut "github.com/pierback/bchain-qlearning/pkg/utils"
)

var (
	db     *nutsdb.DB
	bucket string
	err    error
)

//ID any struct that needs to persist should implement this function defined
//in Entity interface.
func (c Qlearner) ID() (jsonField string, value interface{}) {
	value = c.Usr
	jsonField = "usr"
	return
}

type Qlearner struct {
	Usr string `json:"usr"`
	Qv  qlvals `json:"qv"`
}

type qlvals struct {
	Qtable  string `json:"qt"`
	Epsilon string `json:"ep"`
	Negs    int    `json:"negs"`
	Wa      []int  `json:"Wk_negs"`
}

var (
	jsonFile *os.File
)

//StartDB inits db
func StartDB() {
	jsonFile, _ = os.Create("./Users.json")
}

//SaveQl saves qt and current epsilon val
func SaveQl(usr string, qt []byte, ep string, ng int, wa []int) {
	qlvs := &qlvals{Qtable: string(qt[:]), Epsilon: ep}

	_, filename, _, _ := runtime.Caller(0)
	bvglJSON := createSCJson(qlvs)
	file := usr + "-ql.json"
	bvglJSONDir := path.Join(path.Dir(filename), file)

	err := ioutil.WriteFile(bvglJSONDir, bvglJSON, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
	}

	err = ut.PostFile(bvglJSONDir, os.Getenv("UPID"))
	ut.PrintError(err)

	err = os.Remove(bvglJSONDir)
}

func GetQl(usr string) map[string]interface{} {
	var result map[string]interface{} = ut.DownloadFile(usr + "-ql.json")
	return result
}

func createSCJson(qv *qlvals) []byte {
	output, err := json.MarshalIndent(&qv, "", "\t\t")
	if err != nil {
		log.Println("Error marshalling to JSON:", err)
	}
	return output
}
