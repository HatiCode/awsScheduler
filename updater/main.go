package main

import (
	"fmt"

	"github.com/HatiCode/awsScheduler/updater/cmd"
	"github.com/HatiCode/awsScheduler/updater/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	fmt.Println("testing S3")
	cmd.ListS3()

	fmt.Println("Listing SQS")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String("eu-central-1")},
		Profile: "scheduler-test",
	}))
	// utils.CreateSQS(sess, "test-queue")
	list := utils.ListSQS(sess, "test")
	for i, l := range list {
		fmt.Printf("%d: %s\n", i, l)
	}
}
