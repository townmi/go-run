package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type SQL struct {
	NAME     string
	USER     string
	PASSWORD string
	HOST     string
	PORT     string
	DATABASE string
}

type ENV struct {
	PATH string
	PORT string
	SQL  SQL `json:"SQL"`
}

var Env ENV

func CheckError(err error) {
	if err != nil {
		panic("open file failed!")
	}
}

func init() {

	file, _ := exec.LookPath(os.Args[0])

	path, _ := filepath.Abs(file + "/../")

	fmt.Printf("path: %v", path+"/env.json")

	envFile, err := os.Open(path + "/env.json")
	CheckError(err)

	defer envFile.Close()

	envString, err := ioutil.ReadAll(envFile)
	CheckError(err)

	errJson := json.Unmarshal([]byte(envString), &Env)
	CheckError(errJson)

}
