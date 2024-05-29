// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

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

// EventUtilsAddressArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsAddressArrayKeyValue struct {
	Key   string
	Value []common.Address
}

// EventUtilsAddressItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsAddressItems struct {
	Items      []EventUtilsAddressKeyValue
	ArrayItems []EventUtilsAddressArrayKeyValue
}

// EventUtilsAddressKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsAddressKeyValue struct {
	Key   string
	Value common.Address
}

// EventUtilsBoolArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBoolArrayKeyValue struct {
	Key   string
	Value []bool
}

// EventUtilsBoolItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBoolItems struct {
	Items      []EventUtilsBoolKeyValue
	ArrayItems []EventUtilsBoolArrayKeyValue
}

// EventUtilsBoolKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBoolKeyValue struct {
	Key   string
	Value bool
}

// EventUtilsBytes32ArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytes32ArrayKeyValue struct {
	Key   string
	Value [][32]byte
}

// EventUtilsBytes32Items is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytes32Items struct {
	Items      []EventUtilsBytes32KeyValue
	ArrayItems []EventUtilsBytes32ArrayKeyValue
}

// EventUtilsBytes32KeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytes32KeyValue struct {
	Key   string
	Value [32]byte
}

// EventUtilsBytesArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytesArrayKeyValue struct {
	Key   string
	Value [][]byte
}

// EventUtilsBytesItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytesItems struct {
	Items      []EventUtilsBytesKeyValue
	ArrayItems []EventUtilsBytesArrayKeyValue
}

// EventUtilsBytesKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytesKeyValue struct {
	Key   string
	Value []byte
}

// EventUtilsEventLogData is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsEventLogData struct {
	AddressItems EventUtilsAddressItems
	UintItems    EventUtilsUintItems
	IntItems     EventUtilsIntItems
	BoolItems    EventUtilsBoolItems
	Bytes32Items EventUtilsBytes32Items
	BytesItems   EventUtilsBytesItems
	StringItems  EventUtilsStringItems
}

// EventUtilsIntArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsIntArrayKeyValue struct {
	Key   string
	Value []*big.Int
}

// EventUtilsIntItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsIntItems struct {
	Items      []EventUtilsIntKeyValue
	ArrayItems []EventUtilsIntArrayKeyValue
}

// EventUtilsIntKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsIntKeyValue struct {
	Key   string
	Value *big.Int
}

// EventUtilsStringArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsStringArrayKeyValue struct {
	Key   string
	Value []string
}

// EventUtilsStringItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsStringItems struct {
	Items      []EventUtilsStringKeyValue
	ArrayItems []EventUtilsStringArrayKeyValue
}

// EventUtilsStringKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsStringKeyValue struct {
	Key   string
	Value string
}

// EventUtilsUintArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsUintArrayKeyValue struct {
	Key   string
	Value []*big.Int
}

// EventUtilsUintItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsUintItems struct {
	Items      []EventUtilsUintKeyValue
	ArrayItems []EventUtilsUintArrayKeyValue
}

// EventUtilsUintKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsUintKeyValue struct {
	Key   string
	Value *big.Int
}

// AbisMetaData contains all meta data concerning the Abis contract.
var AbisMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRoleStore\",\"name\":\"_roleStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"eventNameHash\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.AddressItems\",\"name\":\"addressItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structEventUtils.UintKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.UintItems\",\"name\":\"uintItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"internalType\":\"structEventUtils.IntKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.IntItems\",\"name\":\"intItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BoolItems\",\"name\":\"boolItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.Bytes32Items\",\"name\":\"bytes32Items\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"value\",\"type\":\"bytes[]\"}],\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BytesItems\",\"name\":\"bytesItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structEventUtils.StringKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.StringItems\",\"name\":\"stringItems\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structEventUtils.EventLogData\",\"name\":\"eventData\",\"type\":\"tuple\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"eventNameHash\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.AddressItems\",\"name\":\"addressItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structEventUtils.UintKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.UintItems\",\"name\":\"uintItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"internalType\":\"structEventUtils.IntKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.IntItems\",\"name\":\"intItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BoolItems\",\"name\":\"boolItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.Bytes32Items\",\"name\":\"bytes32Items\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"value\",\"type\":\"bytes[]\"}],\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BytesItems\",\"name\":\"bytesItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structEventUtils.StringKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.StringItems\",\"name\":\"stringItems\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structEventUtils.EventLogData\",\"name\":\"eventData\",\"type\":\"tuple\"}],\"name\":\"EventLog1\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"eventNameHash\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"topic2\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.AddressItems\",\"name\":\"addressItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structEventUtils.UintKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.UintItems\",\"name\":\"uintItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"internalType\":\"structEventUtils.IntKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.IntItems\",\"name\":\"intItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BoolItems\",\"name\":\"boolItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.Bytes32Items\",\"name\":\"bytes32Items\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"value\",\"type\":\"bytes[]\"}],\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BytesItems\",\"name\":\"bytesItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structEventUtils.StringKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.StringItems\",\"name\":\"stringItems\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structEventUtils.EventLogData\",\"name\":\"eventData\",\"type\":\"tuple\"}],\"name\":\"EventLog2\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitDataLog1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitDataLog2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic3\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitDataLog3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic3\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic4\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitDataLog4\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.AddressItems\",\"name\":\"addressItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structEventUtils.UintKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.UintItems\",\"name\":\"uintItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"internalType\":\"structEventUtils.IntKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.IntItems\",\"name\":\"intItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BoolItems\",\"name\":\"boolItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.Bytes32Items\",\"name\":\"bytes32Items\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"value\",\"type\":\"bytes[]\"}],\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BytesItems\",\"name\":\"bytesItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structEventUtils.StringKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.StringItems\",\"name\":\"stringItems\",\"type\":\"tuple\"}],\"internalType\":\"structEventUtils.EventLogData\",\"name\":\"eventData\",\"type\":\"tuple\"}],\"name\":\"emitEventLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.AddressItems\",\"name\":\"addressItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structEventUtils.UintKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.UintItems\",\"name\":\"uintItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"internalType\":\"structEventUtils.IntKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.IntItems\",\"name\":\"intItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BoolItems\",\"name\":\"boolItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.Bytes32Items\",\"name\":\"bytes32Items\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"value\",\"type\":\"bytes[]\"}],\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BytesItems\",\"name\":\"bytesItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structEventUtils.StringKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.StringItems\",\"name\":\"stringItems\",\"type\":\"tuple\"}],\"internalType\":\"structEventUtils.EventLogData\",\"name\":\"eventData\",\"type\":\"tuple\"}],\"name\":\"emitEventLog1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic2\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.AddressItems\",\"name\":\"addressItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structEventUtils.UintKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.UintItems\",\"name\":\"uintItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"internalType\":\"structEventUtils.IntKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.IntItems\",\"name\":\"intItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BoolItems\",\"name\":\"boolItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.Bytes32Items\",\"name\":\"bytes32Items\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"value\",\"type\":\"bytes[]\"}],\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BytesItems\",\"name\":\"bytesItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structEventUtils.StringKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.StringItems\",\"name\":\"stringItems\",\"type\":\"tuple\"}],\"internalType\":\"structEventUtils.EventLogData\",\"name\":\"eventData\",\"type\":\"tuple\"}],\"name\":\"emitEventLog2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roleStore\",\"outputs\":[{\"internalType\":\"contractRoleStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AbisABI is the input ABI used to generate the binding from.
// Deprecated: Use AbisMetaData.ABI instead.
var AbisABI = AbisMetaData.ABI

// Abis is an auto generated Go binding around an Ethereum contract.
type Abis struct {
	AbisCaller     // Read-only binding to the contract
	AbisTransactor // Write-only binding to the contract
	AbisFilterer   // Log filterer for contract events
}

// AbisCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbisCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbisTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbisTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbisFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbisFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbisSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbisSession struct {
	Contract     *Abis             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbisCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbisCallerSession struct {
	Contract *AbisCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbisTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbisTransactorSession struct {
	Contract     *AbisTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbisRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbisRaw struct {
	Contract *Abis // Generic contract binding to access the raw methods on
}

// AbisCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbisCallerRaw struct {
	Contract *AbisCaller // Generic read-only contract binding to access the raw methods on
}

// AbisTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbisTransactorRaw struct {
	Contract *AbisTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbis creates a new instance of Abis, bound to a specific deployed contract.
func NewAbis(address common.Address, backend bind.ContractBackend) (*Abis, error) {
	contract, err := bindAbis(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abis{AbisCaller: AbisCaller{contract: contract}, AbisTransactor: AbisTransactor{contract: contract}, AbisFilterer: AbisFilterer{contract: contract}}, nil
}

// NewAbisCaller creates a new read-only instance of Abis, bound to a specific deployed contract.
func NewAbisCaller(address common.Address, caller bind.ContractCaller) (*AbisCaller, error) {
	contract, err := bindAbis(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbisCaller{contract: contract}, nil
}

// NewAbisTransactor creates a new write-only instance of Abis, bound to a specific deployed contract.
func NewAbisTransactor(address common.Address, transactor bind.ContractTransactor) (*AbisTransactor, error) {
	contract, err := bindAbis(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbisTransactor{contract: contract}, nil
}

// NewAbisFilterer creates a new log filterer instance of Abis, bound to a specific deployed contract.
func NewAbisFilterer(address common.Address, filterer bind.ContractFilterer) (*AbisFilterer, error) {
	contract, err := bindAbis(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbisFilterer{contract: contract}, nil
}

// bindAbis binds a generic wrapper to an already deployed contract.
func bindAbis(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbisMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abis *AbisRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abis.Contract.AbisCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abis *AbisRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abis.Contract.AbisTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abis *AbisRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abis.Contract.AbisTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abis *AbisCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abis.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abis *AbisTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abis.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abis *AbisTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abis.Contract.contract.Transact(opts, method, params...)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Abis *AbisCaller) RoleStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abis.contract.Call(opts, &out, "roleStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Abis *AbisSession) RoleStore() (common.Address, error) {
	return _Abis.Contract.RoleStore(&_Abis.CallOpts)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Abis *AbisCallerSession) RoleStore() (common.Address, error) {
	return _Abis.Contract.RoleStore(&_Abis.CallOpts)
}

// EmitDataLog1 is a paid mutator transaction binding the contract method 0xf9d5c0ea.
//
// Solidity: function emitDataLog1(bytes32 topic1, bytes data) returns()
func (_Abis *AbisTransactor) EmitDataLog1(opts *bind.TransactOpts, topic1 [32]byte, data []byte) (*types.Transaction, error) {
	return _Abis.contract.Transact(opts, "emitDataLog1", topic1, data)
}

// EmitDataLog1 is a paid mutator transaction binding the contract method 0xf9d5c0ea.
//
// Solidity: function emitDataLog1(bytes32 topic1, bytes data) returns()
func (_Abis *AbisSession) EmitDataLog1(topic1 [32]byte, data []byte) (*types.Transaction, error) {
	return _Abis.Contract.EmitDataLog1(&_Abis.TransactOpts, topic1, data)
}

// EmitDataLog1 is a paid mutator transaction binding the contract method 0xf9d5c0ea.
//
// Solidity: function emitDataLog1(bytes32 topic1, bytes data) returns()
func (_Abis *AbisTransactorSession) EmitDataLog1(topic1 [32]byte, data []byte) (*types.Transaction, error) {
	return _Abis.Contract.EmitDataLog1(&_Abis.TransactOpts, topic1, data)
}

// EmitDataLog2 is a paid mutator transaction binding the contract method 0xdda0db32.
//
// Solidity: function emitDataLog2(bytes32 topic1, bytes32 topic2, bytes data) returns()
func (_Abis *AbisTransactor) EmitDataLog2(opts *bind.TransactOpts, topic1 [32]byte, topic2 [32]byte, data []byte) (*types.Transaction, error) {
	return _Abis.contract.Transact(opts, "emitDataLog2", topic1, topic2, data)
}

// EmitDataLog2 is a paid mutator transaction binding the contract method 0xdda0db32.
//
// Solidity: function emitDataLog2(bytes32 topic1, bytes32 topic2, bytes data) returns()
func (_Abis *AbisSession) EmitDataLog2(topic1 [32]byte, topic2 [32]byte, data []byte) (*types.Transaction, error) {
	return _Abis.Contract.EmitDataLog2(&_Abis.TransactOpts, topic1, topic2, data)
}

// EmitDataLog2 is a paid mutator transaction binding the contract method 0xdda0db32.
//
// Solidity: function emitDataLog2(bytes32 topic1, bytes32 topic2, bytes data) returns()
func (_Abis *AbisTransactorSession) EmitDataLog2(topic1 [32]byte, topic2 [32]byte, data []byte) (*types.Transaction, error) {
	return _Abis.Contract.EmitDataLog2(&_Abis.TransactOpts, topic1, topic2, data)
}

// EmitDataLog3 is a paid mutator transaction binding the contract method 0xb3ac1c38.
//
// Solidity: function emitDataLog3(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes data) returns()
func (_Abis *AbisTransactor) EmitDataLog3(opts *bind.TransactOpts, topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, data []byte) (*types.Transaction, error) {
	return _Abis.contract.Transact(opts, "emitDataLog3", topic1, topic2, topic3, data)
}

// EmitDataLog3 is a paid mutator transaction binding the contract method 0xb3ac1c38.
//
// Solidity: function emitDataLog3(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes data) returns()
func (_Abis *AbisSession) EmitDataLog3(topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, data []byte) (*types.Transaction, error) {
	return _Abis.Contract.EmitDataLog3(&_Abis.TransactOpts, topic1, topic2, topic3, data)
}

// EmitDataLog3 is a paid mutator transaction binding the contract method 0xb3ac1c38.
//
// Solidity: function emitDataLog3(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes data) returns()
func (_Abis *AbisTransactorSession) EmitDataLog3(topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, data []byte) (*types.Transaction, error) {
	return _Abis.Contract.EmitDataLog3(&_Abis.TransactOpts, topic1, topic2, topic3, data)
}

// EmitDataLog4 is a paid mutator transaction binding the contract method 0xee288ce8.
//
// Solidity: function emitDataLog4(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4, bytes data) returns()
func (_Abis *AbisTransactor) EmitDataLog4(opts *bind.TransactOpts, topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte, data []byte) (*types.Transaction, error) {
	return _Abis.contract.Transact(opts, "emitDataLog4", topic1, topic2, topic3, topic4, data)
}

// EmitDataLog4 is a paid mutator transaction binding the contract method 0xee288ce8.
//
// Solidity: function emitDataLog4(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4, bytes data) returns()
func (_Abis *AbisSession) EmitDataLog4(topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte, data []byte) (*types.Transaction, error) {
	return _Abis.Contract.EmitDataLog4(&_Abis.TransactOpts, topic1, topic2, topic3, topic4, data)
}

// EmitDataLog4 is a paid mutator transaction binding the contract method 0xee288ce8.
//
// Solidity: function emitDataLog4(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4, bytes data) returns()
func (_Abis *AbisTransactorSession) EmitDataLog4(topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte, data []byte) (*types.Transaction, error) {
	return _Abis.Contract.EmitDataLog4(&_Abis.TransactOpts, topic1, topic2, topic3, topic4, data)
}

// EmitEventLog is a paid mutator transaction binding the contract method 0x906c49f6.
//
// Solidity: function emitEventLog(string eventName, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Abis *AbisTransactor) EmitEventLog(opts *bind.TransactOpts, eventName string, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Abis.contract.Transact(opts, "emitEventLog", eventName, eventData)
}

// EmitEventLog is a paid mutator transaction binding the contract method 0x906c49f6.
//
// Solidity: function emitEventLog(string eventName, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Abis *AbisSession) EmitEventLog(eventName string, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Abis.Contract.EmitEventLog(&_Abis.TransactOpts, eventName, eventData)
}

// EmitEventLog is a paid mutator transaction binding the contract method 0x906c49f6.
//
// Solidity: function emitEventLog(string eventName, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Abis *AbisTransactorSession) EmitEventLog(eventName string, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Abis.Contract.EmitEventLog(&_Abis.TransactOpts, eventName, eventData)
}

// EmitEventLog1 is a paid mutator transaction binding the contract method 0x24de01e4.
//
// Solidity: function emitEventLog1(string eventName, bytes32 topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Abis *AbisTransactor) EmitEventLog1(opts *bind.TransactOpts, eventName string, topic1 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Abis.contract.Transact(opts, "emitEventLog1", eventName, topic1, eventData)
}

// EmitEventLog1 is a paid mutator transaction binding the contract method 0x24de01e4.
//
// Solidity: function emitEventLog1(string eventName, bytes32 topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Abis *AbisSession) EmitEventLog1(eventName string, topic1 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Abis.Contract.EmitEventLog1(&_Abis.TransactOpts, eventName, topic1, eventData)
}

// EmitEventLog1 is a paid mutator transaction binding the contract method 0x24de01e4.
//
// Solidity: function emitEventLog1(string eventName, bytes32 topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Abis *AbisTransactorSession) EmitEventLog1(eventName string, topic1 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Abis.Contract.EmitEventLog1(&_Abis.TransactOpts, eventName, topic1, eventData)
}

// EmitEventLog2 is a paid mutator transaction binding the contract method 0x63d16363.
//
// Solidity: function emitEventLog2(string eventName, bytes32 topic1, bytes32 topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Abis *AbisTransactor) EmitEventLog2(opts *bind.TransactOpts, eventName string, topic1 [32]byte, topic2 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Abis.contract.Transact(opts, "emitEventLog2", eventName, topic1, topic2, eventData)
}

// EmitEventLog2 is a paid mutator transaction binding the contract method 0x63d16363.
//
// Solidity: function emitEventLog2(string eventName, bytes32 topic1, bytes32 topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Abis *AbisSession) EmitEventLog2(eventName string, topic1 [32]byte, topic2 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Abis.Contract.EmitEventLog2(&_Abis.TransactOpts, eventName, topic1, topic2, eventData)
}

// EmitEventLog2 is a paid mutator transaction binding the contract method 0x63d16363.
//
// Solidity: function emitEventLog2(string eventName, bytes32 topic1, bytes32 topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Abis *AbisTransactorSession) EmitEventLog2(eventName string, topic1 [32]byte, topic2 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Abis.Contract.EmitEventLog2(&_Abis.TransactOpts, eventName, topic1, topic2, eventData)
}

// AbisEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Abis contract.
type AbisEventLogIterator struct {
	Event *AbisEventLog // Event containing the contract specifics and raw log

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
func (it *AbisEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbisEventLog)
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
		it.Event = new(AbisEventLog)
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
func (it *AbisEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbisEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbisEventLog represents a EventLog event raised by the Abis contract.
type AbisEventLog struct {
	MsgSender     common.Address
	EventName     string
	EventNameHash common.Hash
	EventData     EventUtilsEventLogData
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0x7e3bde2ba7aca4a8499608ca57f3b0c1c1c93ace63ffd3741a9fab204146fc9a.
//
// Solidity: event EventLog(address msgSender, string eventName, string indexed eventNameHash, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Abis *AbisFilterer) FilterEventLog(opts *bind.FilterOpts, eventNameHash []string) (*AbisEventLogIterator, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}

	logs, sub, err := _Abis.contract.FilterLogs(opts, "EventLog", eventNameHashRule)
	if err != nil {
		return nil, err
	}
	return &AbisEventLogIterator{contract: _Abis.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0x7e3bde2ba7aca4a8499608ca57f3b0c1c1c93ace63ffd3741a9fab204146fc9a.
//
// Solidity: event EventLog(address msgSender, string eventName, string indexed eventNameHash, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Abis *AbisFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *AbisEventLog, eventNameHash []string) (event.Subscription, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}

	logs, sub, err := _Abis.contract.WatchLogs(opts, "EventLog", eventNameHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbisEventLog)
				if err := _Abis.contract.UnpackLog(event, "EventLog", log); err != nil {
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

// ParseEventLog is a log parse operation binding the contract event 0x7e3bde2ba7aca4a8499608ca57f3b0c1c1c93ace63ffd3741a9fab204146fc9a.
//
// Solidity: event EventLog(address msgSender, string eventName, string indexed eventNameHash, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Abis *AbisFilterer) ParseEventLog(log types.Log) (*AbisEventLog, error) {
	event := new(AbisEventLog)
	if err := _Abis.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbisEventLog1Iterator is returned from FilterEventLog1 and is used to iterate over the raw logs and unpacked data for EventLog1 events raised by the Abis contract.
type AbisEventLog1Iterator struct {
	Event *AbisEventLog1 // Event containing the contract specifics and raw log

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
func (it *AbisEventLog1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbisEventLog1)
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
		it.Event = new(AbisEventLog1)
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
func (it *AbisEventLog1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbisEventLog1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbisEventLog1 represents a EventLog1 event raised by the Abis contract.
type AbisEventLog1 struct {
	MsgSender     common.Address
	EventName     string
	EventNameHash common.Hash
	Topic1        [32]byte
	EventData     EventUtilsEventLogData
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventLog1 is a free log retrieval operation binding the contract event 0x137a44067c8961cd7e1d876f4754a5a3a75989b4552f1843fc69c3b372def160.
//
// Solidity: event EventLog1(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Abis *AbisFilterer) FilterEventLog1(opts *bind.FilterOpts, eventNameHash []string, topic1 [][32]byte) (*AbisEventLog1Iterator, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}
	var topic1Rule []interface{}
	for _, topic1Item := range topic1 {
		topic1Rule = append(topic1Rule, topic1Item)
	}

	logs, sub, err := _Abis.contract.FilterLogs(opts, "EventLog1", eventNameHashRule, topic1Rule)
	if err != nil {
		return nil, err
	}
	return &AbisEventLog1Iterator{contract: _Abis.contract, event: "EventLog1", logs: logs, sub: sub}, nil
}

// WatchEventLog1 is a free log subscription operation binding the contract event 0x137a44067c8961cd7e1d876f4754a5a3a75989b4552f1843fc69c3b372def160.
//
// Solidity: event EventLog1(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Abis *AbisFilterer) WatchEventLog1(opts *bind.WatchOpts, sink chan<- *AbisEventLog1, eventNameHash []string, topic1 [][32]byte) (event.Subscription, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}
	var topic1Rule []interface{}
	for _, topic1Item := range topic1 {
		topic1Rule = append(topic1Rule, topic1Item)
	}

	logs, sub, err := _Abis.contract.WatchLogs(opts, "EventLog1", eventNameHashRule, topic1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbisEventLog1)
				if err := _Abis.contract.UnpackLog(event, "EventLog1", log); err != nil {
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

// ParseEventLog1 is a log parse operation binding the contract event 0x137a44067c8961cd7e1d876f4754a5a3a75989b4552f1843fc69c3b372def160.
//
// Solidity: event EventLog1(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Abis *AbisFilterer) ParseEventLog1(log types.Log) (*AbisEventLog1, error) {
	event := new(AbisEventLog1)
	if err := _Abis.contract.UnpackLog(event, "EventLog1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbisEventLog2Iterator is returned from FilterEventLog2 and is used to iterate over the raw logs and unpacked data for EventLog2 events raised by the Abis contract.
type AbisEventLog2Iterator struct {
	Event *AbisEventLog2 // Event containing the contract specifics and raw log

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
func (it *AbisEventLog2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbisEventLog2)
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
		it.Event = new(AbisEventLog2)
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
func (it *AbisEventLog2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbisEventLog2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbisEventLog2 represents a EventLog2 event raised by the Abis contract.
type AbisEventLog2 struct {
	MsgSender     common.Address
	EventName     string
	EventNameHash common.Hash
	Topic1        [32]byte
	Topic2        [32]byte
	EventData     EventUtilsEventLogData
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventLog2 is a free log retrieval operation binding the contract event 0x468a25a7ba624ceea6e540ad6f49171b52495b648417ae91bca21676d8a24dc5.
//
// Solidity: event EventLog2(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, bytes32 indexed topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Abis *AbisFilterer) FilterEventLog2(opts *bind.FilterOpts, eventNameHash []string, topic1 [][32]byte, topic2 [][32]byte) (*AbisEventLog2Iterator, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}
	var topic1Rule []interface{}
	for _, topic1Item := range topic1 {
		topic1Rule = append(topic1Rule, topic1Item)
	}
	var topic2Rule []interface{}
	for _, topic2Item := range topic2 {
		topic2Rule = append(topic2Rule, topic2Item)
	}

	logs, sub, err := _Abis.contract.FilterLogs(opts, "EventLog2", eventNameHashRule, topic1Rule, topic2Rule)
	if err != nil {
		return nil, err
	}
	return &AbisEventLog2Iterator{contract: _Abis.contract, event: "EventLog2", logs: logs, sub: sub}, nil
}

// WatchEventLog2 is a free log subscription operation binding the contract event 0x468a25a7ba624ceea6e540ad6f49171b52495b648417ae91bca21676d8a24dc5.
//
// Solidity: event EventLog2(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, bytes32 indexed topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Abis *AbisFilterer) WatchEventLog2(opts *bind.WatchOpts, sink chan<- *AbisEventLog2, eventNameHash []string, topic1 [][32]byte, topic2 [][32]byte) (event.Subscription, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}
	var topic1Rule []interface{}
	for _, topic1Item := range topic1 {
		topic1Rule = append(topic1Rule, topic1Item)
	}
	var topic2Rule []interface{}
	for _, topic2Item := range topic2 {
		topic2Rule = append(topic2Rule, topic2Item)
	}

	logs, sub, err := _Abis.contract.WatchLogs(opts, "EventLog2", eventNameHashRule, topic1Rule, topic2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbisEventLog2)
				if err := _Abis.contract.UnpackLog(event, "EventLog2", log); err != nil {
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

// ParseEventLog2 is a log parse operation binding the contract event 0x468a25a7ba624ceea6e540ad6f49171b52495b648417ae91bca21676d8a24dc5.
//
// Solidity: event EventLog2(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, bytes32 indexed topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Abis *AbisFilterer) ParseEventLog2(log types.Log) (*AbisEventLog2, error) {
	event := new(AbisEventLog2)
	if err := _Abis.contract.UnpackLog(event, "EventLog2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
