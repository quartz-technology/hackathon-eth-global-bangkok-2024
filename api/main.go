package main

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/quartz-technology/hackathon-eth-global-bangkok-2024/pkg/api"
	"github.com/quartz-technology/hackathon-eth-global-bangkok-2024/pkg/euler"
)

func main() {
	_ = godotenv.Load(".env")

	ctx := context.Background()

	r := gin.New()

	client, err := ethclient.DialContext(ctx, os.Getenv("ETH_RPC_URL"))
	if err != nil {
		log.Fatalf("failed to connect to the Ethereum client: %v", err)
	}

	em, err := euler.NewEulerManager(
		ctx,
		client,
		common.HexToAddress(os.Getenv("EULER_VAULT_FACTORY")),
		common.HexToAddress(os.Getenv("USDC_ADDRESS")),
	)
	if err != nil {
		log.Fatalf("failed to create Euler manager: %v", err)
	}

	api.Routes(r, em)

	r.Run(":8080")
}