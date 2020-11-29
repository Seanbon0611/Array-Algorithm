package main

import "fmt"

/*
Write a function that accepts two arrays
The function should return true if every value in the array has a corresponding value squared in the second array the frequency of values must be the same
*/

func same(arr1 []float64, arr2 []float64) bool {
	//Check if arrays are the same length. If they are not return false
	if len(arr1) != len(arr2) {
		fmt.Println("false, arrays are not of the same length")
		return false
	}
	//loop through arrays and see if array 2 contains a value equal to the squared value of the first array
	firstArrayMap := make(map[float64]float64)
	secondArrayMap := make(map[float64]float64)
	for _, num := range arr1 {
		if firstArrayMap[num] > 0 {
			firstArrayMap[num] += 1
		} else {
			firstArrayMap[num] = 1
		}
	}
	for _, num := range arr2 {
		if secondArrayMap[num] > 0 {
			secondArrayMap[num] += 1
		} else {
			secondArrayMap[num] = 1
		}
	}
  //iterate through first map
	for key, _ := range firstArrayMap {
    //Checking to see if second map contains a value thats squared the first map's key
		_, ok := secondArrayMap[key*key]
    //if does not contain the correlating squared element return false
		if !ok {
			fmt.Println("false, arrays do not contain correlating squared elements")
			return false
		}
    //Check to see if both maps are of the same frequency if not return false
		if secondArrayMap[key*key] != firstArrayMap[key] {
      fmt.Println("false, arrays do not have the same frequency")
			return false
		}
	}
	//return true
	fmt.Println("true")
	return true
}

func main() {
	array1 := []float64{1, 2, 3}
	array2 := []float64{4, 1, 9}
	same(array1, array2)
}
