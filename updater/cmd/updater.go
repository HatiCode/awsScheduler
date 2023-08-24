package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Create 2 databases : rds, ec2
// List existing accounts in organization
// for each database, create a collection using account number
// Scan resource type in each account and add as document to corresponding collection & database
// Format document with name - uptime - downtime - updated time

func ListS3() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default config, is your AWS default profile set ?")
		fmt.Println(err)
		return
	}
	s3Client := s3.NewFromConfig(cfg)
	count := 10
	fmt.Printf("Let's list up to %v buckets :\n", count)
	result, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		fmt.Printf("Couldn't list buckets for your account. Here's why: %v\n", err)
		return
	}
	if len(result.Buckets) == 0 {
		fmt.Println("You don't have any buckets!")
	} else {
		if count > len(result.Buckets) {
			count = len(result.Buckets)
		}
		for _, bucket := range result.Buckets[:count] {
			bucketTime := *bucket.CreationDate
			creationTime := bucketTime.Add(time.Hour * 2)
			fmt.Printf("\t%v %v\n", *bucket.Name, creationTime.Format(time.RFC850))
		}
	}
}
