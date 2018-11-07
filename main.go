package main

import (
	"go-bito4ik-multisig/multisig"

	"os"

	"gopkg.in/alecthomas/kingpin.v1"
)

// Kingpin configurations for command-line subcommands and their respective flags.
var (
	app = kingpin.New("go-bitcoin-multisig", "A Bitcoin multisig builder built in Go")

	cmdAddress           = app.Command("address", "Generate a multisig P2SH address with M-of-N requirements and set of public keys.")
	cmdCoin              = cmdAddress.Flag("coin-name", "Coin Name").Required().String()
	cmdAddressM          = cmdAddress.Flag("m", "M, the minimum number of keys needed to spend Bitcoin in M-of-N multisig transaction.").Required().Int()
	cmdAddressN          = cmdAddress.Flag("n", "N, the total number of possible keys that can be used to spend Bitcoin in M-of-N multisig transaction.").Required().Int()
	cmdAddressPublicKeys = cmdAddress.Flag("public-keys", "Comma separated list of private keys to sign with. Whitespace is stripped and quotes may be placed around keys. Eg. key1,key2,\"key3\"").PlaceHolder("PUBLIC-KEYS(Comma separated)").Required().String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case cmdAddress.FullCommand():
		multisig.OutputAddress(*cmdCoin, *cmdAddressM, *cmdAddressN, *cmdAddressPublicKeys)
	}
}
