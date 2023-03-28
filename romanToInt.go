package main

func RomanToInt(s string) int {
	//MDCCCXXXIV
	simpleRomans := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var res int
	for i := 0; i < len(s)-1; i++ {
		if simpleRomans[string(s[i])] < simpleRomans[string(s[i+1])] {
			res -= simpleRomans[string(s[i])]
		} else {
			res += simpleRomans[string(s[i])]
		}
	}
	return res + simpleRomans[string(s[len(s)-1])]
}
