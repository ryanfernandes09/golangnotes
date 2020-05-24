package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("\tMENU")
	fmt.Println("0. playground")
	fmt.Println("1. arraySlice")
	fmt.Println("2. weirdMaps")
	fmt.Println("3. canadianProvincesSlice")
	fmt.Println("4. simpleMap")
	fmt.Println("5. structs")
	fmt.Println("6. embeddedStruct")
	fmt.Println("7. anonymousStruct")
	fmt.Println("8. interfaceAndReceiver")
	fmt.Println("9. anonymousFunc")
	fmt.Println("10. returningFunctions")
	fmt.Println("11. callback")
	fmt.Println("12. closure")
	fmt.Println("13. deferredFunc")
	fmt.Println("14. sharingFunctions")
	fmt.Println("15. ...")
	fmt.Println("16. ...")
	fmt.Println("17. ...")
	fmt.Println("18. ...")
	fmt.Println("19. ...")
	fmt.Println("20. ...")
	fmt.Println("21. ...")
	fmt.Println("22. ...")
	fmt.Println("23. Exit")
	fmt.Printf("Enter your choice: ")

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	switch {
	case input.Text() == "0":
		playground()
	case input.Text() == "1":
		arraySlice()
	case input.Text() == "2":
		weirdMaps()
	case input.Text() == "3":
		canadianProvincesSlice()
	case input.Text() == "4":
		simpleMap()
	case input.Text() == "5":
		structs()
	case input.Text() == "6":
		embeddedStruct()
	case input.Text() == "7":
		anonymousStruct()
	case input.Text() == "8":
		interfaceAndReceiver()
	case input.Text() == "9":
		anonymousFunc()
	case input.Text() == "10":
		returningFunctions()
	case input.Text() == "11":
		callback()
	case input.Text() == "12":
		closure()
	case input.Text() == "13":
		deferredFunc()
	case input.Text() == "14":
		sharingFunctions()
	case input.Text() == "15":
		maintenancePage()
	case input.Text() == "16":
		maintenancePage()
	case input.Text() == "17":
		maintenancePage()
	case input.Text() == "18":
		maintenancePage()
	case input.Text() == "19":
		maintenancePage()
	case input.Text() == "20":
		maintenancePage()
	case input.Text() == "21":
		maintenancePage()
	case input.Text() == "22":
		maintenancePage()
	case input.Text() == "23":
		os.Exit(0)
	default:
		fmt.Println("Please enter a valid choice.")
	}
}
