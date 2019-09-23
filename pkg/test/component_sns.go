package test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
	"net/http"
	"time"
)

var SnsClient *sns.SNS

func getSnsClient(port int) *sns.SNS {
	if SnsClient != nil {
		return SnsClient
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
		logErr(err, "could not create sns client: %s")
	}

	SnsClient = sns.New(sess)

	return SnsClient
}

type snsConfig struct {
	Port           int    `mapstructure:"port"`
	SqsEndpoint    string `mapstructure:"sqs_endpoint"`
	LambdaEndpoint string `mapstructure:"lambda_endpoint"`
}

func runSnsContainer(name string, config configInput) {
	wait.Add(1)
	go doRunSns(name, config)
}

func doRunSns(name string, configMap configInput) {
	defer wait.Done()
	defer log.Printf("%s component of type %s is ready", name, "sns")

	config := &snsConfig{}
	unmarshalConfig(configMap, config)

	runContainer("gosoline_test_sns_"+name, ContainerConfig{
		Repository: "localstack/localstack",
		Tag:        "0.10.3",
		Env: []string{
			"SERVICES=sns",
			"SQS_BACKEND=" + config.SqsEndpoint,
			"LAMBDA_BACKEND=" + config.LambdaEndpoint,
		},
		PortBindings: PortBinding{
			"4575/tcp": fmt.Sprint(config.Port),
		},
		HealthCheck: func() error {
			snsClient := getSnsClient(config.Port)

			_, err := snsClient.ListTopics(&sns.ListTopicsInput{})

			return err
		},
	})
}
