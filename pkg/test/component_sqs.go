package test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
	"net/http"
	"time"
)

var SqsClient *sqs.SQS

func getSqsClient(port int) *sqs.SQS {
	if SqsClient != nil {
		return SqsClient
	}

	host := fmt.Sprintf("http://localhost:%d", port)

	config := &aws.Config{
		Region:   aws.String(endpoints.EuCentral1RegionID),
		Endpoint: aws.String(host),
		HTTPClient: &http.Client{
			Timeout: 1 * time.Minute,
		},
	}

	sess, err := session.NewSession(config)

	if err != nil {
		logErr(err, "could not create sqs client: %s")
	}

	SqsClient = sqs.New(sess)

	return SqsClient
}

type sqsConfig struct {
	Port int `mapstructure:"port"`
}

func runSqsContainer(name string, config configInput) {
	wait.Add(1)
	go doRunSqs(name, config)
}

func doRunSqs(name string, configMap configInput) {
	defer wait.Done()
	defer log.Printf("%s component of type %s is ready", name, "sqs")

	config := &sqsConfig{}
	unmarshalConfig(configMap, config)

	runContainer("gosoline-test-sqs-"+name, ContainerConfig{
		Repository: "localstack/localstack",
		Tag:        "0.10.3",
		Env: []string{
			"SERVICES=sqs",
		},
		PortBindings: PortBinding{
			"4576/tcp": fmt.Sprint(config.Port),
		},
		HealthCheck: func() error {
			sqsClient := getSqsClient(config.Port)

			_, err := sqsClient.ListQueues(&sqs.ListQueuesInput{})

			return err
		},
	})
}
