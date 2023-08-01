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

func SendMsg(sess *session.Session, queueUrl string, msgBody string) error {
	svc := sqs.New(sess)

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    &queueUrl,
		MessageBody: aws.String(msgBody),
	})

	return err
}

func GetMsg(sess *session.Session, queueUrl string, maxMessages int) (*sqs.ReceiveMessageOutput, error) {
	svc := sqs.New(sess)

	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            &queueUrl,
		MaxNumberOfMessages: aws.Int64(int64(maxMessages)),
	})

	if err != nil {
		return nil, err
	}

	return msgResult, nil
}

func DeleteMsg(sess *session.Session, queueUrl string, messageHandle *string) error {
	svc := sqs.New(sess)

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &queueUrl,
		ReceiptHandle: messageHandle,
	})

	return err
}
