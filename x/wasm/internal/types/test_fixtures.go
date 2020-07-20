package types

import (
	"bytes"
	"crypto/sha256"

	"github.com/tendermint/tendermint/libs/rand"
)

func GenesisFixture(mutators ...func(*GenesisState)) GenesisState {
	const (
		numCodes     = 2
		numContracts = 2
		numSequences = 2
	)

	fixture := GenesisState{
		Codes:     make([]Code, numCodes),
		Contracts: make([]Contract, numContracts),
		Sequences: make([]Sequence, numSequences),
	}
	for i := 0; i < numCodes; i++ {
		fixture.Codes[i] = CodeFixture()
	}
	for i := 0; i < numContracts; i++ {
		fixture.Contracts[i] = ContractFixture()
	}
	for i := 0; i < numSequences; i++ {
		fixture.Sequences[i] = Sequence{
			IDKey: rand.Bytes(5),
			Value: uint64(i),
		}
	}
	for _, m := range mutators {
		m(&fixture)
	}
	return fixture
}

func CodeFixture(mutators ...func(*Code)) Code {
	wasmCode := rand.Bytes(100)
	codeHash := sha256.Sum256(wasmCode)
	anyAddress := make([]byte, 20)

	fixture := Code{
		CodeID: 1,
		CodeInfo: CodeInfo{
			CodeHash: codeHash[:],
			Creator:  anyAddress,
		},
		CodesBytes: wasmCode,
	}

	for _, m := range mutators {
		m(&fixture)
	}
	return fixture
}

func CodeInfoFixture(mutators ...func(*CodeInfo)) CodeInfo {
	wasmCode := bytes.Repeat([]byte{0x1}, 10)
	codeHash := sha256.Sum256(wasmCode)
	anyAddress := make([]byte, 20)
	fixture := CodeInfo{
		CodeHash: codeHash[:],
		Creator:  anyAddress,
		Source:   "https://example.com",
		Builder:  "my/builder:tag",
	}
	for _, m := range mutators {
		m(&fixture)
	}
	return fixture
}

func ContractFixture(mutators ...func(*Contract)) Contract {
	anyAddress := make([]byte, 20)
	fixture := Contract{
		ContractAddress: anyAddress,
		ContractInfo:    ContractInfoFixture(),
		ContractState:   []Model{{Key: []byte("anyKey"), Value: []byte("anyValue")}},
	}

	for _, m := range mutators {
		m(&fixture)
	}
	return fixture
}

func ContractInfoFixture(mutators ...func(*ContractInfo)) ContractInfo {
	anyAddress := make([]byte, 20)
	fixture := ContractInfo{
		CodeID:  1,
		Creator: anyAddress,
		Label:   "any",
		Created: &AbsoluteTxPosition{BlockHeight: 1, TxIndex: 1},
	}

	for _, m := range mutators {
		m(&fixture)
	}
	return fixture
}
