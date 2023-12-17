// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/962/problem/C
// https://codeforces.com/problemset/status/962/problem/C
func Test_cf962C(t *testing.T) {
	testCases := [][2]string{
		{
			`8314`,
			`2`,
		},
		{
			`625`,
			`0`,
		},
		{
			`333`,
			`-1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf962C)
}