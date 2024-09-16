package eventbus

import (
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

func (s *AWSEventBus) TransmissionMessage(payload string, topic string) error {
	params := &sns.PublishInput{
		Message:  aws.String(payload),
		TopicArn: aws.String(fmt.Sprintf("arn:aws:sns:%s:837217772820:%s", os.Getenv("AWS_REGION"), topic)),
	}

	result, err := s.sns.Publish(params)
	if err != nil {
		fmt.Println("Error publishing message: ", err.Error())
		return err
	}

	fmt.Println("Message ", *result.MessageId, " published.")

	return nil
}

func (s *AWSEventBus) Subscribe(topic string, callback func(messageBody string)) error {
	queueURL := fmt.Sprintf("https://sqs.%s.amazonaws.com/837217772820/%s-sqs", os.Getenv("AWS_REGION"), topic)

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

				// Procesa el mensaje
				callback(payload.Message)

				// Elimina el mensaje de la cola después de procesarlo
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
	return nil
}
