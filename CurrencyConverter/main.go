package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	var i float64
	fmt.Print("Enter Amount in LKR : ")
	for {
		_, err := fmt.Scan(&s)
		i, err = strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Println("Enter a valid Amount")
		} else {
			usd := fmt.Sprintf("%.2f", i*0.0052)
			fmt.Println("Your amount in USD : ", usd)
			break
		}
	}

}
