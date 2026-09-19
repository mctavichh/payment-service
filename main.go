package main

import (
	"fmt"
	"jobprojects/payments"
	"jobprojects/payments/methods"
	"time"

	"github.com/k0kubun/pp"
)

func main() {
	method := methods.NewPaypal()
	paymentModule := payments.NewPaymentModule(&method)

	paymentModule.Pay("Бургер", 5)
	paymentModule.Pay("BMW 330i", 30_000)
	paymentModule.Pay("Rust", 11)
	idPizza := paymentModule.Pay("Пицца", 6)
	paymentModule.Cancel(idPizza)

	allinfo := paymentModule.AllInfo()

	fmt.Println("Вся сумма оплат:")
	fmt.Println("before")
	time.Sleep(500 * time.Millisecond)
	pp.Println(allinfo)
	infoPizza := paymentModule.Info(idPizza)
	pp.Println(infoPizza)

}
