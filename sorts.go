package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	// confirmedIndex keeps track of what is sorted to get rid of redundant comparisons
	var temp, confirmedIndex int;
	for n > 1 {
		confirmedIndex = 0
		for i := 1; i < n; i++ {
			if(arr[i - 1] > arr[i]) {
				temp = arr[i - 1];
        			arr[i - 1] = arr[i];
        			arr[i] = temp;

        			confirmedIndex = i;
			}
		}
		n = confirmedIndex;
	}
}

func shuffleSort(arr []int) {
	beginIndex := 1
	endIndex := len(arr)
	// newBegin and newInt keep track of what is sorted to get rid of redundant comparisons
	var temp, newBegin, newEnd int
	
	for beginIndex <= endIndex {
		newBegin = endIndex
		newEnd = beginIndex
		
		for i := beginIndex; i < endIndex; i++ {
			if(arr[i - 1] > arr[i]) {
				temp = arr[i - 1];
       				arr[i - 1] = arr[i];
       				arr[i] = temp;

        			newEnd = i;
			}
		}
		endIndex = newEnd - 1;
		
		for i := endIndex; i >= beginIndex; i-- {
			if(arr[i - 1] > arr[i]) {
				temp = arr[i - 1];
       				arr[i - 1] = arr[i];
       				arr[i] = temp;

        			newBegin = i;
			}
		}
		beginIndex = newBegin + 1;
	}
}

func main() {
	fmt.Println("Bubble sort")
	array := []int{3,5,2,9,6,1,8}
	fmt.Println(array)
	bubbleSort(array)
	fmt.Println(array)

	fmt.Println("Shuffle sort")
	array = []int{3,5,2,9,6,1,8}
	fmt.Println(array)
	shuffleSort(array)
	fmt.Println(array)
}