package test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
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
