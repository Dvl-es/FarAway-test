import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
require('dotenv').config();

const {
  DEPLOYER_PRIVATE_KEY,
} = process.env;

const config: HardhatUserConfig = {
  solidity: "0.8.20",
  networks: {
    sepolia: {
      url: `https://rpc.ankr.com/eth_sepolia`,
      chainId: 11155111,
      accounts: [`${DEPLOYER_PRIVATE_KEY}`]
    },
  },
  etherscan: {
    apiKey: {
      sepolia: process.env.SEPOLIA_API_KEY,
    }
  },
};

export default config;
