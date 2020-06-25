package head

import "fmt"

type MaxArrayHeap struct {
	data []int
}

func (mah *MaxArrayHeap) GetSize() int {
	return len(mah.data)
}

func (mah *MaxArrayHeap) IsEmpty() bool {
	return mah.GetSize() == 0
}

func (mah *MaxArrayHeap) parent(index int) int {
	if index == 0 {
		panic("index-0 doest't have parent.")
	}
	return (index - 1) / 2
}

func (mah *MaxArrayHeap) leftChild(index int) int {
	return index*2 + 1
}

func (mah *MaxArrayHeap) rightChild(index int) int {
	return index*2 + 2
}

func (mah *MaxArrayHeap) Add(e int) {
	mah.data = append(mah.data, e)
	mah.siftUp(len(mah.data) - 1)
}

func (mah *MaxArrayHeap) ExtraMax() int {
	max := mah.data[0]
	mah.swap(0, len(mah.data)-1)
	mah.data = mah.data[0 : len(mah.data)-1]
	mah.siftDown(0)
	return max
}

func (mah *MaxArrayHeap) siftUp(index int) {
	for index > 0 && mah.data[mah.parent(index)] < mah.data[index] {
		mah.swap(mah.parent(index), index)
		index = mah.parent(index)
	}
}

func (mah *MaxArrayHeap) siftDown(index int) {
	for mah.leftChild(index) < mah.GetSize() {
		max := mah.leftChild(index)
		right := mah.rightChild(index)
		if right < mah.GetSize() {
			if mah.data[right] > mah.data[max] {
				max = right
			}
		}
		if mah.data[index] >= mah.data[max] {
			break
		}
		mah.swap(index, max)
		index = max
	}
}

func (mah *MaxArrayHeap) swap(i, j int) {
	if i < 0 || i >= mah.GetSize() || j < 0 || j >= mah.GetSize() {
		return
	}
	mah.data[i], mah.data[j] = mah.data[j], mah.data[i]
}

func (mah *MaxArrayHeap) heapify() {
	for i := mah.parent(len(mah.data) - 1); i >= 0; i-- {
		mah.siftDown(i)
	}
}

func (mah *MaxArrayHeap) Print() {
	fmt.Println(mah.data)
}

func NewMaxArrayHeap(nums []int) *MaxArrayHeap {
	mah := new(MaxArrayHeap)
	mah.data = nums
	mah.heapify()
	return mah
}
