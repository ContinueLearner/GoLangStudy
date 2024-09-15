package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

//多态
func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())
}

func main() {
	cat := Cat{"Red"}
	dog := Dog{"Green"}
	showAnimal(&cat)
	showAnimal(&dog)
}
