1. Develop a smart contract(s) on Solidity for deploying an NFT collection (ERC721) with several arguments (a name, 
symbol). The smart contract should emit the following events:
   a. CollectionCreated(address collection, string name, string symbol)
   b. TokenMinted(address collection, address recipient, uint256 tokenId, string tokenURI)
2. Develop a simple backend server with an in-memory storage to handle emitted events and serve them via HTTP.
3. Develop a frontend demo application that interacts with this smart contract and provides the following functionality:
   a. Create a new NFT collection with a specified name and symbol (from the user input);
   b. Mints a new NFT item with a specified collection address (only created at 3.a), token id, token URI.
   Estimated time: ~ 60 minutes.
