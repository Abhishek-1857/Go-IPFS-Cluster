package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
	shell "github.com/ipfs/go-ipfs-api"
)

type Users struct {
	Cluster string `json:"cluster"`
}

func main() {
	var user Users
    buf := new(bytes.Buffer)
	//Clusters
	india := shell.NewShell("localhost:5001")
	pakistan := shell.NewShell("localhost:5002")

	// tsd := &Data{
	// 	Id:    1,
	// 	Value: 123,
	// }
	content, _ := os.ReadFile("KycInfo.json")
	toMarshal, _ := os.Open("KycInfo.json")
	data, _ := ioutil.ReadAll(toMarshal)
	json.Unmarshal(data, &user)
	// tsdBin, _ := json.Marshal(tsd)
	if user.Cluster == "india" {
		reader := bytes.NewReader(content)
		cid, err := india.Add(reader)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err)
			os.Exit(1)
		}
		fmt.Printf("added to indian cluster %s\n", cid)

		data, err := india.Cat(cid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	buf.ReadFrom(data)
	newStr:= buf.String()
	buf.Reset()
	fmt.Println(newStr)


	} else if user.Cluster == "pakistan" {
		reader := bytes.NewReader(content)
		cid, err := pakistan.Add(reader)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err)
			os.Exit(1)
		}
		fmt.Printf("added to pakistan cluster %s\n", cid)


		data, err := pakistan.Cat(cid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	buf.ReadFrom(data)
	newStr:= buf.String()
	buf.Reset()
	fmt.Println(newStr)

	}

	// content2, _ := os.ReadFile("bawa2.txt")
	// reader2 := bytes.NewReader(content2)

	// cid, err := india.Add(reader)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %s", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("added %s\n", cid)

	// cid2, err := pakistan.Add(reader2)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %s", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("added %s\n", cid2)

	// Get the data from IPFS and output the contents

	// data1, err := india.Cat(cid)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %s", err)
	// 	os.Exit(1)
	// }

	// data2, err := pakistan.Cat(cid2)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %s", err)
	// 	os.Exit(1)
	// }

	// ...so we convert it to a string by passing it through
	// a buffer first. A 'costly' but useful process.
	// https://golangcode.com/convert-io-readcloser-to-a-string/
	// buf := new(bytes.Buffer)

	// buf.ReadFrom(data1)
	// newStr1 := buf.String()
	// buf.Reset()
	// fmt.Println(newStr1)

	// buf.ReadFrom(data2)
	// newStr2 := buf.String()
	// buf.Reset()
	// fmt.Println(newStr2)

	// res := &TimeSeriesDatum{}
	// json.Unmarshal([]byte(newStr), &res)
	// fmt.Println(res)

}
