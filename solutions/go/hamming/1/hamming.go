package hamming

import "errors"

func Distance(a, b string) (int, error) {
    dist := 0

    if len(a) != len(b) {
        return dist, errors.New("Not Allowed")
    }
    
	for i := range a {
        if a[i] != b[i] {
            dist++
        }
    }

    return dist, nil
}
