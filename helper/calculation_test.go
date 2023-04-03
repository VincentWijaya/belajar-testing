package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestFailSum(t *testing.T) {
// 	result := Sum(10, 10)

// 	// if result != 10 {
// 	// 	t.Fatal("Result is not 10")
// 	// 	//stop
// 	// }

// 	require.Equal(t, 10, result, "Result shoud be 20")

// 	fmt.Println("TestFailSum akan berhenti")
// }

// func TestSum(t *testing.T) {
// 	result := Sum(10, 10)

// 	require.Equal(t, 20, result, "Result shoud be 20")
// }

func TestSum(t *testing.T) {
	tests := []struct {
		request  int
		expected int
		errorMsg string
	}{
		{
			request:  Sum(1, 1),
			expected: 2,
			errorMsg: "Result should be 2",
		},
		{
			request:  Sum(1, 2),
			expected: 3,
			errorMsg: "Result should be 3",
		},
		{
			request:  Sum(10, 2),
			expected: 12,
			errorMsg: "Result should be 12",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require.Equal(t, test.expected, test.request, test.errorMsg)
		})
	}
}
