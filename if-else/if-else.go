/**
	Summary:
	The existing program is used to restrict access to a resource based on day of the week and user role. Currently, the program allows any yser to access the resource. Implement the functionnality needed to grant resource access using any combination of ìf`, `else if` and `else`.

	Requirements:
	* Use the accessGranted() and accessDenied() functions to display informational messages.
	* Access at any time: Admin, Manager
	* Access weekends: Contractor
	* Access weekdays: Member
	* Access Mondays, Wednesdays, and Fridays: Guest
**/

package main

import "fmt"

const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func main() {
	// Day and Role. Change these to check your work.
	today, role := Sunday, Manager

	isWeekEnd := today == Saturday || today == Sunday
	isWeek := today == Monday || today == Tuesday || today == Wednesday || today == Thursday || today == Friday
	isNotGuestDay := today == Tuesday || today == Thursday || isWeekEnd

	if role == Guest && isNotGuestDay {
		accessDenied()
	} else if role == Contractor && isWeek {
		accessDenied()
	} else if role == Member && isWeekEnd {
		accessDenied()
	} else {
		accessGranted()
	}
}
