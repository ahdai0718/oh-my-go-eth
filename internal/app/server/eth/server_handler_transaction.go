package eth

import (
	"net/http"

	"github.com/ahdai0718/oh-my-go-eth/internal/app/server/eth/store"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

//
// @Summary get single eth transaction with hash
// @Description
// @Accept	json
// @Param	tx_hash		path    	string		true	"eth transaction hash"
// @Produce	json
// @Success 200 		{string} 	string	"ok"
// @Failure 400 		{string} 	string	"bad request"
// @Failure 404 		{string} 	string	"not found"
// @Router /api/v1/eth/transaction/{tx_hash} [get]
func ServerHandlerTransaction(ctx *gin.Context) {
	txHash := ctx.Param("tx_hash")

	transaction, err := store.DefaultStorer().GetTransactionByHash(txHash)
	if err != nil {
		glog.Error(err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	transactionLogList, err := store.DefaultStorer().GetTransactionLogListByBlockTxHash(transaction.TxHash)
	if err != nil {
		glog.Error(err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	transaction.Logs = transactionLogList

	ctx.JSON(http.StatusOK, transaction)
}
