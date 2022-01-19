package iot

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iot"
)

var AWSClient *iot.Client

func Connect() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	AWSClient = iot.NewFromConfig(cfg)
}

func GetAWSClient() *iot.Client {
	return AWSClient
}
