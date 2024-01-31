// Package needed for all Go projects, this allows the compiler to know how to bundle the application
package main

//need to import functions from Go standard library to perform specific actions
import (
	//"booking-app/helper"
	"fmt"
	"sync" //wait groups
	"time" //sleep
	//"strings"
	//"strconv" convert from int to string, etc.
)

// Variables declared outside main, are accessible to all functions, main or otherwise.
const conferenceTickets int = 50

var conferenceName string = "Go Conference" //:= is the 'sugar' variable assignment operator to let Go infer the type of variable
var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0) //make([]map is to make a list of maps)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName    string
	lastName     string
	email        string
	numOfTickets uint
}

var wg = sync.WaitGroup{}

// Main function needed for Go compiler to know where to start
func main() {

	//var declares and assigns variable [name] = value
	//conferenceName := "Go Conference" //:= is the 'sugar' variable assignment operator to let Go infer the type of variable
	//const conferenceTickets int = 50
	//var remainingTickets uint = 50
	//var bookings = [50]string{"Nana", "Nicole"} this is to declare an array with size 50 and values already known.
	//var bookings [50]string //delcare array with size 50 without known values
	//bookings[0] = "Nana" add value to '0' index in array
	//var bookings []string //Slice declaration, dynamic array not needing a size.
	//bookings := []string{} //Sugar assignment := for array/slices

	//demonstrating printing data types
	//fmt.Printf("ConferenceName type: %T, remainingTickets is %T, conferenceTickets is %T.\n", conferenceName, remainingTickets, conferenceTickets)

	//greet users!
	greetUsers()

	//for loop is the only loop, no while, no do while, etc.
	//for {} is infinite loop
	//for { removed for goroutine/wait/concurrency

	fmt.Printf("variable value for Remaining tickets: %v.\n", remainingTickets)   //value of remainingTickets variable
	fmt.Printf("variable pointer for remaining tickets %v.\n", &remainingTickets) //& = Pointer to location of value, returns the location of that value in memory

	firstName, lastName, email, userTickets := getUserInput()
	//bookings[0] = firstName + " " + lastName array addition to index 0

	isValidName, isValidEmail, isValidTicketNum := validateUserInput(firstName, lastName, email, userTickets)
	//isValidName, isValidEmail, isValidTicketNum := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets) //use our built helper package and function with capital letter to import to code.

	if isValidName && isValidEmail && isValidTicketNum {

		//book tickets func
		bookTicket(userTickets, email, firstName, lastName)

		wg.Add(1)                                              //adds (n) threads for (n) func/go routines needed. we need 1 for sendTicket so only 1 is used.
		go sendTicket(userTickets, firstName, lastName, email) //go keyword adds goroutine/concurrency

		firstNames := getFirstNames()
		fmt.Printf("First names of all booked users: %v.\n", firstNames)

		//create bool that evaluates true or false on a variable.
		var noTicketsRemaining bool = remainingTickets == 0

		//check bool variable, to decide if true or false and if true, conference is fully booked and application breaks(closes).
		if noTicketsRemaining {
			fmt.Printf("Remaining tickets = %v. This conference is fully booked, please come back next year.\n", remainingTickets)
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("Your first or last name entered is too short.")
		}

		if !isValidEmail {
			fmt.Println("Your email address is not a valid email address.")
		}

		if !isValidTicketNum {
			fmt.Println("The number of tickets you requested is invalid.")
		}
	}
	//switch case example
	//city := "London"
	//
	//switch city {
	//case "New York":
	//code here
	//case "Singapore":
	//code here
	//default:
	//report no values correct
	//}
	wg.Wait() //waits for any threads not part of main.
}

// function with variables passed as parameters
func greetUsers() {

	//fmt package using the print function from that package
	//Printf allows for variable formatting
	//Println allows for print with newlines

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

// function that has slice params and slice return
// func name(params) return value {}
func getFirstNames() []string {
	firstNames := []string{}
	//to loop over a data structure(array in this example) you need for index, new varName := range array name.
	//to ignore one of the required unused items(index for example), you place an underscore(_) to signify it is unused.
	for _, booking := range bookings {
		//firstNames = append(firstNames, booking["firstName"]) map example for getting firstname
		firstNames = append(firstNames, booking.firstName)

	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	//declare non static variable
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, email string, firstName string, lastName string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user, make() creates the map empty
	//var userData = make(map[string]string) map example
	var userData = UserData{ //struct example
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		numOfTickets: userTickets,
	}

	//map key[]:value pairs for user
	//userData["firstName"] = firstName
	//userData["lastName"] = lastName
	//userData["email"] = email
	//userData["numOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("User %v %v booked %v tickets. Email confirmation sent to: %v.\n", firstName, lastName, userTickets, email)

	//fmt.Printf("First index of bookings: %v.\n", bookings[0])
	//fmt.Printf("Bookings array/slice type: %T.\n", bookings)
	//fmt.Printf("Bookings array/slice length: %v.\n", len(bookings))
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

	fmt.Printf("List of all current bookings: %v\n", bookings)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) //stored value of what would have been printed to the console. fmt.Sprintf

	fmt.Println("#################")
	fmt.Printf("Sending Ticket:\n%v\nto email: %v\n", ticket, email)
	fmt.Println("#################")

	wg.Done() //clears thread from waitlist
}
