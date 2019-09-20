package test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/hashicorp/go-multierror"
	"log"
)

type localstackConfig struct {
	SqsPort int `mapstructure:"sqs_port"`
	SnsPort int `mapstructure:"sns_port"`
}

func runLocalstackContainer(name string, config configInput) {
	wait.Add(1)
	go doRunLocalstack(name, config)
}

func doRunLocalstack(name string, configMap configInput) {
	defer wait.Done()
	defer log.Printf("%s component of type %s is ready", name, "localstack")

	config := &localstackConfig{}
	unmarshalConfig(configMap, config)

	runContainer("gosoline_test_localstack", ContainerConfig{
		Repository: "localstack/localstack",
		Tag:        "0.10.3",
		Env: []string{
			"SERVICES=sns,sqs",
		},
		PortBindings: PortBinding{
			"4575/tcp": fmt.Sprint(config.SnsPort),
			"4576/tcp": fmt.Sprint(config.SqsPort),
		},
		HealthCheck: func() error {
			snsClient := getSnsClient(config.SnsPort)
			sqsClient := getSqsClient(config.SqsPort)

			_, errSns := snsClient.ListTopics(&sns.ListTopicsInput{})
			_, errSqs := sqsClient.ListQueues(&sqs.ListQueuesInput{})

			err := multierror.Append(errSns, errSqs)
			return err.ErrorOrNil()
		},
	})
}
