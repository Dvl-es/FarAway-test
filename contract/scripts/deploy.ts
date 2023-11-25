import { ethers } from "hardhat";

async function main() {
  const nftManager = await ethers.deployContract("FarAwayNFTManager");

  await nftManager.waitForDeployment();

  console.log(
    `NFT manager deployed to ${nftManager.target}`
  );
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
