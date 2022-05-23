package main
import "fmt"

func name(nombre, apellido1, apellido2 string) (string, string, string){
	return nombre, apellido1, apellido2
}

func main() {
	a, b, c := name("michael", "Alexis", "Chacon")
	fmt.Println(a, b, c)
}