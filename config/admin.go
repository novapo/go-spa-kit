package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// AdminConfig holds config data for the backend
type AdminConfig struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// ReadAdminConf reads admin login data from backend and returns it as struct
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
