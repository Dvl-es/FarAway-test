import {loadFixture,} from "@nomicfoundation/hardhat-toolbox/network-helpers";
import {anyValue} from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import {expect} from "chai";
import {ethers} from "hardhat";
import {EventLog, Log} from "ethers";
import {FarAwayNFT, FarAwayNFT__factory} from "../typechain-types";

const BASE_URI = "base_uri/";
const NFT_NAME = "nft_name";
const NFT_SYMBOL = "nft_symbol";

const TOKEN_ID = "1";
const TOKEN_ID_2 = "2";

describe("NFT manager unit tests", function () {

    function extractCollectionAddress(logs: Array<EventLog | Log>): string | null {
        for (const event of logs) {
            if ("eventName" in event && event.eventName == "CollectionCreated") {
                return event.args[0];
            }
        }
        return null;
    }

    async function getCollectionContract(address: string, factory: FarAwayNFT__factory): Promise<FarAwayNFT> {
        return (factory.attach(address) as FarAwayNFT);
    }

    async function deployNFTManager() {

        const [owner, otherAccount] = await ethers.getSigners();

        const NFTManager = await ethers.getContractFactory("FarAwayNFTManager");
        const nftManager = await NFTManager.deploy();

        const NFT = await ethers.getContractFactory("FarAwayNFT");

        return {nftManager, owner, otherAccount, NFT};
    }

    async function deployCollection() {

        const [owner, otherAccount] = await ethers.getSigners();

        const NFTManager = await ethers.getContractFactory("FarAwayNFTManager");
        const nftManager = await NFTManager.deploy();

        const NFT = await ethers.getContractFactory("FarAwayNFT");

        const tx = await nftManager.deployCollection(NFT_NAME, NFT_SYMBOL, BASE_URI);
        const receipt = await tx.wait();
        const address = extractCollectionAddress(receipt!.logs);
        const collection = await getCollectionContract(address!, NFT);

        return {nftManager, owner, otherAccount, collection};
    }

    describe("Collection", function () {

        describe("Deploy collection", function () {

            it("Should emit event on deploy", async function () {
                const {nftManager, owner, otherAccount} = await loadFixture(
                    deployNFTManager
                );

                await expect(nftManager.deployCollection(NFT_NAME, NFT_SYMBOL, BASE_URI))
                    .to.emit(nftManager, "CollectionCreated")
                    .withArgs(anyValue, NFT_NAME, NFT_SYMBOL);
            });

            it("Should have address in event", async function () {
                const {nftManager, owner, otherAccount} = await loadFixture(
                    deployNFTManager
                );

                const tx = await nftManager.deployCollection(NFT_NAME, NFT_SYMBOL, BASE_URI);
                const receipt = await tx.wait();

                const collectionAddress = extractCollectionAddress(receipt!.logs)
                expect(collectionAddress).not.null;
            });

            it("Should have correct params", async function () {
                const {nftManager, owner, otherAccount, NFT} = await loadFixture(
                    deployNFTManager
                );

                const tx = await nftManager.deployCollection(NFT_NAME, NFT_SYMBOL, BASE_URI);
                const receipt = await tx.wait();

                const collectionAddress = extractCollectionAddress(receipt!.logs);
                const collection = await getCollectionContract(collectionAddress!, NFT);

                expect((await collection.name())).to.equal(NFT_NAME);
                expect((await collection.symbol())).to.equal(NFT_SYMBOL);
            });
        });
    });

    describe("Token", function () {

        describe("Mint", function () {

            it("Should emit correct event", async function () {
                const {nftManager, owner, otherAccount, collection} = await loadFixture(
                    deployCollection
                );

                await expect(nftManager.mint(collection.target, otherAccount.address, TOKEN_ID))
                    .to.emit(nftManager, "TokenMinted")
                    .withArgs(collection.target, otherAccount.address, TOKEN_ID, BASE_URI + TOKEN_ID);
            });

            it("Should not mint same token twice", async function () {
                const {nftManager, owner, otherAccount, collection} = await loadFixture(
                    deployCollection
                );

                await nftManager.mint(collection.target, otherAccount.address, TOKEN_ID);

                await expect(nftManager.mint(collection.target, otherAccount.address, TOKEN_ID))
                    .to.be.revertedWithCustomError(collection, "ERC721InvalidSender");
            });

            it("Should be mintable twice", async function () {
                const {nftManager, owner, otherAccount, collection} = await loadFixture(
                    deployCollection
                );

                await nftManager.mint(collection.target, otherAccount.address, TOKEN_ID);

                await expect(nftManager.mint(collection.target, otherAccount.address, TOKEN_ID_2))
                    .to.emit(nftManager, "TokenMinted")
                    .withArgs(collection.target, otherAccount.address, TOKEN_ID_2, BASE_URI + TOKEN_ID_2);
            });
        });
    });
});
