package main

import "fmt"

func main() {
	Greet()
	PassParams("Manideep") //used the fun from other file named as params.go

}

func Greet() {
	fmt.Println("Called from Greet Function")
}
