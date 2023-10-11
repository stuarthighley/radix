package radixsort

func FirstNonRepeating(in string) rune {

	skipping := false
	for i, v := range in {
		if i >= len(in)-1 {
			break
		}
		if in[i+1] != in[i] {
			if !skipping {
				return v
			}
			skipping = false
		} else {
			skipping = !skipping
		}
	}

	return 0
}
