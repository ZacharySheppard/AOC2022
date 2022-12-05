package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntMinHeap) Pop() any {
	tmp := *h
	n := len(tmp)
	x := tmp[n-1]
	*h = tmp[0 : len(tmp)-1]
	return x
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	topThreeElves := IntMinHeap{0, 0, 0}
	heap.Init(&topThreeElves)

	current := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			count, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			current += count
			continue
		}

		if topThreeElves.Len() < 3 {
			heap.Push(&topThreeElves, current)
		} else {
			if current > topThreeElves[0] {
				heap.Pop(&topThreeElves)
				heap.Push(&topThreeElves, current)
			}
		}
		current = 0
	}

	if current > topThreeElves[0] {
		heap.Pop(&topThreeElves)
		heap.Push(&topThreeElves, current)
	}

	topThreeTotal := 0
	for topThreeElves.Len() > 0 {
		topThreeTotal += heap.Pop(&topThreeElves).(int)
	}

	fmt.Println("The sum of the three elves carrying the most calories is: " + fmt.Sprint(topThreeTotal))
}
