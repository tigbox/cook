package main

import (
	"cook/cmd/leetcode"
	"fmt"
)

func main() {
	problems := leetcode.GetSortedProblemsInstance()
	for _, v := range problems {
		fmt.Println(v)
	}
	//problems := leetcode.GetProblemsJson()
	//fmt.Println(problems)
	//for _, v := range problems {
	//	fmt.Println(v.Stat.QuestionID, v.Stat.QuestionTitle, v.PathName,v.Stat.QuestionTitleSlup)
	//err := leetcode.UrlPath("https://leetcode.com/problems/" + v.Stat.QuestionTitleSlup )
	//if err != nil {
	//	fmt.Println(v.Stat.QuestionID, v.Stat.QuestionTitle, v.PathName,v.Stat.QuestionTitleSlup)
	//	fmt.Println(err.Error())
	//}
	//}

	//	生成Problem 目录
	// 执行路径是在项目根目录中 go run cmd/main.go
	start := 1
	for i := 0; i < 10; i++ {
		leetcode.MakeDirFromTo(problems, start, start+99)
		start += 100
	}
	// leetcode.MakeDirFromTo(problems, 1401, 1500)
	// leetcode.MakeDirFromTo(problems, 1501, 1600)
	//leetcode.MakeDirFromTo(problems, 1601, 1700)
	//leetcode.MakeDirFromTo(problems, 1701, 1800)
	//leetcode.MakeDirFromTo(problems, 1901, 2000)
	//leetcode.MakeDirFromTo(problems, 2001, 2100)
	//leetcode.MakeDirFromTo(problems, 2101, 2200)
	//leetcode.MakeDirFromTo(problems, 2201, 2300)
	//leetcode.MakeDirFromTo(problems, 2301, 2400)
	//leetcode.MakeDirFromTo(problems, 2401, 2500)
	//leetcode.MakeDirFromTo(problems, 2901, 3000)

	// leetcode.MakeDir(problems)

	// leetcode.GetReadmeTemplateBuffer()

	//	GitBook
	// leetcode.MakeGitbookSummary(problems)

	//	sitemap
	// s := sitemap.New(problems)
	// fmt.Println(s)
}
