package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"sync"
	"time"

	"gopkg.in/gomail.v2"
)

const otpLength = 6
const otpExpiryDuration = 5 * time.Minute

var otpStore = struct {
	sync.RWMutex
	otps map[string]otpData
}{
	otps: make(map[string]otpData),
}

type otpData struct {
	OTP    string
	Expiry time.Time
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/submit-email", submitEmailHandler)
	mux.HandleFunc("/verify-otp", verifyOtpHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "submit_email.html")
}

func submitEmailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		otp := generateOTP(otpLength)
		err := sendOTP(email, otp)
		if err != nil {
			response := map[string]string{
				"status":  "error",
				"message": fmt.Sprintf("Failed to send OTP: %v", err),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Store the OTP with expiry
		otpStore.Lock()
		otpStore.otps[email] = otpData{
			OTP:    otp,
			Expiry: time.Now().Add(otpExpiryDuration),
		}
		otpStore.Unlock()

		// Redirect to verify OTP page with email as query parameter
		http.Redirect(w, r, "/verify-otp?email="+email, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func verifyOtpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		otp := r.FormValue("otp")

		otpStore.RLock()
		storedOtpData, exists := otpStore.otps[email]
		otpStore.RUnlock()

		if !exists || time.Now().After(storedOtpData.Expiry) {
			response := map[string]string{
				"status":  "error",
				"message": "OTP expired or does not exist",
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		if storedOtpData.OTP == otp {
			response := map[string]string{
				"status":  "success",
				"message": "OTP verified successfully",
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			response := map[string]string{
				"status":  "error",
				"message": "Invalid OTP",
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	} else if r.Method == "GET" {
		http.ServeFile(w, r, "verify_otp.html")
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func generateOTP(length int) string {
	const charset = "0123456789"
	otp := make([]byte, length)
	for i := range otp {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		otp[i] = charset[num.Int64()]
	}
	return string(otp)
}

func sendOTP(email, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "parthdv2111@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", fmt.Sprintf("Your OTP code is: %s", otp))

	d := gomail.NewDialer("smtp.gmail.com", 587, "parthdv2111@gmail.com", "dwvntkfkeqnejzco")

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
