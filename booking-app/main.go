package main

import (
	"fmt"
	"time"
	//"strconv"
)

const conferencetickets = 50

var conferenceName = "Go conference"
var Ticketsremaining uint = 50
var bookings = make([]userdata, 0)

type userdata struct {
	userName     string
	userlastName string
	useremail    string
	usertickets  uint
}

func main() {
	for {
		greetuser()
		userName, userlastName, useremail, usertickets := getuserinput()
		validname, validemail, validticketnumber := validateUserInput(userName, userlastName, useremail, usertickets, Ticketsremaining)
		if validname && validemail && validticketnumber {
			bookticket(usertickets, Ticketsremaining, userName, userlastName, useremail)
			go sendticket(usertickets, userName, userlastName, useremail)
			fmt.Printf("The lenght of Slice:%v\n", len(bookings))
			fmt.Printf("The Slice type: %T\n", bookings)
			getfirstnames()
			if Ticketsremaining == 0 {
				fmt.Println("The tickets of go conference has been sold out come back next year")
				break
			}

		}
		if !validname {
			fmt.Println("The firstname or lastname is to short.")
		}
		if !validname {
			fmt.Println("The email does not contain @.")

		}
		if !validticketnumber {
			fmt.Println("The tickets no invalid.")

		}

	}

}

func greetuser() {
	fmt.Printf("Welcome to %v \n", conferenceName)
	fmt.Printf("we have Total %v tickets and %v ticket are still avilable\n", conferencetickets, Ticketsremaining)
	fmt.Println("Get yours ticket here to join the session")
}

func getuserinput() (string, string, string, uint) {
	var userName string
	var userlastName string
	var useremail string
	var usertickets uint
	fmt.Printf("Enter the Name:\n")
	fmt.Scan(&userName)
	fmt.Printf("Enter the LastName:\n")
	fmt.Scan(&userlastName)
	fmt.Printf("Enter the Email:\n")
	fmt.Scan(&useremail)
	fmt.Printf("Enter the Tickets No:\n")
	fmt.Scan(&usertickets)
	return userName, userlastName, useremail, usertickets
}

func getfirstnames() {
	firstnames := []string{}
	for _, booking := range bookings {
		firstnames = append(firstnames, booking.userName)
		fmt.Printf("the first name of bookings:%v\n", firstnames)

	}
}

func bookticket(usertickets uint, conferencetickets uint, userName string, userlastName string, useremail string) {
	Ticketsremaining = conferencetickets - usertickets
	/*creating map for userdata
	var userdata = make(map[string]string)
	userdata["firstname"] = userName
	userdata["userlastname"] = userlastName
	userdata["useremail"] = useremail
	userdata["usertickets"] = strconv.FormatUint(uint64(usertickets), 10)*/
	var userdata = userdata{
		userName:     userName,
		userlastName: userlastName,
		useremail:    useremail,
		usertickets:  usertickets,
	}
	bookings = append(bookings, userdata)
	fmt.Printf("Thank You %v %v for booking %v tickets you will recieve confirmation mail on %v\n", userName, userlastName, usertickets, useremail)
	fmt.Printf("Remaining tickets for go conferene: %v\n", Ticketsremaining)
	fmt.Printf("list of bookings is %v\n", bookings)
}

func sendticket(usertickets uint, userName string, userlastName string, useremail string) {
	time.Sleep(40 * time.Second)
	var Ticket = fmt.Sprintf("The %v tickets to %v %v ", usertickets, userName, userlastName)
	fmt.Println("################")
	fmt.Printf("sending Tickets:%v\n at %v email address ", Ticket, useremail)
	fmt.Println("################")
}
