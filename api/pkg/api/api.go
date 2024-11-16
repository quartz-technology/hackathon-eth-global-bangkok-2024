package api

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/quartz-technology/hackathon-eth-global-bangkok-2024/pkg/euler"
	log "github.com/sirupsen/logrus"
)

func Routes(r *gin.Engine, em *euler.EulerManager) {
	r.GET("/markets", func(c *gin.Context) {
		evaults, err := em.GetEVaults()
		if err != nil {
			log.Errorf("failed to get evaults: %v", err)
			c.JSON(500, gin.H{"error": "failed to get evaults"})
		}

		vaultsInfos, err := em.GetEVaultsInfos(c, evaults)
		c.JSON(200, vaultsInfos)
		return
	})

	r.GET("/operator", func(c *gin.Context) {
		c.JSON(200, common.HexToAddress(os.Getenv("EURLER_OPERATOR")))
		return
	})
}
