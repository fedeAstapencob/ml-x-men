package application

import "ml-x-men/internal/domain"

var (
	axisX = [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	axisY = [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
)

func (i interactor) IsMutant(matrix [][]byte) (bool, error) {
	isMutant := searchWordInMatrix(matrix)
	return isMutant, nil
}

func (i interactor) GetByDna(dna string) (*domain.Person, error) {
	return i.storage.PersonGetByDna(dna)

}
func (i interactor) PersonCreate(dna string, isMutant bool) (*domain.Person, error) {
	return i.storage.PersonCreate(dna, isMutant)

}

func searchWordInMatrix(matrix [][]byte) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if searchStringInMatrix(matrix, i, j, matrix[i][j], 4) {
				return true
			}
		}
	}
	return false
}

func searchStringInMatrix(matrix [][]byte, row int, col int, letter byte, searchLength int) bool {

	// search same value in all possible directions
	for i := 0; i < 8; i++ {
		matchedCount := row + axisX[i]
		rowDirection := row + axisX[i]
		columnDirection := col + axisY[i]
		for matchedCount = 1; matchedCount < searchLength; matchedCount++ {
			// break when the variables are out of the matrix
			if rowDirection < 0 || columnDirection < 0 || rowDirection >= len(matrix) || columnDirection >= len(matrix[0]) {
				break
			}
			// the char in the given position doesn't match the searched char
			if matrix[rowDirection][columnDirection] != letter {
				break
			}
			// move to the next position
			rowDirection += axisX[i]
			columnDirection += axisY[i]
		}

		if matchedCount == searchLength {
			return true
		}
	}
	return false
}
