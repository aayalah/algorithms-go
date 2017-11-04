package numbers

import (
	"strconv"
	"strings"
)

type ComplexNumber struct {
	A int
	B int
}

func (c *ComplexNumber) NewComplexNumber(s string) {
	indOfPlusSign := strings.Index(s, "+")
	a, _ := strconv.ParseInt(s[:indOfPlusSign], 10, 64)
	indOfI := strings.Index(s, "i")
	b, _ := strconv.ParseInt(s[indOfPlusSign+1:indOfI], 10, 64)
	c.A = int(a)
	c.B = int(b)
}

func (c *ComplexNumber) Multiply(p *ComplexNumber) string {
	a := c.A*p.A + c.B*p.B*-1
	b := c.A*p.B + c.B*p.A

	astr := strconv.Itoa(a)
	bstr := strconv.Itoa(b)
	res := astr + "+" + bstr + "i"
	return res
}

func complexNumberMultiply(a string, b string) string {
	c1 := new(ComplexNumber)
	c2 := new(ComplexNumber)

	c1.NewComplexNumber(a)
	c2.NewComplexNumber(b)
	return c1.Multiply(c2)
}

func countBits(num int) []int {
	res := make([]int, num+1)
	if num == 0 {
		return []int{0}
	}

	if num == 1 {
		return []int{0, 1}
	}

	res[0] = 0
	res[1] = 1
	res[2] = 1

	lastPowerOfTwo := 2

	for i := 3; i <= num; i += 1 {
		diff := i - lastPowerOfTwo
		if diff == lastPowerOfTwo {
			res[i] = 1
			lastPowerOfTwo = i
		} else {
			res[i] = res[lastPowerOfTwo] + res[diff]
		}
	}
	return res
}
