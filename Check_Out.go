// Justin Rodriguez
// CS 3364 HW 5
// Shin
package main

import (
	"fmt"
	"math/rand"
)

var deskClerkInterface DeskClerk_Interface
var checkoutManager CheckOutManager
var room Room
var bill Bill
var cardReaderInterface CardReader_Interface
var bankInterface Bank_Interface

// Check Out Manager Interfaces and Entity Classes
//*****************************************************************************************************************************************//
// Interfaces
type DeskClerk_Interface struct {
}

type Bank_Interface struct {
}

type CardReader_Interface struct {
}

type CheckOutManager struct {
}

type ReceiptPrinter_Interface struct {
}

type BillPrinter_Interface struct {
}

//Entity Classes (Room, Bill, and Customer). Fields are global.
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

// Interface/Entity Methods
//*********************************************************************************************************************************//
func (Desk DeskClerk_Interface) CheckOut(roomNo int) Bill {
	return checkoutManager.RequestBill(room.RoomNo)
}

// Message display for bill printed and credit card denied.
func (Desk DeskClerk_Interface) DisplayMessage(billPrinted bool) {
	if billPrinted {
		fmt.Println("Credit Card Accepted. Bill Printed.")
	} else {
		fmt.Println("Credit Card Declined.")
	}
}

// bank logic approves and charges credit card.
func (bankInterface Bank_Interface) ChargeCreditCard(cardNumber, total int) (bool, int) {
	return true, rand.Int()
}

// Customer enters credit card as form of payment.
func (card CardReader_Interface) ReadCard() int {
	return rand.Int()
}

// Customer requests bill via desk clerk by giving the room number.
func (checkout CheckOutManager) RequestBill(roomNumber int) Bill {
	return ReadBill(roomNumber)
}

// Customer pays via credit card which is processed by the checkout manager.
func (checkout CheckOutManager) PayByCreditCard(cardNumber int) {
	return
}

// Once credit card is approved, checkout manager requests the receipt to be printed.
func PrintReciept(cardNumber, total, reference int) {
	fmt.Printf("Reciept: CardNum - %d | Total - $%d | Reference - %d \n", cardNumber, total, reference)
}

// Print bill to user
func PrintBill(bill Bill) bool {
	fmt.Printf("%s | ReferenceNo - %d | RoomNo - %d | Total - $%d\n", bill.Bill, bill.ReferenceNo, bill.RoomNo, bill.Total)
	return true
}

// Once the customer has paid the bill, release the room and set to 'Available'
func ReleaseRoom(roomNumber int) bool {
	room.Status = "Available"
	return true
}

// Once the customer requests to check out, read the bill and send back to desk clerk
func ReadBill(roomNumber int) Bill {
	return bill
}

// Read the total of the room that was entered, and return back to desk clerk
func (bill Bill) readTotal(roomNumber int) int {
	return bill.Total
}

// Once credit card is approved the checkout manager will send for a reference number to be generated for the receipt.
func (bill *Bill) updateReference(roomNumber, referenceNumber int) {
	bill.ReferenceNo = referenceNumber
	return
}

// Delete the customer after
func(customer Customer) DeleteCustomer(roomNumber int) bool {
	fmt.Println("Customer has been deleted.")
	return true
}

//Check Out Manager Logic
//**************************************************************************************************************************************//
func main() {
	// creates the room and bill classes
	room = Room{
		RoomNo: 103,
		Status: "Occupied"}
	bill = Bill{
		RoomNo: 103,
		Bill:   "Sam's Bill",
		Total:  200,
	}
	// logic for checkout manager and contol flow
	if room.Status != "Available"{
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

	} else{
		fmt.Printf("Room %d does not have an occupiant. Is Available.",room.RoomNo)
	}

}
