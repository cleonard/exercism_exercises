package hamming

import "errors"

func Distance(a, b string) (int, error) {
	iterLength := len(a)
	if iterLength != len(b) {
		return 0, errors.New("Args of unequal lengths.")
	}
	count := 0
	for i := range iterLength {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
