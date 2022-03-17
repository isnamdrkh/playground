package main

import "fmt"

// Check Point:
// Cashier
// - Input: Nama Menu, Add Again
// - Output: Total Harga yang harus dibayar

// Contoh:
// Menu makanan:
// 1 . Menu: bakso ,  Price : 20000
// 2 . Menu: burger ,  Price : 15000
// 3 . Menu: sate ,  Price : 25000
// 4 . Menu: sosis ,  Price : 20000
// 5 . Menu: soto ,  Price : 25000

// Input:
// - Form:
//   - Masukan nama menu yang mau dipesan: bakso

// Output:
// Menu telah ditambahkan ke Ordered Menu:
// Menu: bakso ,  Price : 20000

// Input:
// - Ingin memesan menu lain?(yes/no): if no (break) --> if yes (add again)

// Output:
// Data Order Updated:
//  - Menu: sate ,  Price : 25000
//  - Menu: burger ,  Price : 15000
//  - Menu: sosis ,  Price : 20000
//  - Total harga makanan yang harus anda bayar:  60000

func main() {
	foodMenu := map[string]int64{
		"bakso":  20000,
		"burger": 15000,
		"sate":   25000,
		"sosis":  20000,
		"soto":   25000,
	}

	orderMenu := make(map[string]int64)

	fmt.Println("Menu Makanan: ")
	count := 0
	var addAgain string

	for {
		for menu, price := range foodMenu {
			count++
			fmt.Println("- Menu: ", menu, ",", "Price:", price)
		}

		menu := ""
		fmt.Println("Form: ")
		fmt.Println("Masukan nama menu yang mau dipesan: ")
		fmt.Scan(&menu)

		if foodMenu[menu] != 0 {
			orderMenu[menu]++
		}

		fmt.Println("Ingin memesan menu lain?(yes/no): ")
		fmt.Scan(&addAgain)

		if addAgain == "no" {
			break
		}
	}

	var total int64
	for menu, qty := range orderMenu {
		total += foodMenu[menu] * qty
	}

	for menu, qty := range orderMenu {
		count++
		fmt.Println(count, ".Menu: ", menu, ", Quantity: ", qty)
	}

	fmt.Println("Total Price: Rp.", total)

}
