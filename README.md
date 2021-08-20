# cross-chain-contract
PChain cross-chain Bridge Contract

## Generate abi files instruction

Download go-ethereum, choose the preferred version and build the executable abigen file.

```sh
abigen --sol ./contracts/PBridge.sol --pkg pbridge > ./pbridge_abi/pbridge_abi.go
```