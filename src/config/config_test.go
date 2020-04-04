/**
* @Author: zy
* @Date: 2020/04/04 15:00
 */
package config

import (
	"testing"
)

var yamlConfig = "../../config/config.yaml"

func TestGetConfigFile(t *testing.T) {
	err := InitConfig(yamlConfig)
	if err != nil {
		t.Error(err)
	}

	file := GetConfigFile()
	if file == "" {
		t.Error("get configFile error")
	}
	t.Log("configFile:", file)
}

func TestGetRestfulListenAddress(t *testing.T) {
	err := InitConfig(yamlConfig)
	if err != nil {
		t.Error(err)
	}

	address := GetRestfulListenAddress()
	if address == "" {
		t.Error("get channelanme error")
	}
	t.Log("address:", address)
}

func TestGetFabricDocker(t *testing.T) {
	err := InitConfig(yamlConfig)
	if err != nil {
		t.Error(err)
	}

	peers, err := GetPeers()
	if err != nil {
		t.Error(err)
	}
	t.Log("peers:", peers)
}
