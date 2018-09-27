package main

import (
	"fmt"
	"github.com/zbo14/envoke/common"
	"os/exec"
	"reflect"

	"encoding/json"
	cl "github.com/bigchaindb/go-bigchaindb-driver/pkg/client"
	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"io/ioutil"
)

const (
	URL = "https://test.bigchaindb.com"
)

func postTx(postData map[string]interface{}) {
	fmt.Sprint(postData)
	fmt.Printf("type of is %s", reflect.TypeOf(postData))
	cmd := exec.Command("node", "index.js", string(common.MustMarshalJSON(postData)))
	out, err := cmd.CombinedOutput()
	fmt.Sprint(out)
	fmt.Sprint(err)
}
func getTx() []txn.Transaction {
	headerBytes, err := ioutil.ReadFile("../fixtures/dbd_key.json")
	print(err)

	var h map[string]string
	err = json.Unmarshal(headerBytes, &h)
	cfg := cl.ClientConfig{
		Url:     "https://test.bigchaindb.com/api/v1/",
		Headers: h,
	}
	client, _ := cl.New(cfg)
	txns, err := client.ListTransactions("1", "CREATE")
	return txns
}
