package main

import (
	"fmt"
	"time"

	"github.com/HatiCode/awsScheduler/updater/cmd"
	"github.com/HatiCode/awsScheduler/updater/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sqsName = "test-queue"

func main() {
	fmt.Println("testing S3")
	cmd.ListS3()

	// List SQS Queues
	fmt.Println("Listing SQS")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String("eu-central-1")},
		Profile: "scheduler-test",
	}))

	list := utils.ListSQS(sess, "test")
	// If list is empty, create a new SQS Queue
	if len(list) == 0 {
		fmt.Println("No SQS queues found")
		fmt.Println("Creating SQS Queue...")
		newSqs := utils.CreateSQS(sess, sqsName, "sqs.json")
		fmt.Printf("New queue %d created\n", newSqs)
		time.Sleep(30 * time.Second)
	}

	fmt.Println(list)

	// queue := list[0]
	// // Send test message to queue
	// fmt.Printf("Sending test message to %s\n", queue)
	// err := utils.SendMsg(sess, queue, "Hello, world!")
	// if err != nil {
	// 	fmt.Println("Error in sending test message")
	// }
	// fmt.Println("Test message sent")

	// // Print test message from queue
	// fmt.Printf("Reading Test message from %s\n", queue)
	// msg, err := utils.GetMsg(sess, queue, 1)
	// if err != nil {
	// 	fmt.Println("Error in reading Test message")
	// 	fmt.Println(err)
	// }
	// fmt.Println(msg.Messages[0].Body)

	// time.Sleep(20 * time.Second)

	// // Delete test message from queue
	// fmt.Printf("Deleting Test message from %s\n", queue)
	// receiptHandle := msg.Messages[0].ReceiptHandle
	// utils.DeleteMsg(sess, queue, receiptHandle)
	// fmt.Printf("Deleted Test message from %s\n", queue)
}
