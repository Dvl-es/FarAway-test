// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

import "./interfaces/IFarAwayNFTManager.sol";
import "./interfaces/IFarAwayNFT.sol";

contract FarAwayNFT is ERC721, Ownable, IFarAwayNFT {

    string _uri;

    constructor(string memory uri, string memory name, string memory symbol) Ownable(msg.sender) ERC721(name, symbol) {
        _uri = uri;
    }

    function mint(address to, uint256 tokenId) external onlyOwner() returns(string memory){
        _safeMint(to, tokenId);

        return tokenURI(tokenId);
    }

    function _baseURI() internal view override returns (string memory) {
        return _uri;
    }
}
