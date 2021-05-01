package application

import "ml-x-men/internal/domain"

var (
	axisX = [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	axisY = [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	words = [4]string{"AAAA", "TTTT", "CCCC", "GGGG"}
)

func (i interactor) IsMutant(matrix [][]byte) (bool, error) {
	isMutant := false

	for _, val := range words {
		isMutant = searchWordInMatrix(matrix, val)
		if isMutant {
			break
		}
	}
	return isMutant, nil
}

func (i interactor) GetByDna(dna string) (*domain.Person, error) {
	return i.storage.PersonGetByDna(dna)

}
func (i interactor) PersonCreate(dna string, isMutant bool) (*domain.Person, error) {
	return i.storage.PersonCreate(dna, isMutant)

}

func searchWordInMatrix(matrix [][]byte, word string) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if searchStringInMatrix(matrix, i, j, word) {
				return true
			}
		}
	}
	return false
}

func searchStringInMatrix(matrix [][]byte, row int, col int, word string) bool {
	if matrix[row][col] != word[0] {
		return false
	}

	wordLength := len(word)

	// search same value in all possible directions
	for i := 0; i < 8; i++ {
		matchedCount := row + axisX[i]
		rowDirection := row + axisX[i]
		columnDirection := col + axisY[i]
		for matchedCount = 1; matchedCount < wordLength; matchedCount++ {
			// break when the variables are out of the matrix
			if rowDirection < 0 || columnDirection < 0 || rowDirection >= len(matrix) || columnDirection >= len(matrix[0]) {
				break
			}
			// the char in the given position doesn't match the searched char
			if matrix[rowDirection][columnDirection] != word[matchedCount] {
				break
			}
			// move to the next position
			rowDirection += axisX[i]
			columnDirection += axisY[i]
		}

		if matchedCount == wordLength {
			return true
		}
	}
	return false
}
