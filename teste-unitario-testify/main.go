package main

import "fmt"

func main() {
	// A "Production" Example
    fmt.Println("Tutorial teste unitário com mock usando testify")

    messageService := message.MessageServiceImp{}
    myService := message.MyService{messageService}
    myService.ChargeCustomer(100)

}