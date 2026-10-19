package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	iterLength := len(a)
	if iterLength != len(b) {
		return 0, fmt.Errorf("Args of unequal lengths.")
	}
	count := 0
	for i := 0; i < iterLength; i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
