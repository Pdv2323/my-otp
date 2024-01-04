package main

import (
	"fmt"
	"log"

	"github.com/Pdv2323/otp/email"
	"github.com/Pdv2323/otp/generate"
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

}
