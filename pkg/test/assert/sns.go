package assert

import (
	"github.com/applike/gosoline/pkg/test"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/stretchr/testify/assert"
	"testing"
)

func SnsTopicExists(t *testing.T, topicArn string) {
	assert.NotNil(t, test.SnsClient)
	getTopicAttributesOutput, err := test.SnsClient.GetTopicAttributes(&sns.GetTopicAttributesInput{
		TopicArn: &topicArn,
	})

	assert.NotNil(t, getTopicAttributesOutput)
	assert.NoError(t, err)
}
