package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const numberOfPizzas = 10

var (
	pizzasMade   int
	pizzasFailed int
	total        int
)

type Poducer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Poducer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}
func makePiza(pizzanumber int) *PizzaOrder {
	pizzanumber++
	if pizzanumber <= numberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recieved order %d \n", pizzanumber)
		ran := rand.Intn(12) + 1
		msg := ""
		success := false
		if ran < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++
		fmt.Printf("Making pizza %d. it will take %d seconds.... \n", pizzanumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)
		if ran <= 2 {
			msg = fmt.Sprintf("** We ran out of ingredients for piza %d", pizzanumber)
		} else if ran <= 4 {
			msg = fmt.Sprintf("** The cook had an accident for piza %d", pizzanumber)

		} else {
			success = true
			msg = fmt.Sprintf("we succesifully cooked %d", pizzanumber)
		}
		return &PizzaOrder{
			pizzaNumber: pizzanumber,
			message:     msg,
			success:     success,
		}
	}
	return &PizzaOrder{
		pizzaNumber: pizzanumber,
	}
}
func pizzeria(pizzamaker *Poducer) {
	// keep track of which pizza we are p=making
	var i = 0
	// run forrever until we get quit
	// try to make opizzas
	for {
		currentPizza := makePiza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we try to make a pizza
			case pizzamaker.data <- *currentPizza:
				//  we want to quit the program
			case quitChan := <-pizzamaker.quit:

				close(pizzamaker.data)
				close(quitChan)
				return
			}
		}

	}
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out the message
	color.Cyan("The Pizzeria is open for business")
	color.Cyan("----------------------------------")
	// create a producer
	pizzaJob := &Poducer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	// run a producer in the background
	go pizzeria(pizzaJob)
	// create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= numberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order %d is is out for delivery", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("the customer is really mad!")
			}
		} else {
			color.Cyan("done making pizzas")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing the channel", err)
			}
		}
	}
	// print out the ending message
	color.Cyan("-----------------")
	color.Cyan("Done for the day")
	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts", pizzasMade, pizzasFailed, total)
	switch {
	case pizzasFailed > 9:
		color.Red("an aweful day")
	case pizzasFailed >= 5:
		color.Blue("it was ok")
	default:
		color.Green("it was a great day")
	}
}
