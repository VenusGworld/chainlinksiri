// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf_beacon

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

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

type ECCArithmeticG1Point struct {
	P [2]*big.Int
}

type KeyDataStructKeyData struct {
	PublicKey []byte
	Hashes    [][32]byte
}

type VRFBeaconReportReport struct {
	Outputs            []VRFBeaconTypesVRFOutput
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	RecentBlockHeight  uint64
	RecentBlockHash    [32]byte
}

type VRFBeaconTypesCallback struct {
	RequestID      *big.Int
	NumWords       uint16
	Requester      common.Address
	Arguments      []byte
	GasAllowance   *big.Int
	SubID          *big.Int
	GasPrice       *big.Int
	WeiPerUnitLink *big.Int
}

type VRFBeaconTypesCostedCallback struct {
	Callback VRFBeaconTypesCallback
	Price    *big.Int
}

type VRFBeaconTypesOutputServed struct {
	Height            uint64
	ConfirmationDelay *big.Int
	ProofG1X          *big.Int
	ProofG1Y          *big.Int
}

type VRFBeaconTypesVRFOutput struct {
	BlockHeight       uint64
	ConfirmationDelay *big.Int
	VrfOutput         ECCArithmeticG1Point
	Callbacks         []VRFBeaconTypesCostedCallback
}

var VRFBeaconMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"contractIVRFCoordinatorProducerAPI\",\"name\":\"coordinator\",\"type\":\"address\"},{\"internalType\":\"contractDKG\",\"name\":\"keyProvider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"}],\"name\":\"CalldataLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotAcceptPayeeship\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateSigners\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"providedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"onchainHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorWrong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectNumberOfFaultyOracles\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTransmitters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numPayees\",\"type\":\"uint256\"}],\"name\":\"IncorrectNumberOfPayees\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"expectedNumSignatures\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"rsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ssLength\",\"type\":\"uint256\"}],\"name\":\"IncorrectNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredBalance\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPayee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"keyProvider\",\"type\":\"address\"}],\"name\":\"KeyInfoMustComeFromProvider\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LeftGasExceedsInitialGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeOwnerOrBillingAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"numFaultyOracles\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"numSigners\",\"type\":\"uint256\"}],\"name\":\"NumberOfFaultyOraclesTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"}],\"name\":\"OnchainConfigHasWrongLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"OnlyActiveSigners\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"OnlyActiveTransmitters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCurrentPayee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"existingPayee\",\"type\":\"address\"}],\"name\":\"PayeeAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"repeatedSignerAddress\",\"type\":\"address\"}],\"name\":\"RepeatedSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"repeatedTransmitterAddress\",\"type\":\"address\"}],\"name\":\"RepeatedTransmitter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reportHash\",\"type\":\"bytes32\"}],\"name\":\"SeenReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numSigners\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numTransmitters\",\"type\":\"uint256\"}],\"name\":\"SignersTransmittersMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maxOracles\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"providedOracles\",\"type\":\"uint256\"}],\"name\":\"TooManyOracles\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ocrVersion\",\"type\":\"uint64\"}],\"name\":\"UnknownConfigVersion\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"BillingAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maximumGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"observationPayment\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"transmissionPayment\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"BillingSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"name\":\"OraclePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"proofG1X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofG1Y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"OutputsServed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"requestIDs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint96[]\",\"name\":\"subBalances\",\"type\":\"uint96[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"subIDs\",\"type\":\"uint256[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAllowance\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSubBalance\",\"type\":\"uint256\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"}],\"name\":\"RandomnessRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSubBalance\",\"type\":\"uint256\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"acceptPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"vrfOutput\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"}],\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.CostedCallback[]\",\"name\":\"callbacks\",\"type\":\"tuple[]\"}],\"internalType\":\"structVRFBeaconTypes.VRFOutput[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"recentBlockHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVRFBeaconReport.Report\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"exposeType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBilling\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"maximumGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"observationPayment\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transmissionPayment\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBillingAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_coordinator\",\"outputs\":[{\"internalType\":\"contractIVRFCoordinatorProducerAPI\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_link\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"kd\",\"type\":\"tuple\"}],\"name\":\"keyGenerated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newKeyRequested\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"owedPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_keyID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_keyProvider\",\"outputs\":[{\"internalType\":\"contractDKG\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"maximumGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"observationPayment\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transmissionPayment\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"setBilling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"}],\"name\":\"setBillingAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"}],\"name\":\"setPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"transferPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162004d9f38038062004d9f8339810160408190526200003491620001c7565b8181858581813380600081620000915760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c457620000c48162000103565b5050506001600160a01b03918216608052811660a052600f80546001600160a01b03191695909116949094179093555060c05250620002219350505050565b336001600160a01b038216036200015d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000088565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6001600160a01b0381168114620001c457600080fd5b50565b60008060008060808587031215620001de57600080fd5b8451620001eb81620001ae565b6020860151909450620001fe81620001ae565b60408601519093506200021181620001ae565b6060959095015193969295505050565b60805160a05160c051614ac9620002d660003960006104810152600081816103710152818161114e015281816112260152818161131c0152818161141d015281816114bb01528181612216015281816122ee015281816124650152818161270b01528181612b1e01528181612bf601528181613100015261366801526000818161031a015281816112540152818161144a0152818161231c015281816124c301528181612c240152612feb0152614ac96000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c8063b121e147116100f9578063d09dc33911610097578063e53bbc9a11610071578063e53bbc9a146104c7578063eb5dcd6c146104da578063f2fde38b146104ed578063fbffd2c11461050057600080fd5b8063d09dc339146104a3578063d57fc45a146104ab578063e3d0e712146104b457600080fd5b8063bf2732c7116100d3578063bf2732c714610438578063c10753291461044b578063c4c92b371461045e578063cc31f7dd1461047c57600080fd5b8063b121e14714610401578063b1dc65a414610414578063b8be03cd1461042757600080fd5b80637d253aff116101665780638ac28d5a116101405780638ac28d5a146103935780638da5cb5b146103a65780639c849b30146103c4578063afcb95d7146103d757600080fd5b80637d253aff1461031557806381ff70481461033c5780638a1b17721461036c57600080fd5b80632f7527cc116101a25780632f7527cc146102a457806355e48749146102be5780635f27026f146102c857806379ba50971461030d57600080fd5b80630eafb25b146101c9578063181f5a77146101ef578063299372681461022e575b600080fd5b6101dc6101d73660046137b4565b610513565b6040519081526020015b60405180910390f35b604080518082018252600f81527f565246426561636f6e20312e302e300000000000000000000000000000000000602082015290516101e6919061383f565b60025460035460408051610100840467ffffffffffffffff9081168252690100000000000000000085048116602083015271010000000000000000000000000000000000909404841691810191909152918116606083015268010000000000000000900462ffffff16608082015260a0016101e6565b6102ac600881565b60405160ff90911681526020016101e6565b6102c661059f565b005b600f546102e89073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101e6565b6102c661061c565b6102e87f000000000000000000000000000000000000000000000000000000000000000081565b6004546005546040805163ffffffff808516825264010000000090940490931660208401528201526060016101e6565b6102e87f000000000000000000000000000000000000000000000000000000000000000081565b6102c66103a13660046137b4565b610719565b60005473ffffffffffffffffffffffffffffffffffffffff166102e8565b6102c66103d236600461389e565b610782565b6005546006546040805160008152602081019390935263ffffffff909116908201526060016101e6565b6102c661040f3660046137b4565b6109d4565b6102c661042236600461394c565b610acc565b6102c6610435366004613a03565b50565b6102c6610446366004613c47565b610f87565b6102c6610459366004613d14565b611046565b600e5473ffffffffffffffffffffffffffffffffffffffff166102e8565b6101dc7f000000000000000000000000000000000000000000000000000000000000000081565b6101dc6113e0565b6101dc60105481565b6102c66104c2366004613d77565b611573565b6102c66104d5366004613e76565b611dc9565b6102c66104e8366004613ee7565b61201b565b6102c66104fb3660046137b4565b612174565b6102c661050e3660046137b4565b612185565b73ffffffffffffffffffffffffffffffffffffffff811660009081526008602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046bffffffffffffffffffffffff1691810191909152906105875750600092915050565b604001516bffffffffffffffffffffffff1692915050565b600f5473ffffffffffffffffffffffffffffffffffffffff16338114610614576040517f292f4fb500000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff821660248201526044015b60405180910390fd5b506000601055565b60015473ffffffffffffffffffffffffffffffffffffffff16331461069d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161060b565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b73ffffffffffffffffffffffffffffffffffffffff8181166000908152600c6020526040902054163314610779576040517fdce38c2400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61043581612196565b61078a612569565b8281146107cd576040517f36d20459000000000000000000000000000000000000000000000000000000008152600481018490526024810182905260440161060b565b60005b838110156109cd5760008585838181106107ec576107ec613f20565b905060200201602081019061080191906137b4565b9050600084848481811061081757610817613f20565b905060200201602081019061082c91906137b4565b73ffffffffffffffffffffffffffffffffffffffff8084166000908152600c602052604090205491925016801580158161089257508273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b156108e9576040517febdf175600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80861660048301528316602482015260440161060b565b73ffffffffffffffffffffffffffffffffffffffff8481166000908152600c6020526040902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016858316908117909155908316146109b6578273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b360405160405180910390a45b5050505080806109c590613f7e565b9150506107d0565b5050505050565b73ffffffffffffffffffffffffffffffffffffffff8181166000908152600d6020526040902054163314610a34576040517f9d12ec4f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8181166000818152600c602090815260408083208054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217909355600d909452828520805490921690915590519416939092849290917f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b39190a45050565b60005a6040805160c08101825260025460ff808216835267ffffffffffffffff61010083048116602080860191909152690100000000000000000084048216858701527101000000000000000000000000000000000090930481166060850152600354908116608085015262ffffff680100000000000000009091041660a08401523360009081526008835293909320549394509092908c01359116610ba0576040517fb1c1f68e00000000000000000000000000000000000000000000000000000000815233600482015260240161060b565b6005548b3514610bea576005546040517f93df584c00000000000000000000000000000000000000000000000000000000815260048101919091528b35602482015260440161060b565b610bf88a8a8a8a8a8a6125ec565b8151610c05906001613fb6565b60ff1687141580610c165750868514155b15610c6e578151610c28906001613fb6565b6040517ffc33647500000000000000000000000000000000000000000000000000000000815260ff9091166004820152602481018890526044810186905260640161060b565b60008a8a604051610c80929190613fcf565b604051908190038120610c97918e90602001613fdf565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152828252805160209182012083830190925260008084529083018190529092509060005b8a811015610e885760006001858a8460208110610d0457610d04613f20565b610d1191901a601b613fb6565b8f8f86818110610d2357610d23613f20565b905060200201358e8e87818110610d3c57610d3c613f20565b9050602002013560405160008152602001604052604051610d79949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015610d9b573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff811660009081526009602090815290849020838501909452925460ff8082161515808552610100909204169383019390935290955092509050610e61576040517f20fb74ee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260240161060b565b826020015160080260ff166001901b84019350508080610e8090613f7e565b915050610ce5565b5081827e010101010101010101010101010101010101010101010101010101010101011614610ee3576040517fc103be2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060009150819050610f328d826020020135848e8e8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061267c92505050565b600680547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff600888901c161790559092509050610f788483838833612861565b50505050505050505050505050565b600f5473ffffffffffffffffffffffffffffffffffffffff16338114610ff7576040517f292f4fb500000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff8216602482015260440161060b565b81516040516110099190602001613ff3565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101206010555050565b60005473ffffffffffffffffffffffffffffffffffffffff1633148015906111075750600e546040517f6b14daf800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690636b14daf8906110c49033906000903690600401614058565b602060405180830381865afa1580156110e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111059190614088565b155b1561113e576040517fc04ecc2800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006111486129bf565b905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663597d2f3c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156111b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111db91906140aa565b905060006111e982846140c3565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660048301529192506000917f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa15801561129b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112bf91906140aa565b905081811015611305576040517fcf479181000000000000000000000000000000000000000000000000000000008152600481018290526024810183905260440161060b565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663f99b1d688761135561134f86866140d6565b89612aba565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff90921660048301526024820152604401600060405180830381600087803b1580156113c057600080fd5b505af11580156113d4573d6000803e3d6000fd5b50505050505050505050565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116600483015260009182917f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015611491573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114b591906140aa565b905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663597d2f3c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611524573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061154891906140aa565b905060006115546129bf565b90508161156182856140e9565b61156b91906140e9565b935050505090565b888787601f8311156115bb576040517f809fc428000000000000000000000000000000000000000000000000000000008152601f60048201526024810184905260440161060b565b8183146115fe576040517f988a0804000000000000000000000000000000000000000000000000000000008152600481018490526024810183905260440161060b565b611609816003614110565b60ff168311611650576040517ffda9db7800000000000000000000000000000000000000000000000000000000815260ff821660048201526024810184905260440161060b565b61165c8160ff16612ad4565b611664612569565b60006040518060c001604052808f8f80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505081526020018d8d8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050509082525060ff8c1660208083019190915260408051601f8d0183900483028101830182528c8152920191908c908c908190840183828082843760009201919091525050509082525067ffffffffffffffff891660208083019190915260408051601f8a01839004830281018301825289815292019190899089908190840183828082843760009201919091525050509152509050611786612b0e565b600a5460005b8181101561187f576000600a82815481106117a9576117a9613f20565b6000918252602082200154600b805473ffffffffffffffffffffffffffffffffffffffff909216935090849081106117e3576117e3613f20565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff948516835260098252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016905594168252600890529190912080547fffffffffffffffffffffffffffffffffffff0000000000000000000000000000169055508061187781613f7e565b91505061178c565b5061188c600a60006136d5565b611898600b60006136d5565b60005b825151811015611c215760096000846000015183815181106118bf576118bf613f20565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff1682528101919091526040016000205460ff161561196357825180518290811061190c5761190c613f20565b60200260200101516040517f7451f83e00000000000000000000000000000000000000000000000000000000815260040161060b919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b604080518082019091526001815260ff82166020820152835180516009916000918590811061199457611994613f20565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281810192909252604001600090812083518154948401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009095169015157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff161761010060ff90951694909402939093179092558401518051600892919084908110611a4657611a46613f20565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff1682528101919091526040016000205460ff1615611aec5782602001518181518110611a9557611a95613f20565b60200260200101516040517fe8d2989900000000000000000000000000000000000000000000000000000000815260040161060b919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b60405180606001604052806001151581526020018260ff16815260200160006bffffffffffffffffffffffff168152506008600085602001518481518110611b3657611b36613f20565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040908101600020835181549385015194909201516bffffffffffffffffffffffff1662010000027fffffffffffffffffffffffffffffffffffff000000000000000000000000ffff60ff95909516610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff931515939093167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090941693909317919091179290921617905580611c1981613f7e565b91505061189b565b5081518051611c3891600a916020909101906136f3565b506020808301518051611c4f92600b9201906136f3565b506040820151600280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600454640100000000900463ffffffff16611c9f613170565b6004805463ffffffff928316640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff82168117909255600092611ced928116911617600161412c565b905080600460006101000a81548163ffffffff021916908363ffffffff1602179055506000611d4146308463ffffffff16886000015189602001518a604001518b606001518c608001518d60a00151613207565b9050806005819055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e058360055484886000015189602001518a604001518b606001518c608001518d60a00151604051611da39998979695949392919061419a565b60405180910390a1611db58d8d6132b2565b505050505050505050505050505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314801590611e8a5750600e546040517f6b14daf800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690636b14daf890611e479033906000903690600401614058565b602060405180830381865afa158015611e64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e889190614088565b155b15611ec1576040517fc04ecc2800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611ec9612b0e565b6002805467ffffffffffffffff8581167101000000000000000000000000000000000081027fffffffffffffff0000000000000000ffffffffffffffffffffffffffffffffff898416690100000000000000000081027fffffffffffffffffffffffffffffff0000000000000000ffffffffffffffffff8d87166101008102919091167fffffffffffffffffffffffffffffff00000000000000000000000000000000ff909816979097171791909116919091179094556003805462ffffff87166801000000000000000081027fffffffffffffffffffffffffffffffffffffffffff00000000000000000000009092169489169485179190911790915560408051948552602085019590955293830152606082015260808101919091527f49275ddcdfc9c0519b3d094308c8bf675f06070a754ce90c152163cb6e66e8a09060a00160405180910390a15050505050565b73ffffffffffffffffffffffffffffffffffffffff8281166000908152600c602052604090205416331461207b576040517fdce38c2400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811633036120ca576040517fb387a23800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8083166000908152600d6020526040902080548383167fffffffffffffffffffffffff00000000000000000000000000000000000000008216811790925590911690811461216f5760405173ffffffffffffffffffffffffffffffffffffffff8084169133918616907f84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e3836790600090a45b505050565b61217c612569565b610435816132c0565b61218d612569565b610435816133b5565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600860209081526040918290208251606081018452905460ff80821615158084526101008304909116938301939093526201000090046bffffffffffffffffffffffff1692810192909252612205575050565b600061221083610513565b905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663597d2f3c6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561227f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122a391906140aa565b905060006122b182846140c3565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660048301529192506000917f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015612363573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061238791906140aa565b9050818110156123cd576040517fcf479181000000000000000000000000000000000000000000000000000000008152600481018290526024810183905260440161060b565b83156125615773ffffffffffffffffffffffffffffffffffffffff8681166000908152600c602090815260408083205460089092529182902080547fffffffffffffffffffffffffffffffffffff000000000000000000000000ffff16905590517ff99b1d680000000000000000000000000000000000000000000000000000000081529082166004820181905260248201879052917f0000000000000000000000000000000000000000000000000000000000000000169063f99b1d6890604401600060405180830381600087803b1580156124a957600080fd5b505af11580156124bd573d6000803e3d6000fd5b505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c8860405161255791815260200190565b60405180910390a4505b505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146125ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161060b565b565b60006125f9826020614230565b612604856020614230565b612610886101446140c3565b61261a91906140c3565b61262491906140c3565b61262f9060006140c3565b9050368114612673576040517ff7b94f0a0000000000000000000000000000000000000000000000000000000081526004810182905236602482015260440161060b565b50505050505050565b6000806000838060200190518101906126959190614474565b905060006126a6826060015161345d565b90508082608001511461270957608082015160608301516040517faed0afe500000000000000000000000000000000000000000000000000000000815260048101929092526024820183905267ffffffffffffffff16604482015260640161060b565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639cdb000a83600001518460200151856040015186606001516040518563ffffffff1660e01b815260040161277894939291906147bf565b600060405180830381600087803b15801561279257600080fd5b505af11580156127a6573d6000803e3d6000fd5b505050508564ffffffffff167ffc3c7a7927e878a0fca37c904953c3c75cee3ca1d1640184a0ab1c65eec6274333846020015185604001518b604051612841949392919073ffffffffffffffffffffffffffffffffffffffff94909416845277ffffffffffffffffffffffffffffffffffffffffffffffff92909216602084015267ffffffffffffffff166040830152606082015260800190565b60405180910390a281602001518260400151935093505050935093915050565b600061288d3a67ffffffffffffffff86161561287d5785612883565b87604001515b8860200151613540565b90506010360260005a905060006128b68663ffffffff1685858c60a0015162ffffff1686613591565b90506000670de0b6b3a764000077ffffffffffffffffffffffffffffffffffffffffffffffff8a16830273ffffffffffffffffffffffffffffffffffffffff881660009081526008602052604090205460808d01519290910492506201000090046bffffffffffffffffffffffff9081169167ffffffffffffffff168284010190811682111561294c57505050505050506109cd565b73ffffffffffffffffffffffffffffffffffffffff8816600090815260086020526040902080546bffffffffffffffffffffffff90921662010000027fffffffffffffffffffffffffffffffffffff000000000000000000000000ffff9092169190911790555050505050505050505050565b600080600b805480602002602001604051908101604052809291908181526020018280548015612a2557602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116129fa575b505083519394506000925050505b81811015612ab45760086000848381518110612a5157612a51613f20565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054612aa0906201000090046bffffffffffffffffffffffff16856140c3565b935080612aac81613f7e565b915050612a33565b50505090565b600081831015612acb575081612ace565b50805b92915050565b80600003610435576040517fe77dba5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000612b186129bf565b905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663597d2f3c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612b87573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bab91906140aa565b90506000612bb982846140c3565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660048301529192506000917f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015612c6b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c8f91906140aa565b905081811015612cd5576040517fcf479181000000000000000000000000000000000000000000000000000000008152600481018290526024810183905260440161060b565b6000600b805480602002602001604051908101604052809291908181526020018280548015612d3a57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311612d0f575b5050505050905060008151905060008167ffffffffffffffff811115612d6257612d62613a3e565b604051908082528060200260200182016040528015612d8b578160200160208202803683370190505b50905060008267ffffffffffffffff811115612da957612da9613a3e565b604051908082528060200260200182016040528015612dd2578160200160208202803683370190505b5090506000805b848110156130ac57600060086000888481518110612df957612df9613f20565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160029054906101000a90046bffffffffffffffffffffffff166bffffffffffffffffffffffff169050600060086000898581518110612e7f57612e7f613f20565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002080546bffffffffffffffffffffffff9290921662010000027fffffffffffffffffffffffffffffffffffff000000000000000000000000ffff9092169190911790558080156130a2576000600c60008a8681518110612f1057612f10613f20565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080878681518110612f8857612f88613f20565b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505081868681518110612fd557612fd5613f20565b60200260200101818152505084806001019550507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff168a868151811061304957613049613f20565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c8560405161309891815260200190565b60405180910390a4505b5050600101612dd9565b50815181146130bc578082528083525b815115613165576040517f73433a2f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906373433a2f906131379086908690600401614918565b600060405180830381600087803b15801561315157600080fd5b505af1158015610f78573d6000803e3d6000fd5b505050505050505050565b60004661a4b1811480613185575062066eed81145b1561320057606473ffffffffffffffffffffffffffffffffffffffff1663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156131d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131fa91906140aa565b91505090565b4391505090565b6000808a8a8a8a8a8a8a8a8a60405160200161322b99989796959493929190614962565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b6132bc82826135d9565b5050565b3373ffffffffffffffffffffffffffffffffffffffff82160361333f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161060b565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600e5473ffffffffffffffffffffffffffffffffffffffff90811690821681146132bc57600e80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84811691821790925560408051928416835260208301919091527f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912910160405180910390a15050565b60004661a4b1811480613472575062066eed81145b15613530576101008367ffffffffffffffff1661348d613170565b61349791906140d6565b11156134a65750600092915050565b6040517f2b407a8200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152606490632b407a8290602401602060405180830381865afa158015613505573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061352991906140aa565b9392505050565b505067ffffffffffffffff164090565b60008367ffffffffffffffff8416811015613574576002858567ffffffffffffffff160381613571576135716148e9565b04015b613588818467ffffffffffffffff16612aba565b95945050505050565b6000818610156135cd576040517f3fef97df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50909303019091010290565b61010081811461361b578282826040517f418a179b00000000000000000000000000000000000000000000000000000000815260040161060b939291906149f7565b600061362983850185614a1b565b90506040517f8eef585f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690638eef585f9061369d908490600401614a85565b600060405180830381600087803b1580156136b757600080fd5b505af11580156136cb573d6000803e3d6000fd5b5050505050505050565b5080546000825590600052602060002090810190610435919061377d565b82805482825590600052602060002090810192821561376d579160200282015b8281111561376d57825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190613713565b5061377992915061377d565b5090565b5b80821115613779576000815560010161377e565b73ffffffffffffffffffffffffffffffffffffffff8116811461043557600080fd5b6000602082840312156137c657600080fd5b813561352981613792565b60005b838110156137ec5781810151838201526020016137d4565b50506000910152565b6000815180845261380d8160208601602086016137d1565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061352960208301846137f5565b60008083601f84011261386457600080fd5b50813567ffffffffffffffff81111561387c57600080fd5b6020830191508360208260051b850101111561389757600080fd5b9250929050565b600080600080604085870312156138b457600080fd5b843567ffffffffffffffff808211156138cc57600080fd5b6138d888838901613852565b909650945060208701359150808211156138f157600080fd5b506138fe87828801613852565b95989497509550505050565b60008083601f84011261391c57600080fd5b50813567ffffffffffffffff81111561393457600080fd5b60208301915083602082850101111561389757600080fd5b60008060008060008060008060e0898b03121561396857600080fd5b606089018a81111561397957600080fd5b8998503567ffffffffffffffff8082111561399357600080fd5b61399f8c838d0161390a565b909950975060808b01359150808211156139b857600080fd5b6139c48c838d01613852565b909750955060a08b01359150808211156139dd57600080fd5b506139ea8b828c01613852565b999c989b50969995989497949560c00135949350505050565b600060208284031215613a1557600080fd5b813567ffffffffffffffff811115613a2c57600080fd5b820160a0818503121561352957600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613a9057613a90613a3e565b60405290565b604051610100810167ffffffffffffffff81118282101715613a9057613a90613a3e565b60405160a0810167ffffffffffffffff81118282101715613a9057613a90613a3e565b6040516080810167ffffffffffffffff81118282101715613a9057613a90613a3e565b6040516020810167ffffffffffffffff81118282101715613a9057613a90613a3e565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613b6a57613b6a613a3e565b604052919050565b600067ffffffffffffffff821115613b8c57613b8c613a3e565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600067ffffffffffffffff821115613bd257613bd2613a3e565b5060051b60200190565b600082601f830112613bed57600080fd5b81356020613c02613bfd83613bb8565b613b23565b82815260059290921b84018101918181019086841115613c2157600080fd5b8286015b84811015613c3c5780358352918301918301613c25565b509695505050505050565b60006020808385031215613c5a57600080fd5b823567ffffffffffffffff80821115613c7257600080fd5b9084019060408287031215613c8657600080fd5b613c8e613a6d565b823582811115613c9d57600080fd5b8301601f81018813613cae57600080fd5b8035613cbc613bfd82613b72565b8181528987838501011115613cd057600080fd5b818784018883013760008783830101528084525050508383013582811115613cf757600080fd5b613d0388828601613bdc565b948201949094529695505050505050565b60008060408385031215613d2757600080fd5b8235613d3281613792565b946020939093013593505050565b803560ff81168114613d5157600080fd5b919050565b67ffffffffffffffff8116811461043557600080fd5b8035613d5181613d56565b60008060008060008060008060008060c08b8d031215613d9657600080fd5b8a3567ffffffffffffffff80821115613dae57600080fd5b613dba8e838f01613852565b909c509a5060208d0135915080821115613dd357600080fd5b613ddf8e838f01613852565b909a509850889150613df360408e01613d40565b975060608d0135915080821115613e0957600080fd5b613e158e838f0161390a565b9097509550859150613e2960808e01613d6c565b945060a08d0135915080821115613e3f57600080fd5b50613e4c8d828e0161390a565b915080935050809150509295989b9194979a5092959850565b62ffffff8116811461043557600080fd5b600080600080600060a08688031215613e8e57600080fd5b8535613e9981613d56565b94506020860135613ea981613d56565b93506040860135613eb981613d56565b92506060860135613ec981613d56565b91506080860135613ed981613e65565b809150509295509295909350565b60008060408385031215613efa57600080fd5b8235613f0581613792565b91506020830135613f1581613792565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203613faf57613faf613f4f565b5060010190565b60ff8181168382160190811115612ace57612ace613f4f565b8183823760009101908152919050565b828152606082602083013760800192915050565b600082516140058184602087016137d1565b9190910192915050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8416815260406020820152600061358860408301848661400f565b60006020828403121561409a57600080fd5b8151801515811461352957600080fd5b6000602082840312156140bc57600080fd5b5051919050565b80820180821115612ace57612ace613f4f565b81810381811115612ace57612ace613f4f565b818103600083128015838313168383128216171561410957614109613f4f565b5092915050565b60ff818116838216029081169081811461410957614109613f4f565b63ffffffff81811683821601908082111561410957614109613f4f565b600081518084526020808501945080840160005b8381101561418f57815173ffffffffffffffffffffffffffffffffffffffff168752958201959082019060010161415d565b509495945050505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526141ca8184018a614149565b905082810360808401526141de8189614149565b905060ff871660a084015282810360c08401526141fb81876137f5565b905067ffffffffffffffff851660e084015282810361010084015261422081856137f5565b9c9b505050505050505050505050565b8082028115828204841417612ace57612ace613f4f565b8051613d5181613d56565b805161ffff81168114613d5157600080fd5b8051613d5181613792565b600082601f83011261428057600080fd5b815161428e613bfd82613b72565b8181528460208386010111156142a357600080fd5b6142b48260208301602087016137d1565b949350505050565b80516bffffffffffffffffffffffff81168114613d5157600080fd5b600082601f8301126142e957600080fd5b815160206142f9613bfd83613bb8565b82815260059290921b8401810191818101908684111561431857600080fd5b8286015b84811015613c3c57805167ffffffffffffffff8082111561433c57600080fd5b908801907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06040838c038201121561437357600080fd5b61437b613a6d565b878401518381111561438c57600080fd5b8401610100818e03840112156143a157600080fd5b6143a9613a96565b92508881015183526143bd60408201614252565b898401526143cd60608201614264565b60408401526080810151848111156143e457600080fd5b6143f28e8b8385010161426f565b60608501525061440460a082016142bc565b608084015260c081015160a084015260e081015160c084015261010081015160e084015250818152614438604085016142bc565b81890152865250505091830191830161431c565b805177ffffffffffffffffffffffffffffffffffffffffffffffff81168114613d5157600080fd5b60006020828403121561448657600080fd5b815167ffffffffffffffff8082111561449e57600080fd5b9083019060a082860312156144b257600080fd5b6144ba613aba565b8251828111156144c957600080fd5b8301601f810187136144da57600080fd5b80516144e8613bfd82613bb8565b8082825260208201915060208360051b85010192508983111561450a57600080fd5b602084015b838110156146465780518781111561452657600080fd5b850160a0818d037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001121561455a57600080fd5b614562613add565b602082015161457081613d56565b8152604082015161458081613e65565b60208201526040828e037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa00112156145b757600080fd5b6145bf613b00565b8d607f8401126145ce57600080fd5b6145d6613a6d565b808f60a0860111156145e757600080fd5b606085015b60a086018110156146075780518352602092830192016145ec565b50825250604082015260a08201518981111561462257600080fd5b6146318e6020838601016142d8565b6060830152508452506020928301920161450f565b508452506146599150506020840161444c565b602082015261466a60408401614247565b604082015261467b60608401614247565b60608201526080830151608082015280935050505092915050565b600082825180855260208086019550808260051b84010181860160005b848110156147b2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086840301895281516040815181865280518287015286810151606061ffff8216818901528383015193506080915073ffffffffffffffffffffffffffffffffffffffff8416828901528083015193505061010060a081818a01526147446101408a01866137f5565b9284015192945060c06147668a8201856bffffffffffffffffffffffff169052565b9084015160e08a81019190915290840151918901919091529091015161012087015250908501516bffffffffffffffffffffffff169385019390935297830197908301906001016146b3565b5090979650505050505050565b6000608080830181845280885180835260a092508286019150828160051b8701016020808c016000805b8581101561488d578a85037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600187528251805167ffffffffffffffff1686528481015162ffffff16858701526040808201515190849088015b6002821015614861578251815291870191600191909101908701614842565b50505060600151858a0189905261487a868a0182614696565b97850197955050918301916001016147e9565b5050508196506148b88189018c77ffffffffffffffffffffffffffffffffffffffffffffffff169052565b5050505050506148d4604083018567ffffffffffffffff169052565b67ffffffffffffffff83166060830152613588565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60408152600061492b6040830185614149565b82810360208481019190915284518083528582019282019060005b818110156147b257845183529383019391830191600101614946565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526149a98285018b614149565b915083820360808501526149bd828a614149565b915060ff881660a085015283820360c08501526149da82886137f5565b90861660e0850152838103610100850152905061422081856137f5565b604081526000614a0b60408301858761400f565b9050826020830152949350505050565b6000610100808385031215614a2f57600080fd5b83601f840112614a3e57600080fd5b614a46613a96565b908301908085831115614a5857600080fd5b845b83811015614a7b578035614a6d81613e65565b835260209283019201614a5a565b5095945050505050565b6101008101818360005b6008811015614ab357815162ffffff16835260209283019290910190600101614a8f565b5050509291505056fea164736f6c6343000813000a",
}

var VRFBeaconABI = VRFBeaconMetaData.ABI

var VRFBeaconBin = VRFBeaconMetaData.Bin

func DeployVRFBeacon(auth *bind.TransactOpts, backend bind.ContractBackend, link common.Address, coordinator common.Address, keyProvider common.Address, keyID [32]byte) (common.Address, *types.Transaction, *VRFBeacon, error) {
	parsed, err := VRFBeaconMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBeaconBin), backend, link, coordinator, keyProvider, keyID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFBeacon{VRFBeaconCaller: VRFBeaconCaller{contract: contract}, VRFBeaconTransactor: VRFBeaconTransactor{contract: contract}, VRFBeaconFilterer: VRFBeaconFilterer{contract: contract}}, nil
}

type VRFBeacon struct {
	address common.Address
	abi     abi.ABI
	VRFBeaconCaller
	VRFBeaconTransactor
	VRFBeaconFilterer
}

type VRFBeaconCaller struct {
	contract *bind.BoundContract
}

type VRFBeaconTransactor struct {
	contract *bind.BoundContract
}

type VRFBeaconFilterer struct {
	contract *bind.BoundContract
}

type VRFBeaconSession struct {
	Contract     *VRFBeacon
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFBeaconCallerSession struct {
	Contract *VRFBeaconCaller
	CallOpts bind.CallOpts
}

type VRFBeaconTransactorSession struct {
	Contract     *VRFBeaconTransactor
	TransactOpts bind.TransactOpts
}

type VRFBeaconRaw struct {
	Contract *VRFBeacon
}

type VRFBeaconCallerRaw struct {
	Contract *VRFBeaconCaller
}

type VRFBeaconTransactorRaw struct {
	Contract *VRFBeaconTransactor
}

func NewVRFBeacon(address common.Address, backend bind.ContractBackend) (*VRFBeacon, error) {
	abi, err := abi.JSON(strings.NewReader(VRFBeaconABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFBeacon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFBeacon{address: address, abi: abi, VRFBeaconCaller: VRFBeaconCaller{contract: contract}, VRFBeaconTransactor: VRFBeaconTransactor{contract: contract}, VRFBeaconFilterer: VRFBeaconFilterer{contract: contract}}, nil
}

func NewVRFBeaconCaller(address common.Address, caller bind.ContractCaller) (*VRFBeaconCaller, error) {
	contract, err := bindVRFBeacon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCaller{contract: contract}, nil
}

func NewVRFBeaconTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFBeaconTransactor, error) {
	contract, err := bindVRFBeacon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconTransactor{contract: contract}, nil
}

func NewVRFBeaconFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFBeaconFilterer, error) {
	contract, err := bindVRFBeacon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconFilterer{contract: contract}, nil
}

func bindVRFBeacon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFBeaconMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFBeacon *VRFBeaconRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeacon.Contract.VRFBeaconCaller.contract.Call(opts, result, method, params...)
}

func (_VRFBeacon *VRFBeaconRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeacon.Contract.VRFBeaconTransactor.contract.Transfer(opts)
}

func (_VRFBeacon *VRFBeaconRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeacon.Contract.VRFBeaconTransactor.contract.Transact(opts, method, params...)
}

func (_VRFBeacon *VRFBeaconCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeacon.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFBeacon *VRFBeaconTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeacon.Contract.contract.Transfer(opts)
}

func (_VRFBeacon *VRFBeaconTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeacon.Contract.contract.Transact(opts, method, params...)
}

func (_VRFBeacon *VRFBeaconCaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFBeacon.Contract.NUMCONFDELAYS(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFBeacon.Contract.NUMCONFDELAYS(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) GetBilling(opts *bind.CallOpts) (GetBilling,

	error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "getBilling")

	outstruct := new(GetBilling)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaximumGasPrice = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ReasonableGasPrice = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ObservationPayment = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.TransmissionPayment = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.AccountingGas = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_VRFBeacon *VRFBeaconSession) GetBilling() (GetBilling,

	error) {
	return _VRFBeacon.Contract.GetBilling(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) GetBilling() (GetBilling,

	error) {
	return _VRFBeacon.Contract.GetBilling(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) GetBillingAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "getBillingAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) GetBillingAccessController() (common.Address, error) {
	return _VRFBeacon.Contract.GetBillingAccessController(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) GetBillingAccessController() (common.Address, error) {
	return _VRFBeacon.Contract.GetBillingAccessController(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) ICoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "i_coordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) ICoordinator() (common.Address, error) {
	return _VRFBeacon.Contract.ICoordinator(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) ICoordinator() (common.Address, error) {
	return _VRFBeacon.Contract.ICoordinator(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) ILink(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "i_link")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) ILink() (common.Address, error) {
	return _VRFBeacon.Contract.ILink(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) ILink() (common.Address, error) {
	return _VRFBeacon.Contract.ILink(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_VRFBeacon *VRFBeaconSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _VRFBeacon.Contract.LatestConfigDetails(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _VRFBeacon.Contract.LatestConfigDetails(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_VRFBeacon *VRFBeaconSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _VRFBeacon.Contract.LatestConfigDigestAndEpoch(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _VRFBeacon.Contract.LatestConfigDigestAndEpoch(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "linkAvailableForPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) LinkAvailableForPayment() (*big.Int, error) {
	return _VRFBeacon.Contract.LinkAvailableForPayment(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _VRFBeacon.Contract.LinkAvailableForPayment(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) OwedPayment(opts *bind.CallOpts, transmitterAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "owedPayment", transmitterAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _VRFBeacon.Contract.OwedPayment(&_VRFBeacon.CallOpts, transmitterAddress)
}

func (_VRFBeacon *VRFBeaconCallerSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _VRFBeacon.Contract.OwedPayment(&_VRFBeacon.CallOpts, transmitterAddress)
}

func (_VRFBeacon *VRFBeaconCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) Owner() (common.Address, error) {
	return _VRFBeacon.Contract.Owner(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) Owner() (common.Address, error) {
	return _VRFBeacon.Contract.Owner(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) SKeyID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "s_keyID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) SKeyID() ([32]byte, error) {
	return _VRFBeacon.Contract.SKeyID(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) SKeyID() ([32]byte, error) {
	return _VRFBeacon.Contract.SKeyID(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) SKeyProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "s_keyProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) SKeyProvider() (common.Address, error) {
	return _VRFBeacon.Contract.SKeyProvider(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) SKeyProvider() (common.Address, error) {
	return _VRFBeacon.Contract.SKeyProvider(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) SProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "s_provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) SProvingKeyHash() ([32]byte, error) {
	return _VRFBeacon.Contract.SProvingKeyHash(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) SProvingKeyHash() ([32]byte, error) {
	return _VRFBeacon.Contract.SProvingKeyHash(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) TypeAndVersion() (string, error) {
	return _VRFBeacon.Contract.TypeAndVersion(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) TypeAndVersion() (string, error) {
	return _VRFBeacon.Contract.TypeAndVersion(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "acceptOwnership")
}

func (_VRFBeacon *VRFBeaconSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFBeacon.Contract.AcceptOwnership(&_VRFBeacon.TransactOpts)
}

func (_VRFBeacon *VRFBeaconTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFBeacon.Contract.AcceptOwnership(&_VRFBeacon.TransactOpts)
}

func (_VRFBeacon *VRFBeaconTransactor) AcceptPayeeship(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "acceptPayeeship", transmitter)
}

func (_VRFBeacon *VRFBeaconSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.AcceptPayeeship(&_VRFBeacon.TransactOpts, transmitter)
}

func (_VRFBeacon *VRFBeaconTransactorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.AcceptPayeeship(&_VRFBeacon.TransactOpts, transmitter)
}

func (_VRFBeacon *VRFBeaconTransactor) ExposeType(opts *bind.TransactOpts, arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "exposeType", arg0)
}

func (_VRFBeacon *VRFBeaconSession) ExposeType(arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _VRFBeacon.Contract.ExposeType(&_VRFBeacon.TransactOpts, arg0)
}

func (_VRFBeacon *VRFBeaconTransactorSession) ExposeType(arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _VRFBeacon.Contract.ExposeType(&_VRFBeacon.TransactOpts, arg0)
}

func (_VRFBeacon *VRFBeaconTransactor) KeyGenerated(opts *bind.TransactOpts, kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "keyGenerated", kd)
}

func (_VRFBeacon *VRFBeaconSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _VRFBeacon.Contract.KeyGenerated(&_VRFBeacon.TransactOpts, kd)
}

func (_VRFBeacon *VRFBeaconTransactorSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _VRFBeacon.Contract.KeyGenerated(&_VRFBeacon.TransactOpts, kd)
}

func (_VRFBeacon *VRFBeaconTransactor) NewKeyRequested(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "newKeyRequested")
}

func (_VRFBeacon *VRFBeaconSession) NewKeyRequested() (*types.Transaction, error) {
	return _VRFBeacon.Contract.NewKeyRequested(&_VRFBeacon.TransactOpts)
}

func (_VRFBeacon *VRFBeaconTransactorSession) NewKeyRequested() (*types.Transaction, error) {
	return _VRFBeacon.Contract.NewKeyRequested(&_VRFBeacon.TransactOpts)
}

func (_VRFBeacon *VRFBeaconTransactor) SetBilling(opts *bind.TransactOpts, maximumGasPrice uint64, reasonableGasPrice uint64, observationPayment uint64, transmissionPayment uint64, accountingGas *big.Int) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "setBilling", maximumGasPrice, reasonableGasPrice, observationPayment, transmissionPayment, accountingGas)
}

func (_VRFBeacon *VRFBeaconSession) SetBilling(maximumGasPrice uint64, reasonableGasPrice uint64, observationPayment uint64, transmissionPayment uint64, accountingGas *big.Int) (*types.Transaction, error) {
	return _VRFBeacon.Contract.SetBilling(&_VRFBeacon.TransactOpts, maximumGasPrice, reasonableGasPrice, observationPayment, transmissionPayment, accountingGas)
}

func (_VRFBeacon *VRFBeaconTransactorSession) SetBilling(maximumGasPrice uint64, reasonableGasPrice uint64, observationPayment uint64, transmissionPayment uint64, accountingGas *big.Int) (*types.Transaction, error) {
	return _VRFBeacon.Contract.SetBilling(&_VRFBeacon.TransactOpts, maximumGasPrice, reasonableGasPrice, observationPayment, transmissionPayment, accountingGas)
}

func (_VRFBeacon *VRFBeaconTransactor) SetBillingAccessController(opts *bind.TransactOpts, _billingAccessController common.Address) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "setBillingAccessController", _billingAccessController)
}

func (_VRFBeacon *VRFBeaconSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.SetBillingAccessController(&_VRFBeacon.TransactOpts, _billingAccessController)
}

func (_VRFBeacon *VRFBeaconTransactorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.SetBillingAccessController(&_VRFBeacon.TransactOpts, _billingAccessController)
}

func (_VRFBeacon *VRFBeaconTransactor) SetConfig(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "setConfig", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_VRFBeacon *VRFBeaconSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _VRFBeacon.Contract.SetConfig(&_VRFBeacon.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_VRFBeacon *VRFBeaconTransactorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _VRFBeacon.Contract.SetConfig(&_VRFBeacon.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_VRFBeacon *VRFBeaconTransactor) SetPayees(opts *bind.TransactOpts, transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "setPayees", transmitters, payees)
}

func (_VRFBeacon *VRFBeaconSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.SetPayees(&_VRFBeacon.TransactOpts, transmitters, payees)
}

func (_VRFBeacon *VRFBeaconTransactorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.SetPayees(&_VRFBeacon.TransactOpts, transmitters, payees)
}

func (_VRFBeacon *VRFBeaconTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFBeacon *VRFBeaconSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.TransferOwnership(&_VRFBeacon.TransactOpts, to)
}

func (_VRFBeacon *VRFBeaconTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.TransferOwnership(&_VRFBeacon.TransactOpts, to)
}

func (_VRFBeacon *VRFBeaconTransactor) TransferPayeeship(opts *bind.TransactOpts, transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "transferPayeeship", transmitter, proposed)
}

func (_VRFBeacon *VRFBeaconSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.TransferPayeeship(&_VRFBeacon.TransactOpts, transmitter, proposed)
}

func (_VRFBeacon *VRFBeaconTransactorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.TransferPayeeship(&_VRFBeacon.TransactOpts, transmitter, proposed)
}

func (_VRFBeacon *VRFBeaconTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_VRFBeacon *VRFBeaconSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _VRFBeacon.Contract.Transmit(&_VRFBeacon.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_VRFBeacon *VRFBeaconTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _VRFBeacon.Contract.Transmit(&_VRFBeacon.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_VRFBeacon *VRFBeaconTransactor) WithdrawFunds(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "withdrawFunds", recipient, amount)
}

func (_VRFBeacon *VRFBeaconSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFBeacon.Contract.WithdrawFunds(&_VRFBeacon.TransactOpts, recipient, amount)
}

func (_VRFBeacon *VRFBeaconTransactorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFBeacon.Contract.WithdrawFunds(&_VRFBeacon.TransactOpts, recipient, amount)
}

func (_VRFBeacon *VRFBeaconTransactor) WithdrawPayment(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "withdrawPayment", transmitter)
}

func (_VRFBeacon *VRFBeaconSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.WithdrawPayment(&_VRFBeacon.TransactOpts, transmitter)
}

func (_VRFBeacon *VRFBeaconTransactorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.WithdrawPayment(&_VRFBeacon.TransactOpts, transmitter)
}

type VRFBeaconBillingAccessControllerSetIterator struct {
	Event *VRFBeaconBillingAccessControllerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconBillingAccessControllerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconBillingAccessControllerSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconBillingAccessControllerSet)
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

func (it *VRFBeaconBillingAccessControllerSetIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconBillingAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconBillingAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*VRFBeaconBillingAccessControllerSetIterator, error) {

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingAccessControllerSetIterator{contract: _VRFBeacon.contract, event: "BillingAccessControllerSet", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconBillingAccessControllerSet)
				if err := _VRFBeacon.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseBillingAccessControllerSet(log types.Log) (*VRFBeaconBillingAccessControllerSet, error) {
	event := new(VRFBeaconBillingAccessControllerSet)
	if err := _VRFBeacon.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconBillingSetIterator struct {
	Event *VRFBeaconBillingSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconBillingSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconBillingSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconBillingSet)
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

func (it *VRFBeaconBillingSetIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconBillingSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconBillingSet struct {
	MaximumGasPrice     uint64
	ReasonableGasPrice  uint64
	ObservationPayment  uint64
	TransmissionPayment uint64
	AccountingGas       *big.Int
	Raw                 types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterBillingSet(opts *bind.FilterOpts) (*VRFBeaconBillingSetIterator, error) {

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingSetIterator{contract: _VRFBeacon.contract, event: "BillingSet", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchBillingSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingSet) (event.Subscription, error) {

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconBillingSet)
				if err := _VRFBeacon.contract.UnpackLog(event, "BillingSet", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseBillingSet(log types.Log) (*VRFBeaconBillingSet, error) {
	event := new(VRFBeaconBillingSet)
	if err := _VRFBeacon.contract.UnpackLog(event, "BillingSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconConfigSetIterator struct {
	Event *VRFBeaconConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconConfigSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconConfigSet)
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

func (it *VRFBeaconConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterConfigSet(opts *bind.FilterOpts) (*VRFBeaconConfigSetIterator, error) {

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconConfigSetIterator{contract: _VRFBeacon.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconConfigSet)
				if err := _VRFBeacon.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseConfigSet(log types.Log) (*VRFBeaconConfigSet, error) {
	event := new(VRFBeaconConfigSet)
	if err := _VRFBeacon.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconNewTransmissionIterator struct {
	Event *VRFBeaconNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconNewTransmission)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconNewTransmission)
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

func (it *VRFBeaconNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconNewTransmission struct {
	EpochAndRound      *big.Int
	Transmitter        common.Address
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	ConfigDigest       [32]byte
	Raw                types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterNewTransmission(opts *bind.FilterOpts, epochAndRound []*big.Int) (*VRFBeaconNewTransmissionIterator, error) {

	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "NewTransmission", epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconNewTransmissionIterator{contract: _VRFBeacon.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *VRFBeaconNewTransmission, epochAndRound []*big.Int) (event.Subscription, error) {

	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "NewTransmission", epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconNewTransmission)
				if err := _VRFBeacon.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseNewTransmission(log types.Log) (*VRFBeaconNewTransmission, error) {
	event := new(VRFBeaconNewTransmission)
	if err := _VRFBeacon.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOraclePaidIterator struct {
	Event *VRFBeaconOraclePaid

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOraclePaidIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOraclePaid)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconOraclePaid)
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

func (it *VRFBeaconOraclePaidIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOraclePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOraclePaid struct {
	Transmitter common.Address
	Payee       common.Address
	Amount      *big.Int
	LinkToken   common.Address
	Raw         types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*VRFBeaconOraclePaidIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var linkTokenRule []interface{}
	for _, linkTokenItem := range linkToken {
		linkTokenRule = append(linkTokenRule, linkTokenItem)
	}

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOraclePaidIterator{contract: _VRFBeacon.contract, event: "OraclePaid", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *VRFBeaconOraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var linkTokenRule []interface{}
	for _, linkTokenItem := range linkToken {
		linkTokenRule = append(linkTokenRule, linkTokenItem)
	}

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOraclePaid)
				if err := _VRFBeacon.contract.UnpackLog(event, "OraclePaid", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseOraclePaid(log types.Log) (*VRFBeaconOraclePaid, error) {
	event := new(VRFBeaconOraclePaid)
	if err := _VRFBeacon.contract.UnpackLog(event, "OraclePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOutputsServedIterator struct {
	Event *VRFBeaconOutputsServed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOutputsServedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOutputsServed)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconOutputsServed)
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

func (it *VRFBeaconOutputsServedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOutputsServedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOutputsServed struct {
	RecentBlockHeight  uint64
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	OutputsServed      []VRFBeaconTypesOutputServed
	Raw                types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterOutputsServed(opts *bind.FilterOpts) (*VRFBeaconOutputsServedIterator, error) {

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOutputsServedIterator{contract: _VRFBeacon.contract, event: "OutputsServed", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchOutputsServed(opts *bind.WatchOpts, sink chan<- *VRFBeaconOutputsServed) (event.Subscription, error) {

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOutputsServed)
				if err := _VRFBeacon.contract.UnpackLog(event, "OutputsServed", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseOutputsServed(log types.Log) (*VRFBeaconOutputsServed, error) {
	event := new(VRFBeaconOutputsServed)
	if err := _VRFBeacon.contract.UnpackLog(event, "OutputsServed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOwnershipTransferRequestedIterator struct {
	Event *VRFBeaconOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconOwnershipTransferRequested)
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

func (it *VRFBeaconOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFBeaconOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOwnershipTransferRequestedIterator{contract: _VRFBeacon.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOwnershipTransferRequested)
				if err := _VRFBeacon.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFBeaconOwnershipTransferRequested, error) {
	event := new(VRFBeaconOwnershipTransferRequested)
	if err := _VRFBeacon.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOwnershipTransferredIterator struct {
	Event *VRFBeaconOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconOwnershipTransferred)
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

func (it *VRFBeaconOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFBeaconOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOwnershipTransferredIterator{contract: _VRFBeacon.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFBeaconOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOwnershipTransferred)
				if err := _VRFBeacon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseOwnershipTransferred(log types.Log) (*VRFBeaconOwnershipTransferred, error) {
	event := new(VRFBeaconOwnershipTransferred)
	if err := _VRFBeacon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconPayeeshipTransferRequestedIterator struct {
	Event *VRFBeaconPayeeshipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconPayeeshipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconPayeeshipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconPayeeshipTransferRequested)
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

func (it *VRFBeaconPayeeshipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconPayeeshipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconPayeeshipTransferRequested struct {
	Transmitter common.Address
	Current     common.Address
	Proposed    common.Address
	Raw         types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*VRFBeaconPayeeshipTransferRequestedIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var proposedRule []interface{}
	for _, proposedItem := range proposed {
		proposedRule = append(proposedRule, proposedItem)
	}

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconPayeeshipTransferRequestedIterator{contract: _VRFBeacon.contract, event: "PayeeshipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var proposedRule []interface{}
	for _, proposedItem := range proposed {
		proposedRule = append(proposedRule, proposedItem)
	}

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconPayeeshipTransferRequested)
				if err := _VRFBeacon.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParsePayeeshipTransferRequested(log types.Log) (*VRFBeaconPayeeshipTransferRequested, error) {
	event := new(VRFBeaconPayeeshipTransferRequested)
	if err := _VRFBeacon.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconPayeeshipTransferredIterator struct {
	Event *VRFBeaconPayeeshipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconPayeeshipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconPayeeshipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconPayeeshipTransferred)
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

func (it *VRFBeaconPayeeshipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconPayeeshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconPayeeshipTransferred struct {
	Transmitter common.Address
	Previous    common.Address
	Current     common.Address
	Raw         types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*VRFBeaconPayeeshipTransferredIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconPayeeshipTransferredIterator{contract: _VRFBeacon.contract, event: "PayeeshipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *VRFBeaconPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconPayeeshipTransferred)
				if err := _VRFBeacon.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParsePayeeshipTransferred(log types.Log) (*VRFBeaconPayeeshipTransferred, error) {
	event := new(VRFBeaconPayeeshipTransferred)
	if err := _VRFBeacon.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconRandomWordsFulfilledIterator struct {
	Event *VRFBeaconRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconRandomWordsFulfilled)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconRandomWordsFulfilled)
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

func (it *VRFBeaconRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconRandomWordsFulfilled struct {
	RequestIDs            []*big.Int
	SuccessfulFulfillment []byte
	TruncatedErrorData    [][]byte
	SubBalances           []*big.Int
	SubIDs                []*big.Int
	Raw                   types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts) (*VRFBeaconRandomWordsFulfilledIterator, error) {

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconRandomWordsFulfilledIterator{contract: _VRFBeacon.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFBeaconRandomWordsFulfilled) (event.Subscription, error) {

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconRandomWordsFulfilled)
				if err := _VRFBeacon.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseRandomWordsFulfilled(log types.Log) (*VRFBeaconRandomWordsFulfilled, error) {
	event := new(VRFBeaconRandomWordsFulfilled)
	if err := _VRFBeacon.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconRandomnessFulfillmentRequestedIterator struct {
	Event *VRFBeaconRandomnessFulfillmentRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconRandomnessFulfillmentRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconRandomnessFulfillmentRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconRandomnessFulfillmentRequested)
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

func (it *VRFBeaconRandomnessFulfillmentRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconRandomnessFulfillmentRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconRandomnessFulfillmentRequested struct {
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	GasAllowance           uint32
	GasPrice               *big.Int
	WeiPerUnitLink         *big.Int
	Arguments              []byte
	CostJuels              *big.Int
	NewSubBalance          *big.Int
	Raw                    types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts, requestID []*big.Int) (*VRFBeaconRandomnessFulfillmentRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "RandomnessFulfillmentRequested", requestIDRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconRandomnessFulfillmentRequestedIterator{contract: _VRFBeacon.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconRandomnessFulfillmentRequested, requestID []*big.Int) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "RandomnessFulfillmentRequested", requestIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconRandomnessFulfillmentRequested)
				if err := _VRFBeacon.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseRandomnessFulfillmentRequested(log types.Log) (*VRFBeaconRandomnessFulfillmentRequested, error) {
	event := new(VRFBeaconRandomnessFulfillmentRequested)
	if err := _VRFBeacon.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconRandomnessRedeemedIterator struct {
	Event *VRFBeaconRandomnessRedeemed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconRandomnessRedeemedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconRandomnessRedeemed)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconRandomnessRedeemed)
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

func (it *VRFBeaconRandomnessRedeemedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconRandomnessRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconRandomnessRedeemed struct {
	RequestID *big.Int
	Requester common.Address
	SubID     *big.Int
	Raw       types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterRandomnessRedeemed(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFBeaconRandomnessRedeemedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "RandomnessRedeemed", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconRandomnessRedeemedIterator{contract: _VRFBeacon.contract, event: "RandomnessRedeemed", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchRandomnessRedeemed(opts *bind.WatchOpts, sink chan<- *VRFBeaconRandomnessRedeemed, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "RandomnessRedeemed", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconRandomnessRedeemed)
				if err := _VRFBeacon.contract.UnpackLog(event, "RandomnessRedeemed", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseRandomnessRedeemed(log types.Log) (*VRFBeaconRandomnessRedeemed, error) {
	event := new(VRFBeaconRandomnessRedeemed)
	if err := _VRFBeacon.contract.UnpackLog(event, "RandomnessRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconRandomnessRequestedIterator struct {
	Event *VRFBeaconRandomnessRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconRandomnessRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconRandomnessRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconRandomnessRequested)
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

func (it *VRFBeaconRandomnessRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconRandomnessRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconRandomnessRequested struct {
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	CostJuels              *big.Int
	NewSubBalance          *big.Int
	Raw                    types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, requestID []*big.Int) (*VRFBeaconRandomnessRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "RandomnessRequested", requestIDRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconRandomnessRequestedIterator{contract: _VRFBeacon.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconRandomnessRequested, requestID []*big.Int) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "RandomnessRequested", requestIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconRandomnessRequested)
				if err := _VRFBeacon.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseRandomnessRequested(log types.Log) (*VRFBeaconRandomnessRequested, error) {
	event := new(VRFBeaconRandomnessRequested)
	if err := _VRFBeacon.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetBilling struct {
	MaximumGasPrice     uint64
	ReasonableGasPrice  uint64
	ObservationPayment  uint64
	TransmissionPayment uint64
	AccountingGas       *big.Int
}
type LatestConfigDetails struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}
type LatestConfigDigestAndEpoch struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}

func (_VRFBeacon *VRFBeacon) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VRFBeacon.abi.Events["BillingAccessControllerSet"].ID:
		return _VRFBeacon.ParseBillingAccessControllerSet(log)
	case _VRFBeacon.abi.Events["BillingSet"].ID:
		return _VRFBeacon.ParseBillingSet(log)
	case _VRFBeacon.abi.Events["ConfigSet"].ID:
		return _VRFBeacon.ParseConfigSet(log)
	case _VRFBeacon.abi.Events["NewTransmission"].ID:
		return _VRFBeacon.ParseNewTransmission(log)
	case _VRFBeacon.abi.Events["OraclePaid"].ID:
		return _VRFBeacon.ParseOraclePaid(log)
	case _VRFBeacon.abi.Events["OutputsServed"].ID:
		return _VRFBeacon.ParseOutputsServed(log)
	case _VRFBeacon.abi.Events["OwnershipTransferRequested"].ID:
		return _VRFBeacon.ParseOwnershipTransferRequested(log)
	case _VRFBeacon.abi.Events["OwnershipTransferred"].ID:
		return _VRFBeacon.ParseOwnershipTransferred(log)
	case _VRFBeacon.abi.Events["PayeeshipTransferRequested"].ID:
		return _VRFBeacon.ParsePayeeshipTransferRequested(log)
	case _VRFBeacon.abi.Events["PayeeshipTransferred"].ID:
		return _VRFBeacon.ParsePayeeshipTransferred(log)
	case _VRFBeacon.abi.Events["RandomWordsFulfilled"].ID:
		return _VRFBeacon.ParseRandomWordsFulfilled(log)
	case _VRFBeacon.abi.Events["RandomnessFulfillmentRequested"].ID:
		return _VRFBeacon.ParseRandomnessFulfillmentRequested(log)
	case _VRFBeacon.abi.Events["RandomnessRedeemed"].ID:
		return _VRFBeacon.ParseRandomnessRedeemed(log)
	case _VRFBeacon.abi.Events["RandomnessRequested"].ID:
		return _VRFBeacon.ParseRandomnessRequested(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VRFBeaconBillingAccessControllerSet) Topic() common.Hash {
	return common.HexToHash("0x793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912")
}

func (VRFBeaconBillingSet) Topic() common.Hash {
	return common.HexToHash("0x49275ddcdfc9c0519b3d094308c8bf675f06070a754ce90c152163cb6e66e8a0")
}

func (VRFBeaconConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (VRFBeaconNewTransmission) Topic() common.Hash {
	return common.HexToHash("0xfc3c7a7927e878a0fca37c904953c3c75cee3ca1d1640184a0ab1c65eec62743")
}

func (VRFBeaconOraclePaid) Topic() common.Hash {
	return common.HexToHash("0xd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c")
}

func (VRFBeaconOutputsServed) Topic() common.Hash {
	return common.HexToHash("0xf10ea936d00579b4c52035ee33bf46929646b3aa87554c565d8fb2c7aa549c44")
}

func (VRFBeaconOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VRFBeaconOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VRFBeaconPayeeshipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367")
}

func (VRFBeaconPayeeshipTransferred) Topic() common.Hash {
	return common.HexToHash("0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3")
}

func (VRFBeaconRandomWordsFulfilled) Topic() common.Hash {
	return common.HexToHash("0x8f79f730779e875ce76c428039cc2052b5b5918c2a55c598fab251c1198aec54")
}

func (VRFBeaconRandomnessFulfillmentRequested) Topic() common.Hash {
	return common.HexToHash("0x01872fb9c7d6d68af06a17347935e04412da302a377224c205e672c26e18c37f")
}

func (VRFBeaconRandomnessRedeemed) Topic() common.Hash {
	return common.HexToHash("0x16f3f633197fafab10a5df69e6f3f2f7f20092f08d8d47de0a91c0f4b96a1a25")
}

func (VRFBeaconRandomnessRequested) Topic() common.Hash {
	return common.HexToHash("0xb7933fba96b6b452eb44f99fdc08052a45dff82363d59abaff0456931c3d2459")
}

func (_VRFBeacon *VRFBeacon) Address() common.Address {
	return _VRFBeacon.address
}

type VRFBeaconInterface interface {
	NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error)

	GetBilling(opts *bind.CallOpts) (GetBilling,

		error)

	GetBillingAccessController(opts *bind.CallOpts) (common.Address, error)

	ICoordinator(opts *bind.CallOpts) (common.Address, error)

	ILink(opts *bind.CallOpts) (common.Address, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error)

	OwedPayment(opts *bind.CallOpts, transmitterAddress common.Address) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SKeyID(opts *bind.CallOpts) ([32]byte, error)

	SKeyProvider(opts *bind.CallOpts) (common.Address, error)

	SProvingKeyHash(opts *bind.CallOpts) ([32]byte, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AcceptPayeeship(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error)

	ExposeType(opts *bind.TransactOpts, arg0 VRFBeaconReportReport) (*types.Transaction, error)

	KeyGenerated(opts *bind.TransactOpts, kd KeyDataStructKeyData) (*types.Transaction, error)

	NewKeyRequested(opts *bind.TransactOpts) (*types.Transaction, error)

	SetBilling(opts *bind.TransactOpts, maximumGasPrice uint64, reasonableGasPrice uint64, observationPayment uint64, transmissionPayment uint64, accountingGas *big.Int) (*types.Transaction, error)

	SetBillingAccessController(opts *bind.TransactOpts, _billingAccessController common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	SetPayees(opts *bind.TransactOpts, transmitters []common.Address, payees []common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	TransferPayeeship(opts *bind.TransactOpts, transmitter common.Address, proposed common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	WithdrawFunds(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	WithdrawPayment(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error)

	FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*VRFBeaconBillingAccessControllerSetIterator, error)

	WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingAccessControllerSet) (event.Subscription, error)

	ParseBillingAccessControllerSet(log types.Log) (*VRFBeaconBillingAccessControllerSet, error)

	FilterBillingSet(opts *bind.FilterOpts) (*VRFBeaconBillingSetIterator, error)

	WatchBillingSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingSet) (event.Subscription, error)

	ParseBillingSet(log types.Log) (*VRFBeaconBillingSet, error)

	FilterConfigSet(opts *bind.FilterOpts) (*VRFBeaconConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*VRFBeaconConfigSet, error)

	FilterNewTransmission(opts *bind.FilterOpts, epochAndRound []*big.Int) (*VRFBeaconNewTransmissionIterator, error)

	WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *VRFBeaconNewTransmission, epochAndRound []*big.Int) (event.Subscription, error)

	ParseNewTransmission(log types.Log) (*VRFBeaconNewTransmission, error)

	FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*VRFBeaconOraclePaidIterator, error)

	WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *VRFBeaconOraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error)

	ParseOraclePaid(log types.Log) (*VRFBeaconOraclePaid, error)

	FilterOutputsServed(opts *bind.FilterOpts) (*VRFBeaconOutputsServedIterator, error)

	WatchOutputsServed(opts *bind.WatchOpts, sink chan<- *VRFBeaconOutputsServed) (event.Subscription, error)

	ParseOutputsServed(log types.Log) (*VRFBeaconOutputsServed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFBeaconOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VRFBeaconOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFBeaconOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFBeaconOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VRFBeaconOwnershipTransferred, error)

	FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*VRFBeaconPayeeshipTransferRequestedIterator, error)

	WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error)

	ParsePayeeshipTransferRequested(log types.Log) (*VRFBeaconPayeeshipTransferRequested, error)

	FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*VRFBeaconPayeeshipTransferredIterator, error)

	WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *VRFBeaconPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error)

	ParsePayeeshipTransferred(log types.Log) (*VRFBeaconPayeeshipTransferred, error)

	FilterRandomWordsFulfilled(opts *bind.FilterOpts) (*VRFBeaconRandomWordsFulfilledIterator, error)

	WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFBeaconRandomWordsFulfilled) (event.Subscription, error)

	ParseRandomWordsFulfilled(log types.Log) (*VRFBeaconRandomWordsFulfilled, error)

	FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts, requestID []*big.Int) (*VRFBeaconRandomnessFulfillmentRequestedIterator, error)

	WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconRandomnessFulfillmentRequested, requestID []*big.Int) (event.Subscription, error)

	ParseRandomnessFulfillmentRequested(log types.Log) (*VRFBeaconRandomnessFulfillmentRequested, error)

	FilterRandomnessRedeemed(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFBeaconRandomnessRedeemedIterator, error)

	WatchRandomnessRedeemed(opts *bind.WatchOpts, sink chan<- *VRFBeaconRandomnessRedeemed, requestID []*big.Int, requester []common.Address) (event.Subscription, error)

	ParseRandomnessRedeemed(log types.Log) (*VRFBeaconRandomnessRedeemed, error)

	FilterRandomnessRequested(opts *bind.FilterOpts, requestID []*big.Int) (*VRFBeaconRandomnessRequestedIterator, error)

	WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconRandomnessRequested, requestID []*big.Int) (event.Subscription, error)

	ParseRandomnessRequested(log types.Log) (*VRFBeaconRandomnessRequested, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
