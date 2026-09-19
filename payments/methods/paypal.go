package methods

import (
	"fmt"
	"math/rand"
	"time"
)

type Paypal struct{}

func NewPaypal() Paypal {
	return Paypal{}
}

func (c *Paypal) Pay(usd int) int {
	fmt.Println("Оплата через PayPal")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Сумма оплаты", usd, "долларов")

	return rand.Intn(999) + 100
}

func (c *Paypal) Cancel(id int) {
	fmt.Println("Отмена банковской операции под номером:", id)
}
