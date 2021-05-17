package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {

	queue := flag.String("q", "", "The name of the queue")
	flag.Parse()

	// if *queue == "" {
	// 	fmt.Println("You must supply a queue name (-q QUEUE")
	// 	return
	// }

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	urlResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})

	fmt.Println("URL: " + *urlResult.QueueUrl)

	if err != nil {
		fmt.Println("err = ", err)
	}
}
