package euler

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/quartz-technology/hackathon-eth-global-bangkok-2024/pkg/bindings/EVault"
	"github.com/quartz-technology/hackathon-eth-global-bangkok-2024/pkg/bindings/GenericFactory"
	"github.com/quartz-technology/hackathon-eth-global-bangkok-2024/pkg/u"
)

var ray = new(big.Float).SetFloat64(math.Pow(10, 27))

type VaultsMeta struct {
	address common.Address
	symbol  string
}

type EulerManager struct {
	client *ethclient.Client

	USDCVaults []VaultsMeta
}

var MAX_UINT256, _ = new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)

func NewEulerManager(ctx context.Context, client *ethclient.Client, vaultFactoryAddr common.Address, USDC common.Address) (*EulerManager, error) {
	vaultFactory, err := GenericFactory.NewGenericFactortCaller(vaultFactoryAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create vault factory contract: %v", err)
	}

	vaultList, err := vaultFactory.GetProxyListSlice(&bind.CallOpts{Context: ctx}, new(big.Int), MAX_UINT256)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault list: %v", err)
	}

	USDCVaults := make([]VaultsMeta, 0, len(vaultList))
	for _, addr := range vaultList {
		evault, err := EVault.NewEVaultCaller(addr, client)
		if err != nil {
			return nil, fmt.Errorf("failed to create evault contract: %v", err)
		}

		asset, err := evault.Asset(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, fmt.Errorf("failed to get asset: %v", err)
		}

		if asset == USDC {
			symbol, err := evault.Symbol(&bind.CallOpts{Context: ctx})
			if err != nil {
				return nil, fmt.Errorf("failed to get symbol: %v", err)
			}

			log.Infof("found USDC vault: %v - %v", addr, symbol)
			USDCVaults = append(USDCVaults, VaultsMeta{
				address: addr,
				symbol:  symbol,
			})
		}
	}

	return &EulerManager{
		client:     client,
		USDCVaults: USDCVaults,
	}, nil
}

func (m *EulerManager) GetEVaults() ([]VaultsMeta, error) {
	return m.USDCVaults, nil
}

type EVaultsAPY struct {
	Address common.Address `json:"address"`
	APY     float64        `json:"apy"`
	Symbol  string         `json:"symbol"`
	Balance string         `json:"balance"`
}

func (m *EulerManager) GetEVaultsInfos(ctx context.Context, vaults []VaultsMeta, wallet common.Address) ([]EVaultsAPY, error) {
	res := make([]EVaultsAPY, len(vaults))

	eg, egCtx := errgroup.WithContext(ctx)
	mut := &sync.Mutex{}

	for idx, vault := range vaults {
		eg.Go(func() error {
			evault, err := EVault.NewEVaultCaller(vault.address, m.client)
			if err != nil {
				return fmt.Errorf("failed to create evault contract: %v", err)
			}

			rate, err := evault.InterestRate(&bind.CallOpts{Context: egCtx})
			if err != nil {
				return fmt.Errorf("failed to get interest rate: %v", err)
			}

			supplyAPR := big.NewFloat(0).Mul(
				big.NewFloat(0).Quo(
					new(big.Float).SetInt(rate),
					ray,
				),
				big.NewFloat(u.SECONDS_PER_YEAR),
			)
			fAPR, _ := supplyAPR.Float64()

			APY := u.ToAPY(fAPR)

			fee, err := evault.InterestFee(&bind.CallOpts{Context: egCtx})
			if err != nil {
				return fmt.Errorf("failed to get interest fee: %v", err)
			}

			netAPY := APY * (1 - float64(fee)/10000)

			balanceShare, err := evault.BalanceOf(&bind.CallOpts{Context: egCtx}, wallet)
			if err != nil {
				return fmt.Errorf("failed to get balance: %v", err)
			}
			balance, err := evault.ConvertToAssets(&bind.CallOpts{Context: egCtx}, balanceShare)
			if err != nil {
				return fmt.Errorf("failed to convert to assets: %v", err)
			}

			mut.Lock()
			res[idx] = EVaultsAPY{
				Address: vault.address,
				APY:     netAPY,
				Symbol:  vault.symbol,
				Balance: balance.String(),
			}
			mut.Unlock()

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, fmt.Errorf("failed to get evaults infos: %v", err)
	}

	return res, nil
}
