package main

import (
	"fmt"
	"sort"

	"github.com/Werner-Prbk/my-go-playground/pkg/aoc"
)

type joltAdapter struct {
	jolts         int
	inUse         bool
	offsetToNext  int
	possibilities int
	combinations  int
}

type joltAdapterList []joltAdapter

func toJoltAdapters(s []int) joltAdapterList {
	var ja = make(joltAdapterList, 0, len(s))
	for _, v := range s {
		ja = append(ja,
			joltAdapter{
				jolts:         v,
				inUse:         false,
				offsetToNext:  -1,
				possibilities: -1,
				combinations:  -1,
			})
	}
	return ja
}

func orderForUsage(jal joltAdapterList, lastAdapterOffset int) bool {
	var i = 0
	for ; i < len(jal)-1; i++ {
		jal[i].inUse = true
		jal[i].offsetToNext = jal[i+1].jolts - jal[i].jolts

		if jal[i].offsetToNext > 3 ||
			jal[i].offsetToNext < 0 {
			return false
		}
	}

	// last "adapter" is also in use
	jal[i].inUse = true
	jal[i].offsetToNext = 0

	return true
}

func sumPossibilitiesRecursive(jal joltAdapterList, index int) int {
	if index >= len(jal) {
		return 1
	}

	// partial solution is already available
	if jal[index].combinations != -1 {
		return jal[index].combinations
	}

	var sum = 0

	for i := 1; i <= jal[index].possibilities; i++ {
		sum += sumPossibilitiesRecursive(jal, index+i)
	}
	jal[index].combinations = sum
	return jal[index].combinations
}

func countPossibleArrangements(jal joltAdapterList) int {
	var i = 0
	for ; i < len(jal)-3; i++ {
		jal[i].possibilities = 0
		for j := 1; j <= 3; j++ {
			if jal[i+j].jolts-jal[i].jolts <= 3 {
				jal[i].possibilities++
			}
		}
	}

	// treat "rest"
	if jal[i+2].jolts-jal[i].jolts <= 3 {
		jal[i].possibilities = 2
	} else {
		jal[i].possibilities = 1
	}

	// keep data it clean
	jal[i+1].possibilities = 1
	jal[i+2].possibilities = 1

	return sumPossibilitiesRecursive(jal, 0)
}

func main() {
	var str, errRd = aoc.ReadTextFileAllLines("input.txt")
	aoc.EnsureNoErrorOrPanic(errRd)
	var nums, errConv = aoc.SliceConvertStringToInt(str)
	aoc.EnsureNoErrorOrPanic(errConv)

	// consider chaging outlet (0 jolts) as 1st adapter!
	nums = append(nums, 0)
	sort.Ints(nums)

	// consider built-in adapter also as adapter
	nums = append(nums, nums[len(nums)-1]+3)

	var ja = toJoltAdapters(nums)

	var ok = orderForUsage(ja, 3)

	if !ok {
		fmt.Print("Error while sorting!")
		return
	}

	var offset1, offset3 = 0, 0
	for _, v := range ja {
		if v.offsetToNext == 1 {
			offset1++
		} else if v.offsetToNext == 3 {
			offset3++
		}
	}

	fmt.Printf("Th solution is %v\n", offset1*offset3)

	// part 2
	var possibilities = countPossibleArrangements(ja)

	fmt.Printf("Possible combinations: %v\n", possibilities)

}
