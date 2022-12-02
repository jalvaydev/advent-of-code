package day1

import (
	"container/heap"
	"math"
	"os"
	"strconv"
	"strings"
)

func FindTopCalorieElf(path string) (int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return -1, err
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

	return maxCalories, nil
}

func FindTopKCalorieElves(path string, k int) (int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return -1, err
	}

	dataString := string(data)
	calorieData := strings.Split(dataString, "\n\n")

	h := &IntHeap{}
	heap.Init(h)

	sum := 0

	for _, v := range calorieData {
		vals := strings.Fields(v)
		counter := 0
		for _, val := range vals {
			valInt, _ := strconv.Atoi(val)
			counter += valInt
		}

		if h.Len() < k {
			heap.Push(h, counter)
			sum += counter
		} else {
			if (*h)[0] < counter {
				sum -= heap.Pop(h).(int)
				heap.Push(h, counter)
				sum += counter
			}
		}
	}

	return sum, nil
}
