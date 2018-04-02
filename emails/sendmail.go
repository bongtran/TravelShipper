package emails

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

func SendVerifyEmail(emailAddress string, code string) {

	d := gomail.NewDialer("smtp.gmail.com", 587, "nicetravelshipper@gmail.com", "Nam50tam")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}


	body := "Dear <b>User "+ emailAddress+"</b>, </br> Your verification is: <b>" + code +"</b></br>";
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
		//panic(err)
		println(err.Error())
	}
}
