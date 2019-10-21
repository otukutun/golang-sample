
package main

import "fmt"

func fizz(x int) (result string) {
	 if x % 3 == 0 {
		 return "fizz"
	 } else {
		 return ""
	 }
}

func buzz(x int) (result string) {
	 if x % 5 == 0 {
		 return "buzz"
	 } else {
		 return ""
	 }
}

func fizzBuzz(x int) (result string) {
	return fizz(x) + buzz(x)
}

func main() {
	fmt.Println(fizzBuzz(15))
}
