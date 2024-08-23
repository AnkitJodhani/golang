package main

import "fmt"

type LogLevel int

const (
	INFO LogLevel = iota
	WARN
	ERROR
	DEBUG
)

func changeLogLevel(value LogLevel) {
	fmt.Println("Log Levele is ", value)
}

type orderStatus string

const (
	Confirmed      orderStatus = "Confirmed"
	Waiting        orderStatus = "Waiting"
	OutForDelivery orderStatus = "OutForDelivery"
	Delivered      orderStatus = "Delivered "
)

func changeOrderStatus(status orderStatus) {
	fmt.Println("Order status updated to: ", status)
}

func main() {
	// enumerated type
	changeLogLevel(INFO)
	changeLogLevel(WARN)
	changeOrderStatus(Confirmed)
	changeOrderStatus(Delivered)

}
