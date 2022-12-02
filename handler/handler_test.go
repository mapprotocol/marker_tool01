package handler

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

var endpoint = "https://rpc.maplabs.io"

func Test_getMgrMaintainerAddress(t *testing.T) {
	getMgrMaintainerAddress(endpoint)
}

func Test_setMgrMaintainerAddress(t *testing.T) {
	from := common.HexToAddress("")
	target := common.HexToAddress("")
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}

	setMgrMaintainerAddress(endpoint, from, target, privateKey)
}

func Test_getTargetEpochPayment(t *testing.T) {
	getTargetEpochPayment(endpoint)
}

func Test_setTargetEpochPayment(t *testing.T) {
	from := common.HexToAddress("")
	target := new(big.Int).Mul(big.NewInt(50000), big.NewInt(1e18))
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}

	setTargetEpochPayment(endpoint, from, target, privateKey)
}

func Test_getElectableValidators(t *testing.T) {
	getElectableValidators(endpoint)
}

// INFO [08-26|16:55:35.641] getElectableValidators                   minElectableValidators=1 maxElectableValidators=100
// INFO [08-26|17:00:42.247] getElectableValidators                   minElectableValidators=1 maxElectableValidators=50

func Test_setElectableValidators(t *testing.T) {
	from := common.HexToAddress("")
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}
	min := big.NewInt(1)
	max := big.NewInt(100)
	setElectableValidators(endpoint, from, privateKey, min, max)
}

func Test_getCommissionUpdateDelay(t *testing.T) {
	getCommissionUpdateDelay(endpoint)
}

func Test_setCommissionUpdateDelay(t *testing.T) {
	from := common.HexToAddress("")
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}
	delayBlock := big.NewInt(10)
	setCommissionUpdateDelay(endpoint, from, privateKey, delayBlock)
}

func Test_getUnlockingPeriod(t *testing.T) {
	getUnlockingPeriod(endpoint)
}

func Test_setUnlockingPeriod(t *testing.T) {
	from := common.HexToAddress("")
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}
	period := big.NewInt(900)
	setUnlockingPeriod(endpoint, from, privateKey, period)
}

func Test_getImplAddress(t *testing.T) {
	proxyAddress := common.HexToAddress("")
	getImplAddress(endpoint, proxyAddress)
}

func Test_setImplAddress(t *testing.T) {
	proxyAddress := common.HexToAddress("")
	implAddress := common.HexToAddress("")
	from := common.HexToAddress("")
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}
	setImplAddress(endpoint, from, privateKey, proxyAddress, implAddress)
}

func Test_getParticipationParameters(t *testing.T) {
	governanceAddr := common.HexToAddress("0xcdB66B1e6A07279df98f804d0aCAC86695F4b99e")
	getParticipationParameters(endpoint, governanceAddr)
}

func Test_setBaselineQuorumFactor(t *testing.T) {
	governanceAddr := common.HexToAddress("0xcdB66B1e6A07279df98f804d0aCAC86695F4b99e")
	baselineQuorumFactor, b := new(big.Int).SetString("100000000000000000", 10)
	if !b {
		panic("convert failed")
	}
	getParticipationParameters(endpoint, governanceAddr)
	from := common.HexToAddress("")
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}
	setBaselineQuorumFactor(endpoint, from, privateKey, governanceAddr, baselineQuorumFactor)
}

func Test_getPledgeMultiplierInReward(t *testing.T) {
	getPledgeMultiplierInReward(endpoint)
}

func Test_setPledgeMultiplierInReward(t *testing.T) {
	pledgeMultiplier, b := new(big.Int).SetString("700000000000000000", 10)
	if !b {
		panic("convert failed")
	}
	from := common.HexToAddress("")
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}
	setPledgeMultiplierInReward(endpoint, from, privateKey, pledgeMultiplier)
}
