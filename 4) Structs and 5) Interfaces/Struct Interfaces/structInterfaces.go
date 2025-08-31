package main

import (
	"fmt"
	"math"
	"time"
)

func sendMessage(msg message) (string, int) {
	return msg.getMessage(), len(msg.getMessage()) * 3
}

type message interface {
	getMessage() (string)
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

// Different Example

type shape interface {
	area() float64
	perimeter() float64
}

func getValues(sh shape) (float64, float64) {
	return sh.area(), sh.perimeter()
}


type rectangle struct {
	width, height float64
}
func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Example 2

func getEmployeeInfo(e employee) (string, int) {
	return e.getName(), e.getSalary()
}

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

// MULTIPLE INTERFACES EXAMPLE
type expense interface {
	cost() int
}

// Naming the interface functions signatures
type formatter interface {
	format() (formattedMail string)
}

type email struct {
	isSubscribed bool
	body         string
}

func useEmailUtils(e expense, f formatter) (int, string){
	return e.cost(), f.format()
}

func (e email) cost() int {
	if !e.isSubscribed {
		return 5
	}
	return 2
}

func (e email) format() string {
	if !e.isSubscribed {
		return fmt.Sprintf("%v | Not Subscribed", e.body)
	}
	return fmt.Sprintf("%v | Subscribed", e.body)
}

func main() {
	billDeerBday := birthdayMessage{
		birthdayTime: time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
		recipientName: "Bill Deer",
	}
	message, length := sendMessage(billDeerBday)
	fmt.Println("message: ", message, "length: ", length)

	messagesSent := sendingReport{
		reportName: "newReport",
		numberOfSends: 300,
	}
	message, length = sendMessage(messagesSent)
	fmt.Println("message: ", message, "length: ", length)

	newRectangle := rectangle{
		width: 3,
		height: 4,
	}
	rectArea, rectPerimeter := getValues(newRectangle)
	fmt.Println("rectArea: ", rectArea, "rectPerimeter: ", rectPerimeter)

	newCircle := circle{
		radius: 5,
	}
	circleArea, circlePerimeter := getValues(newCircle)
	fmt.Println("circleArea: ", circleArea, "circlePerimeter: ", circlePerimeter)

	// Example 2
	newContractor := contractor{
		name: "Bob",
		hourlyPay: 100,
		hoursPerYear: 73,
	}
	name, salary := getEmployeeInfo(newContractor)
	fmt.Println("name: ", name, "salary: ", salary)

	newFullTimeEmployee := fullTime{
		name: "FullTime Bob",
		salary: 100000,
	}
	name, salary = getEmployeeInfo(newFullTimeEmployee)
	fmt.Println("name: ", name, "salary: ", salary)

	// MULTIPLE INTERFACES EXAMPLE
	newEmail := email{
		isSubscribed: true,
		body: "Hello... Fuck you",
	}
	cost, format := useEmailUtils(newEmail, newEmail)
	fmt.Println("cost: ", cost, "format: ", format)
}
