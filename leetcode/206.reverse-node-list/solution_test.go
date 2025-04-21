package solution

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SolutionFuncType func(root *Node) *Node

var SolutionFuncList = []SolutionFuncType{
	ReverseNodeList,
}

type Case struct {
	name   string
	input  *Node
	expect *Node
}

var cases = []Case{
	{
		name: "normal case",
		input: &Node{
			Value: 1,
			Next: &Node{
				Value: 2,
				Next: &Node{
					Value: 3,
				},
			},
		},
		expect: &Node{
			Value: 3,
			Next: &Node{
				Value: 2,
				Next: &Node{
					Value: 1,
					Next:  nil,
				},
			},
		},
	},
}

// 虽然测试用例中的数据值一样，也是可以用ast.Equal来判定的
// 毕竟ast.Equal对于指针类型的变量比较的不是内存地址，而是基于所引用的值，在这里就是比较Value和Next
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.input)
				fmt.Println(got)
				fmt.Println(&got)
				fmt.Println(c.expect)
				fmt.Println(&c.expect)
				ast.Equal(c.expect, got,
					"func: %v case: %v", funcName, c.name)
			})
		}
	}
}
