package calculate

import "fmt"

func GetPrimeNumber(number int) int {
	fmt.Println("Received Request from client to calculate package under GetPrimeNumber module", number)
	return number + 10

}
