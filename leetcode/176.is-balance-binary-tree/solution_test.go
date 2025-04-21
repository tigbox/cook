package solution

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SolutionFuncType func(root *Node) bool

var SolutionFuncList = []SolutionFuncType{
	IsBalance,
}

type Case struct {
	name   string
	input  *Node
	expect bool
}

var cases = []Case{
	{
		name: "balance binary tree",
		input: &Node{
			Value: 1,
			Left:  &Node{Value: 2},
			Right: &Node{Value: 3},
		},
		expect: true,
	},
	{
		name:   "nil binary tree",
		input:  nil,
		expect: true,
	},
	{
		name: "not balance binary tree",
		input: &Node{
			Value: 1,
			Left: &Node{
				Value: 2,
				Left:  &Node{Value: 3},
			},
		},
		expect: false,
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
					"func: %v case: %v", funcName, c.name)
			})
		}
	}
}
