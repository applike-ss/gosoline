package test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
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
