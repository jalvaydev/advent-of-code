package day1

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func FindTopCalorieElf(path string) int {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	dataString := string(data)
	calorieData := strings.Split(dataString, "\n\n")
	maxCalories := 0

	for _, v := range calorieData {
		vals := strings.Fields(v)
		counter := 0
		for _, val := range vals {
			valInt, _ := strconv.Atoi(val)
			counter += valInt
		}

		maxCalories = int(math.Max(float64(maxCalories), float64(counter)))
	}

	return maxCalories
}

func FindTopKCalorieElves(path string, k int) int {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	dataString := string(data)
	calorieData := strings.Split(dataString, "\n\n")

	h := &IntHeap{}
	heap.Init(h)

	for _, v := range calorieData {
		vals := strings.Fields(v)
		counter := 0
		for _, val := range vals {
			valInt, _ := strconv.Atoi(val)
			counter += valInt
		}

		if h.Len() < k {
			heap.Push(h, counter)
		} else {
			if (*h)[0] < counter {
				heap.Pop(h)
				heap.Push(h, counter)
			}
		}
	}

	sum := 0
	for _, v := range *h {
		sum += v
	}

	return sum
}
