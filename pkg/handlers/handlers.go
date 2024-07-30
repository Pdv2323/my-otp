package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"reflect"
	"strconv"
	"time"

	"golang.org/x/exp/rand"
)

// type Otp struct {
// 	Otp string
// }

var Otp int

func GenerateOtpHandler(n int) int {
	if n < 1 {
		return 0
		// return string(0)
	}

	min := 1
	for i := 1; i < n; i++ {
		min *= 10
	}
	max := min*10 - 1

	// rand.Seed(time.Now().UnixNano())
	t := time.Now().UnixNano()
	rand.Seed(uint64(t))
	r := rand.Intn(max-min+1) + min
	t1 := reflect.TypeOf(r)
	log.Println(r, t1)
	return r
	// return string(r)

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

func VerifyOtpHandler(otp, newOtp string) bool {
	// fmt.Scanln("Enter the Otp you received : ", &otp)
	if otp != newOtp {
		return false
	}
	return true
}

func RenderIndexPageHandler(w http.ResponseWriter, r *http.Request) {
	// t, err := template.ParseFiles("pkg/templates/index.html")
	// if err != nil {
	// 	fmt.Fprintf(w, "Error: %s", err)
	// 	return
	// }
	// t.Execute(w, nil)
	http.ServeFile(w, r, "pkg/templates/index.html")
}

func RenderVerifyPageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("pkg/templates/verify.html")
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
	t.Execute(w, nil)
}

func SubmitEmailIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Email := r.FormValue("email")
		Otp = GenerateOtpHandler(6)
		t1 := reflect.TypeOf(Otp)
		t2 := reflect.TypeOf(Email)

		log.Println(Otp, t1)
		err := SendEmailHandler(Email, Otp)
		if err != nil {
			Response := map[string]string{
				"status":  "error",
				"message": fmt.Sprintf("Failed to send OTP: %v", err),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response)
			return
		}
		http.Redirect(w, r, "/submit-otp?email="+Email, http.StatusSeeOther)
		log.Println(Otp, t1)
		log.Println(Email, t2)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func SubmitOtpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		OTP := r.FormValue("otp")
		// ok := VerifyOtpHandler(Otp, OTP)
		// log.Println(ok)
		// if ok {
		// 	Response := map[string]string{
		// 		"status":  "success",
		// 		"message": "OTP verified successfully.",
		// 	}
		// 	w.WriteHeader(http.StatusOK)
		// 	w.Header().Set("Content-Type", "application/json")
		// 	json.NewEncoder(w).Encode(Response)

		// } else {
		// 	Response := map[string]string{
		// 		"status":  "error",
		// 		"message": "Invalid OTP",
		// 	}
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	w.Header().Set("Content-Type", "application/json")
		// 	json.NewEncoder(w).Encode(Response)
		// }

		IntOtp, err := strconv.Atoi(OTP)
		if err != nil {
			log.Fatal("Error converting string to int")
		}

		if IntOtp != Otp {
			t1 := reflect.TypeOf(OTP)
			t2 := reflect.TypeOf(Otp)

			log.Println("ottttppps:::::", OTP, t1, Otp, t2)
			Response := map[string]string{
				"status":  "error",
				"message": "Invalid OTP",
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Response)
		} else {
			Response := map[string]string{
				"status":  "success",
				"message": "OTP verified successfully.",
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Response)

		}
	} else if r.Method == "GET" {
		http.ServeFile(w, r, "pkg/templates/verify.html")
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
