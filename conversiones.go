package main
import (
	"fmt"
	"math"
)
// %T = imprimir el tipo de dato
// %v ver el valor del dato
// %d numero en base 10
// %o nomero en vase 8
func main() {
	var x, y int = 3, 4
	var  f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Printf("%T %T %T", x, y, z)
}