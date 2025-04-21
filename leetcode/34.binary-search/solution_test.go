package solution

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SolutionFuncType func(arr []int, find int) int

// var SolutionFuncList = []SolutionFuncType{
// 	FindRight,
// }

type Input struct {
	Arr  []int
	Find int
}

type Case struct {
	funcType SolutionFuncType
	name     string
	input    Input
	expect   int
}

var cases = []Case{
	{
		funcType: FindLeft,
		name:     "normal case: has elements in input array",
		input:    Input{Arr: []int{1, 2, 3, 4, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 66, 7, 8, 9}, Find: 6},
		expect:   5,
	},
	{
		funcType: FindRight,
		name:     "normal case: has elements in input array",
		input:    Input{Arr: []int{1, 2, 3, 4, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 66, 7, 8, 9}, Find: 6},
		expect:   13,
	},
}

func TestFindLeft(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(c.funcType).Pointer()).Name(), ".")[1]
		t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
			got := c.funcType(c.input.Arr, c.input.Find)
			fmt.Println(got)
			ast.Equal(c.expect, got,
				"func: %v case:%v ", funcName, c.name)
		})
	}
}
