package handler

import (
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"io"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"
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
	proxyAddress := common.HexToAddress("0xcdB66B1e6A07279df98f804d0aCAC86695F4b99e")
	implAddress := common.HexToAddress("0x40d1215e14A94be82902C5f1CC2a5d438641E4Ff")
	from := common.HexToAddress("0xB3396Fef0cfC3A5d68b3Ef17e5016815af63102B")
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}
	setImplAddress(endpoint, from, privateKey, proxyAddress, implAddress)
}

func Test_isPendingDeRegisterValidator(t *testing.T) {
	addr := "0x5d643dfb9ae372ce4fdbc80890156e2cd8290846"
	isPendingDeRegisterValidator(endpoint, common.HexToAddress(addr))
	//addrs := []string{"0x75f5a34cEB6CaB0f8e3A8fF9038ba972932F816A",
	//	"0xa4a674C82E65ed0629C9532afD0bfdE9e6ddf6f3",
	//	"0x8D3397d2Bd0496ef5F098d5dFaE858128fA7fB56",
	//	"0xA3251F092CD7a5aFc3681a7e82db47290ce1ba12",
	//	"0x3ddB45aC146bBEa5dd11264e862d25ecd1031E5d",
	//	"0xD0751F3781Bcf3a6D17B0AF6f6c6D0d338D57393",
	//	"0x70FFE25404Cd23C591662C9f8A444A94b233E7C9",
	//	"0xC9509F6140197148e3A76d8877620e6F2b413c92",
	//	"0xA591838676Dcb7f7D7919AA6aEB47af1e2225FAb",
	//	"0x80D73761229F9Bea7F0bC9E6140831D458a5d5b0",
	//	"0xf27303A7dfEdef918DFE4706A61CD1147901dC91",
	//	"0x85b629CA2794aB562c562fb4D51E8db98f6BE5b9"}
	//
	//for _, a := range addrs {
	//	isPendingDeRegisterValidator(endpoint, common.HexToAddress(a))
	//}

}
func Test_getActiveVotesForValidator(t *testing.T) {
	height1, height2 := big.NewInt(2900000), big.NewInt(2950000)
	addr1, addr2 := common.HexToAddress("0x44b39830a0215a0904137c4474927dcfd049acbb"), common.HexToAddress("0xdc9e2ea9c16c75e22b1aa904d6c94ca70d0c57f3")
	getActiveVotesForValidator(endpoint, addr1, height1)
	getActiveVotesForValidator(endpoint, addr1, height2)
	getActiveVotesForValidator(endpoint, addr2, height1)
	getActiveVotesForValidator(endpoint, addr2, height2)
}
func TestBatchTransaction(t *testing.T) {
	validatorFile, voterFile := "validator.csv", "voter.csv"
	validators := loadFilesForValidator2(validatorFile)
	sum0, sum1, sum2 := big.NewInt(0), big.NewInt(0), big.NewInt(0)
	for _, v := range validators {
		sum0 = sum0.Add(sum0, v)
	}
	voters, voter_value := loadFilesForVoter2(voterFile)
	addr0 := common.HexToAddress("0xc052261da7602245558b297c587a8545e67d1109")
	for i, v := range voters {
		sum1 = sum1.Add(sum1, voter_value[i])
		if v == addr0 {
			sum2 = sum2.Add(sum2, voter_value[i])
		}
	}
	fmt.Println("sum0", sum0.String(), "sum1", sum1.String(), "sum2", sum2, toCoin(sum2).String())
	fmt.Println("sum", sum0.Add(sum0, sum1).String())
	fmt.Println("sum3", sum1.Sub(sum1, sum2).String())
}
func TestBatchValidators(t *testing.T) {
	from := common.HexToAddress("0xe05665E26eb7da077B2AAeD5cDe1DB47dE6B4544")
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}
	validatorFile := "validator.csv"
	validators := loadFilesForValidator2(validatorFile)
	sum := big.NewInt(0)
	fmt.Println(len(validators))
	for to, balance := range validators {
		fmt.Println(to, balance)
		sum = sum.Add(sum, balance)
		sendTransaction(endpoint, from, to, privateKey, balance)
		balanceOf(endpoint, to)
		time.Sleep(5 * time.Second)
	}
	fmt.Println(sum.String(), toCoin(sum))
}
func TestBatchVoters(t *testing.T) {
	from := common.HexToAddress("0xe05665E26eb7da077B2AAeD5cDe1DB47dE6B4544")
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}
	voterFile := "voter.csv"
	voters, voter_value := loadFilesForVoter2(voterFile)
	addr0 := common.HexToAddress("0xc052261da7602245558b297c587a8545e67d1109")
	sum, count, count2 := big.NewInt(0), 0, 0
	fmt.Println(len(voters))
	for i, to := range voters {
		balance := voter_value[i]
		fmt.Println("index", i, to, balance)
		if to == addr0 {
			count++
		} else {
			sum = sum.Add(sum, balance)
			if balance.Sign() > 0 {
				count2++
				sendTransaction(endpoint, from, to, privateKey, balance)
				balanceOf(endpoint, to)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}
	fmt.Println("sum", sum.String(), toCoin(sum), "count", count, count2)
}
func loadFilesForValidator2(fileName string) map[common.Address]*big.Int {
	fmt.Println("准备读取文件.....文件名:", fileName)
	fs, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("can not open the file, err is %+v", err))
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	result := make(map[common.Address]*big.Int)
	//针对大文件，一行一行的读取文件
	fmt.Println("加载文件.....")
	pos := 0
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			panic(fmt.Sprintf("can not read, err is %+v", err))
		}
		if err == io.EOF {
			break
		}
		if pos > 0 {
			addr, v := handleRow0(row)
			if _, ok := result[addr]; ok {
				result[addr] = new(big.Int).Add(v, result[addr])
			} else {
				result[addr] = new(big.Int).Set(v)
			}
		}
		pos++
	}
	fmt.Println("加载文件结束.....，读取", pos, "条记录")
	return result
}
func loadFilesForVoter2(fileName string) ([]common.Address, []*big.Int) {
	fmt.Println("准备读取文件.....文件名:", fileName)
	fs, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("can not open the file, err is %+v", err))
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	result0, result1 := make([]common.Address, 0, 0), make([]*big.Int, 0, 0)
	//针对大文件，一行一行的读取文件
	fmt.Println("加载文件.....")
	pos := 0
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			panic(fmt.Sprintf("can not read, err is %+v", err))
		}
		if err == io.EOF {
			break
		}
		if pos > 0 {
			addr, v := handleRow1_0(row)
			result0 = append(result0, addr)
			result1 = append(result1, v)
		}
		pos++
	}
	fmt.Println("加载文件结束.....，读取", pos, "条记录")
	return result0, result1
}

// validator
func handleRow0(record []string) (common.Address, *big.Int) {
	v_addr := common.HexToAddress(record[0])
	value := record[2]
	value = strings.Trim(value, " ")
	r, b := new(big.Int).SetString(value, 10)
	if !b {
		panic(fmt.Errorf("load error,:%s", value))
	}
	return v_addr, r
}

// votor
func handleRow1(record []string) (common.Address, *big.Int) {
	v_addr := common.HexToAddress(record[2])
	value := strings.Trim(record[3], " ")
	d, e := decimal.NewFromString(value)
	if e != nil {
		panic(fmt.Errorf("NewFromString %v", e))
	}
	fmt.Println(v_addr, value, d.String())
	return v_addr, toWei(d.BigFloat())
}
func handleRow1_0(record []string) (common.Address, *big.Int) {
	v_addr := common.HexToAddress(record[2])
	value := strings.Trim(record[4], " ")
	r, b := new(big.Int).SetString(value, 10)
	if !b {
		panic(fmt.Errorf("load error,:%s", value))
	}
	return v_addr, r
}

type validatorInfo struct {
	value1 *big.Int
	value2 *big.Int
}
type voterInfo struct {
	addr  common.Address
	value *big.Int
}

func TestCheckTheCsvData(t *testing.T) {
	validatorFile, voterFile := "validator.csv", "voter.csv"
	validators := loadFilesForValidator(validatorFile)
	voters := loadFilesForVoter(voterFile)
	//addr0 := common.HexToAddress("0xc052261da7602245558b297c587a8545e67d1109")

	for k, v := range validators {
		if vv, ok := voters[k]; ok {
			sum0 := big.NewInt(0)
			for _, v2 := range vv {
				sum0 = sum0.Add(sum0, v2.value)
			}
			if sum0.Cmp(v.value2) != 0 {
				println("index", k.String(), "sum0", sum0.String(), "sum1", v.value2.String())
			}
		} else {
			panic("panic.....")
		}
	}
}
func loadFilesForValidator(fileName string) map[common.Address]*validatorInfo {
	fmt.Println("准备读取文件.....文件名:", fileName)
	fs, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("can not open the file, err is %+v", err))
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	result := make(map[common.Address]*validatorInfo)
	//针对大文件，一行一行的读取文件
	fmt.Println("加载文件.....")
	pos := 0
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			panic(fmt.Sprintf("can not read, err is %+v", err))
		}
		if err == io.EOF {
			break
		}
		if pos > 0 {
			addr, v1, v2 := handleRow2(row)
			if _, ok := result[addr]; ok {
				panic("validator file wrong")
			} else {
				result[addr] = &validatorInfo{
					value1: new(big.Int).Set(v1),
					value2: new(big.Int).Set(v2),
				}
			}
		}
		pos++
	}
	fmt.Println("加载文件结束.....，读取", pos, "条记录")
	return result
}
func loadFilesForVoter(fileName string) map[common.Address][]*voterInfo {
	fmt.Println("准备读取文件.....文件名:", fileName)
	fs, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("can not open the file, err is %+v", err))
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	result := make(map[common.Address][]*voterInfo)
	//针对大文件，一行一行的读取文件
	fmt.Println("加载文件.....")
	pos := 0
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			panic(fmt.Sprintf("can not read, err is %+v", err))
		}
		if err == io.EOF {
			break
		}
		if pos > 0 {
			addr0, addr1, v := handleRow3(row)
			if _, ok := result[addr0]; ok {
				result[addr0] = append(result[addr0], &voterInfo{
					addr:  addr1,
					value: new(big.Int).Set(v),
				})
			} else {
				result[addr0] = []*voterInfo{&voterInfo{
					addr:  addr1,
					value: new(big.Int).Set(v),
				}}
			}
		}
		pos++
	}
	fmt.Println("加载文件结束.....，读取", pos, "条记录")
	return result
}

// validator
func handleRow2(record []string) (common.Address, *big.Int, *big.Int) {
	v_addr := common.HexToAddress(record[0])
	value := strings.Trim(record[2], " ")
	r, b := new(big.Int).SetString(value, 10)
	if !b {
		panic(fmt.Errorf("load error,:%s", value))
	}
	value2 := strings.Trim(record[4], " ")
	r1, b1 := new(big.Int).SetString(value2, 10)
	if !b1 {
		panic(fmt.Errorf("load error,:%s", value2))
	}
	return v_addr, r, r1
}

// votor
func handleRow3(record []string) (common.Address, common.Address, *big.Int) {
	v_addr0 := common.HexToAddress(record[1])
	v_addr := common.HexToAddress(record[2])
	value := strings.Trim(record[4], " ")
	r, b := new(big.Int).SetString(value, 10)
	if !b {
		panic(fmt.Errorf("load error,:%s", value))
	}
	return v_addr0, v_addr, r
}

func Test02(t *testing.T) {
	str := "4.28E-05"
	d, e := decimal.NewFromString(str)
	if e != nil {
		panic(fmt.Errorf("NewFromString %v", e))
	}
	fmt.Println(d.BigInt().String(), d.String())
}
func TestMakeAddress(t *testing.T) {
	for i := 0; i < 3; i++ {
		priv, _ := crypto.GenerateKey()
		privHex := hex.EncodeToString(crypto.FromECDSA(priv))
		fmt.Println(privHex)
		addr := crypto.PubkeyToAddress(priv.PublicKey)
		fmt.Println(addr.String())
	}

	fmt.Println("finish")
}
func Test03(t *testing.T) {
	from := common.HexToAddress("0xd34D198B85B491F8b0E3b5Dac5a8c29D49b227E5")
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}
	addrs := []string{"0x2a0fc7506b248fEA4775004B6c33a348e9AEec69",
		"0xFD7ff0b5f4446ae35468d1e599CaaebcbEfe88B0",
		"0x7979b84dF6aA903e37B994d9198662eFFe38C68b",
		"0x34405391e12E477B6Ad06De15B8e738F23De7A72"}
	tos := make([]common.Address, 0, 0)
	for _, a := range addrs {
		tos = append(tos, common.HexToAddress(a))
	}
	balance := big.NewInt(20 * 1e9)
	for _, to := range tos {
		sendTransaction(endpoint, from, to, privateKey, balance)
		balanceOf(endpoint, to)
		time.Sleep(5 * time.Second)
	}
}
func Test04(t *testing.T) {
	to := common.HexToAddress("0xe05665E26eb7da077B2AAeD5cDe1DB47dE6B4544")
	balanceOf(endpoint, to)
}
func Test_getAccountTotalLockedGold(t *testing.T) {
	addr1 := common.HexToAddress("0x979b8ba0A9ddD4Bf4b71A555A6109ef770F778cB")
	getAccountTotalLockedGold(endpoint, addr1, nil)
}
func Test_getAccountNonvotingLockedGold(t *testing.T) {
	addr1 := common.HexToAddress("0x979b8ba0A9ddD4Bf4b71A555A6109ef770F778cB")
	getAccountNonvotingLockedGold(endpoint, addr1, nil)
}
