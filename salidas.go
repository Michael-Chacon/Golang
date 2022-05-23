package main
import "fmt"

func operacion(suma int) (a, b int) {
	a = suma * 2 / 2
	b = suma - 2
	return
}

func main() {
	fmt.Println(operacion(20))
}