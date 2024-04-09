package Collection

import "fmt"

func Slicing() {

	var sampleSlice = make([]string, 5, 10)

	sampleSlice[0] = "Sri"
	sampleSlice[1] = "Krishna"
	sampleSlice[3] = "Devaraya"
	sampleSlice = append(sampleSlice, "Nagadevaraya")
	fmt.Println(sampleSlice)

	// Merging two Slices

	odd := []int{1, 3, 5, 7, 9}
	even := []int{2, 4, 6, 8, 10}

	// Append will not create a new slice but it creates a view onto the original slice.
	// Also changes made to the sub-slices will affect the original slice.
	completeSet := append(odd, even...) // ... after a variable means unpacking them
	fmt.Println(completeSet)

	// to create a new independent sub-slice
	// make a new slice of length equal to subslice
	// copy elements from original slice
	newSubSlice := make([]string, len(sampleSlice[:3]))
	copy(newSubSlice, sampleSlice[:3])

	// removing an element from the SLice
	removeAtIndex := 3
	sampleSlice = append(sampleSlice[:removeAtIndex], sampleSlice[removeAtIndex+1:]...)
	fmt.Println(sampleSlice)
}
