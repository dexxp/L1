package main

import (
	"fmt"
)

// Интерфейс платежного шлюза
type PaymentGateway interface {
	ProcessPayment(amount float64)
	GetPaymentStatus() string
}

// Структура для системы оплаты кредитной картой
type CreditCardPayment struct {
	amount float64
}

func (c *CreditCardPayment) ProcessPayment(amount float64) {
	c.amount = amount
	fmt.Printf("Обработка платежа кредитной картой на сумму %.2f\n", amount)
}

func (c *CreditCardPayment) GetPaymentStatus() string {
	return fmt.Sprintf("Платеж кредитной картой на сумму %.2f был успешно обработан", c.amount)
}

// Структура для системы оплаты с помощью мобильного кошелька
type MobileWalletPayment struct {
	amount float64
}

func (m *MobileWalletPayment) MakeMobilePayment(amount float64) {
	m.amount = amount
	fmt.Printf("Совершение платежа с помощью мобильного кошелька на сумму %.2f\n", amount)
}

func (m *MobileWalletPayment) CheckMobilePaymentStatus() string {
	return fmt.Sprintf("Платеж с помощью мобильного кошелька на сумму %.2f был успешно совершен", m.amount)
}

// Адаптер для преобразования системы оплаты с помощью мобильного кошелька в систему оплаты кредитной картой
type MobileWalletAdapter struct {
	mobileWallet *MobileWalletPayment
}

func (a *MobileWalletAdapter) ProcessPayment(amount float64) {
	a.mobileWallet.MakeMobilePayment(amount)
}

func (a *MobileWalletAdapter) GetPaymentStatus() string {
	return a.mobileWallet.CheckMobilePaymentStatus()
}

func main() {
	// Используем систему оплаты кредитной картой
	creditCardPayment := &CreditCardPayment{}
	pay(creditCardPayment, 100.50)

	// Используем систему оплаты мобильного кошелька с помощью адаптера
	mobileWalletPayment := &MobileWalletPayment{}
	mobileWalletAdapter := &MobileWalletAdapter{mobileWallet: mobileWalletPayment}
	pay(mobileWalletAdapter, 50.25)
}

// Функция обработки платежа
func pay(paymentGateway PaymentGateway, amount float64) {
	paymentGateway.ProcessPayment(amount)
	fmt.Println(paymentGateway.GetPaymentStatus())
}
