package utils

func BuildMatrixDna(dna []string) [][]byte {
	var matrix [][]byte
	for i := 0; i < len(dna); i++ {
		var charsValue []byte
		dnaChars := dna[i]
		for _, char := range dnaChars {
			charsValue = append(charsValue, byte(char))
		}
		matrix = append(matrix, charsValue)
	}
	return matrix
}
