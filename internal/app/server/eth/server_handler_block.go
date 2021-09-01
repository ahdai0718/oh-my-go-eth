package eth

import (
	"net/http"
	"strconv"

	datastruct "github.com/ahdai0718/oh-my-go-eth/internal/app/server/eth/data_struct"
	"github.com/ahdai0718/oh-my-go-eth/internal/app/server/eth/store"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

//
// @Summary get eth block list
// @Description
// @Accept	json
// @Param	limit	query    	int		true	"return {n} block(s)"
// @Produce	json
// @Success 200 	{string} 	string	"ok"
// @Failure 400 	{string} 	string	"bad request"
// @Failure 404 	{string} 	string	"not found"
// @Router /api/v1/eth/blocks [get]
func ServerHandlerBlockList(ctx *gin.Context) {

	limit, _ := strconv.Atoi(ctx.Query("limit"))

	dataBlockListWrapper := new(datastruct.BlockListWrapper)

	blockList, err := getLatestNBlock(limit)
	if err != nil {
		glog.Error(err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	dataBlockListWrapper.Parse(blockList)

	for _, block := range dataBlockListWrapper.List {
		store.DefaultStorer().AddBlock(&block.Block)
	}

	ctx.JSON(http.StatusOK, dataBlockListWrapper)
}

//
// @Summary get single eth block with specific id
// @Description
// @Accept	json
// @Param	id		path    	int		true	"the eth block id"
// @Produce	json
// @Success 200 	{string} 	string	"ok"
// @Failure 400 	{string} 	string	"bad request"
// @Failure 404 	{string} 	string	"not found"
// @Router /api/v1/eth/blocks/{id} [get]
func ServerHandlerBlock(ctx *gin.Context) {

	id := ctx.Param("id")
	glog.Info("id:", id)

	number, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		glog.Error(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	block, err := getBlockByNumber(number)
	if err != nil {
		glog.Error(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dataBlock := new(datastruct.Block)
	dataBlock.Parse(block)

	ctx.JSON(http.StatusOK, dataBlock)
}
