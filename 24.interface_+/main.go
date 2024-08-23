/*

In Go, an interface is a type that specifies a "set of method signatures" but does not provide the
implementation of those methods. An interface allows you to define behavior without specifying how
that behavior is implemented, which makes your code more flexible and modular.

*/

package main

import "fmt"

type Shape interface {
	area() float32
}

type Circle struct {
	redius float32
}

func (c Circle) area() float32 {
	return (3.14 * c.redius * c.redius)
}

type Rectangle struct {
	width, height float32
}

func (r Rectangle) area() float32 {
	return (r.width * r.height)
}

func areaFinder(a Shape) float32 {
	return a.area()
}

func main() {
	obj1 := Circle{redius: 2}
	obj2 := Rectangle{width: 2, height: 3}

	fmt.Println(areaFinder(obj1))
	fmt.Println(areaFinder(obj2))

}
