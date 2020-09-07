package problem_347

import (
	"container/heap"
)

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int, len(nums))
	for _, num := range nums {
		counter[num]++
	}
	pq := make(PriorityQueue, 0, len(counter))
	for num, priority := range counter {
		if len(pq) < k {
			heap.Push(&pq, &Item{
				value:    num,
				priority: priority,
			})
		} else {
			if pq[0].priority < priority {
				heap.Pop(&pq)
				heap.Push(&pq, &Item{
					value:    num,
					priority: priority,
				})
			}
		}
	}
	result := make([]int, 0, len(pq))
	for len(pq) > 0 {
		result = append(result, heap.Pop(&pq).(*Item).value)
	}
	return result
}
