package homework

// PairInput соответствует объекту nums/target из условия контеста.
type PairInput struct {
	Nums   []int
	Target int
}

// longestPrefix возвращает самый длинный общий префикс строк.
func longestPrefix(words []string) string {
	if (len(words) == 0 || len(words[0]) == 0) {
		return "";
	}
	prefix := []rune(words[0])
	var res int = len(prefix)
	for i := 1; i < len(words); i++ {
		current := []rune(words[i])
		var curPrefix int = 0
		for j := 0; j < len(prefix) && j < len(current); j++ {
			if (prefix[j] == current[j]) {
				curPrefix++
			} else {
				break
			}
		}
		res = min(res, curPrefix)
		if res == 0 {
			return ""
		}

	}
	return string(prefix[:res])
}

// pairIndices возвращает индексы двух разных элементов по возрастанию.
func pairIndices(input PairInput) [2]int {
	var collected = make(map[int]int)
	for i, num := range input.Nums {
		if ind, ok := collected[input.Target - num]; ok {
			return [2]int{ind, i}
		}
		collected[num] = i
	}
	return [2]int{}
}

// hasDuplicate сообщает, встречается ли какое-либо значение повторно.
func hasDuplicate(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}
	var collected = make(map[int]bool)
	for _, num := range nums {
		collected[num] = true
	}
	return len(collected) != len(nums)
}

// compactSorted создаёт независимый слайс уникальных значений.
func compactSorted(nums []int) []int {
	if nums == nil {
		return nil
	}
	result := make([]int, 0, len(nums))

	for _, e := range nums {
		if len(result) == 0 || result[len(result) - 1] != e {
			result = append(result, e)
		}
	}

	return result
}

// balancedBrackets проверяет пары и порядок закрытия скобок.
func balancedBrackets(text string) bool {
	var converter = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	for _, char := range text {
		switch char {
			case '(', '[', '{':
				stack = append(stack, byte(char))
			case ')', ']', '}':
				if len(stack) == 0 || stack[len(stack)-1] != converter[byte(char)] {
					return false
				}
				stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
