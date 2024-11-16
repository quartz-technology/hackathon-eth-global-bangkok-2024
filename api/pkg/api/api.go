package api

import (
	"context"
	"os"
	"slices"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/quartz-technology/hackathon-eth-global-bangkok-2024/pkg/euler"
	"github.com/quartz-technology/hackathon-eth-global-bangkok-2024/pkg/u"
	log "github.com/sirupsen/logrus"
)

var (
	rebalancingUser []common.Address
	addrToHex       = func(addr common.Address) string { return addr.String() }
)

func Routes(r *gin.Engine, em *euler.EulerManager, mut *sync.Mutex) {
	r.GET("/markets", func(c *gin.Context) {
		wallet := c.Query("wallet")
		if wallet == "" {
			c.JSON(400, gin.H{"error": "wallet address is required"})
			return
		}

		vaultsInfos, err := em.GetEVaultsInfos(c, em.USDCVaults, common.HexToAddress(wallet))
		if err != nil {
			log.Errorf("failed to get vaults infos: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, vaultsInfos)
		return
	})

	r.GET("/operator", func(c *gin.Context) {
		c.JSON(200, common.HexToAddress(os.Getenv("EURLER_OPERATOR")))
		return
	})

	r.POST("/rebalancing", func(c *gin.Context) {
		log.Infof("rebalancing user: %v", c.Query("wallet"))
		wallet := c.Query("wallet")
		run, err := strconv.ParseBool(c.Query("run"))
		if err != nil {
			log.Errorf("failed to parse run parameter: %v", err)
			c.JSON(400, gin.H{"error": "valid run parameter is required "})
			return
		}

		if !run {
			mut.Lock()
			rebalancingUser = slices.DeleteFunc(rebalancingUser, func(addr common.Address) bool {
				if addr == common.HexToAddress(wallet) {
					return true
				}
				return false
			})
			mut.Unlock()
			c.JSON(200, gin.H{"message": "user no longer rebalancing"})
			return
		}

		addrToHex := func(addr common.Address) string { return addr.String() }
		if !slices.Contains(u.Map(rebalancingUser, addrToHex), common.HexToAddress(wallet).String()) {
			mut.Lock()
			rebalancingUser = append(rebalancingUser, common.HexToAddress(wallet))
			mut.Unlock()
		}

		c.JSON(200, gin.H{"message": "user rebalancing"})
		return
	})

	r.GET("/rebalancing", func(c *gin.Context) {
		walletS := c.Query("wallet")
		wallet := common.HexToAddress(walletS)

		log.Infof("checking if user is rebalancing: %v %v", rebalancingUser, wallet)
		if slices.Contains(u.Map(rebalancingUser, addrToHex), wallet.String()) {
			c.JSON(200, gin.H{"rebalancing": true})
			return
		}

		c.JSON(200, gin.H{"rebalancing": false})
		return
	})
}

func Rebalance(ctx context.Context, em *euler.EulerManager, mut *sync.Mutex) {
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			mut.Lock()
			log.Infof("rebalancing user: %v", rebalancingUser)
			for _, wallet := range rebalancingUser {
				vaults, err := em.GetEVaultsInfos(ctx, em.USDCVaults, wallet)
				if err != nil {
					log.Errorf("failed to get vaults infos: %v", err)
					continue
				}
				slices.SortFunc(vaults, func(i, j euler.EVaultsAPY) int {
					if i.APY > j.APY {
						return -1
					} else if i.APY < j.APY {
						return 1
					}
					return 0
				})

				for _, vault := range vaults {
					if vault.Address == vaults[0].Address || vault.Balance == "0" {
						continue
					}

					log.Infof("rebalancing %v from %v to %v", wallet, vault.Address, vaults[0].Address)
					if err := em.Rebalance(ctx, wallet, vault.Address, vaults[0].Address); err != nil {
						log.Errorf("failed to rebalance: %v", err)
					}

					time.Sleep(15 * time.Second)
				}

			}
			mut.Unlock()
		}
	}
}
