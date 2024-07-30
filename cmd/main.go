package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parthvinchhi/my-otp/pkg/handlers"
)

// func main() {
// 	var userEmail string
// 	fmt.Print("Enter your email: ")
// 	fmt.Scanln(&userEmail)

// 	otp := handlers.GenerateOtpHandler(6)

// 	err := handlers.SendEmailHandler(userEmail, otp)
// 	if err != nil {
// 		log.Fatalf("Error sending email to %s.", userEmail)
// 	}

// 	fmt.Println("OTP sent successfully! Check your email.")

// 	time.Sleep(1000)

// 	var NewOtp int
// 	fmt.Print("Enter the Otp you received : ")
// 	fmt.Scan(&NewOtp)

// 	v := handlers.VerifyOtpHandler(otp, NewOtp)
// 	fmt.Println(v)

// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", handlers.RenderIndexPageHandler)
// 	mux.HandleFunc("/submit-email", handlers.SubmitEmailIdHandler)
// 	mux.HandleFunc("/submit-otp", handlers.SubmitOTPHandler)

// 	fmt.Println("Starting server on :8888")
// 	http.ListenAndServe(":8888", mux)
// }

func main() {
	r := gin.Default()

	indexFile := "pkg/templates/index.html"
	verifyFile := "pkg/templates/verify.html"

	r.LoadHTMLFiles(indexFile, verifyFile)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, indexFile, nil)
	})
	r.POST("/submit-email", handlers.SubmitEmailHandler)

	r.GET("/submit-otp", func(c *gin.Context) {
		c.HTML(http.StatusOK, verifyFile, nil)
	})
	r.POST("/submit-otp", handlers.SubmitOtpHandler)

	r.Run(":8888")
}
