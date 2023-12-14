package need_approve

// import (
// 	"fmt"
// 	"net/smtp"
// 	"os"
// 	"smile-service/entities"
// )

// func SendEmailTest(msg *entities.DataSendEmail) {
// 	// Sender data.
// 	from := os.Getenv("EMAIL")
// 	password := os.Getenv("EMAIL_PASSWORD")

// 	// Receiver email address.
// 	// Change this to your email...
// 	to := []string{
// 		"teresasophia.322@gmail.com",
// 	}

// 	// smtp server configuration.
// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "587"

// 	// Message.
// 	message := []byte(
// 		msg.Message)

// 	// Authentication.
// 	auth := smtp.PlainAuth("", from, password, smtpHost)

// 	// Sending email.
// 	emailErr := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
// 	if emailErr != nil {
// 		fmt.Println(emailErr)
// 		return
// 	}
// 	fmt.Printf("Email successfully sent to %s", to[0])

// }
import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func VerifyEmail(sess *session.Session, email string) error {
	sesClient := ses.New(sess)
	_, err := sesClient.VerifyEmailIdentity(&ses.VerifyEmailIdentityInput{
		EmailAddress: aws.String(email),
	})

	return err
}

func goVerifyMail() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-west-2"),
		},
	})

	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}

	email := "teresasophia.322@gmail.com"
	err = VerifyEmail(sess, email)
	if err != nil {
		fmt.Printf("Got an error while trying to verify email: %v", err)
		return
	}
}

const CHARSET = "UTF-8"

func SendHTMLEmail(sess *session.Session, toAddresses []*string, htmlText string, sender string, subject string) error {
	sesClient := ses.New(sess)

	_, err := sesClient.SendEmail(&ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: toAddresses,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CHARSET),
					Data:    aws.String(htmlText),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CHARSET),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	})

	return err
}

func SendEmailTest() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-west-2"),
		},
	})

	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}

	toAddresses := []*string{
		aws.String("teresasophia.322@gmail.com"),
	}
	// emailText := "Amazing SES Tutorial"
	htmlText := "<html><head></head><h1 style='text-align:center'>This is the heading</h1><p>Hello, world</p></body></html>"
	sender := "teresasophia.322@gmail.com"
	subject := "Amazing Email Tutorial!'"

	// err = SendEmail(sess, toAddresses, emailText, sender, subject)
	// if err != nil {
	// 	fmt.Printf("Got an error while trying to send email: %v", err)
	// 	return
	// }

	fmt.Println("Sent text email successfully")

	err = SendHTMLEmail(sess, toAddresses, htmlText, sender, subject)
	if err != nil {
		fmt.Printf("Got an error while trying to send email: %v", err)
		return
	}

	fmt.Println("Sent html email successfully")

}
