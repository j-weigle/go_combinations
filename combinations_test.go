package combinations

import (
	"reflect"
	"testing"
)

func TestGetSubset(t *testing.T) {
	myRange := getRange(1, 3)
	result := GetSubset(&myRange, 3)
	want := []int{1, 2}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("error: \nreturn:\t%v\nwant:\t%v", result, []int{1, 2})
	}
}

func TestAll(t *testing.T) {
	testInArrs := [][]int{
		{1, 2, 3},
	}
	tests := []struct {
		name string
		in   []int
		out  [][]int
	}{
		{
			name: "3",
			in:   testInArrs[0],
			out: [][]int{
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := All(&test.in)
			if !reflect.DeepEqual(result, test.out) {
				t.Errorf("error: \nreturn:\t%v\nwant:\t%v", result, test.out)
			}
		})
	}
}

func TestAllQualifying(t *testing.T) {
	myRange := getRange(1, 3)
	result := AllQualifying(&myRange, func(v *[]int) bool {
		sum := 0
		for i := 0; i < len(*v); i++ {
			sum += (*v)[i]
		}
		return sum < 5
	})
	want := [][]int{{1}, {2}, {1, 2}, {3}, {1, 3}}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("error: \nreturn:\t%v\nwant:\t%v", result, want)
	}
}

func TestAllQualifyingPositions(t *testing.T) {
	myRange := getRange(1, 3)
	result := AllQualifyingPositions(&myRange, func(v *[]int) bool {
		sum := 0
		for i := 0; i < len(*v); i++ {
			sum += (*v)[i]
		}
		return sum < 5
	})
	if len(result) != 5 {
		t.Errorf("error: \nreturn:\t%v\nwant:\t%v", len(result), 5)
	}
}

func TestCombinations(t *testing.T) {
	myRange := getRange(1, 3)
	result := Combinations(&myRange, 2)
	want := [][]int{{1, 2}, {1, 3}, {2, 3}}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("error: \nreturn:\t%v\nwant:\t%v", result, want)
	}
}

func TestCombinationPositions(t *testing.T) {
	myRange := getRange(1, 20)
	result := CombinationsPositions(&myRange, 10)
	if len(result) != 184756 {
		t.Errorf("error: \nreturn:\t%v\nwant:\t%v", len(result), 185756)
	}
}

func TestCombinationsQualifyingPositions(t *testing.T) {
	myRange := getRange(1, 3)
	result := CombinationsQualifyingPositions(&myRange, 2, func(v *[]int) bool {
		sum := 0
		for i := 0; i < len(*v); i++ {
			sum += (*v)[i]
		}
		return sum < 5
	})
	if len(result) != 2 {
		t.Errorf("error: \nreturn:\t%v\nwant:\t%v", len(result), 2)
	}
}

func getRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
