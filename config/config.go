package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"fmt"
	"path/filepath"
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