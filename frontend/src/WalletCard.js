import React, {useState} from 'react';
import {Button, TextField} from '@mui/material';
import {ethers} from 'ethers';
import managerAbi from "./abi/IFarAwayNFTManager.json";
import "./index.css";

const NFT_MANAGER_ADDRESS = process.env.REACT_APP_NFT_MANAGER_ADDRESS;//"0xC6aB23e1753B5750dB14852046220Eca5bfaE5DA";
const API_ENDPOINT = process.env.REACT_APP_API_ENDPOINT;//"http://localhost:6565/";

const NftManagerInterface = new ethers.utils.Interface(managerAbi);

const provider = new ethers.providers.Web3Provider(window.ethereum);
const WalletCard = () => {
    const [errorMessage, setErrorMessage] = useState(null);
    const [defaultAccount, setDefaultAccount] = useState(null);
    const [signer, setSigner] = useState(null);
    const [userBalance, setUserBalance] = useState(null);
    const [collection, setCollection] = useState(null);
    const [tokenId, setTokenId] = useState(null);
    const [collectionName, setCollectionName] = useState(null);
    const [collectionSymbol, setCollectionSymbol] = useState(null);
    const [collectionUri, setCollectionUri] = useState(null);
    const [serverResponse, setServerResponse] = useState(null);

    const connectWalletHandler = () => {
        if (window.ethereum) {
            provider.send("eth_requestAccounts", []).then(async () => {
                await accountChangedHandler(provider.getSigner());
            })
        } else {
            setErrorMessage("Please Install Metamask!!!");
        }
    }

    const createCollection = async () => {
        if (!collectionName || !collectionSymbol || !collectionUri) {
            return;
        }
        if (window.ethereum) {
            const contract = new ethers.Contract(
                NFT_MANAGER_ADDRESS,
                managerAbi,
                signer,
            );
            const tx = await contract.deployCollection(
                collectionName,
                collectionSymbol,
                collectionUri,
            ).catch(
                err => console.log(err)
            );
            if (!tx) {
                return;
            }

            const receipt = await tx.wait();

            var parsed = NftManagerInterface.parseLog(receipt.events[1]);
            console.log(parsed);
            setCollection(parsed.args.collection);
            console.log(parsed.args.collection);
        } else {
            setErrorMessage("Not connected!");
        }
    }

    const mintToken = async () => {
        if (!collection || !tokenId) {
            return;
        }
        if (window.ethereum) {
            const contract = new ethers.Contract(
                NFT_MANAGER_ADDRESS,
                managerAbi,
                signer,
            );
            const tx = await contract.mint(
                collection,
                defaultAccount,
                tokenId,
            ).catch(
                err => console.log(err)
            );
            if (!tx) {
                return;
            }

            await tx.wait();
            alert("Token minted");

        } else {
            setErrorMessage("Not connected!");
        }
    }

    const apiCall = async (path) => {
        console.log(process.env.REACT_APP_HOST_IP_ADDRESS);
        const xhr = new XMLHttpRequest();
        xhr.open('GET', API_ENDPOINT + path);
        xhr.onload = function () {
            if (xhr.status === 200) {
                var json = JSON.parse(xhr.responseText);
                setServerResponse(JSON.stringify(json, undefined, 4));
            }
        };
        xhr.send();
    }

    const accountChangedHandler = async (newAccount) => {
        setSigner(newAccount);
        const address = await newAccount.getAddress();
        setDefaultAccount(address);
        const balance = await newAccount.getBalance()
        setUserBalance(ethers.utils.formatEther(balance));
        await getuserBalance(address)
    }
    const getuserBalance = async (address) => {
        const balance = await provider.getBalance(address, "latest")
    }
    return (
        <div className="WalletCard" style={{}}>
            <h3 className="h4">
                FarAway test (SEPOLIA!)
            </h3>
            <Button
                style={{background: defaultAccount ? "#A5CC82" : "white"}}
                onClick={connectWalletHandler}>
                {defaultAccount ? "Connected!" : "Connect"}
            </Button>
            <div className="displayAccount">
                <h4 className="walletAddress">Address: {defaultAccount}</h4>
                <div className="balanceDisplay">
                    <h3>
                        Wallet Amount: {userBalance}
                    </h3>
                </div>
            </div>
            <hr/>
            <h3>Create NFT collection</h3>
            <div className='block'
            >
                <div className='collection' style={{margin: "10px"}}>

                    <TextField sx={{ml: 1}} className="input" id="collectionName" variant="standard" label="Name"
                               onChange={(event) => {
                                   setCollectionName(event.target.value);
                               }}/>

                    <TextField sx={{ml: 1}} className="input" id="collectionSymbol" variant="standard" label="Symbol"
                               onChange={(event) => {
                                   setCollectionSymbol(event.target.value);
                               }}/>

                    <TextField sx={{ml: 1}} className="input" id="collectionUri" variant="standard" label="Token URI"
                               onChange={(event) => {
                                   setCollectionUri(event.target.value);
                               }}/>

                    <Button id="createCollectionBtn" className="createBtn"
                            style={{
                                visibility: defaultAccount ? "visible" : "hidden",
                                background: "#A5CC82",
                                color: "black"
                            }}
                            onClick={() => createCollection()}>
                        CREATE
                    </Button>
                </div>
            </div>
            <h3>Mint token for collection</h3>
            <div className='block'
            >
                <div>
                    <TextField
                        sx={{ml: 1}}
                        className="input"
                        id="collectionAddress"
                        label="Collection address"
                        variant="standard"
                        value={collection ? collection : ''}/>

                    <TextField sx={{ml: 1}} className="input" id="tokenId" variant="standard" label="Token ID"
                               onChange={(event) => {
                                   setTokenId(event.target.value);
                               }}/>

                    <Button id="mintBtn" className="createBtn"
                            style={{
                                visibility: defaultAccount ? "visible" : "hidden",
                                background: "#A5CC82",
                                color: "black",
                                position: "absolute",
                            }}
                            onClick={() => mintToken()}>
                        MINT
                    </Button>
                </div>
            </div>
            <h3>API calls</h3>
            <div className='block'>
                <div className="apiBtn">
                    <Button id="getCollections"
                             style={{
                                 background: "#A5CC82",
                                 color: "black",
                                 position: "absolute",
                             }}
                             onClick={() => apiCall("collections/")}>
                    Get collections
                </Button>
                </div>
                <div className="apiBtn">
                    <Button id="getAllTokens"
                            style={{
                                background: "#A5CC82",
                                color: "black",
                                position: "absolute",
                            }}
                            onClick={() => apiCall("tokens/")}>
                        Get all tokens
                    </Button>
                </div>
                <div className="apiBtn">
                    <Button id="getTokensForCollection"
                            style={{
                                background: "#A5CC82",
                                color: "black",
                                position: "absolute",
                            }}
                            onClick={() => apiCall("collections/" + collection)}>
                        Get tokens for collection
                    </Button>
                </div>
            </div>
            {errorMessage}

            <h4>Response</h4>
            {serverResponse}
        </div>
    )
}
export default WalletCard;
