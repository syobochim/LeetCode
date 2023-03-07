package fizzBuzz

import "testing"
import "reflect"

func Test_fizzBuzz_3(t *testing.T) {
	var result []string = fizzBuzz(3)
	var expected []string = []string{"1", "2", "Fizz"}
	if !reflect.DeepEqual(result, expected) {
		t.Error("結果が違う, 期待する結果 ['1','2','Fizz'], 得られた結果 ", result)
	}
}

func Test_fizzBuzz_5(t *testing.T) {
	var result []string = fizzBuzz(5)
	var expected []string = []string{"1", "2", "Fizz", "4", "Buzz"}
	if !reflect.DeepEqual(result, expected) {
		t.Error("結果が違う, 期待する結果, ", expected, " 得られた結果 ", result)
	}
}

func Test_fizzBuzz_15(t *testing.T) {
	var result []string = fizzBuzz(15)
	var expected []string = []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	if !reflect.DeepEqual(result, expected) {
		t.Error("結果が違う, 期待する結果, ", expected, " 得られた結果 ", result)
	}
}
