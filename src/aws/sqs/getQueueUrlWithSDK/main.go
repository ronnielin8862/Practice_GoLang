package main
//
//import (
//	"github.com/aws/aws-sdk-go/aws"
//	"github.com/aws/aws-sdk-go/aws/session"
//	"github.com/aws/aws-sdk-go/service/sqs"
//)
//
//type GetQueueUrlOutput struct {
//	QueueUrl *string
//}
//
//func main() {
//
//	queue := aws.String("tetetete")
//
//	sess := session.Must(session.NewSessionWithOptions(session.Options{
//		SharedConfigState: session.SharedConfigEnable,
//	}))
//
//	svc := sqs.New(sess)
//
//	getQueueUrlInput := sqs.GetQueueUrlInput{QueueName: queue}
//
//	urlResult, err := svc.GetQueueUrl(getQueueUrlInput *GetQueueUrlInput)(*GetQueueUrlOutput, error)
//
//	var qName string = "MR_Getaway_Test"
//	accId := "082439862144"
//
//	input.QueueName = qName
//	input.QueueOwnerAWSAccountId = accId
//
//	eq, out := c.GetQueueUrlInput(input)
//
//}