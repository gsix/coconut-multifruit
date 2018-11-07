// Package btcutils contains a number of useful Bitcoin related functions originally used in the go-bitcoin-multisig
// project but useful in any general Bitcoin project.
package btcutils

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"

	"golang.org/x/crypto/ripemd160"
)

// Hash160 performs the same operations as OP_HASH160 in Bitcoin Script
// It hashes the given data first with SHA256, then RIPEMD160
func Hash160(data []byte) ([]byte, error) {
	//Does identical function to Script OP_HASH160. Hash once with SHA-256, then RIPEMD-160
	if data == nil {
		return nil, errors.New("Empty bytes cannot be hashed")
	}
	shaHash := sha256.New()
	shaHash.Write(data)
	hash := shaHash.Sum(nil) //SHA256 first
	ripemd160Hash := ripemd160.New()
	ripemd160Hash.Write(hash)
	hash = ripemd160Hash.Sum(nil) //RIPEMD160 second

	return hash, nil
}

// NewMOfNRedeemScript creates a M-of-N Multisig redeem script given m, n and n public keys
func NewMOfNRedeemScript(m int, n int, publicKeys [][]byte) ([]byte, error) {
	//Check we have valid numbers for M and N
	if n < 1 || n > 7 {
		return nil, errors.New("N must be between 1 and 7 (inclusive) for valid, standard P2SH multisig transaction as per Bitcoin protocol.")
	}
	if m < 1 || m > n {
		return nil, errors.New("M must be between 1 and N (inclusive).")
	}
	//Check we have N public keys as necessary.
	if len(publicKeys) != n {
		return nil, errors.New(fmt.Sprintf("Need exactly %d public keys to create P2SH address for %d-of-%d multisig transaction. Only %d keys provided.", n, m, n, len(publicKeys)))
	}
	//Get OP Code for m and n.
	//81 is OP_1, 82 is OP_2 etc.
	//80 is not a valid OP_Code, so we floor at 81
	mOPCode := OP_1 + (m - 1)
	nOPCode := OP_1 + (n - 1)
	//Multisig redeemScript format:
	//<OP_m> <A pubkey> <B pubkey> <C pubkey>... <OP_n> OP_CHECKMULTISIG
	var redeemScript bytes.Buffer
	redeemScript.WriteByte(byte(mOPCode)) //m
	for _, publicKey := range publicKeys {
		//err := CheckPublicKeyIsValid(publicKey)
		//if err != nil {
		//	return nil, err
		//}
		redeemScript.WriteByte(byte(len(publicKey))) //PUSH
		redeemScript.Write(publicKey)                //<pubkey>
	}
	redeemScript.WriteByte(byte(nOPCode)) //n
	redeemScript.WriteByte(byte(OP_CHECKMULTISIG))
	return redeemScript.Bytes(), nil
}
