package main

import (
	"fmt"
	"reflect"
)

type boy struct {
	Name string
	age  int
}

type human interface {
	SayName()
	SayAge()
}

func (this *boy) SayName() {
	fmt.Println(this.Name)
}

func (this *boy) SayAge() {
	fmt.Println(this.age)
}

func reflect_2() {
	var i human
	jown := &boy{
		Name: "jown",
		age:  15,
	}

	i = jown

	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	fmt.Printf("Reflect Value: ", v, "Reflect Type: ", t, "\n")
}

func reflect_1() {
	var circle float64 = 3.1415926
	var icir interface{}

	icir = circle
	fmt.Println("Reflect type: ", reflect.TypeOf(icir))
	fmt.Println("Reflect Value: ", reflect.ValueOf(icir))

	icir2 := reflect.ValueOf(&circle)
	fmt.Println("Reflect Value: ", reflect.ValueOf(icir2))

	icir3 := icir2.Elem()
	fmt.Println("Reflect Value: ", reflect.ValueOf(icir))

	icir3.SetFloat(3.14)
	fmt.Println("Reflect Value: ", reflect.ValueOf(icir))
	fmt.Println("Reflect Value: ", reflect.ValueOf(circle))
}

func main() {
	//reflect_1()
	reflect_2()
}
