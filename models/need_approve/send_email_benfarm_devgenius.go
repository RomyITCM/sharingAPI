package need_approve

import (
	"fmt"
	"net/smtp"
	"smile-service/entities"
)

var (
	authUserName      = "AKIA3WB2X4HJBLZRMNW3"
	authPassword      = "BGWmii43x45fNkxKwtcqNr1gXUdXaZT5XPpIdx6c2PzJ"
	smtpServerAddr    = "email-smtp.ap-southeast-1.amazonaws.com"
	smtpServerPort    = "587"
	destinationEmails = []string{"teresa.sophia@dspt.co.id", "stanley.lie@dspt.co.id"}
	senderEmail       = "no-reply@dspt.co.id"
)

func SendEmailDevGenius(data_email *entities.DataSendEmailBenfarm) {
	fmt.Println("sending emails example")

	msg := []byte("Subject: test email\r\n" +
		"\r\n" + data_email.HtmlBody)

	auth := smtp.PlainAuth("", authUserName, authPassword, smtpServerAddr)

	err := smtp.SendMail(smtpServerAddr+":"+smtpServerPort,
		auth, senderEmail, destinationEmails, msg)

	if err != nil {
		fmt.Printf("Error to sending email: %s", err)
		return
	}

	fmt.Println("email sent success")
}
