package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// lwh
// 2*l*w + 2*w*h + 2*h*l

var total = int64(0)
var ribbon = int64(0)

func main() {

	// open a file
	if file, err := os.Open("sizes.txt"); err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			l, w, h := extractSizes(scanner.Text())
			total += surfaceAreaPlus(l, w, h)

			// Let's do the ribbon now.
			ribbon += volumeOfPresent(l, w, h)
			//fmt.Printf("Ribbon Volume: %d\n", ribbon)

			// Let's find the perimeter of the smallest side.
			perimeter := perimeterOfSmallest(l, w, h)
			ribbon += perimeter
		}
		fmt.Printf("Total Feet of Wrapping Paper: %d\n", total)
		fmt.Printf("Total Feet of Ribbon: %d\n", ribbon)

		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

	} else {
		log.Fatal(err)
	}
}

func perimeterOfSmallest(l, w, h int64) int64 {
	numbers := []int{int(l), int(w), int(h)}
	sortedNumbers := sort.IntSlice(numbers)
	sort.Sort(sortedNumbers)
	perimeter := sortedNumbers[0]*2 + sortedNumbers[1]*2
	//fmt.Printf("Perimeter: %d\n", perimeter)
	return int64(perimeter)
}

func extractSides(l, w, h int64) (int64, int64, int64) {
	side1 := l * w
	side2 := w * h
	side3 := h * l
	return side1, side2, side3
}

func volumeOfPresent(l, w, h int64) int64 {
	volume := l * w * h
	return volume
}

func surfaceAreaPlus(l, w, h int64) int64 {
	side1, side2, side3 := extractSides(l, w, h)
	additionalSide := smallestSide(side1, side2, side3)
	total := 2*side1 + 2*side2 + 2*side3 + additionalSide
	return total
}

func smallestSide(side1, side2, side3 int64) int64 {
	numbers := []int{int(side1), int(side2), int(side3)}
	sortedNumbers := sort.IntSlice(numbers)
	sort.Sort(sortedNumbers)
	//fmt.Printf("Smallest: %d\n", sortedNumbers[0])
	return int64(sortedNumbers[0])
}

func extractSizes(sizes string) (int64, int64, int64) {
	dimensions := strings.Split(sizes, "x")
	length, _ := strconv.ParseInt(dimensions[0], 10, 64)
	width, _ := strconv.ParseInt(dimensions[1], 10, 64)
	height, _ := strconv.ParseInt(dimensions[2], 10, 64)
	return length, width, height
}
