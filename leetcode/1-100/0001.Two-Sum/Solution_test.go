package Solution

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SolutionFuncType func([]int, int) []int

var SolutionFuncList = []SolutionFuncType{
	TwoSum,
	TwoSumMap,
	TwoSumMap1,
}

type Case struct {
	name   string
	nums   []int
	target int
	expect []int
}

var cases = []Case{
	{"test case1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
	{"test case2", []int{3, 2, 4}, 6, []int{1, 2}},
	{"test case3", []int{3, 3}, 6, []int{0, 1}},
}

func TestSolution(t *testing.T) {
	ast := assert.New(t)
	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.nums, c.target)
				ast.Equal(c.expect, got, "func: %v,case:%v", funcName, c.name)
			})
		}
	}
}

// 压力测试
func BenchmarkSolution(b *testing.B) {
	ast := assert.New(b)
	for _, f := range SolutionFuncList {
		for i := 0; i < b.N; i++ {
			funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
			for _, c := range cases {
				b.ResetTimer()
				b.Run(funcName, func(b *testing.B) {
					got := f(c.nums, c.target)
					ast.Equal(c.expect, got,
						"func: %v case: %v ", funcName, c.name)
				})
			}
		}
	}
}
