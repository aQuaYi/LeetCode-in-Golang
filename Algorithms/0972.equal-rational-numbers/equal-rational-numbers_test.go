package problem0972

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	T   string
	ans bool
}{

	{
		"0.(52)",
		"0.5(25)",
		true,
	},

	{
		"0.1666(6)",
		"0.166(66)",
		true,
	},

	{
		"0.9(9)",
		"1.",
		true,
	},

	// 可以有多个 testcase
}

func Test_isRationalEqual(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, isRationalEqual(tc.S, tc.T), "输入:%v", tc)
	}
}

func Benchmark_isRationalEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isRationalEqual(tc.S, tc.T)
		}
	}
}
