// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_b(t *testing.T) {
	t.Log("记得初始化所有全局变量")
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeClassWithFile(t, Constructor, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/biweekly-contest-95/problems/find-consecutive-integers-from-a-data-stream/
