package common

import (
	"encoding/json"
	"os"
)

type Config struct {
	AwsAccessKeyId       string
	AwsSecretAccessKey   string
	AwsDynamoDBTableName string
	AwsRegionName        string
	AwsDynamoDBEndpoint  string
}

func NewConfig(path string) (*Config, error) {
	config := &Config{}
	if err := ParseJsonFile(path, config); err != nil {
		return nil, err
	}

	return config, nil
}

// Loads a json file into target interface.
func ParseJsonFile(filename string, target interface{}) error {
	reader, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer reader.Close()

	decoder := json.NewDecoder(reader)
	return decoder.Decode(target)
}
