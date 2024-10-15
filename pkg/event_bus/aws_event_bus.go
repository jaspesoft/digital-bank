package eventbus

import (
	systemdomain "digital-bank/internal/system/domain"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"os"
)

type AWSEventBus struct {
	sns *sns.SNS
	sqs *sqs.SQS
}

func NewAWSEventBus() *AWSEventBus {
	sess := session.Must(session.NewSession())
	return &AWSEventBus{
		sns: sns.New(sess),
		sqs: sqs.New(sess),
	}
}

func (s *AWSEventBus) Emit(data interface{}, topic systemdomain.Topic) error {
	str, _ := data.(string)
	params := &sns.PublishInput{
		Message:  aws.String(str),
		TopicArn: aws.String(fmt.Sprintf("arn:aws:sns:%s:%s:%s", os.Getenv("AWS_REGION"), os.Getenv("AWS_ACCOUNT"), topic)),
	}

	result, err := s.sns.Publish(params)
	if err != nil {
		fmt.Println("Error publishing message: ", err.Error())
		return err
	}

	fmt.Println("Message ", *result.MessageId, " published.")

	return nil
}

func (s *AWSEventBus) Subscribe(topic systemdomain.Topic, callback func(systemdomain.Message)) {
	queueURL := fmt.Sprintf("https://sqs.%s.amazonaws.com/%s/%s-sqs", os.Getenv("AWS_REGION"), os.Getenv("AWS_ACCOUNT"), topic)

	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		WaitTimeSeconds:     aws.Int64(10),
		MaxNumberOfMessages: aws.Int64(1),
	}

	receiveMessage := func() {
		for {
			data, err := s.sqs.ReceiveMessage(params)
			if err != nil {
				fmt.Printf("Received error while receiving messages: %v\n", err)
				continue
			}

			if len(data.Messages) > 0 {
				message := data.Messages[0]

				var payload struct {
					Message string `json:"Message"`
				}

				if err := json.Unmarshal([]byte(*message.Body), &payload); err != nil {
					fmt.Printf("Error unmarshalling message: %v\n", err)
					continue
				}

				// Process the message
				callback(systemdomain.Message{
					Data:  payload.Message,
					Topic: topic,
				})

				// Delete the message from the queue after processing it
				_, err = s.sqs.DeleteMessage(&sqs.DeleteMessageInput{
					QueueUrl:      aws.String(queueURL),
					ReceiptHandle: message.ReceiptHandle,
				})
				if err != nil {
					fmt.Printf("Error deleting message: %v\n", err)
				}
			}
		}
	}

	go receiveMessage()
}
