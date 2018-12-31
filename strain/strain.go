package strain

type Ints []int
type Lists [][]int
type Strings []string

func (sliceNum Ints) Keep(pred func(int) bool) Ints {
	if sliceNum == nil {
		return nil
	}

	filteredNum := make(Ints, 0)
	for _, num := range sliceNum {
		if pred(num) {
			filteredNum = append(filteredNum, num)
		}
	}

	return filteredNum
}

func (sliceNum Ints) Discard(pred func(int) bool) Ints {
	if sliceNum == nil {
		return nil
	}
	filteredNum := make(Ints, 0)
	for _, num := range sliceNum {
		if !pred(num) {
			filteredNum = append(filteredNum, num)
		}
	}

	return filteredNum
}

func (sliceOfslice Lists) KeepLists(pred func([]int) bool) Lists {
	if sliceOfslice == nil {
		return nil
	}
	filteredSlice := make(Lists, 0)
	for _, slice := range sliceOfslice {
		if pred(slice) {
			filteredSlice = append(filteredSlice, slice)
		}
	}

	return filteredSlice
}

func (slice Strings) KeepStrings(pred func(string) bool) Strings {
	if slice == nil {
		return nil
	}
	filteredStrings := make(Strings, 0)
	for _, string := range slice {
		if pred(string) {
			filteredStrings = append(filteredStrings, string)
		}
	}

	return filteredStrings

}
