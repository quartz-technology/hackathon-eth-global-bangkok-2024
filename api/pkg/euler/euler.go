package euler

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	"github.com/quartz-technology/hackathon-eth-global-bangkok-2024/pkg/bindings/EVault"
	"github.com/quartz-technology/hackathon-eth-global-bangkok-2024/pkg/bindings/GenericFactory"
	"github.com/quartz-technology/hackathon-eth-global-bangkok-2024/pkg/u"
)

var ray = new(big.Float).SetFloat64(math.Pow(10, 27))

type EulerManager struct {
	client *ethclient.Client

	USDCVaults []common.Address
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

	USDCVaults := make([]common.Address, 0, len(vaultList))
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
			log.Infof("found USDC vault: %v", addr)
			USDCVaults = append(USDCVaults, addr)
		}
	}

	return &EulerManager{
		client:     client,
		USDCVaults: USDCVaults,
	}, nil
}

func (m *EulerManager) GetEVaults() ([]common.Address, error) {
	return m.USDCVaults, nil
}

type EVaultsAPY struct {
	Address common.Address `json:"address"`
	APY     float64        `json:"apy"`
}

func (m *EulerManager) GetEVaultsInfos(ctx context.Context, addrs []common.Address) ([]EVaultsAPY, error) {
	res := make([]EVaultsAPY, 0, len(addrs))
	for _, addr := range addrs {
		evault, err := EVault.NewEVaultCaller(addr, m.client)
		if err != nil {
			return nil, fmt.Errorf("failed to create evault contract: %v", err)
		}

		rate, err := evault.InterestRate(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, fmt.Errorf("failed to get interest rate: %v", err)
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

		fee, err := evault.InterestFee(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, fmt.Errorf("failed to get interest fee: %v", err)
		}

		netAPY := APY * (1 - float64(fee)/10000)

		res = append(res, EVaultsAPY{
			Address: addr,
			APY:     netAPY,
		})
	}

	return res, nil
}
