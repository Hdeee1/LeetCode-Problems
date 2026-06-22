package main

func romanToInt(s string) int {
	roma := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var count = 0

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && roma[s[i]] < roma[s[i+1]] {
			count -= roma[s[i]]
		} else {
			count += roma[s[i]]
		}
	} 
	return count
}