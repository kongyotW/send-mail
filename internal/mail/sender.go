package mail

import (
	"log"
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"

	ioread "send-mail/internal/ioread"
)

type mailInfoJSON struct {
	Key string `json:"key"`
	//SenderEmail   string `json:"sender_email"`
	ReceiverEmail string `json:"receiver_email"`
	SubjectEmail  string `json:"subject_email"`
	BodyMessage   string `json:"body_message"`
}

var mySecretKey string = "apibearer"

func SendMail(c *gin.Context) {
	var mailInfo mailInfoJSON
	c.BindJSON(&mailInfo)

	if mailInfo.Key != mySecretKey {
		log.Println("SecretKey No Valid")
		c.JSON(http.StatusOK, gin.H{"status": "error"})
		return
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := mailInfo.SubjectEmail

	email_body := "Subject: " + subject + "\n" + mime + mailInfo.BodyMessage
	// log.Println("EmailID : " + ioread.EmailID)
	// log.Println("EmailPassword : " + ioread.EmailPassword)
	auth := smtp.PlainAuth("", ioread.EmailID, ioread.EmailPassword, "smtp.gmail.com")

	status := smtp.SendMail("smtp.gmail.com:587", auth, ioread.EmailID, []string{mailInfo.ReceiverEmail}, []byte(email_body))

	if status != nil {
		log.Print("Email Sent Wrong")
		log.Printf("Error from SMTP Server: %s", status)
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail"})
	} else {
		log.Print("Email Sent Successfully")
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
