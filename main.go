package main

import (
	"bytes"
	// "encoding/json"
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

type Data struct {
	Id    uint64 `json:"id"`
	Value uint64 `json:"value"`
}

func main() {
	sh := shell.NewShell("localhost:5001")




	// tsd := &Data{
	// 	Id:    1,
	// 	Value: 123,
	// }
	content, _ := os.ReadFile("bawa.txt")
	// tsdBin, _ := json.Marshal(tsd)

	reader := bytes.NewReader(content)

	cid, err := sh.Add(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s\n", cid)


	// Get the data from IPFS and output the contents 
	

	data, err := sh.Cat(cid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	// ...so we convert it to a string by passing it through
	// a buffer first. A 'costly' but useful process.
	// https://golangcode.com/convert-io-readcloser-to-a-string/
	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	newStr := buf.String()
	// res := &TimeSeriesDatum{}
	// json.Unmarshal([]byte(newStr), &res)
	// fmt.Println(res)
	fmt.Println(newStr)
}
