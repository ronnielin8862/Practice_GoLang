package main

import (
	"github.com/aws/aws-sdk-go/aws/client"
)

type SQS struct {
	*client.Client
}

type GetQueueUrlInput struct {
	_                      struct{} `type:"structure"`
	QueueName              string   `type:"string" required:"true"`
	QueueOwnerAWSAccountId string   `type:"string"`
}

type GetQueueUrlOutput struct {
	QueueUrl *string `type:"string"`
}

func (input GetQueueUrlInput) main() {
	c := client.Client
	var qName string = "MR_Getaway_Test"
	accId := "082439862144"

	input.QueueName = qName
	input.QueueOwnerAWSAccountId = accId

	eq, out := c.GetQueueUrlRequest(input)

}

// func (c *SQS) GetQueueUrl(input GetQueueUrlInput) (GetQueueUrlOutput, error) {
// 	req, out := c.GetQueueUrlRequest(input)
// 	return out, req.Send()
// }
