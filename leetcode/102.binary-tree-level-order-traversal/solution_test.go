package solution

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SolutionFuncType func(root *TreeNode) [][]int

var SolutionFuncList = []SolutionFuncType{
	LevelOrder,
}

type Case struct {
	name   string
	input  *TreeNode
	expect [][]int
}

var cases = []Case{
	{
		name:   "root is nil case",
		input:  nil,
		expect: [][]int{},
	},
	{
		name: "normal case", // 正常的case
		input: &TreeNode{
			Value: 1,
			Left: &TreeNode{
				Value: 2,
			},
			Right: &TreeNode{
				Value: 3,
				Left: &TreeNode{
					Value: 4,
				},
			},
		},
		expect: [][]int{{1}, {2, 3}, {4}},
	},
	{
		name: "single side case", // 单侧斜出去的境况
		input: &TreeNode{
			Value: 1,
			Right: &TreeNode{
				Value: 2,
				Right: &TreeNode{
					Value: 3,
					Right: &TreeNode{
						Value: 4,
					},
				},
			},
		},
		expect: [][]int{{1}, {2}, {3}, {4}},
	},
}

func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.input)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
