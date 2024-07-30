package handlers

import (
	"fmt"
	"net/smtp"
	"time"

	"golang.org/x/exp/rand"
)

func GenerateOtpHandler(n int) int {
	if n < 1 {
		return 0
	}

	min := 1
	for i := 1; i < n; i++ {
		min *= 10
	}
	max := min*10 - 1

	// rand.Seed(time.Now().UnixNano())
	t := time.Now().UnixNano()
	rand.Seed(uint64(t))
	return rand.Intn(max-min+1) + min
}

func SendEmailHandler(to string, otp int) error {
	from := "parthdv2111@gmail.com"
	password := "dwvntkfkeqnejzco"
	SmtpServer := "smtp.gmail.com"
	SmtpPort := "587"

	msg := fmt.Sprintf("Subject : OTP Verification\n\n Your OTP is : %d", otp)

	auth := smtp.PlainAuth("", from, password, SmtpServer)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", SmtpServer, SmtpPort),
		auth,
		from,
		[]string{to},
		[]byte(msg),
	)

	return err
}

func VerifyOtpHandler(otp, newOtp int) string {
	// fmt.Scanln("Enter the Otp you received : ", &otp)
	if otp != newOtp {
		return "OTP Incorrect!!"
	}
	return "OTP verified Sucessfully!!"
}
