package graph

// Item represents a node in the priority queue
type Item struct {
	Node     int
	Distance int
	Index    int // Index in the heap
}

// PriorityQueue implements heap.Interface and holds Items
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Update modifies the distance of an Item in the queue
func (pq *PriorityQueue) Update(item *Item, distance int) {
	item.Distance = distance
	Fix(pq, item.Index)
}

// Fix re-establishes the heap ordering after the element at index i has changed its value
func Fix(pq *PriorityQueue, i int) {
	down(pq, i, len(*pq))
	up(pq, i)
}

func up(pq *PriorityQueue, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !(*pq).Less(j, i) {
			break
		}
		(*pq).Swap(i, j)
		j = i
	}
}

func down(pq *PriorityQueue, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && (*pq).Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !(*pq).Less(j, i) {
			break
		}
		(*pq).Swap(i, j)
		i = j
	}
	return i > i0
}
