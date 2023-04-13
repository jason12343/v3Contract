// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v3PositionM

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
)

// INonfungiblePositionManagerCollectParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerCollectParams struct {
	TokenId    *big.Int
	Recipient  common.Address
	Amount0Max *big.Int
	Amount1Max *big.Int
}

// INonfungiblePositionManagerDecreaseLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerDecreaseLiquidityParams struct {
	TokenId    *big.Int
	Liquidity  *big.Int
	Amount0Min *big.Int
	Amount1Min *big.Int
	Deadline   *big.Int
}

// INonfungiblePositionManagerIncreaseLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerIncreaseLiquidityParams struct {
	TokenId        *big.Int
	Amount0Desired *big.Int
	Amount1Desired *big.Int
	Amount0Min     *big.Int
	Amount1Min     *big.Int
	Deadline       *big.Int
}

// INonfungiblePositionManagerMintParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerMintParams struct {
	Token0         common.Address
	Token1         common.Address
	Fee            *big.Int
	TickLower      *big.Int
	TickUpper      *big.Int
	Amount0Desired *big.Int
	Amount1Desired *big.Int
	Amount0Min     *big.Int
	Amount1Min     *big.Int
	Recipient      common.Address
	Deadline       *big.Int
}

// V3PositionMMetaData contains all meta data concerning the V3PositionM contract.
var V3PositionMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenDescriptor_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"DecreaseLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"IncreaseLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount0Max\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Max\",\"type\":\"uint128\"}],\"internalType\":\"structINonfungiblePositionManager.CollectParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"createAndInitializePoolIfNecessary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structINonfungiblePositionManager.DecreaseLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"decreaseLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structINonfungiblePositionManager.IncreaseLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"increaseLiquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"amount0Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structINonfungiblePositionManager.MintParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Owed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Owed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// V3PositionMABI is the input ABI used to generate the binding from.
// Deprecated: Use V3PositionMMetaData.ABI instead.
var V3PositionMABI = V3PositionMMetaData.ABI

// V3PositionM is an auto generated Go binding around an Ethereum contract.
type V3PositionM struct {
	V3PositionMCaller     // Read-only binding to the contract
	V3PositionMTransactor // Write-only binding to the contract
	V3PositionMFilterer   // Log filterer for contract events
}

// V3PositionMCaller is an auto generated read-only Go binding around an Ethereum contract.
type V3PositionMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3PositionMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V3PositionMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3PositionMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V3PositionMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3PositionMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V3PositionMSession struct {
	Contract     *V3PositionM      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V3PositionMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V3PositionMCallerSession struct {
	Contract *V3PositionMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// V3PositionMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V3PositionMTransactorSession struct {
	Contract     *V3PositionMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// V3PositionMRaw is an auto generated low-level Go binding around an Ethereum contract.
type V3PositionMRaw struct {
	Contract *V3PositionM // Generic contract binding to access the raw methods on
}

// V3PositionMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V3PositionMCallerRaw struct {
	Contract *V3PositionMCaller // Generic read-only contract binding to access the raw methods on
}

// V3PositionMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V3PositionMTransactorRaw struct {
	Contract *V3PositionMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV3PositionM creates a new instance of V3PositionM, bound to a specific deployed contract.
func NewV3PositionM(address common.Address, backend bind.ContractBackend) (*V3PositionM, error) {
	contract, err := bindV3PositionM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V3PositionM{V3PositionMCaller: V3PositionMCaller{contract: contract}, V3PositionMTransactor: V3PositionMTransactor{contract: contract}, V3PositionMFilterer: V3PositionMFilterer{contract: contract}}, nil
}

// NewV3PositionMCaller creates a new read-only instance of V3PositionM, bound to a specific deployed contract.
func NewV3PositionMCaller(address common.Address, caller bind.ContractCaller) (*V3PositionMCaller, error) {
	contract, err := bindV3PositionM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V3PositionMCaller{contract: contract}, nil
}

// NewV3PositionMTransactor creates a new write-only instance of V3PositionM, bound to a specific deployed contract.
func NewV3PositionMTransactor(address common.Address, transactor bind.ContractTransactor) (*V3PositionMTransactor, error) {
	contract, err := bindV3PositionM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V3PositionMTransactor{contract: contract}, nil
}

// NewV3PositionMFilterer creates a new log filterer instance of V3PositionM, bound to a specific deployed contract.
func NewV3PositionMFilterer(address common.Address, filterer bind.ContractFilterer) (*V3PositionMFilterer, error) {
	contract, err := bindV3PositionM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V3PositionMFilterer{contract: contract}, nil
}

// bindV3PositionM binds a generic wrapper to an already deployed contract.
func bindV3PositionM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(V3PositionMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3PositionM *V3PositionMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3PositionM.Contract.V3PositionMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3PositionM *V3PositionMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3PositionM.Contract.V3PositionMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3PositionM *V3PositionMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3PositionM.Contract.V3PositionMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3PositionM *V3PositionMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3PositionM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3PositionM *V3PositionMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3PositionM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3PositionM *V3PositionMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3PositionM.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_V3PositionM *V3PositionMCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_V3PositionM *V3PositionMSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _V3PositionM.Contract.DOMAINSEPARATOR(&_V3PositionM.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_V3PositionM *V3PositionMCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _V3PositionM.Contract.DOMAINSEPARATOR(&_V3PositionM.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_V3PositionM *V3PositionMCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_V3PositionM *V3PositionMSession) PERMITTYPEHASH() ([32]byte, error) {
	return _V3PositionM.Contract.PERMITTYPEHASH(&_V3PositionM.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_V3PositionM *V3PositionMCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _V3PositionM.Contract.PERMITTYPEHASH(&_V3PositionM.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_V3PositionM *V3PositionMCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_V3PositionM *V3PositionMSession) WETH9() (common.Address, error) {
	return _V3PositionM.Contract.WETH9(&_V3PositionM.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_V3PositionM *V3PositionMCallerSession) WETH9() (common.Address, error) {
	return _V3PositionM.Contract.WETH9(&_V3PositionM.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_V3PositionM *V3PositionMCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_V3PositionM *V3PositionMSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _V3PositionM.Contract.BalanceOf(&_V3PositionM.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_V3PositionM *V3PositionMCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _V3PositionM.Contract.BalanceOf(&_V3PositionM.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() pure returns(string)
func (_V3PositionM *V3PositionMCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() pure returns(string)
func (_V3PositionM *V3PositionMSession) BaseURI() (string, error) {
	return _V3PositionM.Contract.BaseURI(&_V3PositionM.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() pure returns(string)
func (_V3PositionM *V3PositionMCallerSession) BaseURI() (string, error) {
	return _V3PositionM.Contract.BaseURI(&_V3PositionM.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_V3PositionM *V3PositionMCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_V3PositionM *V3PositionMSession) Factory() (common.Address, error) {
	return _V3PositionM.Contract.Factory(&_V3PositionM.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_V3PositionM *V3PositionMCallerSession) Factory() (common.Address, error) {
	return _V3PositionM.Contract.Factory(&_V3PositionM.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_V3PositionM *V3PositionMCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_V3PositionM *V3PositionMSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _V3PositionM.Contract.GetApproved(&_V3PositionM.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_V3PositionM *V3PositionMCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _V3PositionM.Contract.GetApproved(&_V3PositionM.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_V3PositionM *V3PositionMCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_V3PositionM *V3PositionMSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _V3PositionM.Contract.IsApprovedForAll(&_V3PositionM.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_V3PositionM *V3PositionMCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _V3PositionM.Contract.IsApprovedForAll(&_V3PositionM.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_V3PositionM *V3PositionMCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_V3PositionM *V3PositionMSession) Name() (string, error) {
	return _V3PositionM.Contract.Name(&_V3PositionM.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_V3PositionM *V3PositionMCallerSession) Name() (string, error) {
	return _V3PositionM.Contract.Name(&_V3PositionM.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_V3PositionM *V3PositionMCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_V3PositionM *V3PositionMSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _V3PositionM.Contract.OwnerOf(&_V3PositionM.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_V3PositionM *V3PositionMCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _V3PositionM.Contract.OwnerOf(&_V3PositionM.CallOpts, tokenId)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 tokenId) view returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_V3PositionM *V3PositionMCaller) Positions(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "positions", tokenId)

	outstruct := new(struct {
		Nonce                    *big.Int
		Operator                 common.Address
		Token0                   common.Address
		Token1                   common.Address
		Fee                      *big.Int
		TickLower                *big.Int
		TickUpper                *big.Int
		Liquidity                *big.Int
		FeeGrowthInside0LastX128 *big.Int
		FeeGrowthInside1LastX128 *big.Int
		TokensOwed0              *big.Int
		TokensOwed1              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Operator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token0 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TickLower = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TickUpper = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside0LastX128 = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside1LastX128 = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed0 = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed1 = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 tokenId) view returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_V3PositionM *V3PositionMSession) Positions(tokenId *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _V3PositionM.Contract.Positions(&_V3PositionM.CallOpts, tokenId)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 tokenId) view returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_V3PositionM *V3PositionMCallerSession) Positions(tokenId *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _V3PositionM.Contract.Positions(&_V3PositionM.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_V3PositionM *V3PositionMCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_V3PositionM *V3PositionMSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _V3PositionM.Contract.SupportsInterface(&_V3PositionM.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_V3PositionM *V3PositionMCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _V3PositionM.Contract.SupportsInterface(&_V3PositionM.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_V3PositionM *V3PositionMCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_V3PositionM *V3PositionMSession) Symbol() (string, error) {
	return _V3PositionM.Contract.Symbol(&_V3PositionM.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_V3PositionM *V3PositionMCallerSession) Symbol() (string, error) {
	return _V3PositionM.Contract.Symbol(&_V3PositionM.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_V3PositionM *V3PositionMCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_V3PositionM *V3PositionMSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _V3PositionM.Contract.TokenByIndex(&_V3PositionM.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_V3PositionM *V3PositionMCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _V3PositionM.Contract.TokenByIndex(&_V3PositionM.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_V3PositionM *V3PositionMCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_V3PositionM *V3PositionMSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _V3PositionM.Contract.TokenOfOwnerByIndex(&_V3PositionM.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_V3PositionM *V3PositionMCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _V3PositionM.Contract.TokenOfOwnerByIndex(&_V3PositionM.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_V3PositionM *V3PositionMCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_V3PositionM *V3PositionMSession) TokenURI(tokenId *big.Int) (string, error) {
	return _V3PositionM.Contract.TokenURI(&_V3PositionM.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_V3PositionM *V3PositionMCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _V3PositionM.Contract.TokenURI(&_V3PositionM.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_V3PositionM *V3PositionMCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V3PositionM.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_V3PositionM *V3PositionMSession) TotalSupply() (*big.Int, error) {
	return _V3PositionM.Contract.TotalSupply(&_V3PositionM.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_V3PositionM *V3PositionMCallerSession) TotalSupply() (*big.Int, error) {
	return _V3PositionM.Contract.TotalSupply(&_V3PositionM.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_V3PositionM *V3PositionMTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_V3PositionM *V3PositionMSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V3PositionM.Contract.Approve(&_V3PositionM.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_V3PositionM *V3PositionMTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V3PositionM.Contract.Approve(&_V3PositionM.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) payable returns()
func (_V3PositionM *V3PositionMTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) payable returns()
func (_V3PositionM *V3PositionMSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _V3PositionM.Contract.Burn(&_V3PositionM.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) payable returns()
func (_V3PositionM *V3PositionMTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _V3PositionM.Contract.Burn(&_V3PositionM.TransactOpts, tokenId)
}

// Collect is a paid mutator transaction binding the contract method 0xfc6f7865.
//
// Solidity: function collect((uint256,address,uint128,uint128) params) payable returns(uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMTransactor) Collect(opts *bind.TransactOpts, params INonfungiblePositionManagerCollectParams) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "collect", params)
}

// Collect is a paid mutator transaction binding the contract method 0xfc6f7865.
//
// Solidity: function collect((uint256,address,uint128,uint128) params) payable returns(uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMSession) Collect(params INonfungiblePositionManagerCollectParams) (*types.Transaction, error) {
	return _V3PositionM.Contract.Collect(&_V3PositionM.TransactOpts, params)
}

// Collect is a paid mutator transaction binding the contract method 0xfc6f7865.
//
// Solidity: function collect((uint256,address,uint128,uint128) params) payable returns(uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMTransactorSession) Collect(params INonfungiblePositionManagerCollectParams) (*types.Transaction, error) {
	return _V3PositionM.Contract.Collect(&_V3PositionM.TransactOpts, params)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_V3PositionM *V3PositionMTransactor) CreateAndInitializePoolIfNecessary(opts *bind.TransactOpts, token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "createAndInitializePoolIfNecessary", token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_V3PositionM *V3PositionMSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _V3PositionM.Contract.CreateAndInitializePoolIfNecessary(&_V3PositionM.TransactOpts, token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_V3PositionM *V3PositionMTransactorSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _V3PositionM.Contract.CreateAndInitializePoolIfNecessary(&_V3PositionM.TransactOpts, token0, token1, fee, sqrtPriceX96)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x0c49ccbe.
//
// Solidity: function decreaseLiquidity((uint256,uint128,uint256,uint256,uint256) params) payable returns(uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMTransactor) DecreaseLiquidity(opts *bind.TransactOpts, params INonfungiblePositionManagerDecreaseLiquidityParams) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "decreaseLiquidity", params)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x0c49ccbe.
//
// Solidity: function decreaseLiquidity((uint256,uint128,uint256,uint256,uint256) params) payable returns(uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMSession) DecreaseLiquidity(params INonfungiblePositionManagerDecreaseLiquidityParams) (*types.Transaction, error) {
	return _V3PositionM.Contract.DecreaseLiquidity(&_V3PositionM.TransactOpts, params)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x0c49ccbe.
//
// Solidity: function decreaseLiquidity((uint256,uint128,uint256,uint256,uint256) params) payable returns(uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMTransactorSession) DecreaseLiquidity(params INonfungiblePositionManagerDecreaseLiquidityParams) (*types.Transaction, error) {
	return _V3PositionM.Contract.DecreaseLiquidity(&_V3PositionM.TransactOpts, params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0x219f5d17.
//
// Solidity: function increaseLiquidity((uint256,uint256,uint256,uint256,uint256,uint256) params) payable returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMTransactor) IncreaseLiquidity(opts *bind.TransactOpts, params INonfungiblePositionManagerIncreaseLiquidityParams) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "increaseLiquidity", params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0x219f5d17.
//
// Solidity: function increaseLiquidity((uint256,uint256,uint256,uint256,uint256,uint256) params) payable returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMSession) IncreaseLiquidity(params INonfungiblePositionManagerIncreaseLiquidityParams) (*types.Transaction, error) {
	return _V3PositionM.Contract.IncreaseLiquidity(&_V3PositionM.TransactOpts, params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0x219f5d17.
//
// Solidity: function increaseLiquidity((uint256,uint256,uint256,uint256,uint256,uint256) params) payable returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMTransactorSession) IncreaseLiquidity(params INonfungiblePositionManagerIncreaseLiquidityParams) (*types.Transaction, error) {
	return _V3PositionM.Contract.IncreaseLiquidity(&_V3PositionM.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x88316456.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256) params) payable returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMTransactor) Mint(opts *bind.TransactOpts, params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "mint", params)
}

// Mint is a paid mutator transaction binding the contract method 0x88316456.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256) params) payable returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMSession) Mint(params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _V3PositionM.Contract.Mint(&_V3PositionM.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x88316456.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256) params) payable returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMTransactorSession) Mint(params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _V3PositionM.Contract.Mint(&_V3PositionM.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_V3PositionM *V3PositionMTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_V3PositionM *V3PositionMSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _V3PositionM.Contract.Multicall(&_V3PositionM.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_V3PositionM *V3PositionMTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _V3PositionM.Contract.Multicall(&_V3PositionM.TransactOpts, data)
}

// Permit is a paid mutator transaction binding the contract method 0x7ac2ff7b.
//
// Solidity: function permit(address spender, uint256 tokenId, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3PositionM *V3PositionMTransactor) Permit(opts *bind.TransactOpts, spender common.Address, tokenId *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "permit", spender, tokenId, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0x7ac2ff7b.
//
// Solidity: function permit(address spender, uint256 tokenId, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3PositionM *V3PositionMSession) Permit(spender common.Address, tokenId *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3PositionM.Contract.Permit(&_V3PositionM.TransactOpts, spender, tokenId, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0x7ac2ff7b.
//
// Solidity: function permit(address spender, uint256 tokenId, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3PositionM *V3PositionMTransactorSession) Permit(spender common.Address, tokenId *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3PositionM.Contract.Permit(&_V3PositionM.TransactOpts, spender, tokenId, deadline, v, r, s)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_V3PositionM *V3PositionMTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_V3PositionM *V3PositionMSession) RefundETH() (*types.Transaction, error) {
	return _V3PositionM.Contract.RefundETH(&_V3PositionM.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_V3PositionM *V3PositionMTransactorSession) RefundETH() (*types.Transaction, error) {
	return _V3PositionM.Contract.RefundETH(&_V3PositionM.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_V3PositionM *V3PositionMTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_V3PositionM *V3PositionMSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V3PositionM.Contract.SafeTransferFrom(&_V3PositionM.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_V3PositionM *V3PositionMTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V3PositionM.Contract.SafeTransferFrom(&_V3PositionM.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_V3PositionM *V3PositionMTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_V3PositionM *V3PositionMSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _V3PositionM.Contract.SafeTransferFrom0(&_V3PositionM.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_V3PositionM *V3PositionMTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _V3PositionM.Contract.SafeTransferFrom0(&_V3PositionM.TransactOpts, from, to, tokenId, _data)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3PositionM *V3PositionMTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3PositionM *V3PositionMSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3PositionM.Contract.SelfPermit(&_V3PositionM.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3PositionM *V3PositionMTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3PositionM.Contract.SelfPermit(&_V3PositionM.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3PositionM *V3PositionMTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3PositionM *V3PositionMSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3PositionM.Contract.SelfPermitAllowed(&_V3PositionM.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3PositionM *V3PositionMTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3PositionM.Contract.SelfPermitAllowed(&_V3PositionM.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3PositionM *V3PositionMTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3PositionM *V3PositionMSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3PositionM.Contract.SelfPermitAllowedIfNecessary(&_V3PositionM.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3PositionM *V3PositionMTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3PositionM.Contract.SelfPermitAllowedIfNecessary(&_V3PositionM.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3PositionM *V3PositionMTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3PositionM *V3PositionMSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3PositionM.Contract.SelfPermitIfNecessary(&_V3PositionM.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3PositionM *V3PositionMTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3PositionM.Contract.SelfPermitIfNecessary(&_V3PositionM.TransactOpts, token, value, deadline, v, r, s)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_V3PositionM *V3PositionMTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_V3PositionM *V3PositionMSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _V3PositionM.Contract.SetApprovalForAll(&_V3PositionM.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_V3PositionM *V3PositionMTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _V3PositionM.Contract.SetApprovalForAll(&_V3PositionM.TransactOpts, operator, approved)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_V3PositionM *V3PositionMTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_V3PositionM *V3PositionMSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3PositionM.Contract.SweepToken(&_V3PositionM.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_V3PositionM *V3PositionMTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3PositionM.Contract.SweepToken(&_V3PositionM.TransactOpts, token, amountMinimum, recipient)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_V3PositionM *V3PositionMTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_V3PositionM *V3PositionMSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V3PositionM.Contract.TransferFrom(&_V3PositionM.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_V3PositionM *V3PositionMTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V3PositionM.Contract.TransferFrom(&_V3PositionM.TransactOpts, from, to, tokenId)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_V3PositionM *V3PositionMTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "uniswapV3MintCallback", amount0Owed, amount1Owed, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_V3PositionM *V3PositionMSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _V3PositionM.Contract.UniswapV3MintCallback(&_V3PositionM.TransactOpts, amount0Owed, amount1Owed, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_V3PositionM *V3PositionMTransactorSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _V3PositionM.Contract.UniswapV3MintCallback(&_V3PositionM.TransactOpts, amount0Owed, amount1Owed, data)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_V3PositionM *V3PositionMTransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3PositionM.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_V3PositionM *V3PositionMSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3PositionM.Contract.UnwrapWETH9(&_V3PositionM.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_V3PositionM *V3PositionMTransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3PositionM.Contract.UnwrapWETH9(&_V3PositionM.TransactOpts, amountMinimum, recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V3PositionM *V3PositionMTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3PositionM.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V3PositionM *V3PositionMSession) Receive() (*types.Transaction, error) {
	return _V3PositionM.Contract.Receive(&_V3PositionM.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V3PositionM *V3PositionMTransactorSession) Receive() (*types.Transaction, error) {
	return _V3PositionM.Contract.Receive(&_V3PositionM.TransactOpts)
}

// V3PositionMApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the V3PositionM contract.
type V3PositionMApprovalIterator struct {
	Event *V3PositionMApproval // Event containing the contract specifics and raw log

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
func (it *V3PositionMApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3PositionMApproval)
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
		it.Event = new(V3PositionMApproval)
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
func (it *V3PositionMApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3PositionMApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3PositionMApproval represents a Approval event raised by the V3PositionM contract.
type V3PositionMApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_V3PositionM *V3PositionMFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*V3PositionMApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _V3PositionM.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &V3PositionMApprovalIterator{contract: _V3PositionM.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_V3PositionM *V3PositionMFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *V3PositionMApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _V3PositionM.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3PositionMApproval)
				if err := _V3PositionM.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_V3PositionM *V3PositionMFilterer) ParseApproval(log types.Log) (*V3PositionMApproval, error) {
	event := new(V3PositionMApproval)
	if err := _V3PositionM.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3PositionMApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the V3PositionM contract.
type V3PositionMApprovalForAllIterator struct {
	Event *V3PositionMApprovalForAll // Event containing the contract specifics and raw log

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
func (it *V3PositionMApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3PositionMApprovalForAll)
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
		it.Event = new(V3PositionMApprovalForAll)
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
func (it *V3PositionMApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3PositionMApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3PositionMApprovalForAll represents a ApprovalForAll event raised by the V3PositionM contract.
type V3PositionMApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_V3PositionM *V3PositionMFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*V3PositionMApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _V3PositionM.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &V3PositionMApprovalForAllIterator{contract: _V3PositionM.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_V3PositionM *V3PositionMFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *V3PositionMApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _V3PositionM.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3PositionMApprovalForAll)
				if err := _V3PositionM.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_V3PositionM *V3PositionMFilterer) ParseApprovalForAll(log types.Log) (*V3PositionMApprovalForAll, error) {
	event := new(V3PositionMApprovalForAll)
	if err := _V3PositionM.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3PositionMCollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the V3PositionM contract.
type V3PositionMCollectIterator struct {
	Event *V3PositionMCollect // Event containing the contract specifics and raw log

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
func (it *V3PositionMCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3PositionMCollect)
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
		it.Event = new(V3PositionMCollect)
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
func (it *V3PositionMCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3PositionMCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3PositionMCollect represents a Collect event raised by the V3PositionM contract.
type V3PositionMCollect struct {
	TokenId   *big.Int
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollect is a free log retrieval operation binding the contract event 0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01.
//
// Solidity: event Collect(uint256 indexed tokenId, address recipient, uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMFilterer) FilterCollect(opts *bind.FilterOpts, tokenId []*big.Int) (*V3PositionMCollectIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _V3PositionM.contract.FilterLogs(opts, "Collect", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &V3PositionMCollectIterator{contract: _V3PositionM.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01.
//
// Solidity: event Collect(uint256 indexed tokenId, address recipient, uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMFilterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *V3PositionMCollect, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _V3PositionM.contract.WatchLogs(opts, "Collect", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3PositionMCollect)
				if err := _V3PositionM.contract.UnpackLog(event, "Collect", log); err != nil {
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

// ParseCollect is a log parse operation binding the contract event 0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01.
//
// Solidity: event Collect(uint256 indexed tokenId, address recipient, uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMFilterer) ParseCollect(log types.Log) (*V3PositionMCollect, error) {
	event := new(V3PositionMCollect)
	if err := _V3PositionM.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3PositionMDecreaseLiquidityIterator is returned from FilterDecreaseLiquidity and is used to iterate over the raw logs and unpacked data for DecreaseLiquidity events raised by the V3PositionM contract.
type V3PositionMDecreaseLiquidityIterator struct {
	Event *V3PositionMDecreaseLiquidity // Event containing the contract specifics and raw log

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
func (it *V3PositionMDecreaseLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3PositionMDecreaseLiquidity)
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
		it.Event = new(V3PositionMDecreaseLiquidity)
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
func (it *V3PositionMDecreaseLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3PositionMDecreaseLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3PositionMDecreaseLiquidity represents a DecreaseLiquidity event raised by the V3PositionM contract.
type V3PositionMDecreaseLiquidity struct {
	TokenId   *big.Int
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDecreaseLiquidity is a free log retrieval operation binding the contract event 0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4.
//
// Solidity: event DecreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMFilterer) FilterDecreaseLiquidity(opts *bind.FilterOpts, tokenId []*big.Int) (*V3PositionMDecreaseLiquidityIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _V3PositionM.contract.FilterLogs(opts, "DecreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &V3PositionMDecreaseLiquidityIterator{contract: _V3PositionM.contract, event: "DecreaseLiquidity", logs: logs, sub: sub}, nil
}

// WatchDecreaseLiquidity is a free log subscription operation binding the contract event 0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4.
//
// Solidity: event DecreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMFilterer) WatchDecreaseLiquidity(opts *bind.WatchOpts, sink chan<- *V3PositionMDecreaseLiquidity, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _V3PositionM.contract.WatchLogs(opts, "DecreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3PositionMDecreaseLiquidity)
				if err := _V3PositionM.contract.UnpackLog(event, "DecreaseLiquidity", log); err != nil {
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

// ParseDecreaseLiquidity is a log parse operation binding the contract event 0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4.
//
// Solidity: event DecreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMFilterer) ParseDecreaseLiquidity(log types.Log) (*V3PositionMDecreaseLiquidity, error) {
	event := new(V3PositionMDecreaseLiquidity)
	if err := _V3PositionM.contract.UnpackLog(event, "DecreaseLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3PositionMIncreaseLiquidityIterator is returned from FilterIncreaseLiquidity and is used to iterate over the raw logs and unpacked data for IncreaseLiquidity events raised by the V3PositionM contract.
type V3PositionMIncreaseLiquidityIterator struct {
	Event *V3PositionMIncreaseLiquidity // Event containing the contract specifics and raw log

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
func (it *V3PositionMIncreaseLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3PositionMIncreaseLiquidity)
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
		it.Event = new(V3PositionMIncreaseLiquidity)
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
func (it *V3PositionMIncreaseLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3PositionMIncreaseLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3PositionMIncreaseLiquidity represents a IncreaseLiquidity event raised by the V3PositionM contract.
type V3PositionMIncreaseLiquidity struct {
	TokenId   *big.Int
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncreaseLiquidity is a free log retrieval operation binding the contract event 0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f.
//
// Solidity: event IncreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMFilterer) FilterIncreaseLiquidity(opts *bind.FilterOpts, tokenId []*big.Int) (*V3PositionMIncreaseLiquidityIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _V3PositionM.contract.FilterLogs(opts, "IncreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &V3PositionMIncreaseLiquidityIterator{contract: _V3PositionM.contract, event: "IncreaseLiquidity", logs: logs, sub: sub}, nil
}

// WatchIncreaseLiquidity is a free log subscription operation binding the contract event 0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f.
//
// Solidity: event IncreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMFilterer) WatchIncreaseLiquidity(opts *bind.WatchOpts, sink chan<- *V3PositionMIncreaseLiquidity, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _V3PositionM.contract.WatchLogs(opts, "IncreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3PositionMIncreaseLiquidity)
				if err := _V3PositionM.contract.UnpackLog(event, "IncreaseLiquidity", log); err != nil {
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

// ParseIncreaseLiquidity is a log parse operation binding the contract event 0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f.
//
// Solidity: event IncreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_V3PositionM *V3PositionMFilterer) ParseIncreaseLiquidity(log types.Log) (*V3PositionMIncreaseLiquidity, error) {
	event := new(V3PositionMIncreaseLiquidity)
	if err := _V3PositionM.contract.UnpackLog(event, "IncreaseLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3PositionMTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the V3PositionM contract.
type V3PositionMTransferIterator struct {
	Event *V3PositionMTransfer // Event containing the contract specifics and raw log

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
func (it *V3PositionMTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3PositionMTransfer)
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
		it.Event = new(V3PositionMTransfer)
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
func (it *V3PositionMTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3PositionMTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3PositionMTransfer represents a Transfer event raised by the V3PositionM contract.
type V3PositionMTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_V3PositionM *V3PositionMFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*V3PositionMTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _V3PositionM.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &V3PositionMTransferIterator{contract: _V3PositionM.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_V3PositionM *V3PositionMFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *V3PositionMTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _V3PositionM.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3PositionMTransfer)
				if err := _V3PositionM.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_V3PositionM *V3PositionMFilterer) ParseTransfer(log types.Log) (*V3PositionMTransfer, error) {
	event := new(V3PositionMTransfer)
	if err := _V3PositionM.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
