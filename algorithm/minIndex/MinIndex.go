package main

import "fmt"

func main() {
	var list1 = []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	var list2 = []string{"Piatti", "Burger King", "The Grill at Torrey Pines", "Shogun"}
	result := findRestaurant(list1, list2)
	fmt.Printf("result is %v\n", result)
}

func findRestaurant(list1 []string, list2 []string) []string {
	if len(list1) == 0 || len(list2) == 0 {
		return nil
	}

	//将list转为map
	listMap := make(map[string]int, 100)
	for i := range list1 {
		listMap[list1[i]] = i
	}
	fmt.Printf("listMap:%v\n", listMap)
	var resultSlice = make([]string, 10, 10)
	var count = 10000
	for j := range list2 {

		if _, ok := listMap[list2[j]]; ok {
			//匹配上
			co := listMap[list2[j]] + j
			if co < count {
				count = co
				resultSlice = resultSlice[0:0]
				resultSlice = append(resultSlice, list2[j])
				continue
			}
			if co == count {
				resultSlice = append(resultSlice, list2[j])
			}
		}
	}
	return resultSlice
}
