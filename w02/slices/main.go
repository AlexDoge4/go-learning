package main

import (
	"fmt"
	"os"
	"strconv"
	"unsafe"
)

// Запуск: go run . НОМЕР
// Номера совпадают с номерами упражнений в файле «Слайсы_теория_и_упражнения.md».
// Сначала запиши прогноз, потом запускай.

func main() {
	if len(os.Args) != 2 {
		fmt.Println("использование: go run . НОМЕР")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("номер должен быть числом")
		return
	}
	f, ok := exercises[n]
	if !ok {
		fmt.Println("нет упражнения с таким номером")
		return
	}
	f()
}

// unsafe нужен только упражнению про размер заголовка слайса.
var _ = unsafe.Sizeof(0)

var exercises = map[int]func(){
	1: ex01,
	2: ex02,
	3: ex03,
	4: ex04,
	5: ex05,
	6: ex06,
	7: ex07,
	8: ex08,
	9: ex09,
	10: ex10,
	11: ex11,
	12: ex12,
	13: ex13,
	14: ex14,
	15: ex15,
	16: ex16,
	17: ex17,
	18: ex18,
	19: ex19,
	20: ex20,
	21: ex21,
	22: ex22,
	23: ex23,
	24: ex24,
	25: ex25,
	26: ex26,
	27: ex27,
	28: ex28,
	29: ex29,
	30: ex30,
	31: ex31,
	32: ex32,
	33: ex33,
	34: ex34,
	35: ex35,
	36: ex36,
	37: ex37,
	38: ex38,
	39: ex39,
	40: ex40,
	41: ex41,
	42: ex42,
	43: ex43,
	44: ex44,
	45: ex45,
	46: ex46,
	47: ex47,
	48: ex48,
	49: ex49,
	50: ex50,
	51: ex51,
	52: ex52,
	53: ex53,
	54: ex54,
	55: ex55,
	56: ex56,
	57: ex57,
	58: ex58,
	59: ex59,
	60: ex60,
	61: ex61,
	62: ex62,
	63: ex63,
	64: ex64,
	65: ex65,
	66: ex66,
	67: ex67,
	68: ex68,
	69: ex69,
	70: ex70,
	71: ex71,
	72: ex72,
	73: ex73,
	74: ex74,
	75: ex75,
	76: ex76,
	77: ex77,
	78: ex78,
	79: ex79,
	80: ex80,
	81: ex81,
	82: ex82,
	83: ex83,
	84: ex84,
	85: ex85,
	86: ex86,
	87: ex87,
	88: ex88,
	89: ex89,
	90: ex90,
	91: ex91,
	92: ex92,
	93: ex93,
	94: ex94,
}

// №1. Литерал
func ex01() {
	s := []int{10, 20, 30}
	fmt.Println(len(s), cap(s))
}

// №2. make с одним числом
func ex02() {
	s := make([]int, 4)
	fmt.Println(len(s), cap(s), s)
}

// №3. make с len и cap
func ex03() {
	s := make([]int, 2, 10)
	fmt.Println(len(s), cap(s), s)
}

// №4. make для строк
func ex04() {
	s := make([]string, 3)
	fmt.Println(len(s), s == nil)
	fmt.Printf("%q\n", s)
}

// №5. Объявили, но не создали
func ex05() {
	var s []int
	fmt.Println(len(s), cap(s), s == nil)
	fmt.Printf("%#v\n", s)
}

// №6. Пустой литерал
func ex06() {
	s := []int{}
	fmt.Println(len(s), cap(s), s == nil)
	fmt.Printf("%#v\n", s)
}

// №7. Пустой, но с запасом
func ex07() {
	s := make([]int, 0, 5)
	fmt.Println(len(s), cap(s), s == nil)
}

// №8. Запись по индексу
func ex08() {
	s := []int{1, 2, 3}
	s[1] = 99
	fmt.Println(s)
}

// №9. Индекс против cap
func ex09() {
	s := make([]int, 2, 5)
	i := 3
	fmt.Println(s[i])
}

// №10. Массив — это значение
func ex10() {
	a := [3]int{1, 2, 3}
	b := a
	b[0] = 100
	fmt.Println(a, b)
}

// №11. cap у массива
func ex11() {
	a := [...]int{5, 6, 7, 8}
	fmt.Println(len(a), cap(a))
}

// №12. Слайс — это не копия
func ex12() {
	s := []int{1, 2, 3}
	t := s
	t[0] = 100
	fmt.Println(s, t)
}

// №13. Первый срез
func ex13() {
	s := []int{0, 1, 2, 3, 4, 5}
	t := s[2:4]
	fmt.Println(t, len(t), cap(t))
}

// №14. Умолчания
func ex14() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s[:3], s[3:], s[:])
}

// №15. Расширяем окно вправо
func ex15() {
	s := []int{0, 1, 2, 3, 4, 5}
	t := s[1:3]
	fmt.Println(len(t), cap(t))
	t = t[:4]
	fmt.Println(t, len(t), cap(t))
}

// №16. Срез до конца
func ex16() {
	s := []int{0, 1, 2, 3, 4, 5}
	t := s[4:]
	fmt.Println(t, len(t), cap(t))
}

// №17. Пустой срез внутри
func ex17() {
	s := []int{0, 1, 2, 3, 4, 5}
	t := s[3:3]
	fmt.Println(t, len(t), cap(t), t == nil)
}

// №18. Срез среза
func ex18() {
	s := []int{0, 1, 2, 3, 4, 5}
	t := s[1:4]
	u := t[1:2]
	fmt.Println(u, len(u), cap(u))
}

// №19. s[1:] — это до len, а не до cap
func ex19() {
	s := make([]int, 3, 8)
	t := s[1:]
	fmt.Println(len(t), cap(t))
}

// №20. Верхняя граница больше len
func ex20() {
	s := make([]int, 3, 8)
	t := s[:6]
	fmt.Println(len(t), cap(t), t)
}

// №21. Нижняя и верхняя границы вместе
func ex21() {
	s := make([]int, 3, 8)
	t := s[2:7]
	fmt.Println(len(t), cap(t))
}

// №22. Выход за cap
func ex22() {
	s := make([]int, 3, 8)
	i := 9
	t := s[:i]
	fmt.Println(t)
}

// №23. low > high
func ex23() {
	s := []int{0, 1, 2, 3, 4, 5}
	i := 4
	t := s[i:2]
	fmt.Println(t)
}

// №24. Цепочка срезов
func ex24() {
	s := []int{0, 1, 2, 3, 4, 5}
	t := s[1:5]
	t = t[2:]
	t = t[:1]
	fmt.Println(t, len(t), cap(t))
}

// №25. Срез массива
func ex25() {
	a := [5]int{10, 20, 30, 40, 50}
	s := a[1:3]
	fmt.Println(s, len(s), cap(s))
}

// №26. Запись через подслайс
func ex26() {
	s := []int{1, 2, 3, 4, 5}
	t := s[1:4]
	t[0] = 99
	fmt.Println(s)
	fmt.Println(t)
}

// №27. Два окна с пересечением
func ex27() {
	s := []int{1, 2, 3, 4, 5}
	a := s[:3]
	b := s[2:]
	a[2] = 100
	fmt.Println(a, b)
}

// №28. Запись в «запасную» часть
func ex28() {
	s := []int{1, 2, 3, 4, 5}
	t := s[1:3]
	t = t[:4]
	t[3] = 0
	fmt.Println(s, t)
}

// №29. copy: приёмник больше
func ex29() {
	src := []int{1, 2, 3}
	dst := make([]int, 5)
	n := copy(dst, src)
	fmt.Println(n, dst)
}

// №30. copy: приёмник меньше
func ex30() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 2)
	n := copy(dst, src)
	fmt.Println(n, dst)
}

// №31. copy делает независимую копию
func ex31() {
	a := []int{1, 2, 3}
	b := make([]int, len(a))
	copy(b, a)
	b[0] = 100
	fmt.Println(a, b)
}

// №32. Классическая ошибка с copy
func ex32() {
	a := []int{1, 2, 3}
	var b []int
	copy(b, a)
	fmt.Println(b, len(b), b == nil)
}

// №33. copy внутри одного слайса (сдвиг влево)
func ex33() {
	s := []int{1, 2, 3, 4, 5}
	copy(s, s[2:])
	fmt.Println(s)
}

// №34. copy внутри одного слайса (сдвиг вправо)
func ex34() {
	s := []int{1, 2, 3, 4, 5}
	copy(s[1:], s)
	fmt.Println(s)
}

// №35. Функция меняет элемент
func setFirst(s []int) {
    s[0] = 0
}

func ex35() {
	s := []int{7, 8, 9}
	setFirst(s)
	fmt.Println(s)
}

// №36. Функция меняет заголовок
func reslice(s []int) {
    s = s[1:]
    s[0] = 100
}

func ex36() {
	s := []int{1, 2, 3}
	reslice(s)
	fmt.Println(s, len(s))
}

// №37. Функция подменяет слайс
func replace(s []int) {
    s = []int{9, 9, 9}
    s[0] = 0
}

func ex37() {
	s := []int{1, 2, 3}
	replace(s)
	fmt.Println(s)
}

// №38. Функция получает подслайс
func fill(s []int) {
    for i := range s {
        s[i] = 7
    }
}

func ex38() {
	s := make([]int, 6)
	fill(s[2:4])
	fmt.Println(s)
}

// №39. range: v — копия
func ex39() {
	s := []int{1, 2, 3}
	for _, v := range s {
	    v *= 10
	    fmt.Println(v)
	}
	fmt.Println(s)
}

// №40. range по индексу
func ex40() {
	s := []int{1, 2, 3}
	for i := range s {
	    s[i] *= 10
	}
	fmt.Println(s)
}

// №41. Матрица: одна строка на всех
func ex41() {
	row := make([]int, 3)
	m := [][]int{row, row}
	m[0][0] = 5
	fmt.Println(m)
}

// №42. Матрица: правильно
func ex42() {
	m := make([][]int, 2)
	for i := range m {
	    m[i] = make([]int, 2)
	}
	m[0][1] = 8
	fmt.Println(m, len(m), len(m[0]))
}

// №43. Маленькое окно на большой массив
func ex43() {
	big := make([]int, 1000000)
	small := big[:2]
	fmt.Println(len(small), cap(small))
}

// №44. Три индекса
func ex44() {
	s := []int{0, 1, 2, 3, 4, 5}
	t := s[1:3:4]
	fmt.Println(t, len(t), cap(t))
}

// №45. Три индекса: cap == len
func ex45() {
	s := []int{0, 1, 2, 3, 4, 5}
	t := s[1:3:3]
	fmt.Println(len(t), cap(t))
	t = t[:cap(t)]
	fmt.Println(t)
}

// №46. Три индекса: ограничитель
func ex46() {
	s := []int{0, 1, 2, 3, 4, 5}
	t := s[1:3:4]
	n := 4
	u := t[:n]
	fmt.Println(u)
}

// №47. Третий индекс не копирует
func ex47() {
	s := []int{1, 2, 3, 4, 5}
	a := s[:2:2]
	b := a[:2]
	b[0] = 100
	fmt.Println(s)
}

// №48. Нарезка nil-слайса
func ex48() {
	var s []int
	fmt.Println(len(s), s[:] == nil, s[0:0] == nil)
}

// №49. Индекс в nil-слайсе
func ex49() {
	var s []int
	i := 0
	fmt.Println(s[i])
}

// №50. Есть cap, но нет len
func ex50() {
	s := make([]int, 0, 3)
	i := 0
	s[i] = 1
	fmt.Println(s)
}

// №51. Как «открыть» запас
func ex51() {
	s := make([]int, 0, 3)
	s = s[:1]
	s[0] = 1
	s = s[:3]
	fmt.Println(s)
}

// №52. Данные за окном не исчезают
func ex52() {
	s := []int{1, 2, 3, 4, 5}
	t := s[:2]
	fmt.Println(t)
	t = t[:5]
	fmt.Println(t)
}

// №53. Запись через узкое окно видна в широком
func ex53() {
	s := []int{1, 2, 3, 4, 5}
	t := s[:2]
	t[0] = 10
	u := t[:4]
	fmt.Println(u)
}

// №54. Два перекрывающихся окна и copy
func ex54() {
	s := []int{1, 2, 3, 4, 5, 6}
	a := s[:4]
	b := s[2:]
	copy(b, a)
	fmt.Println(s)
}

// №55. Удаление элемента без append
func ex55() {
	s := []int{1, 2, 3, 4, 5}
	copy(s[1:], s[2:])
	s = s[:len(s)-1]
	fmt.Println(s, len(s), cap(s))
}

// №56. Разворот на месте
func ex56() {
	s := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	    s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}

// №57. Разворот части слайса
func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

func ex57() {
	s := []int{1, 2, 3, 4, 5, 6}
	reverse(s[1:5])
	fmt.Println(s)
}

// №58. Массив и слайс в функции
func modArr(a [3]int) {
    a[0] = 100
}

func modSl(s []int) {
    s[0] = 100
}

func ex58() {
	a := [3]int{1, 2, 3}
	s := []int{1, 2, 3}
	modArr(a)
	modSl(s)
	fmt.Println(a, s)
}

// №59. Слайс на массив и копия массива
func ex59() {
	a := [4]int{1, 2, 3, 4}
	s := a[:]
	s[0] = 100
	b := a
	b[1] = 200
	fmt.Println(a, s, b)
}

// №60. Размер заголовка
func ex60() {
	var s []int
	big := make([]int, 1000)
	var a [10]int
	fmt.Println(unsafe.Sizeof(s), unsafe.Sizeof(big), unsafe.Sizeof(a))
}

// №61. range по массиву и по слайсу
func ex61() {
	a := [3]int{1, 2, 3}
	for i, v := range a {
	    a[2] = 100
	    if i == 2 {
	        fmt.Println("массив:", v)
	    }
	}
	b := [3]int{1, 2, 3}
	s := b[:]
	for i, v := range s {
	    s[2] = 100
	    if i == 2 {
	        fmt.Println("слайс:", v)
	    }
	}
}

// №62. Экзамен 1: цепочка из трёх срезов
func ex62() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	a := s[2:5]
	b := a[1:6]
	c := b[2:4:5]
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
	fmt.Println(c, len(c), cap(c))
	c[0] = 0
	fmt.Println(s)
}

// №63. Экзамен 2: функция возвращает подслайс
func shift(s []int) []int {
    return s[1:]
}

func ex63() {
	s := []int{10, 20, 30, 40}
	t := shift(shift(s))
	t[0] = 0
	fmt.Println(s, t, len(t), cap(t))
}

// №64. Экзамен 3: три окна на один массив
func ex64() {
	s := make([]int, 5)
	for i := range s {
	    s[i] = i * i
	}
	a := s[1:3]
	b := s[2:5]
	a = a[:3]
	b[0] = -1
	fmt.Println(s, a, b)
}

// №65. Экзамен 4: поворот через copy
func rotate(s []int, k int) {
    tmp := make([]int, k)
    copy(tmp, s[:k])
    copy(s, s[k:])
    copy(s[len(s)-k:], tmp)
}

func ex65() {
	s := []int{1, 2, 3, 4, 5, 6}
	rotate(s, 2)
	fmt.Println(s)
}

// №66. Экзамен 5: ограниченное окно
func window(s []int, i, n int) []int {
    return s[i : i+n : i+n]
}

func ex66() {
	s := []int{1, 2, 3, 4, 5, 6}
	w := window(s, 1, 3)
	fmt.Println(w, len(w), cap(w))
	w = w[:cap(w)]
	w[2] = 0
	fmt.Println(s)
}

// №67. Экзамен 6: сетка поверх одного массива
func ex67() {
	data := make([]int, 6)
	grid := [][]int{data[0:3:3], data[3:6:6]}
	grid[1][0] = 7
	grid[0][2] = 5
	fmt.Println(data, grid)
}

// №68. Экзамен 7: nil или не nil
func ex68() {
	var a []int
	b := []int{}
	c := make([]int, 0)
	d := a[:0]
	fmt.Println(a == nil, b == nil, c == nil, d == nil)
	fmt.Println(len(a) == len(b), cap(c) == cap(d))
}

// №69. Экзамен 8: len, cap и panic
func ex69() {
	s := make([]int, 2, 4)
	t := s[1:3]
	fmt.Println(t, len(t), cap(t))
	t[1] = 5
	u := s[:4]
	fmt.Println(u)
	idx := 2
	fmt.Println(t[idx])
}

// №70. append к слайсу
func ex70() {
	s := []int{1, 2, 3}
	s = append(s, 4)
	fmt.Println(s, len(s))
}

// №71. append к nil
func ex71() {
	var s []int
	fmt.Println(s == nil)
	s = append(s, 1, 2)
	fmt.Println(s, len(s), s == nil)
}

// №72. Места хватает: память общая
func ex72() {
	a := make([]int, 3, 10)
	b := append(a, 7)
	b[0] = 1
	fmt.Println(a, b)
	fmt.Println(len(a), len(b))
}

// №73. Места не хватает: новый массив
func ex73() {
	a := []int{1, 2, 3}
	b := append(a, 4)
	b[0] = 100
	fmt.Println(a, b)
}

// №74. Проверка роста cap
func ex74() {
	a := []int{1, 2, 3}
	b := append(a, 4)
	fmt.Println(cap(a), cap(b) > cap(a))
}

// №75. Два append от одного слайса
func ex75() {
	s := make([]int, 2, 5)
	t := append(s, 1)
	u := append(s, 2)
	fmt.Println(t, u)
}

// №76. Два append, но места нет
func ex76() {
	s := make([]int, 2, 2)
	t := append(s, 1)
	u := append(s, 2)
	fmt.Println(t, u)
}

// №77. append к подслайсу затирает соседей
func ex77() {
	s := []int{1, 2, 3, 4, 5}
	t := s[:2]
	t = append(t, 100)
	fmt.Println(s, t)
}

// №78. Третий индекс защищает соседей
func ex78() {
	s := []int{1, 2, 3, 4, 5}
	t := s[:2:2]
	t = append(t, 100)
	fmt.Println(s, t)
}

// №79. append внутри функции
func add(s []int) {
    s = append(s, 99)
}

func ex79() {
	s := make([]int, 2, 10)
	add(s)
	fmt.Println(s, len(s))
	fmt.Println(s[:3])
}

// №80. append: возвращаем слайс из функции
func add2(s []int) []int {
    return append(s, 99)
}

func ex80() {
	s := []int{1}
	s = add2(s)
	fmt.Println(s)
}

// №81. append(a, b...)
func ex81() {
	a := []int{1, 2}
	b := []int{3, 4}
	a = append(a, b...)
	fmt.Println(a, b)
}

// №82. Затирание при нескольких append
func ex82() {
	a := make([]int, 2, 10)
	b := append(a, 1, 2, 3)
	c := append(a, 9)
	fmt.Println(b, c)
}

// №83. Удаление через append
func ex83() {
	s := []int{1, 2, 3, 4, 5}
	s = append(s[:1], s[2:]...)
	fmt.Println(s, len(s), cap(s))
}

// №84. Удаление: что стало с оригиналом
func ex84() {
	orig := []int{1, 2, 3, 4, 5}
	s := append(orig[:1], orig[2:]...)
	fmt.Println(s)
	fmt.Println(orig)
}

// №85. Вставка в середину
func ex85() {
	s := []int{1, 2, 4, 5}
	i := 2
	s = append(s[:i], append([]int{3}, s[i:]...)...)
	fmt.Println(s)
}

// №86. Рост cap у []int{1,2,3}
func ex86() {
	s := []int{1, 2, 3}
	s = append(s, 4)
	fmt.Println(len(s), cap(s))
}

// №87. Как растёт cap по шагам
func ex87() {
	var s []int
	last := cap(s)
	for i := 0; i < 5; i++ {
	    s = append(s, i)
	    if cap(s) != last {
	        fmt.Println("cap изменился:", last, "->", cap(s))
	        last = cap(s)
	    }
	}
}

// №88. Предвыделение: роста нет
func ex88() {
	s := make([]int, 0, 3)
	for i := 0; i < 3; i++ {
	    s = append(s, i*10)
	}
	fmt.Println(s, len(s), cap(s))
}

// №89. Классическая ошибка make + append
func ex89() {
	s := make([]int, 3)
	s = append(s, 1)
	fmt.Println(s, len(s))
}

// №90. Пустая база и два append
func ex90() {
	base := make([]int, 0, 10)
	a := append(base, 1)
	b := append(base, 2)
	fmt.Println(a, b)
}

// №91. Старый слайс видит новые данные
func ex91() {
	s := make([]int, 1, 4)
	t := append(s, 5)
	fmt.Println(s, t, s[:2])
}

// №92. Копирование при росте разрывает связь
func ex92() {
	a := make([]int, 2, 4)
	b := append(a, 1)
	a = append(a, 2, 3, 4)
	a[0] = 7
	fmt.Println(a, b)
}

// №93. Стек на слайсе
func ex93() {
	stack := []int{1, 2, 3}
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println(top, stack, len(stack), cap(stack))
	stack = append(stack, 9)
	fmt.Println(stack)
}

// №94. Фильтр на месте
func ex94() {
	s := []int{1, 2, 3, 4, 5, 6}
	out := s[:0]
	for _, v := range s {
	    if v%2 == 0 {
	        out = append(out, v)
	    }
	}
	fmt.Println(out, s)
}
