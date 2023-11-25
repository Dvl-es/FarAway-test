// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

import "./interfaces/IFarAwayNFTManager.sol";
import "./interfaces/IFarAwayNFT.sol";
import {FarAwayNFT} from "./FarAwayNFT.sol";

contract FarAwayNFTManager is IFarAwayNFTManager {

    function deployCollection(string calldata name, string calldata symbol, string memory uri) external override {
        IERC721 newCollection = new FarAwayNFT(uri, name, symbol);

        emit CollectionCreated(address(newCollection), name, symbol);
    }

    function mint(address collection, address to, uint256 tokenId) external override {
        string memory tokenURI = IFarAwayNFT(collection).mint(to, tokenId);

        emit TokenMinted(collection, to, tokenId, tokenURI);
    }
}
