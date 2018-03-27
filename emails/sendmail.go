package emails

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

func SendVerifyEmail(emailAddress string, code string) {
	d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}


	body := "Dear <b>User</b>, </br> Your verification is: <b>" + code +"User</b>";
	// Send emails using d.
	m := gomail.NewMessage()
	m.SetHeader("From", "info@travelshipper.com")
	m.SetHeader("To", emailAddress)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Travel Shipper Verification")
	m.SetBody("text/html", body)
	//m.Attach("/home/Alex/lolcat.jpg")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}