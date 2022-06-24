package main

import (
	"fmt"
	"math"
)

type cube struct {
	length float64
}

type box struct {
	length float64
	width  float64
	height float64
}

type sphere struct {
	radius float64
}

type ofStructure interface {
	volumn() float64
}

//fuctions for volumn()
func (c cube) volumn() float64 {
	return c.length * c.length * c.length
}

func (b box) volumn() float64 {
	return b.length * b.width * b.height
}

func (s sphere) volumn() float64 {
	return (4 * math.Pi * math.Pow(s.radius, float64(3))) / 3
}

func calculateVolume(kind ofStructure, called string) {
	fmt.Printf("The Volume calculated for our %s is: %f \n", called, kind.volumn())
}

func main() {

	b := box{5, 5, 5}
	c := cube{7}
	s := sphere{5}

	calculateVolume(b, "box")
	calculateVolume(c, "cube")
	calculateVolume(s, "sphere")

}
