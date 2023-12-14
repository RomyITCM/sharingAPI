package need_approve

import (
	"fmt"
	"smile-service/entities"

	//go get -u github.com/aws/aws-sdk-go
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	Sender    = "no-reply@dspt.co.id"
	Recipient = "teresasophia.322@gmail.com"
	Subject   = "Request Freezer Benfarm Customer"
	Charset   = "UTF-8"
	TextBody  = "This email was sent with Amazon SES using the AWS SDK for Go."
)

func SendEmailSES(data_email *entities.DataSendEmailBenfarm) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{aws.String(Recipient)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(Charset),
					Data:    aws.String(data_email.HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(Charset),
					Data:    aws.String(TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(Charset),
				Data:    aws.String(Subject + " " + data_email.CustomerName),
			},
		},
		Source: aws.String(Sender),
	}

	result, err := svc.SendEmail(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

		return
	}

	fmt.Println("Email Sent to address: " + Recipient)
	fmt.Println(result)
}
