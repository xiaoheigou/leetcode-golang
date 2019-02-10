package main

import "fmt"

func main() {
	a := []int{2, 1, 3, 34, 45, 76, 37, 46}
	fmt.Println(a)
	b := heapsort(a)
	fmt.Println(b)
}

// 将数组变成堆数据结构，便利的获取父节点，子节点，但是golnag中还要➕1
func left(i int) int {
	return 2 * i
}
func right(i int) int {
	return 2*i + 1
}

func parent(i int) int {
	return i / 2
}

//maxHeapify
func maxHeapify(a []int, i int) []int {

	l := left(i) + 1
	r := right(i) + 1
	// fmt.Println(i, l, r)
	var largest int
	if l < len(a) && a[l] > a[i] {
		largest = l
		// fmt.Println(a[l], ">", a[i])
	} else {
		largest = i
	}
	if r < len(a) && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		// fmt.Printf("exchange a[%d](%d) with a[%d](%d)", i, a[i], largest, a[largest])
		a[i], a[largest] = a[largest], a[i]
		// recurse
		a = maxHeapify(a, largest)
	}
	// fmt.Println(a)
	return a
}

// buildMaxheap
// start from leaves!!
func buildMaxHeap(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		// if i == 2 {
		// 	fmt.Println(i)
		// }
		a = maxHeapify(a, i)
	}
	return a
}

func heapsort(a []int) []int {
	a = buildMaxHeap(a)
	fmt.Println("start sort array")
	size := len(a)
	for i := size - 1; i >= 1; i-- {
		a[0], a[size-1] = a[i], a[0]
		size--
		maxHeapify(a[:size], 0)
	}
	return a
}
