package ioread

var EmailID = ""
var EmailPassword = ""

func GetEmailServerAuth() {
	EmailID, EmailPassword = DecrptEmailId()
	// fmt.Println("EmailID : " + EmailID)
	// fmt.Println("EmailPassword : " + EmailPassword)
}
