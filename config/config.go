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
}

type ENV struct {
	PATH string
	PORT string
	SQL  SQL `json:"SQL"`
}

var Env ENV

func init() {

	file, _ := exec.LookPath(os.Args[0])

	path, _ := filepath.Abs(file + "/../")

	envFile, err := os.Open(path + "/env.json")

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer envFile.Close()

	envString, err := ioutil.ReadAll(envFile)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	errJson := json.Unmarshal([]byte(envString), &Env)

	if errJson != nil {
		fmt.Printf("err was %v", errJson)
	}

}
