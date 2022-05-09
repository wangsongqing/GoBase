package helpers

// BubbleSort 排序
func BubbleSort(arr *[]int) []int {
	temp := 0
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}

	return *arr
}

func BinaryFind(arr []int, findVal int) int {

	keys := -1
	for k, v := range arr {
		if v == findVal {
			keys = k
			break
		}
	}

	return keys
}
