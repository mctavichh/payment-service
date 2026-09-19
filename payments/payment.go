package payments

type PaymentInfo struct {
	Description string
	Usd         int
	Cancelled   bool
} // создание структуры

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PayMentModule struct {
	paymentsInfo  map[int]PaymentInfo // создание мапы в которой хранятся все проведенные операции
	paymentMethod PaymentMethod 
}

func NewPaymentModule(paymentMethod PaymentMethod) *PayMentModule { // берет с main функцию которая будет удовлетворять интерфейс
	return &PayMentModule{ // // и сохраняет ее внутри пеймент модуля посредством использования указателя
		paymentsInfo:  make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	} // возвращает
}

func (p *PayMentModule) Pay(description string, usd int) int {
	id := p.paymentMethod.Pay(usd) // вызываю интерфейс pay из метода, чтобы закинуть туда сумму оплаты и получить id
	info := PaymentInfo{
		Description: description,
		Usd:         usd,
		Cancelled:   false,
	}

	p.paymentsInfo[id] = info // используем id как название ключа в мапе, а value будет = info
	return id // вернет все это в NewPaymentModule, куда далее NewPaymentModule передаст инфо в мапу
}

func (p *PayMentModule) Cancel(id int) {
	p.paymentMethod.Cancel(id) // 
	info, ok := p.paymentsInfo[id]

	if !ok {
		return
	}
	info.Cancelled = true
	p.paymentsInfo[id] = info // аналогичная ситуация с верхним кодом
}

func (p *PayMentModule) Info(id int) PaymentInfo {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return PaymentInfo{}
	}
	return info

}

func (p *PayMentModule) AllInfo() map[int]PaymentInfo {
	copymap := make(map[int]PaymentInfo, len(p.paymentsInfo))
	for k, v := range p.paymentsInfo {
		copymap[k] = v

	}
	return copymap
}
