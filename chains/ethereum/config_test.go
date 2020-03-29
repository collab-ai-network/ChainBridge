// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ChainSafe/ChainBridge/core"
	"github.com/ethereum/go-ethereum/common"
)

func TestParseChainConfig(t *testing.T) {
	input := core.ChainConfig{
		Name:         "chain",
		Id:           1,
		Endpoint:     "endpoint",
		From:         "0x0",
		KeystorePath: "./keys",
		Insecure:     false,
		Opts: map[string]string{
			"contract": "0x1234",
			"gasLimit": "10",
			"gasPrice": "20",
			"http":     "true",
		},
	}

	out, err := parseChainConfig(&input)

	if err != nil {
		t.Fatal(err)
	}

	expected := Config{
		name:         "chain",
		id:           1,
		endpoint:     "endpoint",
		from:         "0x0",
		keystorePath: "./keys",
		contract:     common.HexToAddress("0x1234"),
		gasLimit:     big.NewInt(10),
		gasPrice:     big.NewInt(20),
		http:         true,
	}

	if !reflect.DeepEqual(&expected, out) {
		t.Fatalf("Output not expected.\n\tExpected: %#v\n\tGot: %#v\n", &expected, out)
	}
}

func TestRequiredOpts(t *testing.T) {
	// No opts provided
	input := core.ChainConfig{
		Id:           0,
		Endpoint:     "endpoint",
		From:         "0x0",
		KeystorePath: "./keys",
		Insecure:     false,
		Opts:         map[string]string{},
	}

	_, err := parseChainConfig(&input)

	if err == nil {
		t.Error("config missing chainId field but no error reported")
	}

	// Empty contract provided
	input = core.ChainConfig{
		Id:           0,
		Endpoint:     "endpoint",
		From:         "0x0",
		KeystorePath: "./keys",
		Insecure:     false,
		Opts:         map[string]string{"contract": ""},
	}

	_, err2 := parseChainConfig(&input)

	if err2 == nil {
		t.Error("config missing chainId field but no error reported")
	}

}