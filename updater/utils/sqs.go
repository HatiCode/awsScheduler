package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func CreateSQS(sess *session.Session, name string, policyName string) (queueUrl *string) {
	svc := sqs.New(sess)

	// TODO Create secrets for policies
	// Policy is passed as json file
	policiesFolder, _ := os.LookupEnv("SCHEDULER_POLICY_PATH")
	policiesPath := filepath.Join(policiesFolder, "/", policyName)
	jsonFile, err := os.Open(policiesPath)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	result, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(name),
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("0"),
			"MessageRetentionPeriod": aws.String("86400"),
			"VisibilityTimeout":      aws.String("3600"),
			"SqsManagedSseEnabled":   aws.String("false"),
			"Policy":                 aws.String(string(byteValue)),
		},
		Tags: map[string]*string{
			"env": aws.String("dev"),
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

func PurgeQueue(sess *session.Session, queueUrl string) error {
	svc := sqs.New(sess)

	_, err := svc.PurgeQueue(&sqs.PurgeQueueInput{
		QueueUrl: aws.String(queueUrl),
	})

	return err
}
