package utils

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func CreateSQS(sess *session.Session, name string) (queueUrl string) {
	svc := sqs.New(sess)

	result, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(name),
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("10"),
			"MessageRetentionPeriod": aws.String("86400"),
		},
	})
	if err != nil {
		fmt.Println("Counld not create queue %v", name)
	}
	return *result.QueueUrl
}

// func CheckQueueExists(sess *session.Session, name string) bool {
// 	svc := sqs.New(sess)

// 	result, err := svc.ListQueues(nil)
// 	if err != nil {
// 		fmt.Println("Couldn't fetch queues")
// 		fmt.Println(err)
// 	}
// }
