package multisig

import (
	"testing"

	"go-bito4ik-multisig/testutils"
)

func TestGenerateAddress(t *testing.T) {
	{
		//Compress public key! Bitcoin 2-of-2
		testM := 2
		testN := 2
		coinName := "Bitcoin"
		testPublicKeys := "0367a032d775cab911829c5e4a8d5acc7cbdb1e0500865f7c6d72713ad53c04507,035b21ca62439b8087b0cd657039b5a1fd558747dbb7f69e8498ac10057f16a84c"
		testAddress := "34ggn1p85RArz2tGwFjHwyRyLPRnVfcfFs"
		testRedeemScriptHex := "52210367a032d775cab911829c5e4a8d5acc7cbdb1e0500865f7c6d72713ad53c0450721035b21ca62439b8087b0cd657039b5a1fd558747dbb7f69e8498ac10057f16a84c52ae"

		P2SHAddress, redeemScriptHex := generateAddress(coinName, testM, testN, testPublicKeys)
		if testAddress != P2SHAddress {
			testutils.CompareError(t, "Generated P2SH address different from expected address.", testAddress, P2SHAddress)
		}
		if testRedeemScriptHex != redeemScriptHex {
			testutils.CompareError(t, "Generated P2SH address different from expected address.", testRedeemScriptHex, redeemScriptHex)
		}
	}
	{
		//2-of-3 multisig test DASH
		testM := 2
		testN := 3
		coinName := "Dash"
		testPublicKeys := "02a862b412ff9e3afd01a2873a02622897f6df92e3fc85597788b898309fec882e,0315617694c9d93f0ce92769e050a6868ffc74d229077379c0af8bfb193c3d351c,0287ce6cf69b85593ce7db801874c9a2fb1b653dbe5dd9ebfa73e98b710af9e9ce"
		testAddress := "7UPpy23Fghj7AqMUCD3pxzR5jzwKb4kCUt"
		testRedeemScriptHex := "522102a862b412ff9e3afd01a2873a02622897f6df92e3fc85597788b898309fec882e210315617694c9d93f0ce92769e050a6868ffc74d229077379c0af8bfb193c3d351c210287ce6cf69b85593ce7db801874c9a2fb1b653dbe5dd9ebfa73e98b710af9e9ce53ae"

		P2SHAddress, redeemScriptHex := generateAddress(coinName, testM, testN, testPublicKeys)
		if testAddress != P2SHAddress {
			testutils.CompareError(t, "Generated P2SH address different from expected address.", testAddress, P2SHAddress)
		}
		if testRedeemScriptHex != redeemScriptHex {
			testutils.CompareError(t, "Generated P2SH address different from expected address.", testRedeemScriptHex, redeemScriptHex)
		}
	}
	{
		//5-of-7 multisig test Bitcoin
		coinName := "Bitcoin"
		testM := 5
		testN := 7
		testPublicKeys := "04c22e4293d1d462eef905e592ad4aff332aa52c3415b824cd85cf594258d92c836fe797187bc2459261e0597c4ef351c5d0c26f7a60165221e221a38e448ad08c,04bb28684dfe23852a7c276827dd448c955007e7ccbfacbf536e13f1097b30430ebec5af0bc001e50d3f0e796d52ba43e3c07337bfed2a842659d51632f2b21d28,048f8551173f8e7414ff0e144899b3f70accd957e6913f5cf877bd576f6c16f0aa67fb9b96e0df10562b4f7ba4060acd22f142329ff83f1d96e27f4e4394adeda2,04aa81def7dda6a4f40be2f3287ee3423f255b07965104a7888df075217c9ee5b3e9e2e70115d43bfecbff8062f8289f5cab3d0ebd96c9f55c85f6147ff3a5e949,04493aa5f89ec34184a235b2c9f608eade1634636f94f64b59419875e15cb86a6d8c708a9d5eda3304cb983b2325a57af881ed75f28179f5f263d7758039b68d89,04dc284f749208d7fec57937bc5e72187b064df7d29b7aa82cae273e9a1c91beae9c510e0fd632a3db272c67db04061ea761d1ed91fdb8ab07e354047c64ce405d,042fc7796f54dd482db20f1bcce584f930ae74d5f27fc8336e2701bd0243d681281810c57e079947ebdfdfc8860ed34b0ba32db82a85249adc7c64ab547d48af64"
		testAddress := "34wgSuG9qtaNEV4MGye9UJcffcFTxnmXSC"
		testRedeemScriptHex := "554104c22e4293d1d462eef905e592ad4aff332aa52c3415b824cd85cf594258d92c836fe797187bc2459261e0597c4ef351c5d0c26f7a60165221e221a38e448ad08c4104bb28684dfe23852a7c276827dd448c955007e7ccbfacbf536e13f1097b30430ebec5af0bc001e50d3f0e796d52ba43e3c07337bfed2a842659d51632f2b21d2841048f8551173f8e7414ff0e144899b3f70accd957e6913f5cf877bd576f6c16f0aa67fb9b96e0df10562b4f7ba4060acd22f142329ff83f1d96e27f4e4394adeda24104aa81def7dda6a4f40be2f3287ee3423f255b07965104a7888df075217c9ee5b3e9e2e70115d43bfecbff8062f8289f5cab3d0ebd96c9f55c85f6147ff3a5e9494104493aa5f89ec34184a235b2c9f608eade1634636f94f64b59419875e15cb86a6d8c708a9d5eda3304cb983b2325a57af881ed75f28179f5f263d7758039b68d894104dc284f749208d7fec57937bc5e72187b064df7d29b7aa82cae273e9a1c91beae9c510e0fd632a3db272c67db04061ea761d1ed91fdb8ab07e354047c64ce405d41042fc7796f54dd482db20f1bcce584f930ae74d5f27fc8336e2701bd0243d681281810c57e079947ebdfdfc8860ed34b0ba32db82a85249adc7c64ab547d48af6457ae"

		P2SHAddress, redeemScriptHex := generateAddress(coinName, testM, testN, testPublicKeys)
		if testAddress != P2SHAddress {
			testutils.CompareError(t, "Generated P2SH address different from expected address.", testAddress, P2SHAddress)
		}
		if testRedeemScriptHex != redeemScriptHex {
			testutils.CompareError(t, "Generated P2SH address different from expected address.", testRedeemScriptHex, redeemScriptHex)
		}
	}
	{
		//5-of-7 multisig test DASH
		coinName := "Dash"
		testM := 5
		testN := 7
		testPublicKeys := "02a862b412ff9e3afd01a2873a02622897f6df92e3fc85597788b898309fec882e,0315617694c9d93f0ce92769e050a6868ffc74d229077379c0af8bfb193c3d351c,0287ce6cf69b85593ce7db801874c9a2fb1b653dbe5dd9ebfa73e98b710af9e9ce,02462c8cad4dd0551b5cb1c3b3b2607e58f0e582f6718e7604484dd2a0ebd04252,03c7646def5f3825add597cba980932fe287c212f443c109894248bcd214549445,026a599ac5fe04f1b4896ed2460a5b3821b6e8ff2351126993e069b0d37b7722bf,025e796ff123da5b1e8863d0e18b1fc25d1c0a8fcd3d60283bd2a516a76d78bdfc"
		testAddress := "7hQrdBoNfqhwM5poxPyvU6FyHuw8cCXUm1"
		testRedeemScriptHex := "552102a862b412ff9e3afd01a2873a02622897f6df92e3fc85597788b898309fec882e210315617694c9d93f0ce92769e050a6868ffc74d229077379c0af8bfb193c3d351c210287ce6cf69b85593ce7db801874c9a2fb1b653dbe5dd9ebfa73e98b710af9e9ce2102462c8cad4dd0551b5cb1c3b3b2607e58f0e582f6718e7604484dd2a0ebd042522103c7646def5f3825add597cba980932fe287c212f443c109894248bcd21454944521026a599ac5fe04f1b4896ed2460a5b3821b6e8ff2351126993e069b0d37b7722bf21025e796ff123da5b1e8863d0e18b1fc25d1c0a8fcd3d60283bd2a516a76d78bdfc57ae"

		P2SHAddress, redeemScriptHex := generateAddress(coinName, testM, testN, testPublicKeys)
		if testAddress != P2SHAddress {
			testutils.CompareError(t, "Generated P2SH address different from expected address.", testAddress, P2SHAddress)
		}
		if testRedeemScriptHex != redeemScriptHex {
			testutils.CompareError(t, "Generated P2SH address different from expected address.", testRedeemScriptHex, redeemScriptHex)
		}
	}
}
