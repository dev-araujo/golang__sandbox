package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expectation := 4

	if sum != expectation {
		t.Errorf("expectation '%d', received '%d'", expectation, sum)
	}
}
func ExampleAdiciona() {
	soma := Add(1, 5)
	fmt.Println(soma)
	// Output: 6
}
