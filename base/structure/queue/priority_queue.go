package queue

type QItem interface {
	Then(obj interface{}) int
}
type PriorityQueue struct {
	data []QItem
}

func (qp *PriorityQueue) GetSize() int {
	return len(qp.data)
}

func (qp *PriorityQueue) IsEmpty() bool {
	return qp.GetSize() != 0
}

func (qp *PriorityQueue) Push(e QItem) {
	qp.data = append(qp.data, e)
	qp.siftUp(len(qp.data) - 1)
}

func (qp *PriorityQueue) Pop() QItem {
	ret := qp.Top()
	qp.swap(0, len(qp.data)-1)
	qp.data = qp.data[:len(qp.data)-1]
	qp.siftDown(0)
	return ret
}

func (qp *PriorityQueue) Top() QItem {
	if qp.GetSize() == 0 {
		panic("Can't findMax when heap")
	}
	return qp.data[0]
}

func (qp *PriorityQueue) Replace(e QItem) QItem {
	ret := qp.Top()
	qp.data[0] = e
	qp.siftDown(0)
	return ret
}

func (qp *PriorityQueue) parent(index int) int {
	if index == 0 {
		panic("index-0 doesn't have parent.")
	}
	return (index - 1) / 2
}

func (qp *PriorityQueue) leftChild(index int) int {
	return index*2 + 1
}

func (qp *PriorityQueue) rightChild(index int) int {
	return index*2 + 2
}

func (qp *PriorityQueue) siftUp(index int) {
	// 首先当前节点必须大于零，因为一直上浮到堆顶，index是0，则退出
	// 每一次都比较当前节点是否大于父亲节点
	for index > 0 && qp.data[qp.parent(index)].Then(qp.data[index]) < 0 {

		// 与父亲节点交换
		qp.swap(index, qp.parent(index))

		// 然后把父亲节点当成当前节点，继续上浮
		index = qp.parent(index)
	}
}

func (qp *PriorityQueue) siftDown(index int) {
	for qp.leftChild(index) < len(qp.data) {

		max := qp.leftChild(index) // 先让左节点左最大值
		right := qp.rightChild(index)

		if right < len(qp.data) {
			// 比较左右节点的大小, 取最大值得索引
			if qp.data[right].Then(qp.data[max]) > 0 {
				max = right
			}
		}

		// 当当前节点的值，比左右节点中最大值得节点都大，则没必要调整了
		if qp.data[index].Then(qp.data[max]) >= 0 {
			break
		}

		// 否则交换当前节点与左右节点中的最大值节点
		qp.swap(index, max)

		// 然后接续下沉
		index = max
	}
}

func (qp *PriorityQueue) swap(i, j int) {
	if i < 0 || i >= len(qp.data) || j < 0 || j >= len(qp.data) {
		panic("index is illegal.")
	}

	t := qp.data[i]
	qp.data[i] = qp.data[j]
	qp.data[j] = t
}

func NewPriorityQueue(capacity int) *PriorityQueue {
	return &PriorityQueue{
		data: make([]QItem, capacity),
	}
}

func NewPriorityQueueFromArr(arr []QItem) *PriorityQueue {
	qp := &PriorityQueue{data: arr}
	for i := qp.parent(len(qp.data) - 1); i >= 0; i-- {
		qp.siftDown(i)
	}
	return qp
}
