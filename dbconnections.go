package main

import (
	"fmt"
	"github.com/zbo14/envoke/common"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"reflect"
)

const (
	URL = "https://test.bigchaindb.com"
	public_key = "D1ACjwzXEyP8s5Yjs6xDhfhd3yM7d92jxnAznysDFC3D"
)

func postTx(postData map[string]interface{}) {
	fmt.Printf("type of is %s", reflect.TypeOf(postData))
	cmd := exec.Command("node", "index.js", string(common.MustMarshalJSON(postData)))
	out, err := cmd.CombinedOutput()
	fmt.Sprint(out)
	fmt.Sprint(err)
}
func getTx(path string, assetId string, publicKey string) string {

	type respArray struct {
		data map[string] interface{}
	}

	fmt.Printf("path is %s \n\n\n", path)
	client := http.Client{}
	request, err := http.NewRequest("GET", "https://test.bigchaindb.com/api/v1" + path, nil)
	request.Header.Add("app_id", "82826d8b")
	request.Header.Add("app_key", "8ea557d8f236db15626c7d68a04eca6b")
	request.Header.Add("Content-Type", "application/json")
	q := request.URL.Query()
	if (assetId != "") {
		q.Add("asset_id", assetId)
	}
	if (publicKey != "") {
		q.Add("public_key", publicKey)
	}
	if err != nil {
		log.Fatalln(err)
	}
	request.URL.RawQuery = q.Encode()

	resp, err := client.Do(request)
	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(request.URL.String())
	fmt.Print("\n\n")
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("result is %s \n\n\n", string(bodyBytes))
	fmt.Printf("error is %s \n\n\n", err)
	//var target respArray
	//json.Unmarshal(bodyBytes, &target)
	return string(bodyBytes)
	//headerBytes, err := ioutil.ReadFile("../fixtures/dbd_key.json")
	//print(err)
	//
	//var h map[string]string
	//err = json.Unmarshal(headerBytes, &h)
	//cfg := cl.ClientConfig{
	//	Url:     "https://test.bigchaindb.com/api/v1/",
	//	Headers: h,
	//}
	//client, _ := cl.New(cfg)
	//txns, err := client.ListOutputs(public_key, false)
	//rtxn, err := client.ListBlocks("c3b4b84a8eacdc389dd0deb010eb467280a9c995b3b98115c0e674de5e25bce0")
	//fmt.Println(rtxn)
	//fmt.Printf("returned data is ")
	//fmt.Println(txns)
	//return  body
}
