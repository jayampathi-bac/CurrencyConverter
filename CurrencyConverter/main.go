package main

import "fmt"

func main() {
	var rupee float64
	fmt.Print("Enter amount in LKR : ")
	fmt.Scan(&rupee)
	fmt.Println(rupee)
	if rupee!=0{
		fmt.Println("Your amount in USD : ", rupee*0.0051)
	}else  {
		fmt.Println("Please Enter a Valid Amount")
	}




}