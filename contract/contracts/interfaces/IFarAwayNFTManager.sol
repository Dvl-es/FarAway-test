// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

interface IFarAwayNFTManager {

    event CollectionCreated(address collection, string name, string symbol);
    event TokenMinted(address collection, address recipient, uint256 tokenId, string tokenURI);

    function deployCollection(string calldata uri, string calldata name, string calldata symbol) external;
    function mint(address collection, address to, uint256 tokenId) external;
}