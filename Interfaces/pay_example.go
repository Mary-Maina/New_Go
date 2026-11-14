package main

import "fmt"

type Payment interface{
	Pay()
}//Now this is the interface. 

type Mpesa struct{}

func (m Mpesa) Pay(){
	fmt.Println("Payment is done through Mpesa")
}
type Paypal struct{}

func (p Paypal) Pay(){
	fmt.Println("Payment is done through Paypal")
}
type Bankcard struct{}

func (b Bankcard) Pay(){
	fmt.Println("Payment done through bank account")
}
func MakePayment(p Payment){
	p.Pay()
}

func main(){
	mpesa := Mpesa{}
	paypal := Paypal{}
	bank := Bankcard{}


	MakePayment(mpesa)
	MakePayment(paypal)
	MakePayment(bank)
}
