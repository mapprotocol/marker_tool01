package handler

import (
	"context"
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
func isPendingDeRegisterValidator(endpoint string, sender common.Address) {
	cli := dial(endpoint)
	parsed := parseABI(ValidatorsABI)
	input := packInput(parsed, "isPendingDeRegisterValidator")
	output := CallContract2(cli, sender, GenesisAddresses["ValidatorsProxy"], input)
	var result bool
	if err := parsed.UnpackIntoInterface(&result, "isPendingDeRegisterValidator", output); err != nil {
		log.Crit("unpack failed", "err", err.Error())
	}
	log.Info("getUnlockingPeriod", "result", result)
}
func getActiveVotesForValidator(endpoint string, addr common.Address, height *big.Int) {
	cli := dial(endpoint)
	parsed := parseABI(ElectionABI)
	input := packInput(parsed, "getActiveVotesForValidator", addr)
	output := CallContract3(cli, GenesisAddresses["ElectionProxy"], input, height)
	var res *big.Int
	if err := parsed.UnpackIntoInterface(&res, "getActiveVotesForValidator", output); err != nil {
		log.Crit("unpack failed", "err", err.Error())
	}
	log.Info("getActiveVotesForValidator", "validator", addr, "height", height.String(), "val", toCoin(res))
}
func sendTransaction(endpoint string, from, to common.Address, privateKey *ecdsa.PrivateKey, value *big.Int) {
	cli := dial(endpoint)
	txHash := sendTransaction0(cli, from, to, value, privateKey)
	getResult(cli, txHash)
}
func balanceOf(endpoint string, to common.Address) {
	cli := dial(endpoint)
	b, e := cli.BalanceAt(context.Background(), to, nil)
	log.Info("balanceOf", "to", to, "balance", b.String(), "coin", toCoin(b), "error", e)
}
func toCoin(val *big.Int) *big.Float {
	BaseBig := big.NewInt(1e18)
	return new(big.Float).Quo(new(big.Float).SetInt(val), new(big.Float).SetInt(BaseBig))
}
func toWei(value *big.Float) *big.Int {
	BaseBig := big.NewInt(1e18)
	base := new(big.Float).SetInt(BaseBig)
	val, _ := new(big.Float).Mul(value, base).Int(big.NewInt(0))
	return val
}
func getAccountTotalLockedGold(endpoint string, addr common.Address, height *big.Int) {
	cli := dial(endpoint)
	parsed := parseABI(LockedGoldABI)
	input := packInput(parsed, "getAccountTotalLockedGold", addr)
	output := CallContract3(cli, GenesisAddresses["LockedGoldProxy"], input, height)
	var res *big.Int
	if err := parsed.UnpackIntoInterface(&res, "getAccountTotalLockedGold", output); err != nil {
		log.Crit("unpack failed", "err", err.Error())
	}
	log.Info("getAccountTotalLockedGold", "addr", addr, "height", height.String(), "val", toCoin(res))
}

func getAccountNonvotingLockedGold(endpoint string, addr common.Address, height *big.Int) {
	cli := dial(endpoint)
	parsed := parseABI(LockedGoldABI)
	input := packInput(parsed, "getAccountNonvotingLockedGold", addr)
	output := CallContract3(cli, GenesisAddresses["LockedGoldProxy"], input, height)
	var res *big.Int
	if err := parsed.UnpackIntoInterface(&res, "getAccountNonvotingLockedGold", output); err != nil {
		log.Crit("unpack failed", "err", err.Error())
	}
	log.Info("getAccountNonvotingLockedGold", "addr", addr, "height", height.String(), "val", toCoin(res))
}

func getBlockGasLimit(endpoint string) {
	cli := dial(endpoint)
	parsed := parseABI(BlockchainParametersABI)
	input := packInput(parsed, "blockGasLimit")
	output := CallContract(cli, GenesisAddresses["BlockchainParametersProxy"], input)
	var res *big.Int
	if err := parsed.UnpackIntoInterface(&res, "blockGasLimit", output); err != nil {
		log.Crit("unpack failed", "err", err.Error())
	}
	log.Info("getBlockGasLimit", "res", toCoin(res), "res0", res)
}
