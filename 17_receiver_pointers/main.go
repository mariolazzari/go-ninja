package main

import "fmt"

func main() {
	myBill := newBill("Mario's bill")
	myBill.updateTip(5)

	myBill.addItem("coffee", 1.99)
	myBill.addItem("pie", 3.99)

	fmt.Println((myBill.format()))
}
