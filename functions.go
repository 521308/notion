package main

import "fmt"

func PrintMain() {
	fmt.Println("MAIN MENU (please select):")
	fmt.Println("1. Buy")
	fmt.Println("2. Fill")
	fmt.Println("3. Take")
	fmt.Println("4. Stat")
	fmt.Println("5. Exit")
}
func PrintStat() {
	fmt.Print("В машине:\n")
}
func PrintTake() {
	fmt.Print("В машине было: ")
}
func PrintBuy() {
	fmt.Print("В меню:\n\t1-Espresso\n\t2-Latte\n\t3-Cappucino")
}
