# coconut-multifruit
Bitcoin, Litecoin, Dash  Multisig Address
1) one step
clone repo

```bash
git@github.com:gsix/coconut-multifruit.git
go get not work ptivate repo :(
```

2) two step build script main.go

```bash
  cd coconut-multifruit
  go build main.go
```

**Example:**

### Generate P2SH Multisig Address

```bash
coconut-multifruit address --coin-name="Coin name" --m=M --n=N --public-keys=PUBLIC-KEYS(Comma separated, Hex format)
```

**Example:** (2-of-2 Multisig)

```bash
coconut-multifruit address --coin-name="Bitcoin" --m 2 --n 2 --public-keys 0367a032d775cab911829c5e4a8d5acc7cbdb1e0500865f7c6d72713ad53c04507,035b21ca62439b8087b0cd657039b5a1fd558747dbb7f69e8498ac10057f16a84c
```

```bash
coconut-multifruit address --coin-name="Dash" --m 2 --n 2 --public-keys 0367a032d775cab911829c5e4a8d5acc7cbdb1e0500865f7c6d72713ad53c04507,035b21ca62439b8087b0cd657039b5a1fd558747dbb7f69e8498ac10057f16a84c
```

```bash
coconut-multifruitaddress --coin-name="LitecoinTest" --m 2 --n 2 --public-keys 0367a032d775cab911829c5e4a8d5acc7cbdb1e0500865f7c6d72713ad53c04507,035b21ca62439b8087b0cd657039b5a1fd558747dbb7f69e8498ac10057f16a84c
```
