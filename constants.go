package main
import "fmt"
const Pi = 3.14
func main() {
	const world = "Michael"
	fmt.Println("Hello", world)
	fmt.Println("Happy ", Pi, "day")

	const estado = true
	fmt.Println("Go rules?", estado)
}