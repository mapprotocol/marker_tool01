package handler

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"os"
)

func init() {
	startLogger()
}

func startLogger() {
	var lvl = log.LvlInfo
	logger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(true)))
	logger.Verbosity(lvl)
	log.Root().SetHandler(logger)
}

func getMgrMaintainerAddress(endpoint string) {
	cli := dial(endpoint)
	parsed := parseABI(EpochRewardsABI)
	input := packInput(parsed, "getMgrMaintainerAddress")
	output := CallContract(cli, GenesisAddresses["EpochRewardsProxy"], input)
	var addr common.Address
	if err := parsed.UnpackIntoInterface(&addr, "getMgrMaintainerAddress", output); err != nil {
		log.Crit("unpack failed", "err", err.Error())
	}
	log.Info("getMgrMaintainerAddress", "address", addr)
}

func setMgrMaintainerAddress(endpoint string, from, target common.Address, privateKey *ecdsa.PrivateKey) {
	cli := dial(endpoint)
	input := packInput(parseABI(EpochRewardsABI), "setMgrMaintainerAddress", target)
	txHash := sendContractTransaction(cli, from, GenesisAddresses["EpochRewardsProxy"], nil, privateKey, input, 0)
	getResult(cli, txHash)
	log.Info("setMgrMaintainerAddress", "address", target)
}

func getTargetEpochPayment(endpoint string) {
	cli := dial(endpoint)
	parsed := parseABI(EpochRewardsABI)
	input := packInput(parsed, "epochPayment")
	output := CallContract(cli, GenesisAddresses["EpochRewardsProxy"], input)
	var value *big.Int
	if err := parsed.UnpackIntoInterface(&value, "epochPayment", output); err != nil {
		log.Crit("unpack failed", "err", err.Error())
	}
	log.Info("getTargetEpochPayment", "value", value)
}

func setTargetEpochPayment(endpoint string, from common.Address, target *big.Int, privateKey *ecdsa.PrivateKey) {
	cli := dial(endpoint)
	input := packInput(parseABI(EpochRewardsABI), "setTargetEpochPayment", target)
	txHash := sendContractTransaction(cli, from, GenesisAddresses["EpochRewardsProxy"], nil, privateKey, input, 0)
	getResult(cli, txHash)
	log.Info("setTargetEpochPayment", "value", target)
}

func getElectableValidators(endpoint string) {
	cli := dial(endpoint)
	parsed := parseABI(ElectionABI)
	input := packInput(parsed, "electableValidators")
	output := CallContract(cli, GenesisAddresses["ElectionProxy"], input)

	var min *big.Int
	var max *big.Int
	resp := []*big.Int{min, max}
	if err := parsed.UnpackIntoInterface(&resp, "electableValidators", output); err != nil {
		log.Crit("unpack failed", "err", err.Error())
	}
	log.Info("getElectableValidators", "minElectableValidators", resp[0], "maxElectableValidators", resp[1])
}

func setElectableValidators(endpoint string, from common.Address, privateKey *ecdsa.PrivateKey, minElectableValidators, maxElectableValidators *big.Int) {
	cli := dial(endpoint)
	input := packInput(parseABI(ElectionABI), "setElectableValidators", minElectableValidators, maxElectableValidators)
	txHash := sendContractTransaction(cli, from, GenesisAddresses["ElectionProxy"], nil, privateKey, input, 0)
	getResult(cli, txHash)
	log.Info("setElectableValidators", "minElectableValidators", minElectableValidators, "maxElectableValidators", maxElectableValidators)
}

func getCommissionUpdateDelay(endpoint string) {
	cli := dial(endpoint)
	parsed := parseABI(ValidatorsABI)
	input := packInput(parsed, "commissionUpdateDelay")
	output := CallContract(cli, GenesisAddresses["ValidatorsProxy"], input)
	var value *big.Int
	if err := parsed.UnpackIntoInterface(&value, "commissionUpdateDelay", output); err != nil {
		log.Crit("unpack failed", "err", err.Error())
	}
	log.Info("getCommissionUpdateDelay", "delayBlock", value)
}

func setCommissionUpdateDelay(endpoint string, from common.Address, privateKey *ecdsa.PrivateKey, delayBlock *big.Int) {
	cli := dial(endpoint)
	input := packInput(parseABI(ValidatorsABI), "setCommissionUpdateDelay", delayBlock)
	txHash := sendContractTransaction(cli, from, GenesisAddresses["ValidatorsProxy"], nil, privateKey, input, 0)
	getResult(cli, txHash)
	log.Info("setCommissionUpdateDelay", "address", from, "delayBlock", delayBlock)
}

func getUnlockingPeriod(endpoint string) {
	cli := dial(endpoint)
	parsed := parseABI(LockedGoldABI)
	input := packInput(parsed, "unlockingPeriod")
	output := CallContract(cli, GenesisAddresses["LockedGoldProxy"], input)
	var period *big.Int
	if err := parsed.UnpackIntoInterface(&period, "unlockingPeriod", output); err != nil {
		log.Crit("unpack failed", "err", err.Error())
	}
	log.Info("getUnlockingPeriod", "period", period)
}

func setUnlockingPeriod(endpoint string, from common.Address, privateKey *ecdsa.PrivateKey, period *big.Int) {
	cli := dial(endpoint)
	input := packInput(parseABI(LockedGoldABI), "setUnlockingPeriod", period)
	txHash := sendContractTransaction(cli, from, GenesisAddresses["LockedGoldProxy"], nil, privateKey, input, 0)
	getResult(cli, txHash)
	log.Info("setCommissionUpdateDelay", "from", from, "period", period)
}

func getImplAddress(endpoint string, proxyAddress common.Address) {
	cli := dial(endpoint)
	parsed := parseABI(ProxyABI)
	input := packInput(parsed, "_getImplementation")
	output := CallContract(cli, proxyAddress, input)
	var implAddress common.Address
	if err := parsed.UnpackIntoInterface(&implAddress, "_getImplementation", output); err != nil {
		log.Crit("unpack failed", "err", err.Error())
	}
	log.Info("getImplAddress", "proxy", proxyAddress, "impl", implAddress)
}

func setImplAddress(endpoint string, from common.Address, privateKey *ecdsa.PrivateKey, proxyAddress, implAddress common.Address) {
	cli := dial(endpoint)
	input := packInput(parseABI(ProxyABI), "_setImplementation", implAddress)
	txHash := sendContractTransaction(cli, from, proxyAddress, nil, privateKey, input, 0)
	getResult(cli, txHash)
	log.Info("setImplAddress", "from", from, "proxy", proxyAddress, "impl", implAddress)
}
