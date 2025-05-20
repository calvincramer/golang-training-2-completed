package main

import (
	"fmt"
	"testing"
)

func sliceEq[T comparable](s1 []T, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestDefer(t *testing.T) {
	if Defer(-104) != 95 {
		t.Fatalf("Defer(-104) expected 95")
	}
	if Defer(578) != 79 {
		t.Fatalf("Defer(578) expected 79")
	}
	if Defer(9) != 67 {
		t.Fatalf("Defer(9) expected 67")
	}
}

func TestNumberCheckerInterface(t *testing.T) {
	if MyInt(50).IsEven() != true {
		t.Fatal("Expected 50 to be even")
	}
	if MyInt(50).IsOdd() != false {
		t.Fatal("Expected 50 to not be odd")
	}
	if MyInt(33).IsEven() != false {
		t.Fatal("Expected 33 to not be even")
	}
	if MyInt(33).IsOdd() != true {
		t.Fatal("Expected 33 to be odd")
	}

	if MyInt(1).IsPrime() != false {
		t.Fatal("Expected 1 to not be prime")
	}
	if MyInt(2).IsPrime() != true {
		t.Fatal("Expected 2 to be prime")
	}
	if MyInt(3).IsPrime() != true {
		t.Fatal("Expected 3 to be prime")
	}
	if MyInt(4).IsPrime() != false {
		t.Fatal("Expected 4 to not be prime")
	}
	if MyInt(5).IsPrime() != true {
		t.Fatal("Expected 5 to be prime")
	}
	if MyInt(9).IsPrime() != false {
		t.Fatal("Expected 9 to not be prime")
	}
	if MyInt(42).IsPrime() != false {
		t.Fatal("Expected 42 to not be prime")
	}
	if MyInt(43).IsPrime() != true {
		t.Fatal("Expected 43 to be prime")
	}
}

func TestMakeHeaderLine(t *testing.T) {
	// int
	if MakeHeaderLine(int(10), "XYZ") != "========== XYZ ==========" {
		t.Fatalf("Wrong header line for (int(10), 'XYZ')")
	}

	// some type whose underlying type is int
	if MakeHeaderLine(MyInt(2), "ABC") != "== ABC ==" {
		t.Fatalf("Wrong header line for (MyInt(2), 'ABC')")
	}
}

func testReader[T FileReader](reader T, t *testing.T) {
	res := reader.Read("no file")
	if len(res) != 1 {
		t.Fatalf("Expected stub read to return one byte")
	}
	if res[0] != 0x0 {
		t.Fatalf("Expected stub read to return null byte")
	}
}

func TestFileReader(t *testing.T) {
	testReader(CSVReader{}, t)
	testReader(JSONReader{}, t)
}

func TestFormatLogMessage(t *testing.T) {
	test := func(got string, exp string) {
		if got != exp {
			t.Fatalf("Got log message: '%s', but expected '%s'", got, exp)
		}
	}
	test(FormatLogMessage("my message", 42, 100), "my message - 0x2A - 0x64")
	test(FormatLogMessage("X", "foo", 3.14), "X - foo - 3.14")
	test(FormatLogMessage("X", 5-2i, []int{1, 2, 3}), "X - (5-2i) - [1 2 3]")
	test(FormatLogMessage("X", nil, map[int]int{5: 42}), "X - <nil> - map[5:42]")
}

func TestChimera(t *testing.T) {
	chimera := Chimera{Lion: Lion{weight: 100}, Goat: Goat{hornLength: 10}, Snake: Snake{}}
	if chimera.Roar() != "Lion of weight 100 roars" {
		t.Fatalf("Chimera fails")
	}
	if chimera.Headbutt() != "Goat with horn size 10 roars" {
		t.Fatalf("Chimera fails")
	}
	if chimera.Hiss() != "Snake hisses" {
		t.Fatalf("Chimera fails")
	}
}

func TestCalcPriceMotorcycle(t *testing.T) {
	m := Motorcycle{coolnessFactor: 3, Vehicle: Vehicle{numWheels: 2, weight: 500}}
	if m.CalcPrice() != 5000 {
		t.Fatalf("Expected motorcycle price to be 5000")
	}
}

func TestCalcPriceSedan(t *testing.T) {
	s := Sedan{engineHorsepower: 300, Vehicle: Vehicle{numWheels: 4, weight: 2000}}
	if s.CalcPrice() != 9000 {
		t.Fatalf("Expected sedan price to be 9000")
	}
}

func TestAdjustPrice(t *testing.T) {
	m := Motorcycle{coolnessFactor: 3, Vehicle: Vehicle{numWheels: 2, weight: 500}}
	s := Sedan{engineHorsepower: 300, Vehicle: Vehicle{numWheels: 4, weight: 2000}}

	if AdjustPrice(m, USA) != 12500 {
		t.Fatalf("Expected motorcycle in USA to be 12500")
	}
	if AdjustPrice(s, CANADA) != 9630 {
		t.Fatalf("Expected sedan in Canada to be 9630")
	}

}

func TestDoubleArr(t *testing.T) {
	arr := [3]int{1, 2, 3}
	DoubleArr(arr)
	if arr[0] != 1 {
		t.Fatalf("Array should not be modified by function when using call by value")
	}
	if arr[1] != 2 {
		t.Fatalf("Array should not be modified by function when using call by value")
	}
	if arr[2] != 3 {
		t.Fatalf("Array should not be modified by function when using call by value")
	}
}

func TestDoubleArrByPtr(t *testing.T) {
	arr := [3]int{1, 2, 3}
	DoubleArrByPtr(&arr)
	if arr[0] != 2 {
		t.Fatalf("Array should have been doubled")
	}
	if arr[1] != 4 {
		t.Fatalf("Array should have been doubled")
	}
	if arr[2] != 6 {
		t.Fatalf("Array should have been doubled")
	}
}

func TestDoubleSlice(t *testing.T) {
	s := []int{1, 2, 3}
	DoubleSlice(s)
	if sliceEq(s, []int{2, 4, 6}) != true {
		t.Fatalf("Slice should have been doubled")
	}
}

func TestCalcArea(t *testing.T) {
	p := Point2D{X: 5, Y: 10}
	if p.CalcArea() != 50 {
		t.Fatalf("Incorrect area, expected 50")
	}

	p = Point2D{X: -5, Y: 2}
	if p.CalcArea() != 10 {
		t.Fatalf("Incorrect area, expected 50")
	}
}

func TestTranspose(t *testing.T) {
	p := Point2D{X: 2, Y: 4}
	p.Transpose()
	if p.X != 4 || p.Y != 2 {
		t.Fatalf("Expected point to be transposed")
	}
}

func TestGotoLabelSumSquares(t *testing.T) {
	if GotoLabelSumSquares(10) != 14 {
		t.Fatalf("Expected sum of squares 1..10 to be 14")
	}
	if GotoLabelSumSquares(50) != 140 {
		t.Fatalf("Expected sum of squares 1..50 to be 140")
	}
	if GotoLabelSumSquares(-10) != 0 {
		fmt.Printf("%d\n", GotoLabelSumSquares(0))
		t.Fatalf("Expected sum of squares 1..-10 to be 0")
	}
}

func TestBreakLabel3D(t *testing.T) {
	{
		matrix := [][][]int{
			{
				{1, 2, 3},
				{4, -1, 9},
				{9, 9, 9},
			},
		}
		if BreakLabel3D(matrix) != 10 {
			t.Fatalf("Expected sum of matrix to be 10. Make sure the entire plane is skipped when a negative number is encountered.")
		}
	}
	{
		matrix := [][][]int{
			{
				{1, 2, 3},
				{4, -1, 9},
				{9, 9, 9},
			},
			{
				{5, 3, -2},
				{9, 9, 9},
				{9, 9, 9},
			},
			{
				{0, 0, 0},
				{4, 2, -7},
				{9, 9, 9},
			},
		}
		if BreakLabel3D(matrix) != 24 {
			t.Fatalf("Expected sum of matrix to be 24. Make sure the entire plane is skipped when a negative number is encountered.")
		}
	}
}

func TestFunkySumLinkedList(t *testing.T) {
	{
		sll := SLL{Val: 4, Next: &SLL{Val: 3, Next: &SLL{Val: 5, Next: &SLL{Val: 6, Next: nil}}}}
		if FunkySumLinkedList(sll) != 108 {
			t.Fatalf("Expected funky sum of linked list 4 -> 3 -> 5 -> 6 to be 108")
		}
	}
	{
		sll := SLL{Val: 4, Next: &SLL{Val: 3, Next: &SLL{Val: 5, Next: &SLL{Val: -7, Next: &SLL{Val: 6, Next: nil}}}}}
		if FunkySumLinkedList(sll) != 48 {
			t.Fatalf("Expected funky sum of linked list 4 -> 3 -> 5 -> -> -7 -> 6 to be 108")
		}
	}
}
