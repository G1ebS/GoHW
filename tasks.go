package lesson2

// FizzBuzz возвращает FizzBuzz, Fizz, Buzz или пустую строку для обычного числа.
// Для n <= 0 возвращает invalid.
func FizzBuzz(n int) string {
	switch  {
		case n <= 0:
			return "invalid"
		case n % 15 == 0:
			return "FizzBuzz"
		case n % 3 == 0:
			return "Fizz"
		case n % 5 == 0:
			return "Buzz"
		default:
			return ""
	}
}

// Sum возвращает сумму. Для пустого среза и nil возвращает 0.
func Sum(values []int64) int64 {
	// TODO: Пройти по значениям циклом.
	if (values == nil || len(values) == 0) {return 0}
	var sum int64 = 0
	for _, value := range values {
		sum += value
	}
	return sum
}

// Min возвращает mint: минимум массива. Для пустого среза и nil возвращает 0 по условию задачи.
func Min(values []int64) (mint int64) {
	if (values == nil || len(values) == 0) {return 0}
	mint = values[0]
	for _, values := range values {
		mint = min(mint, values)
	}
	return
}

// Max возвращает maxt: максимум массива. Для пустого среза и nil возвращает 0 по условию задачи.
func Max(values []int64) (maxt int64) {
	if (values == nil || len(values) == 0) {return 0}
	maxt = values[0]
	for _, values := range values {
		maxt = max(maxt, values)
	}
	return

}

// Average возвращает среднее как float64. Для пустого среза и nil возвращает 0.
func Average(values []int64) float64 {
	// TODO: Выполнить преобразования до деления.
	if (values == nil || len(values) == 0) {return 0}
	var sum float64 = 0
	for _, value := range values {
		sum += float64(value)
	}
	return sum / float64(len(values))
}

// Category возвращает invalid, small, regular или large по сумме в копейках.
func Category(amount int64) string {
	switch {
		case amount < 0:
			return "invalid"
		case amount < 10000:
			return "small"
		case amount < 50000:
			return "regular"
		default:
			return "large"
	}
}

// RuneCount считает кодовые точки Unicode в строке через range.
func RuneCount(text string) (res int) {
	res = 0
	for range text {
		res += 1
	}
	return res
}
