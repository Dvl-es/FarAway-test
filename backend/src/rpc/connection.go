package rpc

import (
	"backend/src/config"
	"backend/src/contracts"
	"backend/src/repository"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type Connection struct {
	logger          *zap.SugaredLogger
	client          *ethclient.Client
	managerContract *contracts.NftManager
	storage         repository.Repository
}

func NewRpcConnection(logger *zap.SugaredLogger, storage repository.Repository) *Connection {
	return &Connection{
		logger:  logger,
		storage: storage,
	}
}

func (c *Connection) Start(appConfig *config.Config) {
	client, err := ethclient.Dial(appConfig.RpcURL)
	if err != nil {
		c.logger.Fatalf("Failed to dial node: %v", err)
	}
	c.client = client

	contract, err := contracts.NewNftManager(common.HexToAddress(appConfig.NftManagerAddress), c.client)
	if err != nil {
		c.logger.Fatalf("Failed to connect with contract: %v", err)
	}
	c.managerContract = contract

	c.subscribe()

	c.logger.Debugf("RPC connection started")
}

func (c *Connection) subscribe() {
	go c.subscribeToCollectionCreated()
	go c.subscribeToTokenMinted()
}

func (c *Connection) subscribeToCollectionCreated() {
	collections := make(chan *contracts.NftManagerCollectionCreated)
	sub, err := c.managerContract.WatchCollectionCreated(nil, collections)
	if err != nil {
		c.logger.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			c.logger.Fatal(err)
		case event := <-collections:
			c.logger.Debugf("Got collection created event")
			c.storage.AddCollectionCreatedEvent(event)
		}
	}
}

func (c *Connection) subscribeToTokenMinted() {
	tokens := make(chan *contracts.NftManagerTokenMinted)
	sub, err := c.managerContract.WatchTokenMinted(nil, tokens)
	if err != nil {
		c.logger.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			c.logger.Fatal(err)
		case event := <-tokens:
			c.logger.Debugf("Got token minted event")
			c.storage.AddTokenMinted(event)
		}
	}
}
