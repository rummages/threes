package main 

import "fmt"

func main() { 
	for i := 1; i <= 100; i++ { //Creates and sets i as 1 ; for i is less than 100 ; special operator '++' is +1 -- is -1
		if i % 3 == 0 {
			fmt.Println(i, "divisible by three")
		} else {
			fmt.Println(i)
		}	
	}
}
