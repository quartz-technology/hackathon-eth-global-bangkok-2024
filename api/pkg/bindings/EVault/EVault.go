// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EVault

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BaseIntegrations is an auto generated low-level Go binding around an user-defined struct.
type BaseIntegrations struct {
	Evc              common.Address
	ProtocolConfig   common.Address
	SequenceRegistry common.Address
	BalanceTracker   common.Address
	Permit2          common.Address
}

// DispatchDeployedModules is an auto generated low-level Go binding around an user-defined struct.
type DispatchDeployedModules struct {
	Initialize       common.Address
	Token            common.Address
	Vault            common.Address
	Borrowing        common.Address
	Liquidation      common.Address
	RiskManager      common.Address
	BalanceForwarder common.Address
	Governance       common.Address
}

// EVaultMetaData contains all meta data concerning the EVault contract.
var EVaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"integrations\",\"type\":\"tuple\",\"internalType\":\"structBase.Integrations\",\"components\":[{\"name\":\"evc\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"protocolConfig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sequenceRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balanceTracker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permit2\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"modules\",\"type\":\"tuple\",\"internalType\":\"structDispatch.DeployedModules\",\"components\":[{\"name\":\"initialize\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"borrowing\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"liquidation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"riskManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balanceForwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"governance\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"EVC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LTVBorrow\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LTVFull\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"borrowLTV\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"liquidationLTV\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"initialLiquidationLTV\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"targetTimestamp\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"rampDuration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LTVLiquidation\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LTVList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MODULE_BALANCE_FORWARDER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MODULE_BORROWING\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MODULE_GOVERNANCE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MODULE_INITIALIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MODULE_LIQUIDATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MODULE_RISKMANAGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MODULE_TOKEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MODULE_VAULT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accountLiquidity\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"liquidation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"collateralValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"liabilityValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accountLiquidityFull\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"liquidation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"collaterals\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"collateralValues\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"liabilityValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accumulatedFeesAssets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"holder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"asset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceForwarderEnabled\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceTrackerAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"borrow\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"caps\",\"inputs\":[],\"outputs\":[{\"name\":\"supplyCap\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"borrowCap\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkAccountStatus\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collaterals\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkLiquidation\",\"inputs\":[{\"name\":\"liquidator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"violator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"maxRepay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxYield\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkVaultStatus\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configFlags\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"convertFees\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convertToAssets\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"convertToShares\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"creator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"debtOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"debtOfExact\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableBalanceForwarder\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableController\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableBalanceForwarder\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeReceiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"flashLoan\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"governorAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hookConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"proxyCreator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"interestAccumulator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"interestFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"interestRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"interestRateModel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidate\",\"inputs\":[{\"name\":\"violator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"repayAssets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minYieldBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"liquidationCoolOffTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxDeposit\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxLiquidationDiscount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxMint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxRedeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxWithdraw\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit2Address\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewDeposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewMint\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewRedeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewWithdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolConfigAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolFeeReceiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolFeeShare\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pullDebt\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"repay\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"repayWithShares\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCaps\",\"inputs\":[{\"name\":\"supplyCap\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"borrowCap\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setConfigFlags\",\"inputs\":[{\"name\":\"newConfigFlags\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeReceiver\",\"inputs\":[{\"name\":\"newFeeReceiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGovernorAdmin\",\"inputs\":[{\"name\":\"newGovernorAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setHookConfig\",\"inputs\":[{\"name\":\"newHookTarget\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newHookedOps\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInterestFee\",\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInterestRateModel\",\"inputs\":[{\"name\":\"newModel\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLTV\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"borrowLTV\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"liquidationLTV\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rampDuration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLiquidationCoolOffTime\",\"inputs\":[{\"name\":\"newCoolOffTime\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxLiquidationDiscount\",\"inputs\":[{\"name\":\"newDiscount\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"skim\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalAssets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalBorrows\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalBorrowsExact\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"touch\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFromMax\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unitOfAccount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"viewDelegate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BalanceForwarderStatus\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Borrow\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConvertFees\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"protocolReceiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"governorReceiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"protocolShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"governorShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DebtSocialized\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EVaultCreated\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovSetCaps\",\"inputs\":[{\"name\":\"newSupplyCap\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newBorrowCap\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovSetConfigFlags\",\"inputs\":[{\"name\":\"newConfigFlags\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovSetFeeReceiver\",\"inputs\":[{\"name\":\"newFeeReceiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovSetGovernorAdmin\",\"inputs\":[{\"name\":\"newGovernorAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovSetHookConfig\",\"inputs\":[{\"name\":\"newHookTarget\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newHookedOps\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovSetInterestFee\",\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovSetInterestRateModel\",\"inputs\":[{\"name\":\"newInterestRateModel\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovSetLTV\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"borrowLTV\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"liquidationLTV\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"initialLiquidationLTV\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"targetTimestamp\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"rampDuration\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovSetLiquidationCoolOffTime\",\"inputs\":[{\"name\":\"newCoolOffTime\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovSetMaxLiquidationDiscount\",\"inputs\":[{\"name\":\"newDiscount\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InterestAccrued\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Liquidate\",\"inputs\":[{\"name\":\"liquidator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"violator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"collateral\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"repayAssets\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"yieldBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PullDebt\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Repay\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultStatus\",\"inputs\":[{\"name\":\"totalShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"totalBorrows\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"accumulatedFees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"cash\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"interestAccumulator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"interestRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"E_AccountLiquidity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_AmountTooLargeToEncode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_BadAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_BadAssetReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_BadBorrowCap\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_BadCollateral\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_BadFee\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_BadMaxLiquidationDiscount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_BadSharesOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_BadSharesReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_BadSupplyCap\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_BorrowCapExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_CheckUnauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_CollateralDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_ConfigAmountTooLargeToEncode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_ControllerDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_DebtAmountTooLargeToEncode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_EmptyError\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_ExcessiveRepayAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_FlashLoanNotRepaid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_Initialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_InsufficientAllowance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_InsufficientAssets\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_InsufficientCash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_InsufficientDebt\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_InvalidLTVAsset\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_LTVBorrow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_LTVLiquidation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_LiquidationCoolOff\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_MinYield\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_NoLiability\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_NoPriceOracle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_NotController\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_NotHookTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_NotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_OperationDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_OutstandingDebt\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_ProxyMetadata\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_RepayTooMuch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_SelfLiquidation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_SelfTransfer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_SupplyCapExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_TransientState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_ViolatorLiquidityDeferred\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_ZeroAssets\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"E_ZeroShares\",\"inputs\":[]}]",
}

// EVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use EVaultMetaData.ABI instead.
var EVaultABI = EVaultMetaData.ABI

// EVault is an auto generated Go binding around an Ethereum contract.
type EVault struct {
	EVaultCaller     // Read-only binding to the contract
	EVaultTransactor // Write-only binding to the contract
	EVaultFilterer   // Log filterer for contract events
}

// EVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type EVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EVaultSession struct {
	Contract     *EVault           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EVaultCallerSession struct {
	Contract *EVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EVaultTransactorSession struct {
	Contract     *EVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type EVaultRaw struct {
	Contract *EVault // Generic contract binding to access the raw methods on
}

// EVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EVaultCallerRaw struct {
	Contract *EVaultCaller // Generic read-only contract binding to access the raw methods on
}

// EVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EVaultTransactorRaw struct {
	Contract *EVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEVault creates a new instance of EVault, bound to a specific deployed contract.
func NewEVault(address common.Address, backend bind.ContractBackend) (*EVault, error) {
	contract, err := bindEVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVault{EVaultCaller: EVaultCaller{contract: contract}, EVaultTransactor: EVaultTransactor{contract: contract}, EVaultFilterer: EVaultFilterer{contract: contract}}, nil
}

// NewEVaultCaller creates a new read-only instance of EVault, bound to a specific deployed contract.
func NewEVaultCaller(address common.Address, caller bind.ContractCaller) (*EVaultCaller, error) {
	contract, err := bindEVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVaultCaller{contract: contract}, nil
}

// NewEVaultTransactor creates a new write-only instance of EVault, bound to a specific deployed contract.
func NewEVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*EVaultTransactor, error) {
	contract, err := bindEVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVaultTransactor{contract: contract}, nil
}

// NewEVaultFilterer creates a new log filterer instance of EVault, bound to a specific deployed contract.
func NewEVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*EVaultFilterer, error) {
	contract, err := bindEVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVaultFilterer{contract: contract}, nil
}

// bindEVault binds a generic wrapper to an already deployed contract.
func bindEVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EVault *EVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVault.Contract.EVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EVault *EVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVault.Contract.EVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EVault *EVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVault.Contract.EVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EVault *EVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EVault *EVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EVault *EVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVault.Contract.contract.Transact(opts, method, params...)
}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EVault *EVaultCaller) EVC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "EVC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EVault *EVaultSession) EVC() (common.Address, error) {
	return _EVault.Contract.EVC(&_EVault.CallOpts)
}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EVault *EVaultCallerSession) EVC() (common.Address, error) {
	return _EVault.Contract.EVC(&_EVault.CallOpts)
}

// LTVBorrow is a free data retrieval call binding the contract method 0xbf58094d.
//
// Solidity: function LTVBorrow(address collateral) view returns(uint16)
func (_EVault *EVaultCaller) LTVBorrow(opts *bind.CallOpts, collateral common.Address) (uint16, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "LTVBorrow", collateral)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LTVBorrow is a free data retrieval call binding the contract method 0xbf58094d.
//
// Solidity: function LTVBorrow(address collateral) view returns(uint16)
func (_EVault *EVaultSession) LTVBorrow(collateral common.Address) (uint16, error) {
	return _EVault.Contract.LTVBorrow(&_EVault.CallOpts, collateral)
}

// LTVBorrow is a free data retrieval call binding the contract method 0xbf58094d.
//
// Solidity: function LTVBorrow(address collateral) view returns(uint16)
func (_EVault *EVaultCallerSession) LTVBorrow(collateral common.Address) (uint16, error) {
	return _EVault.Contract.LTVBorrow(&_EVault.CallOpts, collateral)
}

// LTVFull is a free data retrieval call binding the contract method 0x33708d0c.
//
// Solidity: function LTVFull(address collateral) view returns(uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EVault *EVaultCaller) LTVFull(opts *bind.CallOpts, collateral common.Address) (struct {
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
}, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "LTVFull", collateral)

	outstruct := new(struct {
		BorrowLTV             uint16
		LiquidationLTV        uint16
		InitialLiquidationLTV uint16
		TargetTimestamp       *big.Int
		RampDuration          uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BorrowLTV = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.LiquidationLTV = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.InitialLiquidationLTV = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.TargetTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RampDuration = *abi.ConvertType(out[4], new(uint32)).(*uint32)

	return *outstruct, err

}

// LTVFull is a free data retrieval call binding the contract method 0x33708d0c.
//
// Solidity: function LTVFull(address collateral) view returns(uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EVault *EVaultSession) LTVFull(collateral common.Address) (struct {
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
}, error) {
	return _EVault.Contract.LTVFull(&_EVault.CallOpts, collateral)
}

// LTVFull is a free data retrieval call binding the contract method 0x33708d0c.
//
// Solidity: function LTVFull(address collateral) view returns(uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EVault *EVaultCallerSession) LTVFull(collateral common.Address) (struct {
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
}, error) {
	return _EVault.Contract.LTVFull(&_EVault.CallOpts, collateral)
}

// LTVLiquidation is a free data retrieval call binding the contract method 0xaf5aaeeb.
//
// Solidity: function LTVLiquidation(address collateral) view returns(uint16)
func (_EVault *EVaultCaller) LTVLiquidation(opts *bind.CallOpts, collateral common.Address) (uint16, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "LTVLiquidation", collateral)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LTVLiquidation is a free data retrieval call binding the contract method 0xaf5aaeeb.
//
// Solidity: function LTVLiquidation(address collateral) view returns(uint16)
func (_EVault *EVaultSession) LTVLiquidation(collateral common.Address) (uint16, error) {
	return _EVault.Contract.LTVLiquidation(&_EVault.CallOpts, collateral)
}

// LTVLiquidation is a free data retrieval call binding the contract method 0xaf5aaeeb.
//
// Solidity: function LTVLiquidation(address collateral) view returns(uint16)
func (_EVault *EVaultCallerSession) LTVLiquidation(collateral common.Address) (uint16, error) {
	return _EVault.Contract.LTVLiquidation(&_EVault.CallOpts, collateral)
}

// LTVList is a free data retrieval call binding the contract method 0x6a16ef84.
//
// Solidity: function LTVList() view returns(address[])
func (_EVault *EVaultCaller) LTVList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "LTVList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// LTVList is a free data retrieval call binding the contract method 0x6a16ef84.
//
// Solidity: function LTVList() view returns(address[])
func (_EVault *EVaultSession) LTVList() ([]common.Address, error) {
	return _EVault.Contract.LTVList(&_EVault.CallOpts)
}

// LTVList is a free data retrieval call binding the contract method 0x6a16ef84.
//
// Solidity: function LTVList() view returns(address[])
func (_EVault *EVaultCallerSession) LTVList() ([]common.Address, error) {
	return _EVault.Contract.LTVList(&_EVault.CallOpts)
}

// MODULEBALANCEFORWARDER is a free data retrieval call binding the contract method 0x883e3875.
//
// Solidity: function MODULE_BALANCE_FORWARDER() view returns(address)
func (_EVault *EVaultCaller) MODULEBALANCEFORWARDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "MODULE_BALANCE_FORWARDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEBALANCEFORWARDER is a free data retrieval call binding the contract method 0x883e3875.
//
// Solidity: function MODULE_BALANCE_FORWARDER() view returns(address)
func (_EVault *EVaultSession) MODULEBALANCEFORWARDER() (common.Address, error) {
	return _EVault.Contract.MODULEBALANCEFORWARDER(&_EVault.CallOpts)
}

// MODULEBALANCEFORWARDER is a free data retrieval call binding the contract method 0x883e3875.
//
// Solidity: function MODULE_BALANCE_FORWARDER() view returns(address)
func (_EVault *EVaultCallerSession) MODULEBALANCEFORWARDER() (common.Address, error) {
	return _EVault.Contract.MODULEBALANCEFORWARDER(&_EVault.CallOpts)
}

// MODULEBORROWING is a free data retrieval call binding the contract method 0x14c054bc.
//
// Solidity: function MODULE_BORROWING() view returns(address)
func (_EVault *EVaultCaller) MODULEBORROWING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "MODULE_BORROWING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEBORROWING is a free data retrieval call binding the contract method 0x14c054bc.
//
// Solidity: function MODULE_BORROWING() view returns(address)
func (_EVault *EVaultSession) MODULEBORROWING() (common.Address, error) {
	return _EVault.Contract.MODULEBORROWING(&_EVault.CallOpts)
}

// MODULEBORROWING is a free data retrieval call binding the contract method 0x14c054bc.
//
// Solidity: function MODULE_BORROWING() view returns(address)
func (_EVault *EVaultCallerSession) MODULEBORROWING() (common.Address, error) {
	return _EVault.Contract.MODULEBORROWING(&_EVault.CallOpts)
}

// MODULEGOVERNANCE is a free data retrieval call binding the contract method 0xb4cd541b.
//
// Solidity: function MODULE_GOVERNANCE() view returns(address)
func (_EVault *EVaultCaller) MODULEGOVERNANCE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "MODULE_GOVERNANCE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEGOVERNANCE is a free data retrieval call binding the contract method 0xb4cd541b.
//
// Solidity: function MODULE_GOVERNANCE() view returns(address)
func (_EVault *EVaultSession) MODULEGOVERNANCE() (common.Address, error) {
	return _EVault.Contract.MODULEGOVERNANCE(&_EVault.CallOpts)
}

// MODULEGOVERNANCE is a free data retrieval call binding the contract method 0xb4cd541b.
//
// Solidity: function MODULE_GOVERNANCE() view returns(address)
func (_EVault *EVaultCallerSession) MODULEGOVERNANCE() (common.Address, error) {
	return _EVault.Contract.MODULEGOVERNANCE(&_EVault.CallOpts)
}

// MODULEINITIALIZE is a free data retrieval call binding the contract method 0xad80ad0b.
//
// Solidity: function MODULE_INITIALIZE() view returns(address)
func (_EVault *EVaultCaller) MODULEINITIALIZE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "MODULE_INITIALIZE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEINITIALIZE is a free data retrieval call binding the contract method 0xad80ad0b.
//
// Solidity: function MODULE_INITIALIZE() view returns(address)
func (_EVault *EVaultSession) MODULEINITIALIZE() (common.Address, error) {
	return _EVault.Contract.MODULEINITIALIZE(&_EVault.CallOpts)
}

// MODULEINITIALIZE is a free data retrieval call binding the contract method 0xad80ad0b.
//
// Solidity: function MODULE_INITIALIZE() view returns(address)
func (_EVault *EVaultCallerSession) MODULEINITIALIZE() (common.Address, error) {
	return _EVault.Contract.MODULEINITIALIZE(&_EVault.CallOpts)
}

// MODULELIQUIDATION is a free data retrieval call binding the contract method 0x42895567.
//
// Solidity: function MODULE_LIQUIDATION() view returns(address)
func (_EVault *EVaultCaller) MODULELIQUIDATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "MODULE_LIQUIDATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULELIQUIDATION is a free data retrieval call binding the contract method 0x42895567.
//
// Solidity: function MODULE_LIQUIDATION() view returns(address)
func (_EVault *EVaultSession) MODULELIQUIDATION() (common.Address, error) {
	return _EVault.Contract.MODULELIQUIDATION(&_EVault.CallOpts)
}

// MODULELIQUIDATION is a free data retrieval call binding the contract method 0x42895567.
//
// Solidity: function MODULE_LIQUIDATION() view returns(address)
func (_EVault *EVaultCallerSession) MODULELIQUIDATION() (common.Address, error) {
	return _EVault.Contract.MODULELIQUIDATION(&_EVault.CallOpts)
}

// MODULERISKMANAGER is a free data retrieval call binding the contract method 0x7d5f2e4e.
//
// Solidity: function MODULE_RISKMANAGER() view returns(address)
func (_EVault *EVaultCaller) MODULERISKMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "MODULE_RISKMANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULERISKMANAGER is a free data retrieval call binding the contract method 0x7d5f2e4e.
//
// Solidity: function MODULE_RISKMANAGER() view returns(address)
func (_EVault *EVaultSession) MODULERISKMANAGER() (common.Address, error) {
	return _EVault.Contract.MODULERISKMANAGER(&_EVault.CallOpts)
}

// MODULERISKMANAGER is a free data retrieval call binding the contract method 0x7d5f2e4e.
//
// Solidity: function MODULE_RISKMANAGER() view returns(address)
func (_EVault *EVaultCallerSession) MODULERISKMANAGER() (common.Address, error) {
	return _EVault.Contract.MODULERISKMANAGER(&_EVault.CallOpts)
}

// MODULETOKEN is a free data retrieval call binding the contract method 0x5fa23055.
//
// Solidity: function MODULE_TOKEN() view returns(address)
func (_EVault *EVaultCaller) MODULETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "MODULE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULETOKEN is a free data retrieval call binding the contract method 0x5fa23055.
//
// Solidity: function MODULE_TOKEN() view returns(address)
func (_EVault *EVaultSession) MODULETOKEN() (common.Address, error) {
	return _EVault.Contract.MODULETOKEN(&_EVault.CallOpts)
}

// MODULETOKEN is a free data retrieval call binding the contract method 0x5fa23055.
//
// Solidity: function MODULE_TOKEN() view returns(address)
func (_EVault *EVaultCallerSession) MODULETOKEN() (common.Address, error) {
	return _EVault.Contract.MODULETOKEN(&_EVault.CallOpts)
}

// MODULEVAULT is a free data retrieval call binding the contract method 0xe2f206e5.
//
// Solidity: function MODULE_VAULT() view returns(address)
func (_EVault *EVaultCaller) MODULEVAULT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "MODULE_VAULT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEVAULT is a free data retrieval call binding the contract method 0xe2f206e5.
//
// Solidity: function MODULE_VAULT() view returns(address)
func (_EVault *EVaultSession) MODULEVAULT() (common.Address, error) {
	return _EVault.Contract.MODULEVAULT(&_EVault.CallOpts)
}

// MODULEVAULT is a free data retrieval call binding the contract method 0xe2f206e5.
//
// Solidity: function MODULE_VAULT() view returns(address)
func (_EVault *EVaultCallerSession) MODULEVAULT() (common.Address, error) {
	return _EVault.Contract.MODULEVAULT(&_EVault.CallOpts)
}

// AccountLiquidity is a free data retrieval call binding the contract method 0xa824bf67.
//
// Solidity: function accountLiquidity(address account, bool liquidation) view returns(uint256 collateralValue, uint256 liabilityValue)
func (_EVault *EVaultCaller) AccountLiquidity(opts *bind.CallOpts, account common.Address, liquidation bool) (struct {
	CollateralValue *big.Int
	LiabilityValue  *big.Int
}, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "accountLiquidity", account, liquidation)

	outstruct := new(struct {
		CollateralValue *big.Int
		LiabilityValue  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CollateralValue = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiabilityValue = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AccountLiquidity is a free data retrieval call binding the contract method 0xa824bf67.
//
// Solidity: function accountLiquidity(address account, bool liquidation) view returns(uint256 collateralValue, uint256 liabilityValue)
func (_EVault *EVaultSession) AccountLiquidity(account common.Address, liquidation bool) (struct {
	CollateralValue *big.Int
	LiabilityValue  *big.Int
}, error) {
	return _EVault.Contract.AccountLiquidity(&_EVault.CallOpts, account, liquidation)
}

// AccountLiquidity is a free data retrieval call binding the contract method 0xa824bf67.
//
// Solidity: function accountLiquidity(address account, bool liquidation) view returns(uint256 collateralValue, uint256 liabilityValue)
func (_EVault *EVaultCallerSession) AccountLiquidity(account common.Address, liquidation bool) (struct {
	CollateralValue *big.Int
	LiabilityValue  *big.Int
}, error) {
	return _EVault.Contract.AccountLiquidity(&_EVault.CallOpts, account, liquidation)
}

// AccountLiquidityFull is a free data retrieval call binding the contract method 0xc7b0e3a3.
//
// Solidity: function accountLiquidityFull(address account, bool liquidation) view returns(address[] collaterals, uint256[] collateralValues, uint256 liabilityValue)
func (_EVault *EVaultCaller) AccountLiquidityFull(opts *bind.CallOpts, account common.Address, liquidation bool) (struct {
	Collaterals      []common.Address
	CollateralValues []*big.Int
	LiabilityValue   *big.Int
}, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "accountLiquidityFull", account, liquidation)

	outstruct := new(struct {
		Collaterals      []common.Address
		CollateralValues []*big.Int
		LiabilityValue   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Collaterals = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.CollateralValues = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.LiabilityValue = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AccountLiquidityFull is a free data retrieval call binding the contract method 0xc7b0e3a3.
//
// Solidity: function accountLiquidityFull(address account, bool liquidation) view returns(address[] collaterals, uint256[] collateralValues, uint256 liabilityValue)
func (_EVault *EVaultSession) AccountLiquidityFull(account common.Address, liquidation bool) (struct {
	Collaterals      []common.Address
	CollateralValues []*big.Int
	LiabilityValue   *big.Int
}, error) {
	return _EVault.Contract.AccountLiquidityFull(&_EVault.CallOpts, account, liquidation)
}

// AccountLiquidityFull is a free data retrieval call binding the contract method 0xc7b0e3a3.
//
// Solidity: function accountLiquidityFull(address account, bool liquidation) view returns(address[] collaterals, uint256[] collateralValues, uint256 liabilityValue)
func (_EVault *EVaultCallerSession) AccountLiquidityFull(account common.Address, liquidation bool) (struct {
	Collaterals      []common.Address
	CollateralValues []*big.Int
	LiabilityValue   *big.Int
}, error) {
	return _EVault.Contract.AccountLiquidityFull(&_EVault.CallOpts, account, liquidation)
}

// AccumulatedFees is a free data retrieval call binding the contract method 0x587f5ed7.
//
// Solidity: function accumulatedFees() view returns(uint256)
func (_EVault *EVaultCaller) AccumulatedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "accumulatedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedFees is a free data retrieval call binding the contract method 0x587f5ed7.
//
// Solidity: function accumulatedFees() view returns(uint256)
func (_EVault *EVaultSession) AccumulatedFees() (*big.Int, error) {
	return _EVault.Contract.AccumulatedFees(&_EVault.CallOpts)
}

// AccumulatedFees is a free data retrieval call binding the contract method 0x587f5ed7.
//
// Solidity: function accumulatedFees() view returns(uint256)
func (_EVault *EVaultCallerSession) AccumulatedFees() (*big.Int, error) {
	return _EVault.Contract.AccumulatedFees(&_EVault.CallOpts)
}

// AccumulatedFeesAssets is a free data retrieval call binding the contract method 0xf6e50f58.
//
// Solidity: function accumulatedFeesAssets() view returns(uint256)
func (_EVault *EVaultCaller) AccumulatedFeesAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "accumulatedFeesAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedFeesAssets is a free data retrieval call binding the contract method 0xf6e50f58.
//
// Solidity: function accumulatedFeesAssets() view returns(uint256)
func (_EVault *EVaultSession) AccumulatedFeesAssets() (*big.Int, error) {
	return _EVault.Contract.AccumulatedFeesAssets(&_EVault.CallOpts)
}

// AccumulatedFeesAssets is a free data retrieval call binding the contract method 0xf6e50f58.
//
// Solidity: function accumulatedFeesAssets() view returns(uint256)
func (_EVault *EVaultCallerSession) AccumulatedFeesAssets() (*big.Int, error) {
	return _EVault.Contract.AccumulatedFeesAssets(&_EVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_EVault *EVaultCaller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "allowance", holder, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_EVault *EVaultSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _EVault.Contract.Allowance(&_EVault.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_EVault *EVaultCallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _EVault.Contract.Allowance(&_EVault.CallOpts, holder, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EVault *EVaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EVault *EVaultSession) Asset() (common.Address, error) {
	return _EVault.Contract.Asset(&_EVault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EVault *EVaultCallerSession) Asset() (common.Address, error) {
	return _EVault.Contract.Asset(&_EVault.CallOpts)
}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address account) view returns(bool)
func (_EVault *EVaultCaller) BalanceForwarderEnabled(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "balanceForwarderEnabled", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address account) view returns(bool)
func (_EVault *EVaultSession) BalanceForwarderEnabled(account common.Address) (bool, error) {
	return _EVault.Contract.BalanceForwarderEnabled(&_EVault.CallOpts, account)
}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address account) view returns(bool)
func (_EVault *EVaultCallerSession) BalanceForwarderEnabled(account common.Address) (bool, error) {
	return _EVault.Contract.BalanceForwarderEnabled(&_EVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EVault *EVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EVault *EVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EVault.Contract.BalanceOf(&_EVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EVault *EVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EVault.Contract.BalanceOf(&_EVault.CallOpts, account)
}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EVault *EVaultCaller) BalanceTrackerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "balanceTrackerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EVault *EVaultSession) BalanceTrackerAddress() (common.Address, error) {
	return _EVault.Contract.BalanceTrackerAddress(&_EVault.CallOpts)
}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EVault *EVaultCallerSession) BalanceTrackerAddress() (common.Address, error) {
	return _EVault.Contract.BalanceTrackerAddress(&_EVault.CallOpts)
}

// Caps is a free data retrieval call binding the contract method 0x18e22d98.
//
// Solidity: function caps() view returns(uint16 supplyCap, uint16 borrowCap)
func (_EVault *EVaultCaller) Caps(opts *bind.CallOpts) (struct {
	SupplyCap uint16
	BorrowCap uint16
}, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "caps")

	outstruct := new(struct {
		SupplyCap uint16
		BorrowCap uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SupplyCap = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.BorrowCap = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// Caps is a free data retrieval call binding the contract method 0x18e22d98.
//
// Solidity: function caps() view returns(uint16 supplyCap, uint16 borrowCap)
func (_EVault *EVaultSession) Caps() (struct {
	SupplyCap uint16
	BorrowCap uint16
}, error) {
	return _EVault.Contract.Caps(&_EVault.CallOpts)
}

// Caps is a free data retrieval call binding the contract method 0x18e22d98.
//
// Solidity: function caps() view returns(uint16 supplyCap, uint16 borrowCap)
func (_EVault *EVaultCallerSession) Caps() (struct {
	SupplyCap uint16
	BorrowCap uint16
}, error) {
	return _EVault.Contract.Caps(&_EVault.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_EVault *EVaultCaller) Cash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "cash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_EVault *EVaultSession) Cash() (*big.Int, error) {
	return _EVault.Contract.Cash(&_EVault.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_EVault *EVaultCallerSession) Cash() (*big.Int, error) {
	return _EVault.Contract.Cash(&_EVault.CallOpts)
}

// CheckAccountStatus is a free data retrieval call binding the contract method 0xb168c58f.
//
// Solidity: function checkAccountStatus(address account, address[] collaterals) view returns(bytes4)
func (_EVault *EVaultCaller) CheckAccountStatus(opts *bind.CallOpts, account common.Address, collaterals []common.Address) ([4]byte, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "checkAccountStatus", account, collaterals)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// CheckAccountStatus is a free data retrieval call binding the contract method 0xb168c58f.
//
// Solidity: function checkAccountStatus(address account, address[] collaterals) view returns(bytes4)
func (_EVault *EVaultSession) CheckAccountStatus(account common.Address, collaterals []common.Address) ([4]byte, error) {
	return _EVault.Contract.CheckAccountStatus(&_EVault.CallOpts, account, collaterals)
}

// CheckAccountStatus is a free data retrieval call binding the contract method 0xb168c58f.
//
// Solidity: function checkAccountStatus(address account, address[] collaterals) view returns(bytes4)
func (_EVault *EVaultCallerSession) CheckAccountStatus(account common.Address, collaterals []common.Address) ([4]byte, error) {
	return _EVault.Contract.CheckAccountStatus(&_EVault.CallOpts, account, collaterals)
}

// CheckLiquidation is a free data retrieval call binding the contract method 0x88aa6f12.
//
// Solidity: function checkLiquidation(address liquidator, address violator, address collateral) view returns(uint256 maxRepay, uint256 maxYield)
func (_EVault *EVaultCaller) CheckLiquidation(opts *bind.CallOpts, liquidator common.Address, violator common.Address, collateral common.Address) (struct {
	MaxRepay *big.Int
	MaxYield *big.Int
}, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "checkLiquidation", liquidator, violator, collateral)

	outstruct := new(struct {
		MaxRepay *big.Int
		MaxYield *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxRepay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxYield = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CheckLiquidation is a free data retrieval call binding the contract method 0x88aa6f12.
//
// Solidity: function checkLiquidation(address liquidator, address violator, address collateral) view returns(uint256 maxRepay, uint256 maxYield)
func (_EVault *EVaultSession) CheckLiquidation(liquidator common.Address, violator common.Address, collateral common.Address) (struct {
	MaxRepay *big.Int
	MaxYield *big.Int
}, error) {
	return _EVault.Contract.CheckLiquidation(&_EVault.CallOpts, liquidator, violator, collateral)
}

// CheckLiquidation is a free data retrieval call binding the contract method 0x88aa6f12.
//
// Solidity: function checkLiquidation(address liquidator, address violator, address collateral) view returns(uint256 maxRepay, uint256 maxYield)
func (_EVault *EVaultCallerSession) CheckLiquidation(liquidator common.Address, violator common.Address, collateral common.Address) (struct {
	MaxRepay *big.Int
	MaxYield *big.Int
}, error) {
	return _EVault.Contract.CheckLiquidation(&_EVault.CallOpts, liquidator, violator, collateral)
}

// ConfigFlags is a free data retrieval call binding the contract method 0x2b38a367.
//
// Solidity: function configFlags() view returns(uint32)
func (_EVault *EVaultCaller) ConfigFlags(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "configFlags")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ConfigFlags is a free data retrieval call binding the contract method 0x2b38a367.
//
// Solidity: function configFlags() view returns(uint32)
func (_EVault *EVaultSession) ConfigFlags() (uint32, error) {
	return _EVault.Contract.ConfigFlags(&_EVault.CallOpts)
}

// ConfigFlags is a free data retrieval call binding the contract method 0x2b38a367.
//
// Solidity: function configFlags() view returns(uint32)
func (_EVault *EVaultCallerSession) ConfigFlags() (uint32, error) {
	return _EVault.Contract.ConfigFlags(&_EVault.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_EVault *EVaultCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_EVault *EVaultSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _EVault.Contract.ConvertToAssets(&_EVault.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_EVault *EVaultCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _EVault.Contract.ConvertToAssets(&_EVault.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_EVault *EVaultCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_EVault *EVaultSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _EVault.Contract.ConvertToShares(&_EVault.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_EVault *EVaultCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _EVault.Contract.ConvertToShares(&_EVault.CallOpts, assets)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_EVault *EVaultCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_EVault *EVaultSession) Creator() (common.Address, error) {
	return _EVault.Contract.Creator(&_EVault.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_EVault *EVaultCallerSession) Creator() (common.Address, error) {
	return _EVault.Contract.Creator(&_EVault.CallOpts)
}

// DToken is a free data retrieval call binding the contract method 0xd9d7858a.
//
// Solidity: function dToken() view returns(address)
func (_EVault *EVaultCaller) DToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "dToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DToken is a free data retrieval call binding the contract method 0xd9d7858a.
//
// Solidity: function dToken() view returns(address)
func (_EVault *EVaultSession) DToken() (common.Address, error) {
	return _EVault.Contract.DToken(&_EVault.CallOpts)
}

// DToken is a free data retrieval call binding the contract method 0xd9d7858a.
//
// Solidity: function dToken() view returns(address)
func (_EVault *EVaultCallerSession) DToken() (common.Address, error) {
	return _EVault.Contract.DToken(&_EVault.CallOpts)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_EVault *EVaultCaller) DebtOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "debtOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_EVault *EVaultSession) DebtOf(account common.Address) (*big.Int, error) {
	return _EVault.Contract.DebtOf(&_EVault.CallOpts, account)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_EVault *EVaultCallerSession) DebtOf(account common.Address) (*big.Int, error) {
	return _EVault.Contract.DebtOf(&_EVault.CallOpts, account)
}

// DebtOfExact is a free data retrieval call binding the contract method 0xab49b7f1.
//
// Solidity: function debtOfExact(address account) view returns(uint256)
func (_EVault *EVaultCaller) DebtOfExact(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "debtOfExact", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOfExact is a free data retrieval call binding the contract method 0xab49b7f1.
//
// Solidity: function debtOfExact(address account) view returns(uint256)
func (_EVault *EVaultSession) DebtOfExact(account common.Address) (*big.Int, error) {
	return _EVault.Contract.DebtOfExact(&_EVault.CallOpts, account)
}

// DebtOfExact is a free data retrieval call binding the contract method 0xab49b7f1.
//
// Solidity: function debtOfExact(address account) view returns(uint256)
func (_EVault *EVaultCallerSession) DebtOfExact(account common.Address) (*big.Int, error) {
	return _EVault.Contract.DebtOfExact(&_EVault.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EVault *EVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EVault *EVaultSession) Decimals() (uint8, error) {
	return _EVault.Contract.Decimals(&_EVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EVault *EVaultCallerSession) Decimals() (uint8, error) {
	return _EVault.Contract.Decimals(&_EVault.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_EVault *EVaultCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "feeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_EVault *EVaultSession) FeeReceiver() (common.Address, error) {
	return _EVault.Contract.FeeReceiver(&_EVault.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_EVault *EVaultCallerSession) FeeReceiver() (common.Address, error) {
	return _EVault.Contract.FeeReceiver(&_EVault.CallOpts)
}

// GovernorAdmin is a free data retrieval call binding the contract method 0x6ce98c29.
//
// Solidity: function governorAdmin() view returns(address)
func (_EVault *EVaultCaller) GovernorAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "governorAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernorAdmin is a free data retrieval call binding the contract method 0x6ce98c29.
//
// Solidity: function governorAdmin() view returns(address)
func (_EVault *EVaultSession) GovernorAdmin() (common.Address, error) {
	return _EVault.Contract.GovernorAdmin(&_EVault.CallOpts)
}

// GovernorAdmin is a free data retrieval call binding the contract method 0x6ce98c29.
//
// Solidity: function governorAdmin() view returns(address)
func (_EVault *EVaultCallerSession) GovernorAdmin() (common.Address, error) {
	return _EVault.Contract.GovernorAdmin(&_EVault.CallOpts)
}

// HookConfig is a free data retrieval call binding the contract method 0xcf349b7d.
//
// Solidity: function hookConfig() view returns(address, uint32)
func (_EVault *EVaultCaller) HookConfig(opts *bind.CallOpts) (common.Address, uint32, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "hookConfig")

	if err != nil {
		return *new(common.Address), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// HookConfig is a free data retrieval call binding the contract method 0xcf349b7d.
//
// Solidity: function hookConfig() view returns(address, uint32)
func (_EVault *EVaultSession) HookConfig() (common.Address, uint32, error) {
	return _EVault.Contract.HookConfig(&_EVault.CallOpts)
}

// HookConfig is a free data retrieval call binding the contract method 0xcf349b7d.
//
// Solidity: function hookConfig() view returns(address, uint32)
func (_EVault *EVaultCallerSession) HookConfig() (common.Address, uint32, error) {
	return _EVault.Contract.HookConfig(&_EVault.CallOpts)
}

// InterestAccumulator is a free data retrieval call binding the contract method 0x087a6007.
//
// Solidity: function interestAccumulator() view returns(uint256)
func (_EVault *EVaultCaller) InterestAccumulator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "interestAccumulator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestAccumulator is a free data retrieval call binding the contract method 0x087a6007.
//
// Solidity: function interestAccumulator() view returns(uint256)
func (_EVault *EVaultSession) InterestAccumulator() (*big.Int, error) {
	return _EVault.Contract.InterestAccumulator(&_EVault.CallOpts)
}

// InterestAccumulator is a free data retrieval call binding the contract method 0x087a6007.
//
// Solidity: function interestAccumulator() view returns(uint256)
func (_EVault *EVaultCallerSession) InterestAccumulator() (*big.Int, error) {
	return _EVault.Contract.InterestAccumulator(&_EVault.CallOpts)
}

// InterestFee is a free data retrieval call binding the contract method 0xa75df498.
//
// Solidity: function interestFee() view returns(uint16)
func (_EVault *EVaultCaller) InterestFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "interestFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// InterestFee is a free data retrieval call binding the contract method 0xa75df498.
//
// Solidity: function interestFee() view returns(uint16)
func (_EVault *EVaultSession) InterestFee() (uint16, error) {
	return _EVault.Contract.InterestFee(&_EVault.CallOpts)
}

// InterestFee is a free data retrieval call binding the contract method 0xa75df498.
//
// Solidity: function interestFee() view returns(uint16)
func (_EVault *EVaultCallerSession) InterestFee() (uint16, error) {
	return _EVault.Contract.InterestFee(&_EVault.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_EVault *EVaultCaller) InterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "interestRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_EVault *EVaultSession) InterestRate() (*big.Int, error) {
	return _EVault.Contract.InterestRate(&_EVault.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_EVault *EVaultCallerSession) InterestRate() (*big.Int, error) {
	return _EVault.Contract.InterestRate(&_EVault.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_EVault *EVaultCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_EVault *EVaultSession) InterestRateModel() (common.Address, error) {
	return _EVault.Contract.InterestRateModel(&_EVault.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_EVault *EVaultCallerSession) InterestRateModel() (common.Address, error) {
	return _EVault.Contract.InterestRateModel(&_EVault.CallOpts)
}

// LiquidationCoolOffTime is a free data retrieval call binding the contract method 0x4abdb959.
//
// Solidity: function liquidationCoolOffTime() view returns(uint16)
func (_EVault *EVaultCaller) LiquidationCoolOffTime(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "liquidationCoolOffTime")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LiquidationCoolOffTime is a free data retrieval call binding the contract method 0x4abdb959.
//
// Solidity: function liquidationCoolOffTime() view returns(uint16)
func (_EVault *EVaultSession) LiquidationCoolOffTime() (uint16, error) {
	return _EVault.Contract.LiquidationCoolOffTime(&_EVault.CallOpts)
}

// LiquidationCoolOffTime is a free data retrieval call binding the contract method 0x4abdb959.
//
// Solidity: function liquidationCoolOffTime() view returns(uint16)
func (_EVault *EVaultCallerSession) LiquidationCoolOffTime() (uint16, error) {
	return _EVault.Contract.LiquidationCoolOffTime(&_EVault.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address account) view returns(uint256)
func (_EVault *EVaultCaller) MaxDeposit(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "maxDeposit", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address account) view returns(uint256)
func (_EVault *EVaultSession) MaxDeposit(account common.Address) (*big.Int, error) {
	return _EVault.Contract.MaxDeposit(&_EVault.CallOpts, account)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address account) view returns(uint256)
func (_EVault *EVaultCallerSession) MaxDeposit(account common.Address) (*big.Int, error) {
	return _EVault.Contract.MaxDeposit(&_EVault.CallOpts, account)
}

// MaxLiquidationDiscount is a free data retrieval call binding the contract method 0x4f7e43df.
//
// Solidity: function maxLiquidationDiscount() view returns(uint16)
func (_EVault *EVaultCaller) MaxLiquidationDiscount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "maxLiquidationDiscount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaxLiquidationDiscount is a free data retrieval call binding the contract method 0x4f7e43df.
//
// Solidity: function maxLiquidationDiscount() view returns(uint16)
func (_EVault *EVaultSession) MaxLiquidationDiscount() (uint16, error) {
	return _EVault.Contract.MaxLiquidationDiscount(&_EVault.CallOpts)
}

// MaxLiquidationDiscount is a free data retrieval call binding the contract method 0x4f7e43df.
//
// Solidity: function maxLiquidationDiscount() view returns(uint16)
func (_EVault *EVaultCallerSession) MaxLiquidationDiscount() (uint16, error) {
	return _EVault.Contract.MaxLiquidationDiscount(&_EVault.CallOpts)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address account) view returns(uint256)
func (_EVault *EVaultCaller) MaxMint(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "maxMint", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address account) view returns(uint256)
func (_EVault *EVaultSession) MaxMint(account common.Address) (*big.Int, error) {
	return _EVault.Contract.MaxMint(&_EVault.CallOpts, account)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address account) view returns(uint256)
func (_EVault *EVaultCallerSession) MaxMint(account common.Address) (*big.Int, error) {
	return _EVault.Contract.MaxMint(&_EVault.CallOpts, account)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_EVault *EVaultCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_EVault *EVaultSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _EVault.Contract.MaxRedeem(&_EVault.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_EVault *EVaultCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _EVault.Contract.MaxRedeem(&_EVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_EVault *EVaultCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_EVault *EVaultSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _EVault.Contract.MaxWithdraw(&_EVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_EVault *EVaultCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _EVault.Contract.MaxWithdraw(&_EVault.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EVault *EVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EVault *EVaultSession) Name() (string, error) {
	return _EVault.Contract.Name(&_EVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EVault *EVaultCallerSession) Name() (string, error) {
	return _EVault.Contract.Name(&_EVault.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EVault *EVaultCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EVault *EVaultSession) Oracle() (common.Address, error) {
	return _EVault.Contract.Oracle(&_EVault.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EVault *EVaultCallerSession) Oracle() (common.Address, error) {
	return _EVault.Contract.Oracle(&_EVault.CallOpts)
}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EVault *EVaultCaller) Permit2Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "permit2Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EVault *EVaultSession) Permit2Address() (common.Address, error) {
	return _EVault.Contract.Permit2Address(&_EVault.CallOpts)
}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EVault *EVaultCallerSession) Permit2Address() (common.Address, error) {
	return _EVault.Contract.Permit2Address(&_EVault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_EVault *EVaultCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_EVault *EVaultSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _EVault.Contract.PreviewDeposit(&_EVault.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_EVault *EVaultCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _EVault.Contract.PreviewDeposit(&_EVault.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_EVault *EVaultCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_EVault *EVaultSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _EVault.Contract.PreviewMint(&_EVault.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_EVault *EVaultCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _EVault.Contract.PreviewMint(&_EVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_EVault *EVaultCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_EVault *EVaultSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _EVault.Contract.PreviewRedeem(&_EVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_EVault *EVaultCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _EVault.Contract.PreviewRedeem(&_EVault.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_EVault *EVaultCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_EVault *EVaultSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _EVault.Contract.PreviewWithdraw(&_EVault.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_EVault *EVaultCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _EVault.Contract.PreviewWithdraw(&_EVault.CallOpts, assets)
}

// ProtocolConfigAddress is a free data retrieval call binding the contract method 0x539bd5bf.
//
// Solidity: function protocolConfigAddress() view returns(address)
func (_EVault *EVaultCaller) ProtocolConfigAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "protocolConfigAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolConfigAddress is a free data retrieval call binding the contract method 0x539bd5bf.
//
// Solidity: function protocolConfigAddress() view returns(address)
func (_EVault *EVaultSession) ProtocolConfigAddress() (common.Address, error) {
	return _EVault.Contract.ProtocolConfigAddress(&_EVault.CallOpts)
}

// ProtocolConfigAddress is a free data retrieval call binding the contract method 0x539bd5bf.
//
// Solidity: function protocolConfigAddress() view returns(address)
func (_EVault *EVaultCallerSession) ProtocolConfigAddress() (common.Address, error) {
	return _EVault.Contract.ProtocolConfigAddress(&_EVault.CallOpts)
}

// ProtocolFeeReceiver is a free data retrieval call binding the contract method 0x39a51be5.
//
// Solidity: function protocolFeeReceiver() view returns(address)
func (_EVault *EVaultCaller) ProtocolFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "protocolFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeReceiver is a free data retrieval call binding the contract method 0x39a51be5.
//
// Solidity: function protocolFeeReceiver() view returns(address)
func (_EVault *EVaultSession) ProtocolFeeReceiver() (common.Address, error) {
	return _EVault.Contract.ProtocolFeeReceiver(&_EVault.CallOpts)
}

// ProtocolFeeReceiver is a free data retrieval call binding the contract method 0x39a51be5.
//
// Solidity: function protocolFeeReceiver() view returns(address)
func (_EVault *EVaultCallerSession) ProtocolFeeReceiver() (common.Address, error) {
	return _EVault.Contract.ProtocolFeeReceiver(&_EVault.CallOpts)
}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_EVault *EVaultCaller) ProtocolFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "protocolFeeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_EVault *EVaultSession) ProtocolFeeShare() (*big.Int, error) {
	return _EVault.Contract.ProtocolFeeShare(&_EVault.CallOpts)
}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_EVault *EVaultCallerSession) ProtocolFeeShare() (*big.Int, error) {
	return _EVault.Contract.ProtocolFeeShare(&_EVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EVault *EVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EVault *EVaultSession) Symbol() (string, error) {
	return _EVault.Contract.Symbol(&_EVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EVault *EVaultCallerSession) Symbol() (string, error) {
	return _EVault.Contract.Symbol(&_EVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EVault *EVaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EVault *EVaultSession) TotalAssets() (*big.Int, error) {
	return _EVault.Contract.TotalAssets(&_EVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EVault *EVaultCallerSession) TotalAssets() (*big.Int, error) {
	return _EVault.Contract.TotalAssets(&_EVault.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_EVault *EVaultCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_EVault *EVaultSession) TotalBorrows() (*big.Int, error) {
	return _EVault.Contract.TotalBorrows(&_EVault.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_EVault *EVaultCallerSession) TotalBorrows() (*big.Int, error) {
	return _EVault.Contract.TotalBorrows(&_EVault.CallOpts)
}

// TotalBorrowsExact is a free data retrieval call binding the contract method 0xe388be7b.
//
// Solidity: function totalBorrowsExact() view returns(uint256)
func (_EVault *EVaultCaller) TotalBorrowsExact(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "totalBorrowsExact")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowsExact is a free data retrieval call binding the contract method 0xe388be7b.
//
// Solidity: function totalBorrowsExact() view returns(uint256)
func (_EVault *EVaultSession) TotalBorrowsExact() (*big.Int, error) {
	return _EVault.Contract.TotalBorrowsExact(&_EVault.CallOpts)
}

// TotalBorrowsExact is a free data retrieval call binding the contract method 0xe388be7b.
//
// Solidity: function totalBorrowsExact() view returns(uint256)
func (_EVault *EVaultCallerSession) TotalBorrowsExact() (*big.Int, error) {
	return _EVault.Contract.TotalBorrowsExact(&_EVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EVault *EVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EVault *EVaultSession) TotalSupply() (*big.Int, error) {
	return _EVault.Contract.TotalSupply(&_EVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EVault *EVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _EVault.Contract.TotalSupply(&_EVault.CallOpts)
}

// UnitOfAccount is a free data retrieval call binding the contract method 0x3e833364.
//
// Solidity: function unitOfAccount() view returns(address)
func (_EVault *EVaultCaller) UnitOfAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVault.contract.Call(opts, &out, "unitOfAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnitOfAccount is a free data retrieval call binding the contract method 0x3e833364.
//
// Solidity: function unitOfAccount() view returns(address)
func (_EVault *EVaultSession) UnitOfAccount() (common.Address, error) {
	return _EVault.Contract.UnitOfAccount(&_EVault.CallOpts)
}

// UnitOfAccount is a free data retrieval call binding the contract method 0x3e833364.
//
// Solidity: function unitOfAccount() view returns(address)
func (_EVault *EVaultCallerSession) UnitOfAccount() (common.Address, error) {
	return _EVault.Contract.UnitOfAccount(&_EVault.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EVault *EVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EVault *EVaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVault.Contract.Approve(&_EVault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EVault *EVaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVault.Contract.Approve(&_EVault.TransactOpts, spender, amount)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount, address receiver) returns(uint256)
func (_EVault *EVaultTransactor) Borrow(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "borrow", amount, receiver)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount, address receiver) returns(uint256)
func (_EVault *EVaultSession) Borrow(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.Contract.Borrow(&_EVault.TransactOpts, amount, receiver)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount, address receiver) returns(uint256)
func (_EVault *EVaultTransactorSession) Borrow(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.Contract.Borrow(&_EVault.TransactOpts, amount, receiver)
}

// CheckVaultStatus is a paid mutator transaction binding the contract method 0x4b3d1223.
//
// Solidity: function checkVaultStatus() returns(bytes4)
func (_EVault *EVaultTransactor) CheckVaultStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "checkVaultStatus")
}

// CheckVaultStatus is a paid mutator transaction binding the contract method 0x4b3d1223.
//
// Solidity: function checkVaultStatus() returns(bytes4)
func (_EVault *EVaultSession) CheckVaultStatus() (*types.Transaction, error) {
	return _EVault.Contract.CheckVaultStatus(&_EVault.TransactOpts)
}

// CheckVaultStatus is a paid mutator transaction binding the contract method 0x4b3d1223.
//
// Solidity: function checkVaultStatus() returns(bytes4)
func (_EVault *EVaultTransactorSession) CheckVaultStatus() (*types.Transaction, error) {
	return _EVault.Contract.CheckVaultStatus(&_EVault.TransactOpts)
}

// ConvertFees is a paid mutator transaction binding the contract method 0x2b5335c3.
//
// Solidity: function convertFees() returns()
func (_EVault *EVaultTransactor) ConvertFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "convertFees")
}

// ConvertFees is a paid mutator transaction binding the contract method 0x2b5335c3.
//
// Solidity: function convertFees() returns()
func (_EVault *EVaultSession) ConvertFees() (*types.Transaction, error) {
	return _EVault.Contract.ConvertFees(&_EVault.TransactOpts)
}

// ConvertFees is a paid mutator transaction binding the contract method 0x2b5335c3.
//
// Solidity: function convertFees() returns()
func (_EVault *EVaultTransactorSession) ConvertFees() (*types.Transaction, error) {
	return _EVault.Contract.ConvertFees(&_EVault.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns(uint256)
func (_EVault *EVaultTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "deposit", amount, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns(uint256)
func (_EVault *EVaultSession) Deposit(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.Contract.Deposit(&_EVault.TransactOpts, amount, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns(uint256)
func (_EVault *EVaultTransactorSession) Deposit(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.Contract.Deposit(&_EVault.TransactOpts, amount, receiver)
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EVault *EVaultTransactor) DisableBalanceForwarder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "disableBalanceForwarder")
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EVault *EVaultSession) DisableBalanceForwarder() (*types.Transaction, error) {
	return _EVault.Contract.DisableBalanceForwarder(&_EVault.TransactOpts)
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EVault *EVaultTransactorSession) DisableBalanceForwarder() (*types.Transaction, error) {
	return _EVault.Contract.DisableBalanceForwarder(&_EVault.TransactOpts)
}

// DisableController is a paid mutator transaction binding the contract method 0x869e50c7.
//
// Solidity: function disableController() returns()
func (_EVault *EVaultTransactor) DisableController(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "disableController")
}

// DisableController is a paid mutator transaction binding the contract method 0x869e50c7.
//
// Solidity: function disableController() returns()
func (_EVault *EVaultSession) DisableController() (*types.Transaction, error) {
	return _EVault.Contract.DisableController(&_EVault.TransactOpts)
}

// DisableController is a paid mutator transaction binding the contract method 0x869e50c7.
//
// Solidity: function disableController() returns()
func (_EVault *EVaultTransactorSession) DisableController() (*types.Transaction, error) {
	return _EVault.Contract.DisableController(&_EVault.TransactOpts)
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EVault *EVaultTransactor) EnableBalanceForwarder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "enableBalanceForwarder")
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EVault *EVaultSession) EnableBalanceForwarder() (*types.Transaction, error) {
	return _EVault.Contract.EnableBalanceForwarder(&_EVault.TransactOpts)
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EVault *EVaultTransactorSession) EnableBalanceForwarder() (*types.Transaction, error) {
	return _EVault.Contract.EnableBalanceForwarder(&_EVault.TransactOpts)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5296a431.
//
// Solidity: function flashLoan(uint256 amount, bytes data) returns()
func (_EVault *EVaultTransactor) FlashLoan(opts *bind.TransactOpts, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "flashLoan", amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5296a431.
//
// Solidity: function flashLoan(uint256 amount, bytes data) returns()
func (_EVault *EVaultSession) FlashLoan(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EVault.Contract.FlashLoan(&_EVault.TransactOpts, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5296a431.
//
// Solidity: function flashLoan(uint256 amount, bytes data) returns()
func (_EVault *EVaultTransactorSession) FlashLoan(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EVault.Contract.FlashLoan(&_EVault.TransactOpts, amount, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address proxyCreator) returns()
func (_EVault *EVaultTransactor) Initialize(opts *bind.TransactOpts, proxyCreator common.Address) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "initialize", proxyCreator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address proxyCreator) returns()
func (_EVault *EVaultSession) Initialize(proxyCreator common.Address) (*types.Transaction, error) {
	return _EVault.Contract.Initialize(&_EVault.TransactOpts, proxyCreator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address proxyCreator) returns()
func (_EVault *EVaultTransactorSession) Initialize(proxyCreator common.Address) (*types.Transaction, error) {
	return _EVault.Contract.Initialize(&_EVault.TransactOpts, proxyCreator)
}

// Liquidate is a paid mutator transaction binding the contract method 0xc1342574.
//
// Solidity: function liquidate(address violator, address collateral, uint256 repayAssets, uint256 minYieldBalance) returns()
func (_EVault *EVaultTransactor) Liquidate(opts *bind.TransactOpts, violator common.Address, collateral common.Address, repayAssets *big.Int, minYieldBalance *big.Int) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "liquidate", violator, collateral, repayAssets, minYieldBalance)
}

// Liquidate is a paid mutator transaction binding the contract method 0xc1342574.
//
// Solidity: function liquidate(address violator, address collateral, uint256 repayAssets, uint256 minYieldBalance) returns()
func (_EVault *EVaultSession) Liquidate(violator common.Address, collateral common.Address, repayAssets *big.Int, minYieldBalance *big.Int) (*types.Transaction, error) {
	return _EVault.Contract.Liquidate(&_EVault.TransactOpts, violator, collateral, repayAssets, minYieldBalance)
}

// Liquidate is a paid mutator transaction binding the contract method 0xc1342574.
//
// Solidity: function liquidate(address violator, address collateral, uint256 repayAssets, uint256 minYieldBalance) returns()
func (_EVault *EVaultTransactorSession) Liquidate(violator common.Address, collateral common.Address, repayAssets *big.Int, minYieldBalance *big.Int) (*types.Transaction, error) {
	return _EVault.Contract.Liquidate(&_EVault.TransactOpts, violator, collateral, repayAssets, minYieldBalance)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 amount, address receiver) returns(uint256)
func (_EVault *EVaultTransactor) Mint(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "mint", amount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 amount, address receiver) returns(uint256)
func (_EVault *EVaultSession) Mint(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.Contract.Mint(&_EVault.TransactOpts, amount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 amount, address receiver) returns(uint256)
func (_EVault *EVaultTransactorSession) Mint(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.Contract.Mint(&_EVault.TransactOpts, amount, receiver)
}

// PullDebt is a paid mutator transaction binding the contract method 0xaebde56b.
//
// Solidity: function pullDebt(uint256 amount, address from) returns()
func (_EVault *EVaultTransactor) PullDebt(opts *bind.TransactOpts, amount *big.Int, from common.Address) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "pullDebt", amount, from)
}

// PullDebt is a paid mutator transaction binding the contract method 0xaebde56b.
//
// Solidity: function pullDebt(uint256 amount, address from) returns()
func (_EVault *EVaultSession) PullDebt(amount *big.Int, from common.Address) (*types.Transaction, error) {
	return _EVault.Contract.PullDebt(&_EVault.TransactOpts, amount, from)
}

// PullDebt is a paid mutator transaction binding the contract method 0xaebde56b.
//
// Solidity: function pullDebt(uint256 amount, address from) returns()
func (_EVault *EVaultTransactorSession) PullDebt(amount *big.Int, from common.Address) (*types.Transaction, error) {
	return _EVault.Contract.PullDebt(&_EVault.TransactOpts, amount, from)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 amount, address receiver, address owner) returns(uint256)
func (_EVault *EVaultTransactor) Redeem(opts *bind.TransactOpts, amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "redeem", amount, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 amount, address receiver, address owner) returns(uint256)
func (_EVault *EVaultSession) Redeem(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EVault.Contract.Redeem(&_EVault.TransactOpts, amount, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 amount, address receiver, address owner) returns(uint256)
func (_EVault *EVaultTransactorSession) Redeem(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EVault.Contract.Redeem(&_EVault.TransactOpts, amount, receiver, owner)
}

// Repay is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 amount, address receiver) returns(uint256)
func (_EVault *EVaultTransactor) Repay(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "repay", amount, receiver)
}

// Repay is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 amount, address receiver) returns(uint256)
func (_EVault *EVaultSession) Repay(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.Contract.Repay(&_EVault.TransactOpts, amount, receiver)
}

// Repay is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 amount, address receiver) returns(uint256)
func (_EVault *EVaultTransactorSession) Repay(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.Contract.Repay(&_EVault.TransactOpts, amount, receiver)
}

// RepayWithShares is a paid mutator transaction binding the contract method 0xa9c8eb7e.
//
// Solidity: function repayWithShares(uint256 amount, address receiver) returns(uint256 shares, uint256 debt)
func (_EVault *EVaultTransactor) RepayWithShares(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "repayWithShares", amount, receiver)
}

// RepayWithShares is a paid mutator transaction binding the contract method 0xa9c8eb7e.
//
// Solidity: function repayWithShares(uint256 amount, address receiver) returns(uint256 shares, uint256 debt)
func (_EVault *EVaultSession) RepayWithShares(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.Contract.RepayWithShares(&_EVault.TransactOpts, amount, receiver)
}

// RepayWithShares is a paid mutator transaction binding the contract method 0xa9c8eb7e.
//
// Solidity: function repayWithShares(uint256 amount, address receiver) returns(uint256 shares, uint256 debt)
func (_EVault *EVaultTransactorSession) RepayWithShares(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.Contract.RepayWithShares(&_EVault.TransactOpts, amount, receiver)
}

// SetCaps is a paid mutator transaction binding the contract method 0xd87f780f.
//
// Solidity: function setCaps(uint16 supplyCap, uint16 borrowCap) returns()
func (_EVault *EVaultTransactor) SetCaps(opts *bind.TransactOpts, supplyCap uint16, borrowCap uint16) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "setCaps", supplyCap, borrowCap)
}

// SetCaps is a paid mutator transaction binding the contract method 0xd87f780f.
//
// Solidity: function setCaps(uint16 supplyCap, uint16 borrowCap) returns()
func (_EVault *EVaultSession) SetCaps(supplyCap uint16, borrowCap uint16) (*types.Transaction, error) {
	return _EVault.Contract.SetCaps(&_EVault.TransactOpts, supplyCap, borrowCap)
}

// SetCaps is a paid mutator transaction binding the contract method 0xd87f780f.
//
// Solidity: function setCaps(uint16 supplyCap, uint16 borrowCap) returns()
func (_EVault *EVaultTransactorSession) SetCaps(supplyCap uint16, borrowCap uint16) (*types.Transaction, error) {
	return _EVault.Contract.SetCaps(&_EVault.TransactOpts, supplyCap, borrowCap)
}

// SetConfigFlags is a paid mutator transaction binding the contract method 0xada3d56f.
//
// Solidity: function setConfigFlags(uint32 newConfigFlags) returns()
func (_EVault *EVaultTransactor) SetConfigFlags(opts *bind.TransactOpts, newConfigFlags uint32) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "setConfigFlags", newConfigFlags)
}

// SetConfigFlags is a paid mutator transaction binding the contract method 0xada3d56f.
//
// Solidity: function setConfigFlags(uint32 newConfigFlags) returns()
func (_EVault *EVaultSession) SetConfigFlags(newConfigFlags uint32) (*types.Transaction, error) {
	return _EVault.Contract.SetConfigFlags(&_EVault.TransactOpts, newConfigFlags)
}

// SetConfigFlags is a paid mutator transaction binding the contract method 0xada3d56f.
//
// Solidity: function setConfigFlags(uint32 newConfigFlags) returns()
func (_EVault *EVaultTransactorSession) SetConfigFlags(newConfigFlags uint32) (*types.Transaction, error) {
	return _EVault.Contract.SetConfigFlags(&_EVault.TransactOpts, newConfigFlags)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newFeeReceiver) returns()
func (_EVault *EVaultTransactor) SetFeeReceiver(opts *bind.TransactOpts, newFeeReceiver common.Address) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "setFeeReceiver", newFeeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newFeeReceiver) returns()
func (_EVault *EVaultSession) SetFeeReceiver(newFeeReceiver common.Address) (*types.Transaction, error) {
	return _EVault.Contract.SetFeeReceiver(&_EVault.TransactOpts, newFeeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newFeeReceiver) returns()
func (_EVault *EVaultTransactorSession) SetFeeReceiver(newFeeReceiver common.Address) (*types.Transaction, error) {
	return _EVault.Contract.SetFeeReceiver(&_EVault.TransactOpts, newFeeReceiver)
}

// SetGovernorAdmin is a paid mutator transaction binding the contract method 0x82ebd674.
//
// Solidity: function setGovernorAdmin(address newGovernorAdmin) returns()
func (_EVault *EVaultTransactor) SetGovernorAdmin(opts *bind.TransactOpts, newGovernorAdmin common.Address) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "setGovernorAdmin", newGovernorAdmin)
}

// SetGovernorAdmin is a paid mutator transaction binding the contract method 0x82ebd674.
//
// Solidity: function setGovernorAdmin(address newGovernorAdmin) returns()
func (_EVault *EVaultSession) SetGovernorAdmin(newGovernorAdmin common.Address) (*types.Transaction, error) {
	return _EVault.Contract.SetGovernorAdmin(&_EVault.TransactOpts, newGovernorAdmin)
}

// SetGovernorAdmin is a paid mutator transaction binding the contract method 0x82ebd674.
//
// Solidity: function setGovernorAdmin(address newGovernorAdmin) returns()
func (_EVault *EVaultTransactorSession) SetGovernorAdmin(newGovernorAdmin common.Address) (*types.Transaction, error) {
	return _EVault.Contract.SetGovernorAdmin(&_EVault.TransactOpts, newGovernorAdmin)
}

// SetHookConfig is a paid mutator transaction binding the contract method 0xd1a3a308.
//
// Solidity: function setHookConfig(address newHookTarget, uint32 newHookedOps) returns()
func (_EVault *EVaultTransactor) SetHookConfig(opts *bind.TransactOpts, newHookTarget common.Address, newHookedOps uint32) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "setHookConfig", newHookTarget, newHookedOps)
}

// SetHookConfig is a paid mutator transaction binding the contract method 0xd1a3a308.
//
// Solidity: function setHookConfig(address newHookTarget, uint32 newHookedOps) returns()
func (_EVault *EVaultSession) SetHookConfig(newHookTarget common.Address, newHookedOps uint32) (*types.Transaction, error) {
	return _EVault.Contract.SetHookConfig(&_EVault.TransactOpts, newHookTarget, newHookedOps)
}

// SetHookConfig is a paid mutator transaction binding the contract method 0xd1a3a308.
//
// Solidity: function setHookConfig(address newHookTarget, uint32 newHookedOps) returns()
func (_EVault *EVaultTransactorSession) SetHookConfig(newHookTarget common.Address, newHookedOps uint32) (*types.Transaction, error) {
	return _EVault.Contract.SetHookConfig(&_EVault.TransactOpts, newHookTarget, newHookedOps)
}

// SetInterestFee is a paid mutator transaction binding the contract method 0x60cb90ef.
//
// Solidity: function setInterestFee(uint16 newFee) returns()
func (_EVault *EVaultTransactor) SetInterestFee(opts *bind.TransactOpts, newFee uint16) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "setInterestFee", newFee)
}

// SetInterestFee is a paid mutator transaction binding the contract method 0x60cb90ef.
//
// Solidity: function setInterestFee(uint16 newFee) returns()
func (_EVault *EVaultSession) SetInterestFee(newFee uint16) (*types.Transaction, error) {
	return _EVault.Contract.SetInterestFee(&_EVault.TransactOpts, newFee)
}

// SetInterestFee is a paid mutator transaction binding the contract method 0x60cb90ef.
//
// Solidity: function setInterestFee(uint16 newFee) returns()
func (_EVault *EVaultTransactorSession) SetInterestFee(newFee uint16) (*types.Transaction, error) {
	return _EVault.Contract.SetInterestFee(&_EVault.TransactOpts, newFee)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newModel) returns()
func (_EVault *EVaultTransactor) SetInterestRateModel(opts *bind.TransactOpts, newModel common.Address) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "setInterestRateModel", newModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newModel) returns()
func (_EVault *EVaultSession) SetInterestRateModel(newModel common.Address) (*types.Transaction, error) {
	return _EVault.Contract.SetInterestRateModel(&_EVault.TransactOpts, newModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newModel) returns()
func (_EVault *EVaultTransactorSession) SetInterestRateModel(newModel common.Address) (*types.Transaction, error) {
	return _EVault.Contract.SetInterestRateModel(&_EVault.TransactOpts, newModel)
}

// SetLTV is a paid mutator transaction binding the contract method 0x4bca3d5b.
//
// Solidity: function setLTV(address collateral, uint16 borrowLTV, uint16 liquidationLTV, uint32 rampDuration) returns()
func (_EVault *EVaultTransactor) SetLTV(opts *bind.TransactOpts, collateral common.Address, borrowLTV uint16, liquidationLTV uint16, rampDuration uint32) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "setLTV", collateral, borrowLTV, liquidationLTV, rampDuration)
}

// SetLTV is a paid mutator transaction binding the contract method 0x4bca3d5b.
//
// Solidity: function setLTV(address collateral, uint16 borrowLTV, uint16 liquidationLTV, uint32 rampDuration) returns()
func (_EVault *EVaultSession) SetLTV(collateral common.Address, borrowLTV uint16, liquidationLTV uint16, rampDuration uint32) (*types.Transaction, error) {
	return _EVault.Contract.SetLTV(&_EVault.TransactOpts, collateral, borrowLTV, liquidationLTV, rampDuration)
}

// SetLTV is a paid mutator transaction binding the contract method 0x4bca3d5b.
//
// Solidity: function setLTV(address collateral, uint16 borrowLTV, uint16 liquidationLTV, uint32 rampDuration) returns()
func (_EVault *EVaultTransactorSession) SetLTV(collateral common.Address, borrowLTV uint16, liquidationLTV uint16, rampDuration uint32) (*types.Transaction, error) {
	return _EVault.Contract.SetLTV(&_EVault.TransactOpts, collateral, borrowLTV, liquidationLTV, rampDuration)
}

// SetLiquidationCoolOffTime is a paid mutator transaction binding the contract method 0xaf06d3cf.
//
// Solidity: function setLiquidationCoolOffTime(uint16 newCoolOffTime) returns()
func (_EVault *EVaultTransactor) SetLiquidationCoolOffTime(opts *bind.TransactOpts, newCoolOffTime uint16) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "setLiquidationCoolOffTime", newCoolOffTime)
}

// SetLiquidationCoolOffTime is a paid mutator transaction binding the contract method 0xaf06d3cf.
//
// Solidity: function setLiquidationCoolOffTime(uint16 newCoolOffTime) returns()
func (_EVault *EVaultSession) SetLiquidationCoolOffTime(newCoolOffTime uint16) (*types.Transaction, error) {
	return _EVault.Contract.SetLiquidationCoolOffTime(&_EVault.TransactOpts, newCoolOffTime)
}

// SetLiquidationCoolOffTime is a paid mutator transaction binding the contract method 0xaf06d3cf.
//
// Solidity: function setLiquidationCoolOffTime(uint16 newCoolOffTime) returns()
func (_EVault *EVaultTransactorSession) SetLiquidationCoolOffTime(newCoolOffTime uint16) (*types.Transaction, error) {
	return _EVault.Contract.SetLiquidationCoolOffTime(&_EVault.TransactOpts, newCoolOffTime)
}

// SetMaxLiquidationDiscount is a paid mutator transaction binding the contract method 0xb4113ba7.
//
// Solidity: function setMaxLiquidationDiscount(uint16 newDiscount) returns()
func (_EVault *EVaultTransactor) SetMaxLiquidationDiscount(opts *bind.TransactOpts, newDiscount uint16) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "setMaxLiquidationDiscount", newDiscount)
}

// SetMaxLiquidationDiscount is a paid mutator transaction binding the contract method 0xb4113ba7.
//
// Solidity: function setMaxLiquidationDiscount(uint16 newDiscount) returns()
func (_EVault *EVaultSession) SetMaxLiquidationDiscount(newDiscount uint16) (*types.Transaction, error) {
	return _EVault.Contract.SetMaxLiquidationDiscount(&_EVault.TransactOpts, newDiscount)
}

// SetMaxLiquidationDiscount is a paid mutator transaction binding the contract method 0xb4113ba7.
//
// Solidity: function setMaxLiquidationDiscount(uint16 newDiscount) returns()
func (_EVault *EVaultTransactorSession) SetMaxLiquidationDiscount(newDiscount uint16) (*types.Transaction, error) {
	return _EVault.Contract.SetMaxLiquidationDiscount(&_EVault.TransactOpts, newDiscount)
}

// Skim is a paid mutator transaction binding the contract method 0x8d56c639.
//
// Solidity: function skim(uint256 amount, address receiver) returns(uint256)
func (_EVault *EVaultTransactor) Skim(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "skim", amount, receiver)
}

// Skim is a paid mutator transaction binding the contract method 0x8d56c639.
//
// Solidity: function skim(uint256 amount, address receiver) returns(uint256)
func (_EVault *EVaultSession) Skim(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.Contract.Skim(&_EVault.TransactOpts, amount, receiver)
}

// Skim is a paid mutator transaction binding the contract method 0x8d56c639.
//
// Solidity: function skim(uint256 amount, address receiver) returns(uint256)
func (_EVault *EVaultTransactorSession) Skim(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVault.Contract.Skim(&_EVault.TransactOpts, amount, receiver)
}

// Touch is a paid mutator transaction binding the contract method 0xa55526db.
//
// Solidity: function touch() returns()
func (_EVault *EVaultTransactor) Touch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "touch")
}

// Touch is a paid mutator transaction binding the contract method 0xa55526db.
//
// Solidity: function touch() returns()
func (_EVault *EVaultSession) Touch() (*types.Transaction, error) {
	return _EVault.Contract.Touch(&_EVault.TransactOpts)
}

// Touch is a paid mutator transaction binding the contract method 0xa55526db.
//
// Solidity: function touch() returns()
func (_EVault *EVaultTransactorSession) Touch() (*types.Transaction, error) {
	return _EVault.Contract.Touch(&_EVault.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EVault *EVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EVault *EVaultSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVault.Contract.Transfer(&_EVault.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EVault *EVaultTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVault.Contract.Transfer(&_EVault.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EVault *EVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EVault *EVaultSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVault.Contract.TransferFrom(&_EVault.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EVault *EVaultTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVault.Contract.TransferFrom(&_EVault.TransactOpts, from, to, amount)
}

// TransferFromMax is a paid mutator transaction binding the contract method 0xcbfdd7e1.
//
// Solidity: function transferFromMax(address from, address to) returns(bool)
func (_EVault *EVaultTransactor) TransferFromMax(opts *bind.TransactOpts, from common.Address, to common.Address) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "transferFromMax", from, to)
}

// TransferFromMax is a paid mutator transaction binding the contract method 0xcbfdd7e1.
//
// Solidity: function transferFromMax(address from, address to) returns(bool)
func (_EVault *EVaultSession) TransferFromMax(from common.Address, to common.Address) (*types.Transaction, error) {
	return _EVault.Contract.TransferFromMax(&_EVault.TransactOpts, from, to)
}

// TransferFromMax is a paid mutator transaction binding the contract method 0xcbfdd7e1.
//
// Solidity: function transferFromMax(address from, address to) returns(bool)
func (_EVault *EVaultTransactorSession) TransferFromMax(from common.Address, to common.Address) (*types.Transaction, error) {
	return _EVault.Contract.TransferFromMax(&_EVault.TransactOpts, from, to)
}

// ViewDelegate is a paid mutator transaction binding the contract method 0x1fe8b953.
//
// Solidity: function viewDelegate() payable returns()
func (_EVault *EVaultTransactor) ViewDelegate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "viewDelegate")
}

// ViewDelegate is a paid mutator transaction binding the contract method 0x1fe8b953.
//
// Solidity: function viewDelegate() payable returns()
func (_EVault *EVaultSession) ViewDelegate() (*types.Transaction, error) {
	return _EVault.Contract.ViewDelegate(&_EVault.TransactOpts)
}

// ViewDelegate is a paid mutator transaction binding the contract method 0x1fe8b953.
//
// Solidity: function viewDelegate() payable returns()
func (_EVault *EVaultTransactorSession) ViewDelegate() (*types.Transaction, error) {
	return _EVault.Contract.ViewDelegate(&_EVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 amount, address receiver, address owner) returns(uint256)
func (_EVault *EVaultTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EVault.contract.Transact(opts, "withdraw", amount, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 amount, address receiver, address owner) returns(uint256)
func (_EVault *EVaultSession) Withdraw(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EVault.Contract.Withdraw(&_EVault.TransactOpts, amount, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 amount, address receiver, address owner) returns(uint256)
func (_EVault *EVaultTransactorSession) Withdraw(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EVault.Contract.Withdraw(&_EVault.TransactOpts, amount, receiver, owner)
}

// EVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EVault contract.
type EVaultApprovalIterator struct {
	Event *EVaultApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultApproval represents a Approval event raised by the EVault contract.
type EVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EVault *EVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EVaultApprovalIterator{contract: _EVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EVault *EVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultApproval)
				if err := _EVault.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EVault *EVaultFilterer) ParseApproval(log types.Log) (*EVaultApproval, error) {
	event := new(EVaultApproval)
	if err := _EVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultBalanceForwarderStatusIterator is returned from FilterBalanceForwarderStatus and is used to iterate over the raw logs and unpacked data for BalanceForwarderStatus events raised by the EVault contract.
type EVaultBalanceForwarderStatusIterator struct {
	Event *EVaultBalanceForwarderStatus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultBalanceForwarderStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultBalanceForwarderStatus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultBalanceForwarderStatus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultBalanceForwarderStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultBalanceForwarderStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultBalanceForwarderStatus represents a BalanceForwarderStatus event raised by the EVault contract.
type EVaultBalanceForwarderStatus struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBalanceForwarderStatus is a free log retrieval operation binding the contract event 0xc3e011ddce6181dafb5798a536341c7c601913626c31d31744f91b77b7e2412d.
//
// Solidity: event BalanceForwarderStatus(address indexed account, bool status)
func (_EVault *EVaultFilterer) FilterBalanceForwarderStatus(opts *bind.FilterOpts, account []common.Address) (*EVaultBalanceForwarderStatusIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "BalanceForwarderStatus", accountRule)
	if err != nil {
		return nil, err
	}
	return &EVaultBalanceForwarderStatusIterator{contract: _EVault.contract, event: "BalanceForwarderStatus", logs: logs, sub: sub}, nil
}

// WatchBalanceForwarderStatus is a free log subscription operation binding the contract event 0xc3e011ddce6181dafb5798a536341c7c601913626c31d31744f91b77b7e2412d.
//
// Solidity: event BalanceForwarderStatus(address indexed account, bool status)
func (_EVault *EVaultFilterer) WatchBalanceForwarderStatus(opts *bind.WatchOpts, sink chan<- *EVaultBalanceForwarderStatus, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "BalanceForwarderStatus", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultBalanceForwarderStatus)
				if err := _EVault.contract.UnpackLog(event, "BalanceForwarderStatus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBalanceForwarderStatus is a log parse operation binding the contract event 0xc3e011ddce6181dafb5798a536341c7c601913626c31d31744f91b77b7e2412d.
//
// Solidity: event BalanceForwarderStatus(address indexed account, bool status)
func (_EVault *EVaultFilterer) ParseBalanceForwarderStatus(log types.Log) (*EVaultBalanceForwarderStatus, error) {
	event := new(EVaultBalanceForwarderStatus)
	if err := _EVault.contract.UnpackLog(event, "BalanceForwarderStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the EVault contract.
type EVaultBorrowIterator struct {
	Event *EVaultBorrow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultBorrow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultBorrow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultBorrow represents a Borrow event raised by the EVault contract.
type EVaultBorrow struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750.
//
// Solidity: event Borrow(address indexed account, uint256 assets)
func (_EVault *EVaultFilterer) FilterBorrow(opts *bind.FilterOpts, account []common.Address) (*EVaultBorrowIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "Borrow", accountRule)
	if err != nil {
		return nil, err
	}
	return &EVaultBorrowIterator{contract: _EVault.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750.
//
// Solidity: event Borrow(address indexed account, uint256 assets)
func (_EVault *EVaultFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *EVaultBorrow, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "Borrow", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultBorrow)
				if err := _EVault.contract.UnpackLog(event, "Borrow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBorrow is a log parse operation binding the contract event 0xcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750.
//
// Solidity: event Borrow(address indexed account, uint256 assets)
func (_EVault *EVaultFilterer) ParseBorrow(log types.Log) (*EVaultBorrow, error) {
	event := new(EVaultBorrow)
	if err := _EVault.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultConvertFeesIterator is returned from FilterConvertFees and is used to iterate over the raw logs and unpacked data for ConvertFees events raised by the EVault contract.
type EVaultConvertFeesIterator struct {
	Event *EVaultConvertFees // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultConvertFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultConvertFees)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultConvertFees)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultConvertFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultConvertFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultConvertFees represents a ConvertFees event raised by the EVault contract.
type EVaultConvertFees struct {
	Sender           common.Address
	ProtocolReceiver common.Address
	GovernorReceiver common.Address
	ProtocolShares   *big.Int
	GovernorShares   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterConvertFees is a free log retrieval operation binding the contract event 0x4e16b07cac5fe5604af487e07b1b62efc8bd47477b18839f4688d2cae957f965.
//
// Solidity: event ConvertFees(address indexed sender, address indexed protocolReceiver, address indexed governorReceiver, uint256 protocolShares, uint256 governorShares)
func (_EVault *EVaultFilterer) FilterConvertFees(opts *bind.FilterOpts, sender []common.Address, protocolReceiver []common.Address, governorReceiver []common.Address) (*EVaultConvertFeesIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var protocolReceiverRule []interface{}
	for _, protocolReceiverItem := range protocolReceiver {
		protocolReceiverRule = append(protocolReceiverRule, protocolReceiverItem)
	}
	var governorReceiverRule []interface{}
	for _, governorReceiverItem := range governorReceiver {
		governorReceiverRule = append(governorReceiverRule, governorReceiverItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "ConvertFees", senderRule, protocolReceiverRule, governorReceiverRule)
	if err != nil {
		return nil, err
	}
	return &EVaultConvertFeesIterator{contract: _EVault.contract, event: "ConvertFees", logs: logs, sub: sub}, nil
}

// WatchConvertFees is a free log subscription operation binding the contract event 0x4e16b07cac5fe5604af487e07b1b62efc8bd47477b18839f4688d2cae957f965.
//
// Solidity: event ConvertFees(address indexed sender, address indexed protocolReceiver, address indexed governorReceiver, uint256 protocolShares, uint256 governorShares)
func (_EVault *EVaultFilterer) WatchConvertFees(opts *bind.WatchOpts, sink chan<- *EVaultConvertFees, sender []common.Address, protocolReceiver []common.Address, governorReceiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var protocolReceiverRule []interface{}
	for _, protocolReceiverItem := range protocolReceiver {
		protocolReceiverRule = append(protocolReceiverRule, protocolReceiverItem)
	}
	var governorReceiverRule []interface{}
	for _, governorReceiverItem := range governorReceiver {
		governorReceiverRule = append(governorReceiverRule, governorReceiverItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "ConvertFees", senderRule, protocolReceiverRule, governorReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultConvertFees)
				if err := _EVault.contract.UnpackLog(event, "ConvertFees", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConvertFees is a log parse operation binding the contract event 0x4e16b07cac5fe5604af487e07b1b62efc8bd47477b18839f4688d2cae957f965.
//
// Solidity: event ConvertFees(address indexed sender, address indexed protocolReceiver, address indexed governorReceiver, uint256 protocolShares, uint256 governorShares)
func (_EVault *EVaultFilterer) ParseConvertFees(log types.Log) (*EVaultConvertFees, error) {
	event := new(EVaultConvertFees)
	if err := _EVault.contract.UnpackLog(event, "ConvertFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultDebtSocializedIterator is returned from FilterDebtSocialized and is used to iterate over the raw logs and unpacked data for DebtSocialized events raised by the EVault contract.
type EVaultDebtSocializedIterator struct {
	Event *EVaultDebtSocialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultDebtSocializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultDebtSocialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultDebtSocialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultDebtSocializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultDebtSocializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultDebtSocialized represents a DebtSocialized event raised by the EVault contract.
type EVaultDebtSocialized struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDebtSocialized is a free log retrieval operation binding the contract event 0xe786d0bc2e83bf230ed9895a9c4d7756ab0c6e22eb8a4ff69c161ece76bd36df.
//
// Solidity: event DebtSocialized(address indexed account, uint256 assets)
func (_EVault *EVaultFilterer) FilterDebtSocialized(opts *bind.FilterOpts, account []common.Address) (*EVaultDebtSocializedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "DebtSocialized", accountRule)
	if err != nil {
		return nil, err
	}
	return &EVaultDebtSocializedIterator{contract: _EVault.contract, event: "DebtSocialized", logs: logs, sub: sub}, nil
}

// WatchDebtSocialized is a free log subscription operation binding the contract event 0xe786d0bc2e83bf230ed9895a9c4d7756ab0c6e22eb8a4ff69c161ece76bd36df.
//
// Solidity: event DebtSocialized(address indexed account, uint256 assets)
func (_EVault *EVaultFilterer) WatchDebtSocialized(opts *bind.WatchOpts, sink chan<- *EVaultDebtSocialized, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "DebtSocialized", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultDebtSocialized)
				if err := _EVault.contract.UnpackLog(event, "DebtSocialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDebtSocialized is a log parse operation binding the contract event 0xe786d0bc2e83bf230ed9895a9c4d7756ab0c6e22eb8a4ff69c161ece76bd36df.
//
// Solidity: event DebtSocialized(address indexed account, uint256 assets)
func (_EVault *EVaultFilterer) ParseDebtSocialized(log types.Log) (*EVaultDebtSocialized, error) {
	event := new(EVaultDebtSocialized)
	if err := _EVault.contract.UnpackLog(event, "DebtSocialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the EVault contract.
type EVaultDepositIterator struct {
	Event *EVaultDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultDeposit represents a Deposit event raised by the EVault contract.
type EVaultDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_EVault *EVaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*EVaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EVaultDepositIterator{contract: _EVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_EVault *EVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EVaultDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultDeposit)
				if err := _EVault.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_EVault *EVaultFilterer) ParseDeposit(log types.Log) (*EVaultDeposit, error) {
	event := new(EVaultDeposit)
	if err := _EVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultEVaultCreatedIterator is returned from FilterEVaultCreated and is used to iterate over the raw logs and unpacked data for EVaultCreated events raised by the EVault contract.
type EVaultEVaultCreatedIterator struct {
	Event *EVaultEVaultCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultEVaultCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultEVaultCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultEVaultCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultEVaultCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultEVaultCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultEVaultCreated represents a EVaultCreated event raised by the EVault contract.
type EVaultEVaultCreated struct {
	Creator common.Address
	Asset   common.Address
	DToken  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEVaultCreated is a free log retrieval operation binding the contract event 0x0cd345140b9008a43f99a999a328ece572a0193e8c8bf5f5755585e6f293b85e.
//
// Solidity: event EVaultCreated(address indexed creator, address indexed asset, address dToken)
func (_EVault *EVaultFilterer) FilterEVaultCreated(opts *bind.FilterOpts, creator []common.Address, asset []common.Address) (*EVaultEVaultCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "EVaultCreated", creatorRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &EVaultEVaultCreatedIterator{contract: _EVault.contract, event: "EVaultCreated", logs: logs, sub: sub}, nil
}

// WatchEVaultCreated is a free log subscription operation binding the contract event 0x0cd345140b9008a43f99a999a328ece572a0193e8c8bf5f5755585e6f293b85e.
//
// Solidity: event EVaultCreated(address indexed creator, address indexed asset, address dToken)
func (_EVault *EVaultFilterer) WatchEVaultCreated(opts *bind.WatchOpts, sink chan<- *EVaultEVaultCreated, creator []common.Address, asset []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "EVaultCreated", creatorRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultEVaultCreated)
				if err := _EVault.contract.UnpackLog(event, "EVaultCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEVaultCreated is a log parse operation binding the contract event 0x0cd345140b9008a43f99a999a328ece572a0193e8c8bf5f5755585e6f293b85e.
//
// Solidity: event EVaultCreated(address indexed creator, address indexed asset, address dToken)
func (_EVault *EVaultFilterer) ParseEVaultCreated(log types.Log) (*EVaultEVaultCreated, error) {
	event := new(EVaultEVaultCreated)
	if err := _EVault.contract.UnpackLog(event, "EVaultCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultGovSetCapsIterator is returned from FilterGovSetCaps and is used to iterate over the raw logs and unpacked data for GovSetCaps events raised by the EVault contract.
type EVaultGovSetCapsIterator struct {
	Event *EVaultGovSetCaps // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultGovSetCapsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultGovSetCaps)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultGovSetCaps)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultGovSetCapsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultGovSetCapsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultGovSetCaps represents a GovSetCaps event raised by the EVault contract.
type EVaultGovSetCaps struct {
	NewSupplyCap uint16
	NewBorrowCap uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGovSetCaps is a free log retrieval operation binding the contract event 0xadbdcd178dfddc478805a3703b6cf3b72ca5e78ecebacffe1aad03188cc1cbf4.
//
// Solidity: event GovSetCaps(uint16 newSupplyCap, uint16 newBorrowCap)
func (_EVault *EVaultFilterer) FilterGovSetCaps(opts *bind.FilterOpts) (*EVaultGovSetCapsIterator, error) {

	logs, sub, err := _EVault.contract.FilterLogs(opts, "GovSetCaps")
	if err != nil {
		return nil, err
	}
	return &EVaultGovSetCapsIterator{contract: _EVault.contract, event: "GovSetCaps", logs: logs, sub: sub}, nil
}

// WatchGovSetCaps is a free log subscription operation binding the contract event 0xadbdcd178dfddc478805a3703b6cf3b72ca5e78ecebacffe1aad03188cc1cbf4.
//
// Solidity: event GovSetCaps(uint16 newSupplyCap, uint16 newBorrowCap)
func (_EVault *EVaultFilterer) WatchGovSetCaps(opts *bind.WatchOpts, sink chan<- *EVaultGovSetCaps) (event.Subscription, error) {

	logs, sub, err := _EVault.contract.WatchLogs(opts, "GovSetCaps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultGovSetCaps)
				if err := _EVault.contract.UnpackLog(event, "GovSetCaps", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovSetCaps is a log parse operation binding the contract event 0xadbdcd178dfddc478805a3703b6cf3b72ca5e78ecebacffe1aad03188cc1cbf4.
//
// Solidity: event GovSetCaps(uint16 newSupplyCap, uint16 newBorrowCap)
func (_EVault *EVaultFilterer) ParseGovSetCaps(log types.Log) (*EVaultGovSetCaps, error) {
	event := new(EVaultGovSetCaps)
	if err := _EVault.contract.UnpackLog(event, "GovSetCaps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultGovSetConfigFlagsIterator is returned from FilterGovSetConfigFlags and is used to iterate over the raw logs and unpacked data for GovSetConfigFlags events raised by the EVault contract.
type EVaultGovSetConfigFlagsIterator struct {
	Event *EVaultGovSetConfigFlags // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultGovSetConfigFlagsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultGovSetConfigFlags)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultGovSetConfigFlags)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultGovSetConfigFlagsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultGovSetConfigFlagsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultGovSetConfigFlags represents a GovSetConfigFlags event raised by the EVault contract.
type EVaultGovSetConfigFlags struct {
	NewConfigFlags uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovSetConfigFlags is a free log retrieval operation binding the contract event 0xe7f84c52c0ef295afe77de8cb30516d6f28d50306f979b45776dd1b619ae5ffc.
//
// Solidity: event GovSetConfigFlags(uint32 newConfigFlags)
func (_EVault *EVaultFilterer) FilterGovSetConfigFlags(opts *bind.FilterOpts) (*EVaultGovSetConfigFlagsIterator, error) {

	logs, sub, err := _EVault.contract.FilterLogs(opts, "GovSetConfigFlags")
	if err != nil {
		return nil, err
	}
	return &EVaultGovSetConfigFlagsIterator{contract: _EVault.contract, event: "GovSetConfigFlags", logs: logs, sub: sub}, nil
}

// WatchGovSetConfigFlags is a free log subscription operation binding the contract event 0xe7f84c52c0ef295afe77de8cb30516d6f28d50306f979b45776dd1b619ae5ffc.
//
// Solidity: event GovSetConfigFlags(uint32 newConfigFlags)
func (_EVault *EVaultFilterer) WatchGovSetConfigFlags(opts *bind.WatchOpts, sink chan<- *EVaultGovSetConfigFlags) (event.Subscription, error) {

	logs, sub, err := _EVault.contract.WatchLogs(opts, "GovSetConfigFlags")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultGovSetConfigFlags)
				if err := _EVault.contract.UnpackLog(event, "GovSetConfigFlags", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovSetConfigFlags is a log parse operation binding the contract event 0xe7f84c52c0ef295afe77de8cb30516d6f28d50306f979b45776dd1b619ae5ffc.
//
// Solidity: event GovSetConfigFlags(uint32 newConfigFlags)
func (_EVault *EVaultFilterer) ParseGovSetConfigFlags(log types.Log) (*EVaultGovSetConfigFlags, error) {
	event := new(EVaultGovSetConfigFlags)
	if err := _EVault.contract.UnpackLog(event, "GovSetConfigFlags", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultGovSetFeeReceiverIterator is returned from FilterGovSetFeeReceiver and is used to iterate over the raw logs and unpacked data for GovSetFeeReceiver events raised by the EVault contract.
type EVaultGovSetFeeReceiverIterator struct {
	Event *EVaultGovSetFeeReceiver // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultGovSetFeeReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultGovSetFeeReceiver)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultGovSetFeeReceiver)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultGovSetFeeReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultGovSetFeeReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultGovSetFeeReceiver represents a GovSetFeeReceiver event raised by the EVault contract.
type EVaultGovSetFeeReceiver struct {
	NewFeeReceiver common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovSetFeeReceiver is a free log retrieval operation binding the contract event 0x836a1afef2bc89de2cb4713cc8d312fccf2ff835230721c5f41f13374707413a.
//
// Solidity: event GovSetFeeReceiver(address indexed newFeeReceiver)
func (_EVault *EVaultFilterer) FilterGovSetFeeReceiver(opts *bind.FilterOpts, newFeeReceiver []common.Address) (*EVaultGovSetFeeReceiverIterator, error) {

	var newFeeReceiverRule []interface{}
	for _, newFeeReceiverItem := range newFeeReceiver {
		newFeeReceiverRule = append(newFeeReceiverRule, newFeeReceiverItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "GovSetFeeReceiver", newFeeReceiverRule)
	if err != nil {
		return nil, err
	}
	return &EVaultGovSetFeeReceiverIterator{contract: _EVault.contract, event: "GovSetFeeReceiver", logs: logs, sub: sub}, nil
}

// WatchGovSetFeeReceiver is a free log subscription operation binding the contract event 0x836a1afef2bc89de2cb4713cc8d312fccf2ff835230721c5f41f13374707413a.
//
// Solidity: event GovSetFeeReceiver(address indexed newFeeReceiver)
func (_EVault *EVaultFilterer) WatchGovSetFeeReceiver(opts *bind.WatchOpts, sink chan<- *EVaultGovSetFeeReceiver, newFeeReceiver []common.Address) (event.Subscription, error) {

	var newFeeReceiverRule []interface{}
	for _, newFeeReceiverItem := range newFeeReceiver {
		newFeeReceiverRule = append(newFeeReceiverRule, newFeeReceiverItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "GovSetFeeReceiver", newFeeReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultGovSetFeeReceiver)
				if err := _EVault.contract.UnpackLog(event, "GovSetFeeReceiver", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovSetFeeReceiver is a log parse operation binding the contract event 0x836a1afef2bc89de2cb4713cc8d312fccf2ff835230721c5f41f13374707413a.
//
// Solidity: event GovSetFeeReceiver(address indexed newFeeReceiver)
func (_EVault *EVaultFilterer) ParseGovSetFeeReceiver(log types.Log) (*EVaultGovSetFeeReceiver, error) {
	event := new(EVaultGovSetFeeReceiver)
	if err := _EVault.contract.UnpackLog(event, "GovSetFeeReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultGovSetGovernorAdminIterator is returned from FilterGovSetGovernorAdmin and is used to iterate over the raw logs and unpacked data for GovSetGovernorAdmin events raised by the EVault contract.
type EVaultGovSetGovernorAdminIterator struct {
	Event *EVaultGovSetGovernorAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultGovSetGovernorAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultGovSetGovernorAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultGovSetGovernorAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultGovSetGovernorAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultGovSetGovernorAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultGovSetGovernorAdmin represents a GovSetGovernorAdmin event raised by the EVault contract.
type EVaultGovSetGovernorAdmin struct {
	NewGovernorAdmin common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterGovSetGovernorAdmin is a free log retrieval operation binding the contract event 0x1c145a4cd16d4148579b0f2296884ac4aa47536e4ef10a32e1cdc0dc3dd20ea4.
//
// Solidity: event GovSetGovernorAdmin(address indexed newGovernorAdmin)
func (_EVault *EVaultFilterer) FilterGovSetGovernorAdmin(opts *bind.FilterOpts, newGovernorAdmin []common.Address) (*EVaultGovSetGovernorAdminIterator, error) {

	var newGovernorAdminRule []interface{}
	for _, newGovernorAdminItem := range newGovernorAdmin {
		newGovernorAdminRule = append(newGovernorAdminRule, newGovernorAdminItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "GovSetGovernorAdmin", newGovernorAdminRule)
	if err != nil {
		return nil, err
	}
	return &EVaultGovSetGovernorAdminIterator{contract: _EVault.contract, event: "GovSetGovernorAdmin", logs: logs, sub: sub}, nil
}

// WatchGovSetGovernorAdmin is a free log subscription operation binding the contract event 0x1c145a4cd16d4148579b0f2296884ac4aa47536e4ef10a32e1cdc0dc3dd20ea4.
//
// Solidity: event GovSetGovernorAdmin(address indexed newGovernorAdmin)
func (_EVault *EVaultFilterer) WatchGovSetGovernorAdmin(opts *bind.WatchOpts, sink chan<- *EVaultGovSetGovernorAdmin, newGovernorAdmin []common.Address) (event.Subscription, error) {

	var newGovernorAdminRule []interface{}
	for _, newGovernorAdminItem := range newGovernorAdmin {
		newGovernorAdminRule = append(newGovernorAdminRule, newGovernorAdminItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "GovSetGovernorAdmin", newGovernorAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultGovSetGovernorAdmin)
				if err := _EVault.contract.UnpackLog(event, "GovSetGovernorAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovSetGovernorAdmin is a log parse operation binding the contract event 0x1c145a4cd16d4148579b0f2296884ac4aa47536e4ef10a32e1cdc0dc3dd20ea4.
//
// Solidity: event GovSetGovernorAdmin(address indexed newGovernorAdmin)
func (_EVault *EVaultFilterer) ParseGovSetGovernorAdmin(log types.Log) (*EVaultGovSetGovernorAdmin, error) {
	event := new(EVaultGovSetGovernorAdmin)
	if err := _EVault.contract.UnpackLog(event, "GovSetGovernorAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultGovSetHookConfigIterator is returned from FilterGovSetHookConfig and is used to iterate over the raw logs and unpacked data for GovSetHookConfig events raised by the EVault contract.
type EVaultGovSetHookConfigIterator struct {
	Event *EVaultGovSetHookConfig // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultGovSetHookConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultGovSetHookConfig)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultGovSetHookConfig)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultGovSetHookConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultGovSetHookConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultGovSetHookConfig represents a GovSetHookConfig event raised by the EVault contract.
type EVaultGovSetHookConfig struct {
	NewHookTarget common.Address
	NewHookedOps  uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGovSetHookConfig is a free log retrieval operation binding the contract event 0xabadffb695acdb6863cd1324a91e5c359712b9110a55f9103774e2fb67dedb6a.
//
// Solidity: event GovSetHookConfig(address indexed newHookTarget, uint32 newHookedOps)
func (_EVault *EVaultFilterer) FilterGovSetHookConfig(opts *bind.FilterOpts, newHookTarget []common.Address) (*EVaultGovSetHookConfigIterator, error) {

	var newHookTargetRule []interface{}
	for _, newHookTargetItem := range newHookTarget {
		newHookTargetRule = append(newHookTargetRule, newHookTargetItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "GovSetHookConfig", newHookTargetRule)
	if err != nil {
		return nil, err
	}
	return &EVaultGovSetHookConfigIterator{contract: _EVault.contract, event: "GovSetHookConfig", logs: logs, sub: sub}, nil
}

// WatchGovSetHookConfig is a free log subscription operation binding the contract event 0xabadffb695acdb6863cd1324a91e5c359712b9110a55f9103774e2fb67dedb6a.
//
// Solidity: event GovSetHookConfig(address indexed newHookTarget, uint32 newHookedOps)
func (_EVault *EVaultFilterer) WatchGovSetHookConfig(opts *bind.WatchOpts, sink chan<- *EVaultGovSetHookConfig, newHookTarget []common.Address) (event.Subscription, error) {

	var newHookTargetRule []interface{}
	for _, newHookTargetItem := range newHookTarget {
		newHookTargetRule = append(newHookTargetRule, newHookTargetItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "GovSetHookConfig", newHookTargetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultGovSetHookConfig)
				if err := _EVault.contract.UnpackLog(event, "GovSetHookConfig", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovSetHookConfig is a log parse operation binding the contract event 0xabadffb695acdb6863cd1324a91e5c359712b9110a55f9103774e2fb67dedb6a.
//
// Solidity: event GovSetHookConfig(address indexed newHookTarget, uint32 newHookedOps)
func (_EVault *EVaultFilterer) ParseGovSetHookConfig(log types.Log) (*EVaultGovSetHookConfig, error) {
	event := new(EVaultGovSetHookConfig)
	if err := _EVault.contract.UnpackLog(event, "GovSetHookConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultGovSetInterestFeeIterator is returned from FilterGovSetInterestFee and is used to iterate over the raw logs and unpacked data for GovSetInterestFee events raised by the EVault contract.
type EVaultGovSetInterestFeeIterator struct {
	Event *EVaultGovSetInterestFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultGovSetInterestFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultGovSetInterestFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultGovSetInterestFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultGovSetInterestFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultGovSetInterestFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultGovSetInterestFee represents a GovSetInterestFee event raised by the EVault contract.
type EVaultGovSetInterestFee struct {
	NewFee uint16
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGovSetInterestFee is a free log retrieval operation binding the contract event 0x634a58674e370383703eff32d9d4e4b3d1add94d50e8bcb631b04995d8e47341.
//
// Solidity: event GovSetInterestFee(uint16 newFee)
func (_EVault *EVaultFilterer) FilterGovSetInterestFee(opts *bind.FilterOpts) (*EVaultGovSetInterestFeeIterator, error) {

	logs, sub, err := _EVault.contract.FilterLogs(opts, "GovSetInterestFee")
	if err != nil {
		return nil, err
	}
	return &EVaultGovSetInterestFeeIterator{contract: _EVault.contract, event: "GovSetInterestFee", logs: logs, sub: sub}, nil
}

// WatchGovSetInterestFee is a free log subscription operation binding the contract event 0x634a58674e370383703eff32d9d4e4b3d1add94d50e8bcb631b04995d8e47341.
//
// Solidity: event GovSetInterestFee(uint16 newFee)
func (_EVault *EVaultFilterer) WatchGovSetInterestFee(opts *bind.WatchOpts, sink chan<- *EVaultGovSetInterestFee) (event.Subscription, error) {

	logs, sub, err := _EVault.contract.WatchLogs(opts, "GovSetInterestFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultGovSetInterestFee)
				if err := _EVault.contract.UnpackLog(event, "GovSetInterestFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovSetInterestFee is a log parse operation binding the contract event 0x634a58674e370383703eff32d9d4e4b3d1add94d50e8bcb631b04995d8e47341.
//
// Solidity: event GovSetInterestFee(uint16 newFee)
func (_EVault *EVaultFilterer) ParseGovSetInterestFee(log types.Log) (*EVaultGovSetInterestFee, error) {
	event := new(EVaultGovSetInterestFee)
	if err := _EVault.contract.UnpackLog(event, "GovSetInterestFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultGovSetInterestRateModelIterator is returned from FilterGovSetInterestRateModel and is used to iterate over the raw logs and unpacked data for GovSetInterestRateModel events raised by the EVault contract.
type EVaultGovSetInterestRateModelIterator struct {
	Event *EVaultGovSetInterestRateModel // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultGovSetInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultGovSetInterestRateModel)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultGovSetInterestRateModel)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultGovSetInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultGovSetInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultGovSetInterestRateModel represents a GovSetInterestRateModel event raised by the EVault contract.
type EVaultGovSetInterestRateModel struct {
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterGovSetInterestRateModel is a free log retrieval operation binding the contract event 0xe5f2a795fc5f8baf1b05659293834c88859298226d87422c88624b4c9f4d3a43.
//
// Solidity: event GovSetInterestRateModel(address newInterestRateModel)
func (_EVault *EVaultFilterer) FilterGovSetInterestRateModel(opts *bind.FilterOpts) (*EVaultGovSetInterestRateModelIterator, error) {

	logs, sub, err := _EVault.contract.FilterLogs(opts, "GovSetInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &EVaultGovSetInterestRateModelIterator{contract: _EVault.contract, event: "GovSetInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchGovSetInterestRateModel is a free log subscription operation binding the contract event 0xe5f2a795fc5f8baf1b05659293834c88859298226d87422c88624b4c9f4d3a43.
//
// Solidity: event GovSetInterestRateModel(address newInterestRateModel)
func (_EVault *EVaultFilterer) WatchGovSetInterestRateModel(opts *bind.WatchOpts, sink chan<- *EVaultGovSetInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _EVault.contract.WatchLogs(opts, "GovSetInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultGovSetInterestRateModel)
				if err := _EVault.contract.UnpackLog(event, "GovSetInterestRateModel", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovSetInterestRateModel is a log parse operation binding the contract event 0xe5f2a795fc5f8baf1b05659293834c88859298226d87422c88624b4c9f4d3a43.
//
// Solidity: event GovSetInterestRateModel(address newInterestRateModel)
func (_EVault *EVaultFilterer) ParseGovSetInterestRateModel(log types.Log) (*EVaultGovSetInterestRateModel, error) {
	event := new(EVaultGovSetInterestRateModel)
	if err := _EVault.contract.UnpackLog(event, "GovSetInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultGovSetLTVIterator is returned from FilterGovSetLTV and is used to iterate over the raw logs and unpacked data for GovSetLTV events raised by the EVault contract.
type EVaultGovSetLTVIterator struct {
	Event *EVaultGovSetLTV // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultGovSetLTVIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultGovSetLTV)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultGovSetLTV)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultGovSetLTVIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultGovSetLTVIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultGovSetLTV represents a GovSetLTV event raised by the EVault contract.
type EVaultGovSetLTV struct {
	Collateral            common.Address
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovSetLTV is a free log retrieval operation binding the contract event 0xc69392046c26324e9eee913208811542aabcbde6a41ce9ee3b45473b18eb3c76.
//
// Solidity: event GovSetLTV(address indexed collateral, uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EVault *EVaultFilterer) FilterGovSetLTV(opts *bind.FilterOpts, collateral []common.Address) (*EVaultGovSetLTVIterator, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "GovSetLTV", collateralRule)
	if err != nil {
		return nil, err
	}
	return &EVaultGovSetLTVIterator{contract: _EVault.contract, event: "GovSetLTV", logs: logs, sub: sub}, nil
}

// WatchGovSetLTV is a free log subscription operation binding the contract event 0xc69392046c26324e9eee913208811542aabcbde6a41ce9ee3b45473b18eb3c76.
//
// Solidity: event GovSetLTV(address indexed collateral, uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EVault *EVaultFilterer) WatchGovSetLTV(opts *bind.WatchOpts, sink chan<- *EVaultGovSetLTV, collateral []common.Address) (event.Subscription, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "GovSetLTV", collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultGovSetLTV)
				if err := _EVault.contract.UnpackLog(event, "GovSetLTV", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovSetLTV is a log parse operation binding the contract event 0xc69392046c26324e9eee913208811542aabcbde6a41ce9ee3b45473b18eb3c76.
//
// Solidity: event GovSetLTV(address indexed collateral, uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EVault *EVaultFilterer) ParseGovSetLTV(log types.Log) (*EVaultGovSetLTV, error) {
	event := new(EVaultGovSetLTV)
	if err := _EVault.contract.UnpackLog(event, "GovSetLTV", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultGovSetLiquidationCoolOffTimeIterator is returned from FilterGovSetLiquidationCoolOffTime and is used to iterate over the raw logs and unpacked data for GovSetLiquidationCoolOffTime events raised by the EVault contract.
type EVaultGovSetLiquidationCoolOffTimeIterator struct {
	Event *EVaultGovSetLiquidationCoolOffTime // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultGovSetLiquidationCoolOffTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultGovSetLiquidationCoolOffTime)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultGovSetLiquidationCoolOffTime)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultGovSetLiquidationCoolOffTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultGovSetLiquidationCoolOffTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultGovSetLiquidationCoolOffTime represents a GovSetLiquidationCoolOffTime event raised by the EVault contract.
type EVaultGovSetLiquidationCoolOffTime struct {
	NewCoolOffTime uint16
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovSetLiquidationCoolOffTime is a free log retrieval operation binding the contract event 0xdf4edc1d288e7b3306b287d03fd77b2070b8b308c702bf7297f72d928175dfa5.
//
// Solidity: event GovSetLiquidationCoolOffTime(uint16 newCoolOffTime)
func (_EVault *EVaultFilterer) FilterGovSetLiquidationCoolOffTime(opts *bind.FilterOpts) (*EVaultGovSetLiquidationCoolOffTimeIterator, error) {

	logs, sub, err := _EVault.contract.FilterLogs(opts, "GovSetLiquidationCoolOffTime")
	if err != nil {
		return nil, err
	}
	return &EVaultGovSetLiquidationCoolOffTimeIterator{contract: _EVault.contract, event: "GovSetLiquidationCoolOffTime", logs: logs, sub: sub}, nil
}

// WatchGovSetLiquidationCoolOffTime is a free log subscription operation binding the contract event 0xdf4edc1d288e7b3306b287d03fd77b2070b8b308c702bf7297f72d928175dfa5.
//
// Solidity: event GovSetLiquidationCoolOffTime(uint16 newCoolOffTime)
func (_EVault *EVaultFilterer) WatchGovSetLiquidationCoolOffTime(opts *bind.WatchOpts, sink chan<- *EVaultGovSetLiquidationCoolOffTime) (event.Subscription, error) {

	logs, sub, err := _EVault.contract.WatchLogs(opts, "GovSetLiquidationCoolOffTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultGovSetLiquidationCoolOffTime)
				if err := _EVault.contract.UnpackLog(event, "GovSetLiquidationCoolOffTime", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovSetLiquidationCoolOffTime is a log parse operation binding the contract event 0xdf4edc1d288e7b3306b287d03fd77b2070b8b308c702bf7297f72d928175dfa5.
//
// Solidity: event GovSetLiquidationCoolOffTime(uint16 newCoolOffTime)
func (_EVault *EVaultFilterer) ParseGovSetLiquidationCoolOffTime(log types.Log) (*EVaultGovSetLiquidationCoolOffTime, error) {
	event := new(EVaultGovSetLiquidationCoolOffTime)
	if err := _EVault.contract.UnpackLog(event, "GovSetLiquidationCoolOffTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultGovSetMaxLiquidationDiscountIterator is returned from FilterGovSetMaxLiquidationDiscount and is used to iterate over the raw logs and unpacked data for GovSetMaxLiquidationDiscount events raised by the EVault contract.
type EVaultGovSetMaxLiquidationDiscountIterator struct {
	Event *EVaultGovSetMaxLiquidationDiscount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultGovSetMaxLiquidationDiscountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultGovSetMaxLiquidationDiscount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultGovSetMaxLiquidationDiscount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultGovSetMaxLiquidationDiscountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultGovSetMaxLiquidationDiscountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultGovSetMaxLiquidationDiscount represents a GovSetMaxLiquidationDiscount event raised by the EVault contract.
type EVaultGovSetMaxLiquidationDiscount struct {
	NewDiscount uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGovSetMaxLiquidationDiscount is a free log retrieval operation binding the contract event 0x558a63d245d08220a137de3573129d3921e70e806adccf3a068c4723b9b3322d.
//
// Solidity: event GovSetMaxLiquidationDiscount(uint16 newDiscount)
func (_EVault *EVaultFilterer) FilterGovSetMaxLiquidationDiscount(opts *bind.FilterOpts) (*EVaultGovSetMaxLiquidationDiscountIterator, error) {

	logs, sub, err := _EVault.contract.FilterLogs(opts, "GovSetMaxLiquidationDiscount")
	if err != nil {
		return nil, err
	}
	return &EVaultGovSetMaxLiquidationDiscountIterator{contract: _EVault.contract, event: "GovSetMaxLiquidationDiscount", logs: logs, sub: sub}, nil
}

// WatchGovSetMaxLiquidationDiscount is a free log subscription operation binding the contract event 0x558a63d245d08220a137de3573129d3921e70e806adccf3a068c4723b9b3322d.
//
// Solidity: event GovSetMaxLiquidationDiscount(uint16 newDiscount)
func (_EVault *EVaultFilterer) WatchGovSetMaxLiquidationDiscount(opts *bind.WatchOpts, sink chan<- *EVaultGovSetMaxLiquidationDiscount) (event.Subscription, error) {

	logs, sub, err := _EVault.contract.WatchLogs(opts, "GovSetMaxLiquidationDiscount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultGovSetMaxLiquidationDiscount)
				if err := _EVault.contract.UnpackLog(event, "GovSetMaxLiquidationDiscount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovSetMaxLiquidationDiscount is a log parse operation binding the contract event 0x558a63d245d08220a137de3573129d3921e70e806adccf3a068c4723b9b3322d.
//
// Solidity: event GovSetMaxLiquidationDiscount(uint16 newDiscount)
func (_EVault *EVaultFilterer) ParseGovSetMaxLiquidationDiscount(log types.Log) (*EVaultGovSetMaxLiquidationDiscount, error) {
	event := new(EVaultGovSetMaxLiquidationDiscount)
	if err := _EVault.contract.UnpackLog(event, "GovSetMaxLiquidationDiscount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultInterestAccruedIterator is returned from FilterInterestAccrued and is used to iterate over the raw logs and unpacked data for InterestAccrued events raised by the EVault contract.
type EVaultInterestAccruedIterator struct {
	Event *EVaultInterestAccrued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultInterestAccruedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultInterestAccrued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultInterestAccrued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultInterestAccruedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultInterestAccruedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultInterestAccrued represents a InterestAccrued event raised by the EVault contract.
type EVaultInterestAccrued struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInterestAccrued is a free log retrieval operation binding the contract event 0x5e804d42ae3b860f881d11cb44a4bb1f2f0d5b3d081f5539a32d6f97b629d978.
//
// Solidity: event InterestAccrued(address indexed account, uint256 assets)
func (_EVault *EVaultFilterer) FilterInterestAccrued(opts *bind.FilterOpts, account []common.Address) (*EVaultInterestAccruedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "InterestAccrued", accountRule)
	if err != nil {
		return nil, err
	}
	return &EVaultInterestAccruedIterator{contract: _EVault.contract, event: "InterestAccrued", logs: logs, sub: sub}, nil
}

// WatchInterestAccrued is a free log subscription operation binding the contract event 0x5e804d42ae3b860f881d11cb44a4bb1f2f0d5b3d081f5539a32d6f97b629d978.
//
// Solidity: event InterestAccrued(address indexed account, uint256 assets)
func (_EVault *EVaultFilterer) WatchInterestAccrued(opts *bind.WatchOpts, sink chan<- *EVaultInterestAccrued, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "InterestAccrued", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultInterestAccrued)
				if err := _EVault.contract.UnpackLog(event, "InterestAccrued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInterestAccrued is a log parse operation binding the contract event 0x5e804d42ae3b860f881d11cb44a4bb1f2f0d5b3d081f5539a32d6f97b629d978.
//
// Solidity: event InterestAccrued(address indexed account, uint256 assets)
func (_EVault *EVaultFilterer) ParseInterestAccrued(log types.Log) (*EVaultInterestAccrued, error) {
	event := new(EVaultInterestAccrued)
	if err := _EVault.contract.UnpackLog(event, "InterestAccrued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the EVault contract.
type EVaultLiquidateIterator struct {
	Event *EVaultLiquidate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultLiquidate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultLiquidate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultLiquidate represents a Liquidate event raised by the EVault contract.
type EVaultLiquidate struct {
	Liquidator   common.Address
	Violator     common.Address
	Collateral   common.Address
	RepayAssets  *big.Int
	YieldBalance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0x8246cc71ab01533b5bebc672a636df812f10637ad720797319d5741d5ebb3962.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed violator, address collateral, uint256 repayAssets, uint256 yieldBalance)
func (_EVault *EVaultFilterer) FilterLiquidate(opts *bind.FilterOpts, liquidator []common.Address, violator []common.Address) (*EVaultLiquidateIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var violatorRule []interface{}
	for _, violatorItem := range violator {
		violatorRule = append(violatorRule, violatorItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "Liquidate", liquidatorRule, violatorRule)
	if err != nil {
		return nil, err
	}
	return &EVaultLiquidateIterator{contract: _EVault.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0x8246cc71ab01533b5bebc672a636df812f10637ad720797319d5741d5ebb3962.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed violator, address collateral, uint256 repayAssets, uint256 yieldBalance)
func (_EVault *EVaultFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *EVaultLiquidate, liquidator []common.Address, violator []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var violatorRule []interface{}
	for _, violatorItem := range violator {
		violatorRule = append(violatorRule, violatorItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "Liquidate", liquidatorRule, violatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultLiquidate)
				if err := _EVault.contract.UnpackLog(event, "Liquidate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidate is a log parse operation binding the contract event 0x8246cc71ab01533b5bebc672a636df812f10637ad720797319d5741d5ebb3962.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed violator, address collateral, uint256 repayAssets, uint256 yieldBalance)
func (_EVault *EVaultFilterer) ParseLiquidate(log types.Log) (*EVaultLiquidate, error) {
	event := new(EVaultLiquidate)
	if err := _EVault.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultPullDebtIterator is returned from FilterPullDebt and is used to iterate over the raw logs and unpacked data for PullDebt events raised by the EVault contract.
type EVaultPullDebtIterator struct {
	Event *EVaultPullDebt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultPullDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultPullDebt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultPullDebt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultPullDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultPullDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultPullDebt represents a PullDebt event raised by the EVault contract.
type EVaultPullDebt struct {
	From   common.Address
	To     common.Address
	Assets *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPullDebt is a free log retrieval operation binding the contract event 0xe6d0bfd9025bf59969101a13cf02e3ba2811b533816c47d7155546c7c8a1048f.
//
// Solidity: event PullDebt(address indexed from, address indexed to, uint256 assets)
func (_EVault *EVaultFilterer) FilterPullDebt(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVaultPullDebtIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "PullDebt", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVaultPullDebtIterator{contract: _EVault.contract, event: "PullDebt", logs: logs, sub: sub}, nil
}

// WatchPullDebt is a free log subscription operation binding the contract event 0xe6d0bfd9025bf59969101a13cf02e3ba2811b533816c47d7155546c7c8a1048f.
//
// Solidity: event PullDebt(address indexed from, address indexed to, uint256 assets)
func (_EVault *EVaultFilterer) WatchPullDebt(opts *bind.WatchOpts, sink chan<- *EVaultPullDebt, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "PullDebt", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultPullDebt)
				if err := _EVault.contract.UnpackLog(event, "PullDebt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePullDebt is a log parse operation binding the contract event 0xe6d0bfd9025bf59969101a13cf02e3ba2811b533816c47d7155546c7c8a1048f.
//
// Solidity: event PullDebt(address indexed from, address indexed to, uint256 assets)
func (_EVault *EVaultFilterer) ParsePullDebt(log types.Log) (*EVaultPullDebt, error) {
	event := new(EVaultPullDebt)
	if err := _EVault.contract.UnpackLog(event, "PullDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the EVault contract.
type EVaultRepayIterator struct {
	Event *EVaultRepay // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultRepay)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultRepay)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultRepay represents a Repay event raised by the EVault contract.
type EVaultRepay struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a594231.
//
// Solidity: event Repay(address indexed account, uint256 assets)
func (_EVault *EVaultFilterer) FilterRepay(opts *bind.FilterOpts, account []common.Address) (*EVaultRepayIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "Repay", accountRule)
	if err != nil {
		return nil, err
	}
	return &EVaultRepayIterator{contract: _EVault.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a594231.
//
// Solidity: event Repay(address indexed account, uint256 assets)
func (_EVault *EVaultFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *EVaultRepay, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "Repay", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultRepay)
				if err := _EVault.contract.UnpackLog(event, "Repay", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRepay is a log parse operation binding the contract event 0x5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a594231.
//
// Solidity: event Repay(address indexed account, uint256 assets)
func (_EVault *EVaultFilterer) ParseRepay(log types.Log) (*EVaultRepay, error) {
	event := new(EVaultRepay)
	if err := _EVault.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EVault contract.
type EVaultTransferIterator struct {
	Event *EVaultTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultTransfer represents a Transfer event raised by the EVault contract.
type EVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EVault *EVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVaultTransferIterator{contract: _EVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EVault *EVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultTransfer)
				if err := _EVault.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EVault *EVaultFilterer) ParseTransfer(log types.Log) (*EVaultTransfer, error) {
	event := new(EVaultTransfer)
	if err := _EVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultVaultStatusIterator is returned from FilterVaultStatus and is used to iterate over the raw logs and unpacked data for VaultStatus events raised by the EVault contract.
type EVaultVaultStatusIterator struct {
	Event *EVaultVaultStatus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultVaultStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultVaultStatus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultVaultStatus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultVaultStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultVaultStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultVaultStatus represents a VaultStatus event raised by the EVault contract.
type EVaultVaultStatus struct {
	TotalShares         *big.Int
	TotalBorrows        *big.Int
	AccumulatedFees     *big.Int
	Cash                *big.Int
	InterestAccumulator *big.Int
	InterestRate        *big.Int
	Timestamp           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterVaultStatus is a free log retrieval operation binding the contract event 0x80b61abbfc5f73cfe5cf93cec97a69ed20643dc6c6f1833b05a1560aa164e24c.
//
// Solidity: event VaultStatus(uint256 totalShares, uint256 totalBorrows, uint256 accumulatedFees, uint256 cash, uint256 interestAccumulator, uint256 interestRate, uint256 timestamp)
func (_EVault *EVaultFilterer) FilterVaultStatus(opts *bind.FilterOpts) (*EVaultVaultStatusIterator, error) {

	logs, sub, err := _EVault.contract.FilterLogs(opts, "VaultStatus")
	if err != nil {
		return nil, err
	}
	return &EVaultVaultStatusIterator{contract: _EVault.contract, event: "VaultStatus", logs: logs, sub: sub}, nil
}

// WatchVaultStatus is a free log subscription operation binding the contract event 0x80b61abbfc5f73cfe5cf93cec97a69ed20643dc6c6f1833b05a1560aa164e24c.
//
// Solidity: event VaultStatus(uint256 totalShares, uint256 totalBorrows, uint256 accumulatedFees, uint256 cash, uint256 interestAccumulator, uint256 interestRate, uint256 timestamp)
func (_EVault *EVaultFilterer) WatchVaultStatus(opts *bind.WatchOpts, sink chan<- *EVaultVaultStatus) (event.Subscription, error) {

	logs, sub, err := _EVault.contract.WatchLogs(opts, "VaultStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultVaultStatus)
				if err := _EVault.contract.UnpackLog(event, "VaultStatus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVaultStatus is a log parse operation binding the contract event 0x80b61abbfc5f73cfe5cf93cec97a69ed20643dc6c6f1833b05a1560aa164e24c.
//
// Solidity: event VaultStatus(uint256 totalShares, uint256 totalBorrows, uint256 accumulatedFees, uint256 cash, uint256 interestAccumulator, uint256 interestRate, uint256 timestamp)
func (_EVault *EVaultFilterer) ParseVaultStatus(log types.Log) (*EVaultVaultStatus, error) {
	event := new(EVaultVaultStatus)
	if err := _EVault.contract.UnpackLog(event, "VaultStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the EVault contract.
type EVaultWithdrawIterator struct {
	Event *EVaultWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVaultWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EVaultWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVaultWithdraw represents a Withdraw event raised by the EVault contract.
type EVaultWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_EVault *EVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*EVaultWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EVault.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EVaultWithdrawIterator{contract: _EVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_EVault *EVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *EVaultWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EVault.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVaultWithdraw)
				if err := _EVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_EVault *EVaultFilterer) ParseWithdraw(log types.Log) (*EVaultWithdraw, error) {
	event := new(EVaultWithdraw)
	if err := _EVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
