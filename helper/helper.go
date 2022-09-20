package helper

import "strings"

func Validateuserinput(firstname string, lastname string, email string, tickets uint, remainingtickets uint) (bool, bool, bool) {

	isvalidname := len(firstname) >= 2 && len(lastname) >= 2
	isvalidemail := strings.Contains(email, "@")
	isvalidtickets := tickets > 0 && tickets <= remainingtickets
	return isvalidname, isvalidemail, isvalidtickets
}
