package notifymodels

import "fmt"

const (
	_ = iota
	EmailNotification
	SMSNotification
	VoiceCallNotification
)

type EmployeeDetails struct {
	Name               string
	EmailId            string
	Mobileno           string
	NotificationType   int
	NotificationSystem Notifier
}

type Notifier interface {
	Notify(pMessage string, pEmployee EmployeeDetails) (bool, error)
}

// Email Notifier
type EmailNotifier struct{}

func (en EmailNotifier) Notify(pMessage string, pEmployee EmployeeDetails) (bool, error) {
	var success bool = true
	fmt.Printf("Sending Email Notification to %s on %s\n", pEmployee.Name, pEmployee.EmailId)
	fmt.Printf("Message : %s\n", pMessage)
	return success, nil
}

// SMS Notifier
type SMSNotifier struct{}

func (en SMSNotifier) Notify(pMessage string, pEmployee EmployeeDetails) (bool, error) {
	var success bool = true
	fmt.Printf("Sending SMS Notification to %s on %s\n", pEmployee.Name, pEmployee.Mobileno)
	fmt.Printf("Message : %s\n", pMessage)
	return success, nil
}

// Voice Notification
type VoiceNotifier struct{}

func (en VoiceNotifier) Notify(pMessage string, pEmployee EmployeeDetails) (bool, error) {
	var success bool = true
	fmt.Printf("Sending Voice Notification to %s on %s\n", pEmployee.Name, pEmployee.Mobileno)
	fmt.Printf("Message : %s\n", pMessage)
	return success, nil
}
