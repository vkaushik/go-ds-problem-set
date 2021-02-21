package p1

// Problem:
// Given a string S, find length of the longest substring with all distinct characters.
// For example, for input "abca", the output is 3 as "abc" is the longest substring with all distinct characters.

// Solution:
// LenOfLongestSubStWithUniqueChars returns length of the longest substring with all distinct characters.
// TimeComplexity: O(n), SpaceComplexity: O(n), where n is the length of input string
// Take two indexes named lefthHand and rightHand. The invariant property is that: "The characters between leftHand and rightHand are always unique." 
// Both leftHand and rightHand keep moving from left to right.
// The moment righHand finds a char that already exists between leftHand and rightHand, discard all the chars till the position just next of the duplicate char,
// i.e. move left hand just after the previous location where the duplicate char (of the char at rightHand).
func LenOfLongestSubStWithUniqueChars(str string) int {
	l := len(str)
	if l <= 1 {
		return l
	}

	charToIndex := make(map[rune]int)
	indexToChar := []rune(str)
	leftHand, rightHand := 0, 0
	maxSubStrLen, currentSubStrLen := 0, 0

	for rightHand < l {
		ch := indexToChar[rightHand]
		if duplicateCharIndex, isDuplicate := charToIndex[ch]; isDuplicate { // it's a repeating char
			for leftHand < duplicateCharIndex { // remove all chars till the duplicate chars from map
				staleChar := indexToChar[leftHand]
				delete(charToIndex, staleChar)
				leftHand++
			}
			leftHand++ // new leftHand index is just after the repeated char
		}
		charToIndex[ch] = rightHand

		currentSubStrLen = rightHand - leftHand + 1
		if currentSubStrLen > maxSubStrLen {
			maxSubStrLen = currentSubStrLen
		}
		rightHand++
	}

	return maxSubStrLen
}
