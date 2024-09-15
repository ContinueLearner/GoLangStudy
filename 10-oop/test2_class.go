package main

import "fmt"

type Hero struct {
	Name  string
	Age   int
	level int
}

func (this *Hero) Show() {
	fmt.Println("Name:", this.Name)
	fmt.Println("Age:", this.Age)
	fmt.Println("level:", this.level)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(name string) {
	this.Name = name
}
func main() {

	hero := Hero{Name: "zhang3", Age: 100}
	hero.Show()

	hero.SetName("lisi")
	hero.Show()

	hero.GetName()
}
