package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

var ConferenceName = "Go Conference"

const conferencetickets = 50

var remainingtickets uint = 50
var bookings = make([]userdata, 0)

type userdata struct {
	firstname string
	lastname  string
	email     string
	tickets   uint
}

func main() {

	greetusers()

	for {

		firstname, lastname, email, tickets := getuserinput()

		isvalidname, isvalidemail, isvalidtickets := helper.Validateuserinput(firstname, lastname, email, tickets, remainingtickets)

		if isvalidname && isvalidemail && isvalidtickets {

			bookticket(tickets, firstname, lastname, email)
			go sendticket(tickets, firstname, lastname, email)

			firstNames := getfirstnames()
			fmt.Printf("The first names of bookings are : %v\n", firstNames)

			if remainingtickets == 0 {

				fmt.Println("All tickets Sold")
				break

			}

		} else {

			if !isvalidname {

				fmt.Println("First name or Last name you entered is too short")

			}

			if !isvalidemail {

				fmt.Println("Your email doesn't contain @ sign")

			}

			if !isvalidtickets {

				fmt.Println("Number of tickets you entered is invalid")

			}

		}

	}

}

func Validateuserinput(firstname, lastname, email string, tickets, remainingtickets uint) {
	panic("unimplemented")
}

func greetusers() {

	fmt.Printf("Welcome to %v booking application\n", ConferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferencetickets, remainingtickets)
	fmt.Println("Get your tickets here to attend")

}

func getfirstnames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking)

		firstNames = append(firstNames, booking.firstname)
	}

	return firstNames

}

func getuserinput() (string, string, string, uint) {

	var firstname string
	var lastname string
	var email string
	var tickets uint

	fmt.Println("Enter first name : ")
	fmt.Scan(&firstname)

	fmt.Println("Enter last name : ")
	fmt.Scan(&lastname)

	fmt.Println("Enter email : ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets : ")
	fmt.Scan(&tickets)

	return firstname, lastname, email, tickets
}

func bookticket(tickets uint, firstname string, lastname string, email string) {

	remainingtickets = remainingtickets - tickets

	var userdata = userdata{

		firstname: firstname,
		lastname:  lastname,
		email:     email,
		tickets:   tickets,
	}

	bookings = append(bookings, userdata)
	fmt.Printf("List of Bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation at %v\n", firstname, lastname, tickets, email)
	fmt.Printf("Remaining tickets : %v\n", remainingtickets)

}

func sendticket(tickets uint, firstname string, lastname string, email string) {

	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", tickets, firstname, lastname)
	fmt.Println("######################")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("#######################")
}
