package utils

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func CreateSQS(sess *session.Session, name string) (queueUrl *string) {
	svc := sqs.New(sess)

	result, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(name),
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("10"),
			"MessageRetentionPeriod": aws.String("86400"),
		},
	})
	if err != nil {
		fmt.Printf("Couldn't not create queue %v\n", name)
		fmt.Println(err)
	}
	return result.QueueUrl
}

func ListSQS(sess *session.Session) {
	svc := sqs.New(sess)

	result, err := svc.ListQueues(nil)
	if err != nil {
		fmt.Println("Couldn't list queues")
		fmt.Println(err)
	}
	for i, url := range result.QueueUrls {
		fmt.Printf("%d: %v\n", i, *url)
	}
}
