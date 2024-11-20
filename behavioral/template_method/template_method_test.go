package template_method

import (
	"fmt"
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	o := AbstractIOtp{&Sms{}}
	o.genAndSendOTP(4)

	fmt.Println("")
	o = AbstractIOtp{&Email{}}
	o.genAndSendOTP(5)

	// SMS: generating random otp 1234
	// SMS: saving otp: 1234 to cache
	// SMS: sending sms: SMS OTP for login is 1234

	// EMAIL: generating random otp 1234
	// EMAIL: saving otp: 1234 to cache
	// EMAIL: sending email: EMAIL OTP for login is 1234
}
