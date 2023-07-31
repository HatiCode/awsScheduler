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

func ListSQS(sess *session.Session, name string) (queueList []string) {
	svc := sqs.New(sess)
	var urlList []string

	result, err := svc.ListQueues(&sqs.ListQueuesInput{
		QueueNamePrefix: aws.String(name),
	})
	if err != nil {
		fmt.Println("Couldn't list queues")
		fmt.Println(err)
	}
	for _, url := range result.QueueUrls {
		urlList = append(urlList, *url)
	}
	return urlList
}
