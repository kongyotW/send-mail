package main

import (
	// ioutil "ct-mail/internal/ioutil"
	IOReader "send-mail/internal/ioread"
	Mailer "send-mail/internal/mail"

	// Passwer "ct-mail/internal/passw"
	"github.com/gin-gonic/gin"
)

func main() {
	// IOReader.EncryptByAdmin()
	IOReader.GetEmailServerAuth()

	router := gin.Default()
	router.POST("/mail/send", Mailer.SendMail)
	router.Run(":25003")
}
