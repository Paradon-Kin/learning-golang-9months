package room

import "fmt"

type Room struct {
	Number     int
	price      float64
	isOccupied bool
	guestName  string
}

func NewRoom(no int, price float64) *Room {
	return &Room{
		Number:     no,
		price:      price,
		isOccupied: false,
	}
}

func (r *Room) CheckIn(name string) {
	if !r.isOccupied {
		r.guestName = name
		r.isOccupied = true
		fmt.Printf("Check in success|Guess %s (Room number %d)|Price %.2f\n", r.guestName, r.Number, r.price)
	} else {
		fmt.Println("Room Occupied")
	}
}

func (r *Room) CheckOut(q int) {
	r.guestName = ""
	r.isOccupied = false

	totalPrice := r.price * float64(q)
	fmt.Printf("--- Receipt | %d Nights | Total Price %.2f(THB) ---\n", q, totalPrice)
	fmt.Println("Check Out success")
}

func (r *Room) GetStatus() bool {
	return r.isOccupied
}
