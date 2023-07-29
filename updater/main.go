package main

import (
	"fmt"

	"github.com/HatiCode/awsScheduler/updater/cmd"
	"github.com/HatiCode/awsScheduler/updater/utils"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	fmt.Println("testing S3")
	cmd.ListS3()

	fmt.Println("Creating SQS")
	sess := session.Must(session.NewSession())
	utils.CreateSQS(sess, "test-queue")
}
