package feibo

import "errors"

func FeiBo(n int) ([]int, error) {
	if n < 2 || n > 100 {
		return nil, errors.New("n should be in [2, 100]")
	}

	feiBoList := []int{1, 1}
	for i := 2; i < n; i++ {
		feiBoList = append(feiBoList, feiBoList[i-2]+feiBoList[i-1])
	}
	return feiBoList, nil
}
