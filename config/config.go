package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"fmt"
	"path/filepath"
	"net/http"
	"bytes"
	"strconv"
)

type SQL struct {
	NAME     string `json:"name"`
	USER     string `json:"user"`
	PASSWORD string `json:"password"`
	HOST     string `json:"host"`
	PORT     string `json:"port"`
	DATABASE string `json:"database"`
	LOCAL    string `json:"local"`
}

type ENV struct {
	PATH string `json:"path"`
	PORT string `json:"port"`
	SQL  SQL `json:"SQL"`
}

var Env ENV

func CheckError(err error, errString string) {
	if err != nil {
		fmt.Printf(errString)
		// panic(err)
	}
}

func SetCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type\n")
}

func ByteToHex(data []byte) string {

	buffer := new(bytes.Buffer)
	for _, b := range data {

		s := strconv.FormatInt(int64(b & 0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(s)
	}

	return buffer.String()
}

func init() {

	file, _ := exec.LookPath(os.Args[0])

	path, _ := filepath.Abs(file + "/../")

	envFile, err := os.Open(path + "/env.json")
	CheckError(err, "open env.json fail")

	defer envFile.Close()

	envString, err := ioutil.ReadAll(envFile)
	CheckError(err, "read envFile fail")

	errJson := json.Unmarshal([]byte(envString), &Env)
	CheckError(errJson, "Json Unmarshal fail")

}