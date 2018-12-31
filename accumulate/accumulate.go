package accumulate

//Accumulate given a string collection and an operation to perform on each element of the collection,
//returns a new collection containing the result of applying that operation to each element of
//the input collection
func Accumulate(sequence []string, fn func(string) string) []string {
	finalSequence := make([]string, 0, len(sequence))
	for _, element := range sequence {

		finalSequence = append(finalSequence, fn(element))
	}

	return finalSequence
}
