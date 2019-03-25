package database

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/Jeffail/gabs"
	"github.com/xujiajun/nutsdb"
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
}

var (
	jsonFile *os.File
)

//StartDB inits db
func StartDB() {
	jsonFile, _ = os.Create("./Users.json")
}

//SaveQl saves qt and current epsilon val
func SaveQl(usr string, qt []byte, ep string) {
	/* byteValue, _ := ioutil.ReadAll(jsonFile)
		jsonParsedObj, _ := gabs.ParseJSON(byteValue)
	  jsonObj, _ := gabs.Consume(jsonParsedObj)
	*/
	jsonObj := gabs.New()

	qlvs := &qlvals{Qtable: string(qt[:]), Epsilon: ep}
	jsonObj.Set(qlvs, usr)

	log.Println(jsonObj.String())

	// write to JSON file

	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write([]byte(jsonObj.String()))
	jsonFile.Close()
	log.Println("JSON data written to ", jsonFile.Name())
}

func GetQl(usr string) (string, string) {
	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonParsed, _ := gabs.ParseJSON(byteValue)
	value, _ := jsonParsed.Search(usr).Data().(qlvals)

	return value.Qtable, value.Epsilon
}
