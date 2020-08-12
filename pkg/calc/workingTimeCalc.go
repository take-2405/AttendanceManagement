package calc

import (
	"github.com/gin-gonic/gin"
	"log"
)

func WorkingTimeCalc()gin.HandlerFunc {
	return func(c *gin.Context) {
		// リクエストヘッダからX-API-TOKEN(認証トークン)を取得
		key := "X-API-TOKEN"
		token := c.GetHeader(key)
		if len(token) == 0 {
			log.Println("X-API-TOKEN is empty")
			return
		}

			//データを持ってくる
	}
}