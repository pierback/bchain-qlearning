package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	dbj "github.com/sonyarouje/simdb/db"
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
	driver *dbj.Driver
)

func StartJsonDB() {
	driver, err = dbj.New("data")
	if err != nil {
		panic(err)
	}

	//GET ALL qlvals
	//opens the customer json file and filter all the customers with name sarouje.
	//AsEntity takes a pointer to qlvals array and fills the result to it.
	//we can loop through the customers array and retireve the data.
	/* var customers []qlvals
	err = driver.Open(qlvals{}).Where("name", "=", "sarouje").Get().AsEntity(&customers)
	if err != nil {
		panic(err)
	} */
}

//StartDB inits db
func StartDB() {
	// Open the database located in the /tmp/nutsdb directory.
	// It will be created if it doesn't exist.
	opt := nutsdb.DefaultOptions
	fileDir := "/tmp/nutsdb_example"

	files, _ := ioutil.ReadDir(fileDir)
	for _, f := range files {
		name := f.Name()
		if name != "" {
			err := os.Remove(fileDir + "/" + name)
			if err != nil {
				panic(err)
			}
		}
	}
	opt.Dir = fileDir
	opt.SegmentSize = 1024 * 1024 // 1MB
	db, err = nutsdb.Open(opt)
	if err != nil {
		panic(err)
	}
	bucket = "users"

}

//SaveQl saves qt and current epsilon val
func SaveQl(usr string, qt []byte, ep string) {
	qlvs := &qlvals{Qtable: string(qt[:]), Epsilon: ep}

	jsonData, err := json.Marshal(qlvs)
	if err != nil {
		fmt.Println("error on marshalling qlvs", qlvs)
	}

	fmt.Println("DB push data")

	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			key := []byte(usr)
			val := jsonData
			return tx.RPush(bucket, key, val)
		}); err != nil {
		log.Fatal(err)
	}

	q := Qlearner{
		Usr: usr,
		Qv: qlvals{
			Qtable:  string(qt[:]),
			Epsilon: ep,
		},
	}

	//creates a new qlvals file inside the directory passed as the parameter to New()
	//if the qlvals file already exist then insert operation will add the customer data to the array
	err = driver.Insert(q)
	if err != nil {
		panic(err)
	}

	if err := db.View(
		func(tx *nutsdb.Tx) error {
			key := []byte(usr)
			item, err := tx.LPeek(bucket, key)
			if err != nil {
				return err
			}

			vals := qlvals{}
			json.Unmarshal(item, &vals)
			fmt.Printf("inserted vals epsi: %s qt: %s \n", string(vals.Epsilon[:]), string(vals.Qtable[:]))
			return nil
		}); err != nil {
		log.Fatal(err)
	}
}
