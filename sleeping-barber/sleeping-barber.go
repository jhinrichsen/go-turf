/*
Sleeping barber problem as described by Wikipedia:

The analogy is based upon a hypothetical barber shop with one barber.
The barber has one barber chair and a waiting room with a number of chairs in it.
When the barber finishes cutting a customer's hair, he dismisses the customer and then goes to the waiting room to see if there are other customers waiting.
If there are, he brings one of them back to the chair and cuts his hair.
If there are no other customers waiting, he returns to his chair and sleeps in it.

Each customer, when he arrives, looks to see what the barber is doing.
If the barber is sleeping, then the customer wakes him up and sits in the chair.
If the barber is cutting hair, then the customer goes to the waiting room.
If there is a free chair in the waiting room, the customer sits in it and waits his turn.
If there is no free chair, then the customer leaves.
*/

package main

import (
	"log"
	"math/rand"
	"time"
)

type Customer struct {
	Id int
}

func barber(customer chan Customer, waitingRoom chan Customer) {
	for {
		log.Println("Barber is sleeping")
		var c Customer
		select {
			case c = <-customer:
				log.Printf("Customer c %s entered\n", c)
		}
		select {
			case c = <-customer:
				log.Printf("Customer c %s entered\n", c)
			case c = <-waitingRoom:
				log.Printf("Picked next customer %s from waiting room\n", c)
			default:
		}
		log.Printf("Starting to cut hair for customer %d\n", c.Id)
		haircut(c)
	}
}

func haircut(c Customer) {
	start := time.Now()
/*
	for i := rand.Intn(10000); i > 0; i-- {
		// NOP
	}
*/
	duration := rand.Intn(10)
	time.Sleep(time.Duration(duration) * time.Second)
	log.Printf("Haircut for customer %d took %d\n", c.Id, time.Since(start))
}

func main() {
	barberChannel := make(chan Customer, 1)
	waitingRoomChannel := make(chan Customer, 10)
	go barber(barberChannel, waitingRoomChannel)

	for counter := 0; counter < 20; counter++ {
		customer := new(Customer)
		customer.Id = counter
		log.Printf("New customer %d\n", customer)
		// First: try to go to the barber
		select {
			case barberChannel <- *customer:
			default:
				log.Printf("Barber is busy, customer %d is going to the waiting room\n", customer.Id)
				select {
					case waitingRoomChannel <- *customer:
					default:
						log.Printf("Waiting room is full, customer %d is returning another day\n", customer.Id)
				}
		}
		duration := rand.Intn(10)
		time.Sleep(time.Duration(duration) * time.Second)
	}
}

