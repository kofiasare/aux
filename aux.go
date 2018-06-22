// Package aux is library that provides useful
// functional programming helpers
package aux

import (
	"fmt"
	"reflect"
	"time"
)

// Loop iterates `i` times and calls block `b`on
// each iteration. `i=0` implies an infinite loop
func Loop(i int, b interface{}) {
	if !isFunctionWithSingleArg(b) {
		panic("Second parameter must be a function with one argument of type int")
	}

	// loop i times
	if i != 0 {
		for n := 0; n < i; n++ {
			argVal(b).Call([]reflect.Value{argVal(n)})
		}
		return
	}

	// infinite loop
	c := 0
	for {
		argVal(b).Call([]reflect.Value{argVal(c)})
		c++
	}
}

// Factorial returns the factorial of `n`
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// Fibonacci returns the `int[]` fibonacci series
// of `n`
func Fibonacci(n int) (r []int) {
	if n > 0 {
		a, b := 0, 1
		Loop(n, func(i int) {
			r = append(r, b)
			a, b = b, (a + b)
		})
		return
	}
	return
}

// Next returns the successor of i
func Next(i interface{}) interface{} {

	if isInteger(i) {
		return argVal(i).Interface().(int) + 1
	}

	if isString(i) {
		c := []byte(i.(string))
		l := c[len(c)-1]

		if string(l) != "z" {
			return string(l + 1)
		}

		return "a"
	}

	panic("Parameter must be an integer or a string")
}

// Prime returns true if `n` is a prime number otherwise
// false
func Prime(n int) bool {

	const start = 2
	if n < start {
		return false
	}

	// factors of n
	for _, i := range Range(start, n) {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

// Range return a list containing an arithmetic progression of integers.
// `Range(stop)`
// `Range(start, stop)`
// `Range(start, stop, step)`
func Range(args ...int) (n []int) {
	switch i := len(args); {
	case i < 1:
		panic("Range expected at least 1 arguments, got 0")
	case i > 3:
		panic(fmt.Sprintf("Range expected at most 3 arguments, got %d", i))
	case i == 1:
		Loop(args[0], func(i int) {
			n = append(n, i)
		})
		return
	case i == 2:
		for i := args[0]; i < args[1]; i++ {
			n = append(n, i)
		}
		return
	case i == 3:
		step := args[2]
		if step == 0 {
			panic("Range(start, stop, step) step argument must not be zero")
		}

		if step < 0 {
			for i := args[0]; i > args[1]; i += step {
				n = append(n, i)
			}
		} else {
			for i := args[0]; i < args[1]; i += step {
				n = append(n, i)
			}
		}
		return

	default:
		return
	}
}

// CountDown counts down from start `s` observing delay `d` and
// calls block `b` for each iteration
func CountDown(s int, d time.Duration, b interface{}) {

	if !isFunctionWithSingleArg(b) {
		panic("Third parameter must be a function with one argument of type int")
	}

	Loop(s, func(i int) {
		argVal(b).Call([]reflect.Value{argVal(s - i)})
		time.Sleep(d)
	})

}

// Every calls block `b` at every interval `i`in its
// own goroutine simillar to Javascript's setInterval()
func Every(i time.Duration, b interface{}) {

	if !isFunction(b) {
		panic(" Second parameter must be a function")
	}

	// loop forever but observe delay
	go Loop(0, func(int) {
		time.Sleep(i)
		argVal(b).Call(nil)
	})
}
