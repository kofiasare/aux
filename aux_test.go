package aux

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestLoop(t *testing.T) {

	// task definition
	task := func(c int) {
		fmt.Println("Hello, Task running")
	}

	Loop(5, task)
}

func TestFactorial(t *testing.T) {
	testData := []struct {
		testName string
		n, r     int
	}{
		{"test 5!", 5, 120},
		{"test 4!", 4, 24},
		{"test 3!", 3, 6},
		{"test 2!", 2, 2},
		{"test 1!", 1, 1},
		{"test 0!", 0, 1},
	}

	for _, d := range testData {
		t.Run(d.testName, func(t *testing.T) {
			r := Factorial(d.n)
			if r != d.r {
				t.Errorf("%d! should equal %d, but got %d", d.n, d.r, r)
			}
		})

	}
}

func TestFibonacci(t *testing.T) {
	testData := []struct {
		testName string
		n        int
		r        []int
	}{
		{"test fib(5)", 5, []int{1, 1, 2, 3, 5}},
	}

	for _, d := range testData {
		t.Run(d.testName, func(t *testing.T) {
			if r := Fibonacci(d.n); !reflect.DeepEqual(r, d.r) {
				t.Errorf("Fibonacci(%d) should equal %v, but got %v", d.n, d.r, r)
			}
		})
	}
}

func TestNext(t *testing.T) {
	testData := []struct {
		testName string
		n        interface{}
		s        interface{}
	}{
		{"test successor of 1", 1, 2},
		{"test successor of 2", 2, 3},
		//{"test successor of a", "a", "b"},
	}

	for _, d := range testData {
		t.Run(d.testName, func(t *testing.T) {
			if r := Next(d.n); !reflect.DeepEqual(r, d.s) {
				t.Errorf("Expect successor of (%v) to be %v, but got %v", d.n, d.s, r)
			}
		})
	}
}

func TestCountDown(t *testing.T) {
	type args struct {
		s int
		d time.Duration
		b interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CountDown(tt.args.s, tt.args.d, tt.args.b)
		})
	}
}

func TestEvery(t *testing.T) {
	type args struct {
		i time.Duration
		b interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Every(tt.args.i, tt.args.b)
		})
	}
}

func TestPrime(t *testing.T) {
	testData := []struct {
		testName string
		number   int
		isPrime  bool
	}{
		{"Test 1 is Prime", 1, false},
		{"Test 2 is Prime", 2, true},
		{"Test 3 is Prime", 3, true},
		{"Test 4 is Prime", 4, false},
		{"Test 0 is Prime", 0, false},
	}

	for _, d := range testData {
		t.Run(d.testName, func(t *testing.T) {
			if r := Prime(d.number); r != d.isPrime {
				t.Errorf("Expect  IsPrime(%v)  to be  %v, but got %v", d.number, d.isPrime, r)
			}
		})
	}
}

func TestRange(t *testing.T) {
	tests := []struct {
		name  string
		args  int
		wantN []int
	}{
		{"Test Range(20)", 20, []int{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			10, 11, 12, 13, 14, 15, 16, 17,
			18, 19,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := Range(tt.args); !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("Range(%v) = %v, want %v", tt.args, gotN, tt.wantN)
			}
		})
	}
}
