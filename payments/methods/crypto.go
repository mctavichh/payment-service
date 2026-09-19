package methods

import (
	"fmt"
	"math/rand"
	"time"
)

type Crypto struct{}

func NewCrypto() Crypto {
	return Crypto{}
}

func (c *Crypto) Pay(usd int) int {
	fmt.Println("Оплата криптовалютой")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Сумма оплаты", usd, "долларов")

	return rand.Int()
}

func (c *Crypto) Cancel(id int) {
	fmt.Println("Отмена крипто-операции под номером:", id)
}
