package verify

func VerifyOtp(otp, newOtp int) string {
	// fmt.Scanln("Enter the Otp you received : ", &otp)
	if otp != newOtp {
		return "OTP Incorrect!!"
	}
	return "OTP verified Sucessfully!!"
}
