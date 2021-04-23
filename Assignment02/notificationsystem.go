package main

import (
	"fmt"
	"notificationsystem/notifymodels"
)

var emailNotifier notifymodels.EmailNotifier = notifymodels.EmailNotifier{}
var smsNotifier notifymodels.SMSNotifier = notifymodels.SMSNotifier{}
var voiceNotifier notifymodels.VoiceNotifier = notifymodels.VoiceNotifier{}

func main() {
	var employees [5]notifymodels.EmployeeDetails
	var message string

	fmt.Println("Notificaition System")

	// Employee 01 - Email Notification
	employees[0] = NewEmployee("NaveenGM", "naveen.gm.blr@gmail.com", "9945417417", notifymodels.EmailNotification)

	// Employee 02 - SMS Notification
	employees[1] = NewEmployee("Rekha NG", "rekha.ng@gmail.com", "9886570099", notifymodels.SMSNotification)

	// Employee 03 - Voice Notification
	employees[2] = NewEmployee("Dhruv NG", "dhruv.ng@gmail.com", "1234567890", notifymodels.VoiceCallNotification)

	// Employee 04 - Voice Notification
	employees[3] = NewEmployee("Nanjundarao", "nanjundarao.j@gmail.com", "2222233333", notifymodels.SMSNotification)

	// Employee 05 - Voice Notification
	employees[4] = NewEmployee("Sita NJ", "sita.nj@gmail.com", "4444455555", notifymodels.EmailNotification)

	message = "Do not forget to use mask when you go out."
	fmt.Println("Sending Notificayions to all employees")
	for i := 0; i < 5; i++ {
		fmt.Println("\nNotifying Employee :", i)
		employees[i].NotificationSystem.Notify(message, employees[i])
	}

}

func NewEmployee(pName string, pEmailId string, pMobileNo string, pNotificationType int) notifymodels.EmployeeDetails {
	var emp notifymodels.EmployeeDetails

	emp.Name = pName
	emp.EmailId = pEmailId
	emp.Mobileno = pMobileNo

	switch pNotificationType {
	case notifymodels.EmailNotification:
		emp.NotificationSystem = emailNotifier
	case notifymodels.SMSNotification:
		emp.NotificationSystem = smsNotifier
	case notifymodels.VoiceCallNotification:
		emp.NotificationSystem = voiceNotifier
	}

	return emp
}
