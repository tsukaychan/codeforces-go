// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_d(t *testing.T) {
	targetCaseNum := -1
	if err := testutil.RunLeetCodeFuncWithFile(t, longestCycle, "d.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-304/problems/longest-cycle-in-a-graph/