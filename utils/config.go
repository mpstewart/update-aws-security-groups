package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Holds all state needed for the script
type Config struct {
	AWSProfile string `json:"awsProfile"`
	Hostname   string `json:"hostname"`
	HomeIP     string `json:"homeIP"`
	GroupID    string `json:"groupID"`
	Region     string `json:"region"`
	Ports      []Port `json:"ports"`
}

// Provide singleton access for the config
var configStash *Config

func GetConfig() (c *Config) {
	if configStash == nil {
		configStash = unmarshalConfig()
	}

	c = configStash

	return
}

// Retrieve the config from disk
func unmarshalConfig() *Config {

	file, err := os.Open("/etc/update-aws-security-groups/config")
	if err != nil {
		Logger.Fatalln(err)
	}
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)

	var config Config

	err = json.Unmarshal(bytes, &config)

	if err != nil {
		Logger.Fatalln(err)
	}

	return &config
}

// Write back out to disk
func (c *Config) Write() {
	bytes, err := json.MarshalIndent(c, "", "  ")

	if err != nil {
		Logger.Printf("Error writing config to disk:\n%s", err)
	} else {
		ioutil.WriteFile("config", bytes, 0644)
	}

	return
}
