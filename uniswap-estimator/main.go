package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"math/big"
	"net/http"
	"os"
	"uniswap-estimator/eth"
	"uniswap-estimator/uniswap"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := gin.Default()
	ethURL := os.Getenv("ETH_RPC_URL")
	client, err := eth.NewEthClient(ethURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum node: %v", err)
	}

	router.GET("/estimate", func(c *gin.Context) {
		pool := c.Query("pool")
		src := c.Query("src")
		dst := c.Query("dst")
		srcAmountStr := c.Query("src_amount")

		srcAmount := new(big.Int)
		if _, ok := srcAmount.SetString(srcAmountStr, 10); !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid src_amount"})
			return
		}

		reserveIn, reserveOut, err := client.GetReserves(pool, src, dst)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		dstAmount := uniswap.CalculateAmountOut(srcAmount, reserveIn, reserveOut)

		c.JSON(http.StatusOK, gin.H{"dst_amount": dstAmount.String()})
	})

	err = router.Run(":1337")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
