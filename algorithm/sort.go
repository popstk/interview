package algorithm

func QSort(data []int) {
	i, j := 0, len(data)-1
	pivot := data[0]

	for i < j {
		for ; i < j && data[j] >= pivot; {
			j--
		}
		if i < j {
			data[i] = data[j]
			i++
		}

		for ; i < j && data[i] <= pivot; {
			i++
		}
		if i < j {
			data[j] = data[i]
			j--
		}
	}
	data[i] = pivot

	if i > 0 {
		QSort(data[:i])
	}
	if len(data)-1-(i+1) > 0 {
		QSort(data[i+1:])
	}
}

func shiftDown(data []int, i int) {
	n := len(data) - 1
	pos := i
	for i <= n {
		if 2*i+1 <= n && data[2*i+1] > data[pos] {
			pos = 2*i + 1
		}
		if 2*i+2 <= n && data[2*i+2] > data[pos] {
			pos = 2*i + 2
		}
		if i == pos {
			break
		}
		data[i], data[pos] = data[pos], data[i]
		i = pos
	}
}

func HeapSort(data []int) {
	// left
	for i := len(data) / 2; i >= 0; i-- {
		shiftDown(data, i)
	}

	for i := len(data) - 1; i > 0; i-- {
		data[i], data[0] = data[0], data[i]
		shiftDown(data[:i], 0)
	}

}

func MergeSort(data []int) {
	n := len(data)
	middle := n / 2
	if middle > 0 {
		MergeSort(data[:middle])
	}
	if n-1-middle > 0 {
		MergeSort(data[middle:])
	}

	temp := make([]int, len(data))
	i, j, k := 0, middle, 0
	for ; i < middle && j < len(data); k++ {
		if data[i] < data[j] {
			temp[k] = data[i]
			i++
		} else {
			temp[k] = data[j]
			j++
		}
	}
	for ; i < middle; i++ {
		temp[k] = data[i]
		k++
	}

	for ; j < len(data); j++ {
		temp[k] = data[j]
		k++
	}
	copy(data, temp)
}
