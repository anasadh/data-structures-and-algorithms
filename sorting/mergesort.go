// In computer science, merge sort (also commonly spelled as mergesort) is an efficient, general-purpose,
// and comparison-based sorting algorithm. Most implementations produce a stable sort,
// which means that the order of equal elements is the same in the input and output.
// Merge sort is a divide-and-conquer algorithm that was invented by John von Neumann in 1945.
// A detailed description and analysis of bottom-up merge sort appeared in a report by Goldstine and von Neumann as early as 1948.
// Conceptually, a merge sort works as follows:

// Divide the unsorted list into n sublists, each containing one element (a list of one element is considered sorted).
// Repeatedly merge sublists to produce new sorted sublists until there is only one sublist remaining. This will be the sorted list
// Source(https://en.wikipedia.org/wiki/Merge_sort)

package main

func MergeSort(Arr []int) []int {
	if len(Arr) <= 1 {
		return Arr
	}
	middle := len(Arr) / 2
	left := MergeSort(A[:middle])
	right := MergeSort(A[middle:])
	return Merge(left, right)
}