package practice

import "testing"

func eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func eq2(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !eq(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Errorf("Sum([1 2 3]) = %d, ожидали 6", got)
	}
	if got := Sum(nil); got != 0 {
		t.Errorf("Sum(nil) = %d, ожидали 0", got)
	}
}

func TestLast(t *testing.T) {
	v, ok := Last([]int{4, 5, 6})
	if v != 6 || !ok {
		t.Errorf("Last([4 5 6]) = %d, %v; ожидали 6, true", v, ok)
	}
	v, ok = Last(nil)
	if v != 0 || ok {
		t.Errorf("Last(nil) = %d, %v; ожидали 0, false", v, ok)
	}
}

func TestSwap(t *testing.T) {
	s := []int{1, 2, 3, 4}
	Swap(s, 0, 3)
	if !eq(s, []int{4, 2, 3, 1}) {
		t.Errorf("после Swap(s, 0, 3) получили %v", s)
	}
}

func TestFirstN(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	r := FirstN(s, 2)
	if !eq(r, []int{1, 2}) {
		t.Errorf("FirstN(s, 2) = %v", r)
	}
	r[0] = 100
	if s[0] != 100 {
		t.Errorf("FirstN должен возвращать подслайс (общая память), а не копию")
	}
	if got := FirstN([]int{1, 2, 3}, 10); !eq(got, []int{1, 2, 3}) {
		t.Errorf("FirstN(s, 10) = %v", got)
	}
	if got := FirstN([]int{1, 2, 3}, 0); len(got) != 0 {
		t.Errorf("FirstN(s, 0) = %v", got)
	}
	if got := FirstN([]int{1, 2, 3}, -1); len(got) != 0 {
		t.Errorf("FirstN(s, -1) = %v", got)
	}
}

func TestClone(t *testing.T) {
	s := []int{1, 2, 3}
	c := Clone(s)
	if !eq(c, s) {
		t.Errorf("Clone = %v, ожидали %v", c, s)
	}
	c[0] = 100
	if s[0] != 1 {
		t.Errorf("Clone должен делать независимую копию")
	}
	if Clone(nil) != nil {
		t.Errorf("Clone(nil) должен вернуть nil")
	}
}

func TestReverse(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	Reverse(s)
	if !eq(s, []int{5, 4, 3, 2, 1}) {
		t.Errorf("Reverse = %v", s)
	}
	s2 := []int{1, 2, 3, 4}
	Reverse(s2)
	if !eq(s2, []int{4, 3, 2, 1}) {
		t.Errorf("Reverse (чётная длина) = %v", s2)
	}
	Reverse(nil)
}

func TestMiddle(t *testing.T) {
	s := []int{1, 2, 3, 4}
	r := Middle(s)
	if !eq(r, []int{2, 3}) {
		t.Errorf("Middle([1 2 3 4]) = %v", r)
	}
	r[0] = 100
	if s[1] != 100 {
		t.Errorf("Middle должен возвращать подслайс (общая память)")
	}
	for _, in := range [][]int{nil, {1}, {1, 2}} {
		if got := Middle(in); len(got) != 0 {
			t.Errorf("Middle(%v) = %v, ожидали пустой слайс", in, got)
		}
	}
}

func TestChunks(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	r := Chunks(s, 3)
	if !eq2(r, [][]int{{1, 2, 3}, {4, 5, 6}, {7}}) {
		t.Fatalf("Chunks(s, 3) = %v", r)
	}
	for i, c := range r {
		if cap(c) != len(c) {
			t.Errorf("кусок %d: cap = %d, ожидали %d (нужен третий индекс)", i, cap(c), len(c))
		}
	}
	r[0][0] = 100
	if s[0] != 100 {
		t.Errorf("куски должны быть подслайсами s (общая память)")
	}
	if got := Chunks(nil, 3); len(got) != 0 {
		t.Errorf("Chunks(nil, 3) = %v", got)
	}
	if got := Chunks([]int{1, 2}, 0); got != nil {
		t.Errorf("Chunks(s, 0) должен вернуть nil")
	}
}

func TestFilter(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	even := func(v int) bool { return v%2 == 0 }
	r := Filter(s, even)
	if !eq(r, []int{2, 4, 6}) {
		t.Errorf("Filter = %v", r)
	}
	if !eq(s, []int{1, 2, 3, 4, 5, 6}) {
		t.Errorf("Filter изменил исходный слайс: %v", s)
	}
	if got := Filter([]int{1, 3}, even); len(got) != 0 {
		t.Errorf("Filter без подходящих = %v", got)
	}
}

func TestRemove(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	r := Remove(s, 1)
	if !eq(r, []int{1, 3, 4, 5}) {
		t.Errorf("Remove(s, 1) = %v", r)
	}
	if cap(r) != 5 {
		t.Errorf("Remove не должен выделять новый массив: cap = %d, ожидали 5", cap(r))
	}
	if got := Remove([]int{1, 2, 3}, 2); !eq(got, []int{1, 2}) {
		t.Errorf("Remove(последний) = %v", got)
	}
	if got := Remove([]int{1, 2, 3}, 3); !eq(got, []int{1, 2, 3}) {
		t.Errorf("Remove(вне диапазона) = %v", got)
	}
	if got := Remove([]int{1, 2, 3}, -1); !eq(got, []int{1, 2, 3}) {
		t.Errorf("Remove(-1) = %v", got)
	}
}

func TestInsert(t *testing.T) {
	if got := Insert([]int{1, 2, 4, 5}, 2, 3); !eq(got, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Insert в середину = %v", got)
	}
	if got := Insert([]int{1, 2}, 0, 0); !eq(got, []int{0, 1, 2}) {
		t.Errorf("Insert в начало = %v", got)
	}
	if got := Insert([]int{1, 2}, 2, 9); !eq(got, []int{1, 2, 9}) {
		t.Errorf("Insert в конец = %v", got)
	}
	if got := Insert(nil, 0, 1); !eq(got, []int{1}) {
		t.Errorf("Insert(nil, 0, 1) = %v", got)
	}
	if got := Insert([]int{1, 2}, 5, 9); !eq(got, []int{1, 2}) {
		t.Errorf("Insert вне диапазона = %v", got)
	}
}

func TestUnique(t *testing.T) {
	s := []int{1, 1, 2, 3, 3, 3, 4}
	r := Unique(s)
	if !eq(r, []int{1, 2, 3, 4}) {
		t.Fatalf("Unique = %v", r)
	}
	if &r[0] != &s[0] {
		t.Errorf("Unique должен работать на месте (тот же массив)")
	}
	if got := Unique(nil); len(got) != 0 {
		t.Errorf("Unique(nil) = %v", got)
	}
	if got := Unique([]int{5, 5, 5}); !eq(got, []int{5}) {
		t.Errorf("Unique([5 5 5]) = %v", got)
	}
}

func TestRotate(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	Rotate(s, 2)
	if !eq(s, []int{3, 4, 5, 1, 2}) {
		t.Errorf("Rotate(s, 2) = %v", s)
	}
	s2 := []int{1, 2, 3, 4, 5}
	Rotate(s2, 7)
	if !eq(s2, []int{3, 4, 5, 1, 2}) {
		t.Errorf("Rotate(s, 7) = %v", s2)
	}
	s3 := []int{1, 2, 3}
	Rotate(s3, 0)
	if !eq(s3, []int{1, 2, 3}) {
		t.Errorf("Rotate(s, 0) = %v", s3)
	}
	Rotate(nil, 3)
}

func TestMerge(t *testing.T) {
	a := []int{1, 4, 6}
	b := []int{2, 3, 7, 8}
	r := Merge(a, b)
	if !eq(r, []int{1, 2, 3, 4, 6, 7, 8}) {
		t.Fatalf("Merge = %v", r)
	}
	if cap(r) != len(a)+len(b) {
		t.Errorf("cap(результата) = %d, ожидали %d (нужен make с предвыделением)", cap(r), len(a)+len(b))
	}
	if !eq(a, []int{1, 4, 6}) || !eq(b, []int{2, 3, 7, 8}) {
		t.Errorf("Merge изменил входные слайсы")
	}
	if got := Merge(nil, []int{1, 2}); !eq(got, []int{1, 2}) {
		t.Errorf("Merge(nil, b) = %v", got)
	}
	if got := Merge(nil, nil); len(got) != 0 {
		t.Errorf("Merge(nil, nil) = %v", got)
	}
}

func TestMatrix(t *testing.T) {
	m := Matrix(2, 3)
	if len(m) != 2 || len(m[0]) != 3 || len(m[1]) != 3 {
		t.Fatalf("размеры матрицы неверны: %v", m)
	}
	m[0][0] = 1
	if m[1][0] != 0 {
		t.Errorf("строки матрицы должны быть независимы")
	}
	if got := Matrix(0, 5); len(got) != 0 {
		t.Errorf("Matrix(0, 5) = %v", got)
	}
}

func TestTranspose(t *testing.T) {
	m := [][]int{{1, 2, 3}, {4, 5, 6}}
	r := Transpose(m)
	if !eq2(r, [][]int{{1, 4}, {2, 5}, {3, 6}}) {
		t.Errorf("Transpose = %v", r)
	}
	if !eq2(m, [][]int{{1, 2, 3}, {4, 5, 6}}) {
		t.Errorf("Transpose изменил исходную матрицу")
	}
	if got := Transpose(nil); len(got) != 0 {
		t.Errorf("Transpose(nil) = %v", got)
	}
}

func TestStack(t *testing.T) {
	var st Stack
	if _, ok := st.Pop(); ok {
		t.Errorf("Pop на пустом стеке должен вернуть false")
	}
	if _, ok := st.Peek(); ok {
		t.Errorf("Peek на пустом стеке должен вернуть false")
	}
	st.Push(1)
	st.Push(2)
	st.Push(3)
	if st.Len() != 3 {
		t.Errorf("Len = %d, ожидали 3", st.Len())
	}
	if v, ok := st.Peek(); v != 3 || !ok {
		t.Errorf("Peek = %d, %v; ожидали 3, true", v, ok)
	}
	if st.Len() != 3 {
		t.Errorf("Peek не должен удалять элемент")
	}
	for _, want := range []int{3, 2, 1} {
		if v, ok := st.Pop(); v != want || !ok {
			t.Errorf("Pop = %d, %v; ожидали %d, true", v, ok, want)
		}
	}
	if st.Len() != 0 {
		t.Errorf("после трёх Pop Len = %d", st.Len())
	}
}
