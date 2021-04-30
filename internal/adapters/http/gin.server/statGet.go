package gin_server

import (
	"github.com/gin-gonic/gin"
	"ml-x-men/internal/adapters/json_formatter"
	"net/http"
)

func (rH RouterHandler) statsGet(c *gin.Context){
	log := rH.log(rH.MethodAndPath(c))
	mutantCount,humanCount,ratio, err := rH.ucHandler.StatsGetMutantVsHuman()

	if err != nil {
		log(err, "Error getting stats")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK,json_formatter.NewStatResp(mutantCount,humanCount,ratio))
}