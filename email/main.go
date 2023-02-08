package main

import (
	"fmt"

	"github.com/go-mail/mail"
)

func main() {
	OwnerEmail := "example@gmail.com"
	// OwnerEmail := os.Getenv("OwnerEmail")
	// WebsiteLink := os.Getenv("WebsiteLink")
	// Phone := os.Getenv("Phone")
	m := mail.NewMessage()
	customeremail := "example2@gmail.com"
	name := "chantos"
	amount := 5600
	Phone := "0729308456"
	WebsiteLink := "https://nillavee.co.ke"

	m.SetHeader("From", "example@gmail.com")

	// m.SetHeader("To", "example2@gmail.com", "example@gmail.com")
	m.SetHeader("To", customeremail)

	m.SetAddressHeader("Cc", OwnerEmail, "Nillavees Cakes and Patries")

	m.SetHeader("Subject", "Your Order at Nillavee was Successful!")

	m.SetBody("text/html", "<h3>Hello  "+name+"<br /> Thank you for shopping with us!</h3><br />Your Order of Ksh"+fmt.Sprintf("%.2f", float64(amount))+" is being processed <br />We'll respond shortly.<br><br>Thanks!<br>"+OwnerEmail+"<br />Phone:"+Phone+"<br />Website: <a href='"+WebsiteLink+"'>"+WebsiteLink+"</a>")

	m.Attach("logo.png")

	d := mail.NewDialer("smtp.gmail.com", 587, "example@gmail.com", "secret")

	// Send the email to Kate, Noah and Oliver.

	if err := d.DialAndSend(m); err != nil {

		panic(err)

	}
	// m := mail.NewMessage()

	// m.SetHeader("From", "example@gmail.com")

	// m.SetHeader("To", "example2@gmail.com", "example@gmail.com")

	// m.SetAddressHeader("Cc", "example2s@gmail.com", "chantos")

	// m.SetHeader("Subject", "Hello!")

	// m.SetBody("text/html", "Hello <b>Kate</b> and <i>Noah</i>!")

	// m.Attach("logo.png")

	// d := mail.NewDialer("smtp.gmail.com", 587, "example@gmail.com", "clndtmuohtfnynmi")

	// // Send the email to Kate, Noah and Oliver.

	// if err := d.DialAndSend(m); err != nil {

	// 	panic(err)

	// }

}

// package main

// import (
// 	"fmt"
// 	"net/smtp"
// )

// func main() {
// 	//set the email
// 	email := "example2@gmail.com"
// 	subject := "AKGL ALI is the best :D"
// 	body := "This is the body part you send whatever you want in the body."
// 	err := sendEmail(email, subject, body)
// 	if err != nil {
// 		//if you get password error please make sure your password is correct.
// 		//Or you have followed all the steps
// 		fmt.Println(err)
// 		return
// 	}
// }
// func sendEmail(email, subject, body string) error {
// 	// err := godotenv.Load() //by default, it is .env, so we don't have to write
// 	// if err != nil {
// 	// 	fmt.Println("Error is occurred  on .env file please check")
// 	// 	return err
// 	// }
// 	//    from := os.Getenv("MAIL")
// 	//    password := os.Getenv("PASSWD")
// 	from := "example@gmail.com"
// 	password := "clndtmuohtfnynmi"
// 	host := "smtp.gmail.com"
// 	port := 587

// 	// Connect to the SMTP server.
// 	auth := smtp.PlainAuth("", from, password, host)
// 	to := []string{email}
// 	msg := []byte("To: " + email + "\r\n" +
// 		"Subject: " + subject + "\r\n" +
// 		"\r\n" +
// 		body + "\r\n")
// 	err := smtp.SendMail(fmt.Sprintf("%s:%d", host, port), auth, from, to, msg)
// 	fmt.Println("email is successfully sent to " + email)
// 	return err
// }
