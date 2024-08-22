/*
🔥 Note: if you don't specify value of the variable then go will initilize with ZERO value

🔸 for int   -> 0
🔸 for float -> 0.0
🔸 for sting -> ""
🔸 for bool  -> false

*/

package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nenosecond precision
}

// reciver - to create method for struct
func (o *order) changeStatus(status string) string {
	o.status = status
	return o.status
}
func (o *order) changeAmount(amount float32) float32 {
	o.amount = amount
	return o.amount
}

func main() {
	//  🔥 Creating struct object
	myOrder1 := order{
		id:     "1",
		amount: 34.23,
		status: "Waiting",
	}
	myOrder2 := order{
		id:     "2",
		amount: 234.3432,
		status: "Waiting",
	}
	myOrder1.createdAt = time.Now()
	myOrder2.createdAt = time.Now()
	fmt.Println(myOrder1)
	fmt.Println(myOrder2)

	myOrder1.changeAmount(1000.00)    // calling methods
	myOrder2.changeAmount(1000.00)    // calling methods
	myOrder1.changeStatus("Executed") // calling methods
	myOrder2.changeStatus("Executed") // calling methods

	fmt.Println(myOrder1)
	fmt.Println(myOrder2)

	// fmt.Println(myOrder1.amount)
	// fmt.Println(myOrder2.amount)
}
