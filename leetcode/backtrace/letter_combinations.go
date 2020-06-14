package backtrace

/**
 * digits是数字字符串
 * s(digits)代表digits所能代表的字母字符串
 * s(digits[0...n-1] = letter(digits[0])+s(digits[1...n-1])
 * 					 = letter(digits[0])+letter(digits[1])+s(digits[2...n-1)
 * 					 = ...
 */

type LetterMap struct {
	res       []string
	letterMap []string
}

func NewLetterMap() *LetterMap {
	return &LetterMap{
		res: make([]string, 0),
		letterMap: []string{
			" ",
			"",
			"abc",
			"def",
			"ghi",
			"jkl",
			"mno",
			"pqrs",
			"tuv",
			"wxyz",
		},
	}
}

func (lm *LetterMap) findCombination(digits, s string, index int) {
	if index == len(digits) {
		lm.res = append(lm.res, s)
		return
	}
	c := digits[index]
	if c < '0' || c > '9' || c == '1' {
		panic("invalid argument")
	}
	letters := lm.letterMap[c-'0']
	for i := 0; i < len(letters); i++ {
		lm.findCombination(digits, s+string(letters[i]), index+1)
	}
	return
}

//leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	lm := NewLetterMap()
	if digits == "" {
		return lm.res
	}
	lm.findCombination(digits, "", 0)
	return lm.res
}
