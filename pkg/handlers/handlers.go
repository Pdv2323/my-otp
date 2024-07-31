package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Otp int

func generateOtpHandler(n int) int {
	if n < 1 {
		return 0
		// return string(0)
	}

	min := 1
	for i := 1; i < n; i++ {
		min *= 10
	}
	max := min*10 - 1
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

func SubmitEmailHandler(c *gin.Context) {
	email := c.PostForm("email")
	Otp = generateOtpHandler(6)
	err := SendEmailHandler(email, Otp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Failed to send OTP: %v", err),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/submit-otp?email="+email)
}

func SubmitOtpHandler(c *gin.Context) {
	otp := c.PostForm("otp")

	IntOtp, err := strconv.Atoi(otp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to convert String otp to Int otp.",
		})
	}

	if IntOtp == Otp {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "OTP verified Successfully.",
		})
	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status":  "error",
			"message": "Wrong OTP. Please enter correct OTP.",
		})

		return
	}
}
