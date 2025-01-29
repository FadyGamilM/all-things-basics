# all-things-basics
All-Things-Basics is simply a repo where I apply some basic stuff studying algorithms, data strcutres, system design, maybe some cs concepts to be always sharp

# when Can we customize the `binary-search` algorithm ? 
- Probelm Description : 
> you have array of stairs, and some crystals, and at a given level of the stairs, the crystals will start to crash, so can you find the first stair that the crystal will start to crash from, but do it as optimal as you can?

- Answer : 
so the shape of the input is array of booleans once we found a bool true, this is the stair where any crystal crash from it and moving forward.
> - if we did a linear search this is o(n)
> - if we did binary search, we will find a true point, but we might have a true points before it, so we will have to go backword untill we hit the 1st true point and this will be o(n) + o(logn) = o(n)
> - we can do binary search jumbs but with o(sqrt(n)) and if we found a true point, or we break from the array, we can just go backwrd only one jumb = sqrt(n) indexes in the array to find the first one so the complexity will be o(sqrt(N)) + o(sqrt(N)) = 

code : 
```go 

// i will apply the binary search (customized) because it is a sorted array (falses then true)
func findFirstStairIndexWhereCrystalsStartCrashing(stairs []bool) int {
	jumbSize := int(math.Floor(math.Sqrt(float64(len(stairs)))))
	var i = 0
	for ; i < len(stairs); i += jumbSize {
		if stairs[i] == true {
			break
		}
	}

	// once we broke the searching sqrt loop / or we ended the entire jumbs without finding true position, we can go backword only one jumb to check if we have a true poistions or not
	for j := i - jumbSize; j < i; j++ {
		if stairs[j] == true {
			return j
		}
	}

	return -1
}
```

testing : 
```go
	log.Println(
		findFirstStairIndexWhereCrystalsStartCrashing([]bool{
			false, false, false, false,
			false, false, false, false,
			false, false, true, true,
			true, true, true, true,
		}),
	)
```