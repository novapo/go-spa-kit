package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type AdminConfig struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func ReadAdminConf() (*AdminConfig, error) {
	adnConf := &AdminConfig{}
	byteArr, err := ioutil.ReadFile("./admin/admin.json")
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(byteArr)
	decoder := json.NewDecoder(buffer)
	if err := decoder.Decode(adnConf); err != nil {
		return nil, err
	}
	return adnConf, nil
}
