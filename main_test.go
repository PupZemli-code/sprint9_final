package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func Test_generateRandomElements_size(t *testing.T) {
	// проверка: длина = 0

	size := 0
	randValues, err := generateRandomElements(size)
	require.Error(t, err)
	require.Len(t, randValues, 0)

	// проверка: длина = 1

	size = 1
	randValues, err = generateRandomElements(size)
	require.NoError(t, err)
	require.Len(t, randValues, 1)

	// проверка: длина = 10

	size = 10
	randValues, err = generateRandomElements(size)
	require.NoError(t, err)
	require.Len(t, randValues, 10)
}

func Test_maximum_GiveEmptySlice(t *testing.T) {
	// проверка: len slice = nil

	var randValues []int
	value, err := maximum(randValues)
	require.Error(t, err)
	require.Equal(t, value, 0)

	// проверка: lin slice = 0

	randValues = make([]int, 0)
	value, err = maximum(randValues)
	require.Error(t, err)
	require.Equal(t, value, 0)

	// проверка: len slice = 1

	randValues = append(randValues, 1)
	value, err = maximum(randValues)
	require.Error(t, err)
	require.Equal(t, value, 0)

	// проверка: len slice = 9 (корректный запрос)

	randValues = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	value, err = maximum(randValues)
	require.NoError(t, err)
	require.Equal(t, value, 9)
}

func Test_maxChunks_SizeSlice(t *testing.T) {

	// проверка: len slice = nil

	var randValues []int
	value, err := maxChunks(randValues)
	require.Error(t, err)
	require.Equal(t, value, 0)

	// проверка: lin slice = 0

	randValues = make([]int, 0)
	value, err = maxChunks(randValues)
	require.Error(t, err)
	require.Equal(t, value, 0)

	// проверка: len slice = 1

	randValues = append(randValues, 1)
	value, err = maxChunks(randValues)
	require.Error(t, err)
	require.Equal(t, value, 0)

	// проверка: len slice = 80 (корректный запрос)

	slice := make([]int, 80)
	for i := 0; i < 80; i++ {
		slice[i] = i
	}
	maxVal, _ := maxChunks(slice)
	require.Equal(t, maxVal, 79)
}
