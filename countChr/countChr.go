package countChr

func CntChr() map[string]int {
	cnt_str := "abababcbrr"
	m := make(map[string]int)

	for _, chr := range cnt_str {
		m[string(chr)]++
	}
	return m
}
