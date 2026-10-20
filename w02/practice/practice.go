package practice

// Задачи по слайсам. Реализуй функции вместо заглушек (// TODO).
// Проверка: go test   (одна задача: go test -run TestSum)
// Правила: без пакета slices и без sort — только встроенные средства: len, cap, copy, append, нарезка.

// Задача 1. Sum
// Возвращает сумму элементов. Для пустого и nil-слайса — 0.
func Sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}

	return sum // TODO
}

// Задача 2. Last
// Возвращает последний элемент и true. Если слайс пуст — 0, false. Не должна паниковать.
func Last(s []int) (int, bool) {

	if len(s) == 0 {
		return 0, false
	}
	return s[len(s)-1], true // TODO
}

// Задача 3. Swap
// Меняет местами элементы s[i] и s[j] на месте.
func Swap(s []int, i, j int) {
	// 0 1 2 3 4
	// l5
	//5
	if ((len(s) - 1) < i) || ((len(s) - 1) < j) {
		return
	}
	s[i], s[j] = s[j], s[i]
}

// Задача 4. FirstN
// Первые n элементов без копирования (подслайс).
// Если n > len(s) — весь s.
// Если n ≤ 0 — пустой слайс.
// Изменение результата должно быть видно в s.
func FirstN(s []int, n int) []int {
	if n <= 0 {
		return s[0:0]
	} else if n > len(s) {
		return s[:len(s)]
	}

	return s[:n]
}

// Задача 5. Clone
// Независимая копия: изменение копии не должно менять s. Clone(nil) возвращает nil.
func Clone(s []int) []int {
	if s == nil {
		var answer []int
		return answer
	}
	answer := make([]int, len(s), cap(s))
	copy(answer, s)
	return answer
}

// Задача 6. Reverse
// Разворачивает слайс на месте (без нового слайса).
func Reverse(s []int) {
	// 0 1 2 3
	//l4
	temp := 0
	for i := 0; i < len(s)/2; i++ {
		temp = s[i]
		s[i] = s[len(s)-1-i]
		s[len(s)-1-i] = temp
	}
}

// Задача 7. Middle
// Всё, кроме первого и последнего элемента, без копирования.
// Для слайсов длиной ≤ 2 (и для nil) — пустой слайс, без паники.
// Пример: [1 2 3 4] → [2 3].

func Middle(s []int) []int {
	if len(s) <= 2 {
		var answer []int
		return answer
	}
	return s[1 : len(s)-1]
}

// Задача 8. Chunks
// Режет s на куски длиной size (последний может быть короче).
// Куски — подслайсы s (память общая), но у каждого cap == len,
// чтобы append к куску не затирал соседний. size ≤ 0 → nil.
// Пример: [1 2 3 4 5 6 7], 3 → [[1 2 3] [4 5 6] [7]].
func Chunks(s []int, size int) [][]int {
	if size <= 0 {
		return nil
	}

	var length int
	if len(s)%size == 0 {
		length = len(s) / size
	} else {
		length = len(s)/size + 1
	}

	answer := make([][]int, length)
	for i := 0; i < len(s); i += size { // шаг size
		end := i + size // длина куска size
		if end > len(s) {
			end = len(s) // последний кусок короче
		}
		answer[i/size] = s[i:end:end] // номер куска, а не позиция
	}
	return answer
}

// Задача 9. Filter
// Новый слайс из элементов,
// для которых keep вернула true.
// Исходный s не изменять.
func Filter(s []int, keep func(int) bool) []int {
	var answer []int
	for i, _ := range s {
		if keep(s[i]) {
			answer = append(answer, s[i])
		}
	}
	return answer
}

// Задача 10. Remove
// Удаляет элемент с индексом i, сохраняя порядок,
// без выделения нового массива.
// Возвращает укороченный слайс (cap прежний).
// Если i вне диапазона — вернуть s без изменений.
func Remove(s []int, i int) []int {
	// 0 1 2 3 4 l5 c5//3
	// 0 1 2 4 _ l4 c5

	if i > len(s)-1 || i < 0 {
		return s
	}
	s = append(s[:i], s[i+1:]...)

	return s
}

// Задача 11. Insert
// Вставляет v на позицию i (0 ≤ i ≤ len(s)),
// сдвигая остальное вправо. Возвращает новый слайс.
// Если i вне диапазона — вернуть s без изменений.
// Пример: Insert([1 2 4 5], 2, 3) → [1 2 3 4 5].
func Insert(s []int, i int, v int) []int {
	if i < 0 || i > len(s) {
		return s
	}
	tail := make([]int, len(s[i:]))
	copy(tail, s[i:])

	answer := append(s[:i], v)
	answer = append(answer, tail...)
	return answer
}

// Задача 12. Unique
// Во отсортированном слайсе убирает подряд идущие дубликаты на месте (без нового массива): результат — s[:k]. Пример: [1 1 2 3 3 3 4] → [1 2 3 4].
func Unique(s []int) []int {
	return s // TODO
}

// Задача 13. Rotate
// Циклический сдвиг влево на k на месте. k ≥ 0 может быть больше длины. Пустой слайс — без паники. Пример: [1 2 3 4 5], 2 → [3 4 5 1 2]. Подсказка: три разворота.
func Rotate(s []int, k int) {
	// TODO
}

// Задача 14. Merge
// Сливает два отсортированных слайса в новый отсортированный. Результат должен быть создан через make([]int, 0, len(a)+len(b)), то есть cap(результата) == len(a)+len(b) (проверяется тестом). Исходные слайсы не менять.
func Merge(a, b []int) []int {
	return nil // TODO
}

// Задача 15. Matrix
// Матрица rows × cols из нулей; строки должны быть независимы (изменение m[0][0] не меняет m[1][0]).
func Matrix(rows, cols int) [][]int {
	return nil // TODO
}

// Задача 16. Transpose
// Транспонирование прямоугольной матрицы в новую матрицу. Пустой вход → пустой результат. Пример: [[1 2 3] [4 5 6]] → [[1 4] [2 5] [3 6]].
func Transpose(m [][]int) [][]int {
	return nil // TODO
}

// Задача 17. Stack
// Стек целых на слайсе. Нулевое значение Stack должно быть готово к работе (var st Stack; st.Push(1)). Методы: Push(v int), Pop() (int, bool) (на пустом — 0, false), Peek() (int, bool) (посмотреть верх, не удаляя), Len() int.
type Stack struct {
	data []int
}

func (s *Stack) Push(v int) {
	// TODO
}

func (s *Stack) Pop() (int, bool) {
	return 0, false // TODO
}

func (s *Stack) Peek() (int, bool) {
	return 0, false // TODO
}

func (s *Stack) Len() int {
	return 0 // TODO
}
