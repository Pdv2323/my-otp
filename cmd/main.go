package main

import (
	"fmt"
	"log"
	"time"

	"github.com/parthvinchhi/my-otp/pkg/handlers"
)

func main() {
	var userEmail string
	fmt.Print("Enter your email: ")
	fmt.Scanln(&userEmail)

	otp := handlers.GenerateOtpHandler(6)

	err := handlers.SendEmailHandler(userEmail, otp)
	if err != nil {
		log.Fatalf("Error sending email to %s.", userEmail)
	}

	fmt.Println("OTP sent successfully! Check your email.")

	time.Sleep(1000)

	var NewOtp int
	fmt.Print("Enter the Otp you received : ")
	fmt.Scan(&NewOtp)

	v := handlers.VerifyOtpHandler(otp, NewOtp)
	fmt.Println(v)

}
