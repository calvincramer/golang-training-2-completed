package main

import (
	"fmt"
	"strconv"
)

// TODO: implement the following method **only** by using defer. The **first** operation is provided
// Step 1: if the number is divisible by 13 multiply it by 7
// Step 2: if the number is negative then take its absolute value
// Step 3: convert the number to a string and strip off the left-most character. If the resulting
// string is the empty string then make it "0". Convert the result back to an integer.
// Step 4: set the 0th, 1st, and 6th bit to 1
//
// For example: Defer(-104) -> (step 1) -728 -> (step 2) 728 -> (step 3) 28 -> (step 4) 95
func Defer(n int) (ret int) {
	// Step 4
	defer func() {
		ret = ret | 0x43
	}()
	// Step 3
	defer func() {
		str := strconv.Itoa(ret)
		if len(str) > 0 {
			str = str[1:]
		}
		if len(str) == 0 {
			str = "0"
		}
		ret, _ = strconv.Atoi(str)
	}()
	// Step 2
	defer func() {
		if ret < 0 {
			ret = -ret
		}
	}()
	// Step 1
	defer func() {
		if ret%13 == 0 {
			ret = ret * 7
		}
	}()
	ret = n
	return
}

// TODO: implement the NumberChecker interface for the `MyInt` type.
type MyInt int

func (n MyInt) IsEven() bool {
	return n%2 == 0
}

func (n MyInt) IsOdd() bool {
	return !n.IsEven()
}

// Hints:
// - a number is prime if its only even divisors are 1 and itself
// - this function will only be tested with very small numbers
// - any number less than 2 is not considered prime
func (n MyInt) IsPrime() bool {
	if n < 2 {
		return false
	}
	for d := 2; MyInt(d) < n; d++ {
		if n%MyInt(d) == 0 {
			return false
		}
	}
	return true
}

// TODO: implement the following function which takes both `int` and `MyInt` using approximation constraints (`~T`)
// The function should create a string in the following format:
// "==== ${msg} ===="
// where each "====" bar is has width characters
func MakeHeaderLine[T ~int](width T, msg string) string {
	str := ""
	var i T = 1
	for ; i <= width; i++ {
		str = str + "="
	}
	str += " "
	str += msg
	str += " "
	i = 1
	for ; i <= width; i++ {
		str = str + "="
	}
	return str
}

// TODO: make an interface that is the set of file readers CSVReader, JSONReader.
// These two concrete types already exist are are implemented in util
type FileReader interface {
	CSVReader | JSONReader
	Read(filePath string) []byte
}

// TODO: construct a message that will be added to the log in the following format:
// "${msg} - ${obj1} - ${obj2}"
// If `obj1` or `obj2` are `int`, print them in hexadecimal (use "0x%X"), otherwise use the "%v" print
// format verb to convert `obj1` and `obj2` to strings
func FormatLogMessage(msg string, obj1 any, obj2 any) string {
	var obj1Str string
	if obj1Int, ok := obj1.(int); ok {
		obj1Str = fmt.Sprintf("0x%X", obj1Int)
	} else {
		obj1Str = fmt.Sprintf("%v", obj1)
	}

	var obj2Str string
	if obj2Int, ok := obj2.(int); ok {
		obj2Str = fmt.Sprintf("0x%X", obj2Int)
	} else {
		obj2Str = fmt.Sprintf("%v", obj2)
	}
	return fmt.Sprintf("%s - %s - %s", msg, obj1Str, obj2Str)
}

// TODO: make a struct that embeds the `Lion`, `Goat`, and `Snake` structs, which are all defined in
// util
type Chimera struct {
	Lion
	Goat
	Snake
}

// TODO: make an interface that embeds the three interfaces: `Roarer`, `Headbutter`, and `Hisser`.
// These interfaces and structs that implement the interfaces are defined in util.
type ChimeraI interface {
	Roarer
	Headbutter
	Hisser
}

// TODO: implement the `Priceable` interface (in util) on both `Sedan` and `Motorcycle`.
// Use pointer receivers.
// The price for motorcycles should be: (numWheels * 500) + (weight * 2) + (coolnessFactor * 1000)
// The price for sedans should be: (numWheels * 500) + (weight * 2) + (engineHorsepower * 10)
func (moto Motorcycle) CalcPrice() uint {
	return (uint(moto.numWheels) * 500) + (moto.weight * 2) + (moto.coolnessFactor * 1000)
}
func (sedan Sedan) CalcPrice() uint {
	return (uint(sedan.numWheels) * 500) + (sedan.weight * 2) + (sedan.engineHorsepower * 10)
}

// TODO: implement the function which adjusts price based on country
// Add the following import tax to the price of the vehicle based on the country:
// - USA: 150%
// - EU: 10%
// - CANADA: 7%
// Hint: convert the vehicle price to a float, the adjust price, then cast to `uint` (rounding down)
func AdjustPrice(p Priceable, territory Territory) uint {
	price := float64(p.CalcPrice())
	switch territory {
	case USA:
		price += price * 1.5
	case EU:
		price += price * 0.1
	case CANADA:
		price += price * 0.07
	}
	return uint(price)
}

// TODO: double the elements in arr. This is an attempt to modify the caller's `arr`. This will be
// unsuccessful though, as golang will pass a copy of the array (pass by value)
func DoubleArr(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
}

// TODO: double the elements in arr. This should modify the caller's `arr` successfully.
func DoubleArrByPtr(arr *[3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
}

// TODO: double the elements in the slice. Slices are passed by reference automatically unlike
// arrays, which is why we don't explicitly need the a pointer to the slice.
func DoubleSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] *= 2
	}
}

// TODO: implement a method on the receiver type (p Point2D) that calculates the area of a rectangle
// bounded by the points (0, 0) and (p.X, p.Y)
// Hint: area should always be non-negative
func (p Point2D) CalcArea() float64 {
	area := p.X * p.Y
	if area < 0 {
		area = -area
	}
	return area
}

// TODO: implement a method on receiver type (p *Point2D) that swaps the X and Y components
func (p *Point2D) Transpose() {
	p.X, p.Y = p.Y, p.X
}

// TODO: calculate the sum of squares in the range 1 to num (inclusive). Do this using a gotos and
// labels **instead** of a for loop. Branching statements like if else statements are allowed.
func GotoLabelSumSquares(num int) int {
	sum := 0
	// Init
	n := 1
loopStart:
	// Condition
	if n > num {
		goto done
	}
	// Body
	if IsSquare(n) {
		sum += n
	}
	// Increment
	n += 1
	goto loopStart
done:
	return sum
}

// TODO: modify the following function to continue execution at the outermost loop when the
// innermost loop has an element that is negative. Perform this action before the element is added
// to the sum.
// Hint: use a break and label
func BreakLabel3D(matrix [][][]int) int {
	sum := 0
	for _, plane := range matrix {
	goHere:
		for _, row := range plane {
			for _, elem := range row {
				// TODO add condition here
				if elem < 0 {
					break goHere
				}
				sum += elem
			}
		}
	}
	return sum
}

// TODO: modify the following function that calculates the sum of a linked list (in a funky way) to
// stop execution once a node with a negative value is encountered.
// The only modification needed is to:
// add a new case in the switch with a labeled break **inside** the switch statement.
func FunkySumLinkedList(node SLL) int {
	sum := 0
	var trav *SLL = &node
done:
	for {
		switch {
		case trav.Val < 0:
			break done
		case trav.Val%2 == 0:
			sum += trav.Val * 10
		default:
			sum += trav.Val
		}
		if trav.Next == nil {
			break
		}
		trav = trav.Next
	}
	return sum
}
