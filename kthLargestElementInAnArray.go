package leetcodeByGo

// 第k大元素

type Heap struct {
	nums     []int
	heapSize int
}

// 计算父节点在数组中的索引
func (heap *Heap) GetParent(i int) int {
	if i == 0 {
		return i
	}

	return (i - 1) / 2
}

// 计算左子节点在数组中的索引
func (heap *Heap) GetLeft(i int) int {
	return 2*i + 1
}

// 计算右子节点在数组中的索引
func (heap *Heap) GetRight(i int) int {
	return 2*i + 2
}

// 维护最大堆的性质
func (heap *Heap) MaxHeapify(i int) {
	if i >= len(heap.nums) {
		return
	}

	// 计算左子节点的索引
	left := heap.GetLeft(i)

	// 计算右子节点的索引
	right := heap.GetRight(i)

	// largest保存最大值的索引
	largest := -1

	if left < heap.heapSize && heap.nums[left] > heap.nums[i] {
		largest = left
	} else {
		largest = i
	}

	if right < heap.heapSize && heap.nums[right] > heap.nums[largest] {
		largest = right
	}

	if largest != i {
		temp := heap.nums[i]
		heap.nums[i] = heap.nums[largest]
		heap.nums[largest] = temp
		heap.MaxHeapify(largest)
	}
}

// 建堆
func (heap *Heap) BuildHeap() {
	// 计算最后一个非叶子节点的索引
	heap.heapSize = len(heap.nums)
	lastNoLeafNode := len(heap.nums)/2 - 1
	for i := lastNoLeafNode; i >= 0; i-- {
		heap.MaxHeapify(i)
	}
}

// 堆排序
func (heap *Heap) HeapSort() {
	// 堆排序的思路是，每次将待排序区域的最后一个元素与第一个元素交换，然后缩减待排序区域并进行最大堆化
	for i := len(heap.nums) - 1; i >= 1; i-- {
		temp := heap.nums[i]
		heap.nums[i] = heap.nums[0]
		heap.nums[0] = temp

		heap.heapSize--

		heap.MaxHeapify(0)
	}
}

func findKthLargest(nums []int, k int) int {
	heap := Heap{
		nums:     nums,
		heapSize: -1,
	}

	// 建堆
	heap.BuildHeap()

	// 堆排序
	heap.HeapSort()
	return heap.nums[len(heap.nums)-k]
}
