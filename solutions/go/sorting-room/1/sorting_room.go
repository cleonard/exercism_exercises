package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %0.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	number := nb.Number()
	return fmt.Sprintf("This is a box containing the number %0.1f", float64(number))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fn, ok := fnb.(FancyNumber)
	if ok {
		var i, _ = strconv.Atoi(fn.Value())
		return i
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	num := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %0.1f", float64(num))

}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	switch i.(type) {
	case int:
		num, _ := i.(int)
		return DescribeNumber(float64(num))
	case float64:
		f, _ := i.(float64)
		return DescribeNumber(f)
	case NumberBox:
		val, _ := i.(NumberBox)
		return DescribeNumberBox(val)
	case FancyNumberBox:
		val, _ := i.(FancyNumberBox)
		return DescribeFancyNumberBox(val)
	default:
		return "Return to sender"
	}
}
