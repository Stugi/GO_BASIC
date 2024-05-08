package main

import (
	electronic "electronic/phone"
)

func main() {
	applePhone := electronic.NewApplePhone("15pro")
	androidPhone := electronic.NewAndroidPhone("samsung", "galaxy")
	radioPhone := electronic.NewRadioPhone("sony", "xperia", 10)

	printCharacteristics(applePhone)
	printCharacteristics(androidPhone)
	printCharacteristics(radioPhone)
}

func printCharacteristics(phone electronic.Phone) {
	println(phone.Brand())
	println(phone.Model())
	println(phone.Type())
	switch phone.(type) {
	case electronic.Smartphone:
		println(phone.(electronic.Smartphone).OS())
	case electronic.StationPhone:
		println(phone.(electronic.StationPhone).ButtonCount())
	}

	println()
}
