package main

import (
	"fmt"
	"math"
)

type NumberChecker interface {
	IsEven() bool  // determines if a number is even
	IsOdd() bool   // determines if a number is odd
	IsPrime() bool // determines if a number is prime
}

type CSVReader struct{}
type JSONReader struct{}

func (reader CSVReader) Read(filePath string) []byte {
	return []byte{0x0}
}

func (reader JSONReader) Read(filePath string) []byte {
	return []byte{0x0}
}

type Roarer interface {
	Roar() string
}

type Headbutter interface {
	Headbutt() string
}

type Hisser interface {
	Hiss() string
}

type Lion struct {
	weight int
}

func (lion *Lion) Roar() string {
	return fmt.Sprintf("Lion of weight %d roars", lion.weight)
}

type Goat struct {
	hornLength int
}

func (goat *Goat) Headbutt() string {
	return fmt.Sprintf("Goat with horn size %d roars", goat.hornLength)
}

type Snake struct{}

func (snake *Snake) Hiss() string {
	return "Snake hisses"
}

type Vehicle struct {
	numWheels uint8
	weight    uint
}

type Motorcycle struct {
	Vehicle
	coolnessFactor uint // something specific about motorcycles
}

type Sedan struct {
	Vehicle
	engineHorsepower uint
}

type Priceable interface {
	CalcPrice() uint
}

type Territory int

const (
	USA Territory = iota
	EU
	CANADA
)

type Point2D struct {
	X float64
	Y float64
}

func IsSquare(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}

type SLL struct {
	Val  int
	Next *SLL
}
