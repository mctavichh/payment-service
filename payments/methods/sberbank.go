package methods

import (
	"fmt"
	"math/rand"
	"time"
)

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (c *Bank) Pay(usd int) int {
	fmt.Println("Оплата через банк")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Сумма оплаты", usd, "USDT")

	return rand.Int()
}

func (c *Bank) Cancel(id int) {
	fmt.Println("Отмена банковской операции под номером:", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Банковская операция была успешно отменена")
}
