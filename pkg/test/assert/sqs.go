package assert

import (
	"github.com/applike/gosoline/pkg/mdl"
	"github.com/applike/gosoline/pkg/test"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func SqsQueueExists(t *testing.T, queueName string) {
	assert.NotNil(t, test.SqsClient)
	queueUrlOutput, err := test.SqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})

	assert.NotNil(t, queueUrlOutput)
	assert.NoError(t, err)
}

func SqsQueueContainsMessages(t *testing.T, queueName string, count int) []*sqs.Message {
	assert.NotNil(t, test.SqsClient)
	queueUrlOutput, err := test.SqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})

	assert.NotNil(t, queueUrlOutput)
	assert.NoError(t, err)

	messages, err := test.SqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		MaxNumberOfMessages: mdl.Int64(10),
		QueueUrl:            queueUrlOutput.QueueUrl,
	})

	assert.NotNil(t, messages)
	assert.Len(t, messages.Messages, count)

	return messages.Messages
}
