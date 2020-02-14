package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
)

func GetObjectFromPath(path string, a interface{}) error{
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	file, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &a)
	if err != nil {
		return err
	}

	return nil
}