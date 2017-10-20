// Workshop Golang de Meli
// Clase 2 - Ej3
// Hacer un programa que reciba por parametro una lista de enteros, los almacene en un slice,
// popule un arbol binario y despues imprima los valores en orden ascendente.

package main

import (
	"fmt"
	"os"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func add(t *tree, value int) *tree {

	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func print(t *tree) {

	if t != nil {
		print(t.left)
		fmt.Printf("%d ", t.value)
		print(t.right)
	}
}

func main() {

	var arbol *tree
	valores := os.Args[1:]

	for _, item := range valores {

		if pInt, error := strconv.Atoi(item); error == nil {
			arbol = add(arbol, pInt)

		} else {
			fmt.Printf("No se puede convertir a entero el elemento: %s", item)
		}
	}

	fmt.Println("\n******* Imprimiendo arbol *******")
	print(arbol)
}
