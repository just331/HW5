// Justin Rodriguez
// CS 3364 HW 5
// Shin
package main

import (
	"fmt"
)

var deskClerkInterface DeskClerkInterface
var checkoutManager CheckOutManager
var room Room
var bill Bill
var cardReaderInterface CardReaderInterface
var bankInterface BankInterface

type DeskClerkInterface struct {
}

type BankInterface struct {
}

type CardReaderInterface struct {
}


type CheckOutManager struct {
}

type RecieptPrinterInterface struct {
}

type BillPrinterInterface struct {
}

type Room struct {
	RoomNo int
	Status string
}

type Bill struct {
	RoomNo      int
	Bill        string
	Total       int
	ReferenceNo int
}

type Customer struct {
	RoomNo  int
	Name    string
	PhoneNo string
}
func (DSI DeskClerkInterface) CheckOut(roomNo int) Bill {
	return checkoutManager.RequestBill(room.RoomNo)
}

func (DSI DeskClerkInterface) DisplayMessage(billPrinted bool) { //todo(jrod): BillPrinted Also CCDenied
	if billPrinted {
		fmt.Println("Bill Printed. Congrats")
	} else {
		fmt.Println("Credit Card Declined")
	}
}

func (BI BankInterface) ChargeCreditCard(cardNumber, total int) (bool, int) {
	return true, 8787
}

func (CRI CardReaderInterface) ReadCard() int {
	return 3212321232123
}

func (COM CheckOutManager) RequestBill(roomNumber int) Bill {
	return ReadBill(roomNumber)
}

func (COM CheckOutManager) PayByCreditCard(cardNumber int) {
	return
}

func PrintReciept(cardNumber, total, reference int) {
	fmt.Println("Reciept: ", cardNumber, total, reference)
}

func PrintBill(bill Bill) bool {
	fmt.Println(bill.Bill, bill.ReferenceNo, bill.RoomNo, bill.Total)
	return true
}

func ReleaseRoom(roomNumber int) bool {
	room.Status = "AVAILABLE"
	return true
}
func ReadBill(roomNumber int) Bill {
	return bill /* todo(jrod): make loop for bill matching roomNo */
}
func (bill Bill) readTotal(roomNumber int) int {
	return 0
}

func (bill *Bill) updateReference(roomNumber, referenceNumber int) {
	bill.ReferenceNo = referenceNumber
	return
}

func (cusomter Customer) DeleteCustomer(roomNumber int) bool {
	//todo(jrod): Delete the customer.
	return true
}
func main() {
	room = Room{
		RoomNo: 8,
		Status: "OCCUPIED"}
	bill = Bill{
		RoomNo: 8,
		Bill:   "John's Bill",
		Total:  200,
	}
	outBill := deskClerkInterface.CheckOut(room.RoomNo)
	cardNo := cardReaderInterface.ReadCard()
	total := outBill.readTotal(room.RoomNo)
	chargeRes, referenceNo := bankInterface.ChargeCreditCard(cardNo, total)
	if chargeRes {
		PrintReciept(cardNo, total, referenceNo)
		bill.updateReference(room.RoomNo, referenceNo)
		ReleaseRoom(room.RoomNo)
		bill = ReadBill(room.RoomNo)
		billPrinted := PrintBill(bill)
		deskClerkInterface.DisplayMessage(billPrinted)
	} else {
		deskClerkInterface.DisplayMessage(false)
	}
}
