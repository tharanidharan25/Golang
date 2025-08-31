package main

import "fmt"

/*
	Assignment:
		We've been asked to add a feature that extracts costs for a given day.

		Complete the getDayCosts() function using the append() function. It accepts a slice of cost structs and a day int, and it returns a float64 slice containing that day's costs.

		If there are no costs for that day, return an empty slice.
*/
type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) (res []float64) {
	for i := 0; i < len(costs); i++ {
		cur := costs[i]
		if cur.day == day {
			res = append(res, cur.value)
		}
	}
	return
}


// Assignment for Range
/*
We need to be able to quickly detect bad words in the messages our system sends.

Complete the indexOfFirstBadWord function. If it finds any bad words in the message it should return the index of the first bad word in the msg slice. This will help us filter out naughty words from our messaging system. If no bad words are found, return -1 instead.

Use the range keyword.
*/
func indexOfFirstBadWord(msg []string, badWords []string) int {
	for idx, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return idx
			}
		}
	}
	return -1
}

/*
Assignment for Matrix or 2D Slices:
We support various visualization dashboards on Textio that display message analytics to our users. The UI for our graphs and charts is built on top of a grid system. Let's build some grid logic.

Complete the createMatrix function. It takes a number of rows and columns and returns a 2D slice of integers. The value of each cell is i * j where i and j are the indexes of the row and column respectively. Basically, we're building a multiplication chart.
*/
func createMatrix(rows, cols int) [][]int {
	var res [][]int
	for i := 0; i < rows; i++ {
		var curRow []int
		for j := 0;  j < cols; j++ {
			curRow = append(curRow, i * j)
		}
		res = append(res, curRow)
	}
	return res
}

func main() {
	// Built-in append function is used to dynamically add elements to an array
	// Append Function signature:
	
	/*
		func append(slice []Type, elems ...Type) []Type
	*/
	
	// Note that append() function is a variadic function, so the following are all valid

	/*
		slice = append(slice, oneThing)
		slice = append(slice, firstThing, secondThing)
		slice = append(slice, anotherSlice...)
	*/

	// Append function creates a new slice when the underlying array's capacity has been reached, else it'll use the same array.
	// It's always better to assign the value coming from the append function to the same slice we used inside the append function.
	ip := make([]cost, 10)
	for i := 0; i < 10; i++ {
		ip = append(ip, cost{
			i,
			5.0,
		})
	}
	res := getDayCosts(ip, 5)
	fmt.Println(res)

	// Range keyword:
	// Go provides syntactic sugar to iterate easily over elements of a slice:
	/*
		for INDEX, ELEMENT := range SLICE {}
	*/
	// example:
	fruits := []string{"apple", "banana", "grape"}
	for idx, fruit := range fruits {
		fmt.Println(idx, fruit)
	}

	idxOfBadWord := indexOfFirstBadWord([]string{"hey", "there", "john"}, []string{"crap", "shoot", "frick", "dang"})
	fmt.Println(idxOfBadWord)
	idxOfBadWord = indexOfFirstBadWord([]string{"ugh", "oh", "my", "frick"}, []string{"crap", "shoot", "frick", "dang"})
	fmt.Println(idxOfBadWord)

	// Slices can hold other slices, effectively creating a matrix, or a 2D slice.
	/*
		rows := [][]int{}
	*/
	newMatrix := createMatrix(5, 10)
	fmt.Println(newMatrix)
}