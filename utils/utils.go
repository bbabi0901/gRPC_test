package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-gota/gota/dataframe"
)

var logFn = log.Panic

func HandleErr(err error) {
	if err != nil {
		logFn(err)
	}
}

func CheckCode(resp *http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalf("Status code: %v", resp.StatusCode)
	}
}

func GetResponse(url string, i int) *http.Response {
	resp, err := http.Get(url + fmt.Sprintf("%d", i))
	HandleErr(err)
	CheckCode(resp)

	return resp
}

func CreateDfByCSV(df dataframe.DataFrame, fileName string) {
	f, err := os.Create(fileName)
	HandleErr(err)
	HandleErr(df.WriteCSV(f))
}

func CreateDfByJSON(df dataframe.DataFrame, fileName string) {
	f, err := os.Create(fileName)
	HandleErr(err)
	HandleErr(df.WriteJSON(f))
}

func UnmarshalResponse(i interface{}, resp *http.Response) {
	data, err := ioutil.ReadAll(resp.Body)
	HandleErr(err)

	HandleErr(json.Unmarshal(data, i))
}
