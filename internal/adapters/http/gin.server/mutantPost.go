package gin_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type mutantPostRequest struct {
	Dna []string `json:"dna" binding:"required"`
}

func (rH RouterHandler) mutantPost(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	body := &mutantPostRequest{}
	if err := c.BindJSON(body); err != nil {
		log(err)
		c.Status(http.StatusBadRequest)
		return
	}
	matrixDna := buildMatrixDna(body.Dna)
	isMutant, err := rH.ucHandler.IsMutant(matrixDna)
	if err != nil {
		log(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if isMutant {
		c.JSON(http.StatusOK, gin.H{"message": "The given human is a mutant"})
	} else {
		c.Status(http.StatusForbidden)
		return
	}
}
func buildMatrixDna(dna []string) [][]byte {
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
