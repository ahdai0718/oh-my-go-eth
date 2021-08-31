package eth

import (
	"context"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

//
// @Summary get eth block list
// @Description get eth block list
// @Accept	json
// @Param	limit	query    	int		true	"return {n} block(s)"
// @Produce	json
// @Success 200 	{string} 	string	"ok"
// @Failure 400 	{string} 	string	"bad request"
// @Failure 404 	{string} 	string	"not found"
// @Router /api/v1/eth/blocks [get]
func ServerHandlerBlockList(ctx *gin.Context) {
	client, err := ethclient.Dial(dataSeedURL)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	printBlockInfo(block)

	ctx.JSON(http.StatusOK, block)

}

func ServerHandlerBlock(ctx *gin.Context) {
	client, err := ethclient.Dial(dataSeedURL)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	glog.Info(block)
}
