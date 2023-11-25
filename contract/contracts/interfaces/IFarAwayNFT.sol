// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

interface IFarAwayNFT {

    function mint(address to, uint256 tokenId) external returns(string memory);
}