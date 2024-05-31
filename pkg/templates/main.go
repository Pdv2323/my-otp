package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"sync"
	"time"
)

var (
	otpStore = make(map[string]string)
	mu       sync.Mutex
)

type EmailRequest struct {
	Email string `json:"email"`
}

type OTPRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

func sendEmail(to, otp string) error {
	from := "your-email@example.com"
	password := "your-email-password"
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	msg := []byte("To: " + to + "\r\n" +
		"Subject: Your OTP Code\r\n" +
		"\r\n" +
		"Your OTP is: " + otp + "\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func sendOTPHandler(w http.ResponseWriter, r *http.Request) {
	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	otp := generateOTP()
	if err := sendEmail(req.Email, otp); err != nil {
		http.Error(w, "Failed to send OTP", http.StatusInternalServerError)
		return
	}

	mu.Lock()
	otpStore[req.Email] = otp
	mu.Unlock()

	w.WriteHeader(http.StatusOK)
}

func verifyOTPHandler(w http.ResponseWriter, r *http.Request) {
	var req OTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	if storedOTP, ok := otpStore[req.Email]; ok && storedOTP == req.OTP {
		delete(otpStore, req.Email)
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, "Invalid OTP", http.StatusUnauthorized)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/send-otp", sendOTPHandler)
	http.HandleFunc("/verify-otp", verifyOTPHandler)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
