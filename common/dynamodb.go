package common

import (
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/dynamodb"
	"time"
)

func GetServer(config *Config) (*dynamodb.Server, error) {
	env, err := aws.GetAuth(config.AwsAccessKeyId, config.AwsSecretAccessKey, "", time.Time{})
	if err != nil {
		return nil, err
	}
	return &dynamodb.Server{
		Auth: env,
		Region: aws.Region{
			Name:             config.AwsRegionName,
			DynamoDBEndpoint: config.AwsDynamoDBEndpoint,
		},
	}, nil
}
