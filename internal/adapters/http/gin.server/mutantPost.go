package gin_server

import (
	"github.com/gin-gonic/gin"
	"ml-x-men/internal/adapters/json_formatter"
	"net/http"
	"strings"
)

type mutantPostRequest struct {
	Dna []string `json:"dna" binding:"required"`
}

func (rH RouterHandler) mutantPost(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	body := &mutantPostRequest{}
	if err := c.BindJSON(body); err != nil {
		log(err, "Error parsing request body")
		c.Status(http.StatusBadRequest)
		return
	}
	dnaAsString := strings.Join(body.Dna, ",")
	person, err := rH.ucHandler.GetByDna(dnaAsString)
	if err != nil {
		log(err, "Error getting person by dna")
		c.Status(http.StatusInternalServerError)
		return
	} else if person == nil {
		matrixDna := buildMatrixDna(body.Dna)
		isMutant, err := rH.ucHandler.IsMutant(matrixDna)
		if err != nil {
			log(err, "Error evaluating IsMutant")
			c.Status(http.StatusInternalServerError)
			return
		}

		person, err = rH.ucHandler.PersonCreate(dnaAsString, isMutant)
		if err != nil {
			log(err, "Error creating person")
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	if person.IsMutant {
		c.JSON(http.StatusOK, gin.H{"person": json_formatter.NewPersonResp(*person)})
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
