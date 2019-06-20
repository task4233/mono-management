package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
SearchMonos は, メソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[POST]   /api/v1/search/
*/
func SearchMonos(c *gin.Context) {
	// リクエストに含まれる検索文字列をitemsテーブルから取得
	// データを返却
	c.JSON(http.StatusOK, "search")
}
