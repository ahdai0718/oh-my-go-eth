package eth

import (
	"context"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func ServerHandlerTransaction(ctx *gin.Context) {
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
