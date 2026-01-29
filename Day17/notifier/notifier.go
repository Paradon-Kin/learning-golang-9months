package main

import "fmt"

type Notifier interface {
	Notify(message string)
}

type Email struct {
	Address string
}

type SMS struct {
	PhoneNumber string
}

func (e *Email) Notify(message string) {
	if containsAt(e.Address) {
		fmt.Printf("📧 Sending Email to [%s]: %s\n", e.Address, message)
	} else {
		fmt.Printf("❌ Error: Invalid Email Format (%s)\n", e.Address)
	}
}

func (s *SMS) Notify(message string) {
	if len(s.PhoneNumber) == 10 {
		fmt.Printf("Sending SMS to [%s]:[%s]\n", s.PhoneNumber, message)
	} else {
		fmt.Printf("❌ Error: Invalid Phone Number (Length: %d)\n", len(s.PhoneNumber))
	}
}

func SendAlert(n Notifier, message string) {
	fmt.Println("Sending...")
	n.Notify(message)
}

func containsAt(s string) bool {
	for _, char := range s {
		if char == '@' {
			return true
		}
	}
	return false
}

func main() {
	email := &Email{Address: "paradon.s@psru.ac.th"}
	sms := &SMS{PhoneNumber: "0958970510"}

	SendAlert(email, "We invite you to our company on march")
	SendAlert(sms, "You have a New job")
}
