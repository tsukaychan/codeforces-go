// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/45/I
// https://codeforces.com/problemset/status/45/problem/I
func Test_cf45I(t *testing.T) {
	testCases := [][2]string{
		{
			`5
1 2 -3 3 3`,
			`3 1 2 3`,
		},
		{
			`13
100 100 100 100 100 100 100 100 100 100 100 100 100`,
			`100 100 100 100 100 100 100 100 100 100 100 100 100`,
		},
		{
			`4
-2 -2 -2 -2`,
			`-2 -2 -2 -2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf45I)
}
