package counter

import "fmt"

func Counte(num1 int, num2 int) {

	for i := num1; i < num2; i++ {
		fmt.Printf("%d\n", i)
		fmt.Println("четно : ", i%2 == 0)
	}

}
