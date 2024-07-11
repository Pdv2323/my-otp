package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Pdv2323/my-otp/email"
	"github.com/Pdv2323/my-otp/generate"
	"github.com/Pdv2323/my-otp/verify"
)

func main() {
	var userEmail string
	fmt.Print("Enter your email: ")
	fmt.Scanln(&userEmail)

	otp := generate.GenerateOtp()

	err := email.SendEmail(userEmail, otp)
	if err != nil {
		log.Fatalf("Error sending email to %s.", userEmail)
	}

	fmt.Println("OTP sent successfully! Check your email.")

	time.Sleep(1000)

	var NewOtp int
	fmt.Print("Enter the Otp you received : ")
	fmt.Scan(&NewOtp)

	v := verify.VerifyOtp(otp, NewOtp)
	fmt.Println(v)

}
