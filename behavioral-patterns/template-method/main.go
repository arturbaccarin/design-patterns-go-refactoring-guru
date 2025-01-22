package main

import "fmt"

/*
https://refactoring.guru/design-patterns/template-method

Template Method is a behavioral design pattern that defines
the skeleton of an algorithm in the superclass but lets subclasses
override specific steps of the algorithm without changing its structure.

Imagine that you’re creating a data mining application that analyzes
corporate documents. Users feed the app documents in various formats
(PDF, DOC, CSV), and it tries to extract meaningful data from these
docs in a uniform format.

The first version of the app could work only with DOC files.
In the following version, it was able to support CSV files.
A month later, you “taught” it to extract data from PDF files.

While the code for dealing with various data formats was entirely
different in all classes, the code for data processing and analysis
is almost identical. Wouldn’t it be great to get rid of the code duplication,
leaving the algorithm structure intact?

The Template Method pattern suggests that you break down an algorithm into
a series of steps, turn these steps into methods, and put a series of calls
to these methods inside a single template method. The steps may either be abstract,
or have some default implementation.

As you can see, we’ve got two types of steps:

- abstract steps must be implemented by every subclass
- optional steps already have some default implementation,
but still can be overridden if needed

There’s another type of step, called hooks.
A hook is an optional step with an empty body.
A template method would work even if a hook isn’t overridden.
Usually, hooks are placed before and after crucial steps of algorithms,
providing subclasses with additional extension points for an algorithm.

Template Method is a behavioral design pattern that allows you to
define a skeleton of an algorithm in a base class and let subclasses
override the steps without changing the overall algorithm’s structure.
*/

/*
Example

Let’s consider the example of One Time Password (OTP) functionality.
There are different ways that the OTP can be delivered to a user (SMS, email, etc.).
But irrespective whether it’s an SMS or email OTP, the entire OTP process is the same:

1. Generate a random n digit number.
2. Save this number in the cache for later verification.
3. Prepare the content.
4. Send the notification.

So, we have a scenario where the steps of a particular operation are the same,
but these steps’ implementation may differ. This is an appropriate situation
to consider using the Template Method pattern.
*/

// Template Method
type IOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Otp struct {
	iOtp IOtp
}

func (o *Otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}

// Concrete implementation
type Sms struct {
	Otp
}

func (s *Sms) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *Sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

type Email struct {
	Otp
}

func (s *Email) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *Email) getMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (s *Email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}

func main() {
	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)

}
