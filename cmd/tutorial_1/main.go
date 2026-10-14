package main

import "fmt"

func main() {
	fmt.Println("Hello Word")
	var intNum int16 = 32767
	fmt.Println(intNum)

	var mystring string = "Mani"
	fmt.Println(mystring)

	var concatstring string = "Mani" + " " + "Deep"
	fmt.Println(concatstring)

	// fmt.Println(len("Mani"))
	fullname := "Sachidnanda Manideep\n"
	fmt.Print(fullname)

	x, y := 2, 3
	fmt.Println(x, y)

	//const are the same like variables but we cannot change once we had updated the variables
	const initial string = "Karalapati"
	// initial = "Konakondla" // which is not possbile here

}
