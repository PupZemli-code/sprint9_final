package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) ([]int, error) {
	if size <= 0 {
		return []int{}, fmt.Errorf("invalid function argument: size %d <= 0", size)
	}
	var randValues []int
	src := rand.NewSource(time.Now().Unix())
	for i := 0; i < size; i++ {
		randomNumber := src.Int63()
		randomNumber = randomNumber%int64(size) + 1
		randValues = append(randValues, int(randomNumber))
	}
	return randValues, nil
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int, error) {
	maxValue := 0
	if len(data) <= 1 {
		return 0, fmt.Errorf("incorrect size slice: data = nil or 0 or 1")
	}
	for _, value := range data {
		if maxValue < value {
			maxValue = value
		}
	}
	return maxValue, nil
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) (int, error) {
	if len(data) <= 1 {
		return 0, fmt.Errorf("incorrect size slice: data = nil or 0 or 1")
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	lenSlice := len(data) / CHUNKS
	maxValueSlice := make([]int, 8)
	for i := 0; i < CHUNKS; i++ {
		slice := data[i*lenSlice : i*lenSlice+lenSlice]
		wg.Add(1)
		go func(s []int) {
			defer wg.Done()
			maxVal := 0
			for _, val := range slice {
				if val > maxVal {
					maxVal = val
				}
			}
			mu.Lock()
			maxValueSlice = append(maxValueSlice, maxVal)
			mu.Unlock()
		}(slice)
	}
	wg.Wait()
	maxVal := 0
	for _, val := range maxValueSlice {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal, nil
}

func main() {

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	slice, _ := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now() // начало отсчета
	max, _ := maximum(slice)
	elapsed := time.Since(start) // конец отсчета
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max, _ = maxChunks(slice)
	elapsed = time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
	fmt.Print(max)
}
