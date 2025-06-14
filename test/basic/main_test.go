package basic

import (
	"testing"
)

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 2
	)

	actual := AddOne(input)
	if actual != output {
		t.Errorf("AddOne(%d), output::%d, actual::%d\n", input, output, actual)
	}
	// assert.Equal(t, AddOne(input), output, fmt.Sprintf("AddOne(%d), actual::%d\n", input, output))
	// // assert.NotEqual(t, AddOne(input), output, fmt.Sprintf("AddOne(%d), actual::%d\n", input, output))
	// assert.Nil(t, nil, nil)
}

// func TestRequire(t *testing.T) {
// 	require.Equal(t, 2, 3)
// 	fmt.Println("not executing")
// }

// func TestAssert(t *testing.T) {
// 	assert.Equal(t, 2, 3)
// 	fmt.Println("executing")
// }
