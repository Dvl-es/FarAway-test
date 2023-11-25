# FarAway backend
## Develop a simple backend server with an in-memory storage to handle emitted events and serve them via HTTP.


### Setup envs before run: 
* RPC_URL
* NFT_MANAGER_ADDRESS

#### Generate abi wrapper
abigen --abi src/contracts/abi/IFarAwayNFTManager.json --pkg contracts --type NftManager --out src/contracts/nftmanager.go


