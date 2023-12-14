package need_approve

import (
	"fmt"
	"smile-service/entities"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	Sender1 = "no-reply@dspt.co.id"
	// Recipient1 = "teresa.sophia@dspt.co.id"
	TextBody1 = "This email was sent with Amazon SES using the AWS SDK for Go."
	CharSet   = "UTF-8"
)

func SendEmailHashNode(data_email *entities.DataSendEmailBenfarm) {
	// sess, err := session.NewSessionWithOptions(session.Options{
	//     Profile: "ProfileName",
	// })

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(
				*aws.String("AKIA3WB2X4HJGJTDWPGL"),
				*aws.String("VVmHtwoOR6xUosj9fi0NW4hRY9i2KVIvMFb7KX+j"),
				"", // a token will be created when the session it's used.
			),
		})

	if err != nil {
		panic(err)
	}

	svc := ses.New(sess)

	CCEmail1 := "stanley.lie@dspt.co.id" //[2]string{, }
	CCEmail2 := "ahmad.ari@dspt.co.id"

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{&CCEmail1, &CCEmail2},
			// CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(data_email.Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(data_email.HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(data_email.HtmlBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject + " " + data_email.CustomerName),
			},
		},
		Source: aws.String(Sender1),
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
			fmt.Println(err.Error())
		}

		return

	}

	fmt.Println("Email Sent to address: " + data_email.Recipient)
	fmt.Println(result)
}
