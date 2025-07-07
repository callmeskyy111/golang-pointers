package main

import "fmt"

// other practical example: fmt.Scan(&arg)

func main() {
	age := 29 // regular variable
	agePtr:=&age
	fmt.Println("🧠 Age Pointer (&agePtr):",agePtr) //0xc00008c0a8
	fmt.Println("Age (*agePtr):",*agePtr) //29
	editAgeToAdultYears(agePtr)
	fmt.Println("Adult years:",age) //11
}

func editAgeToAdultYears(age *int){
*age-=18
}