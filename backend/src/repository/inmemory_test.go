package repository

import (
	"backend/src/contracts"
	"backend/src/models"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"log"
	"math/big"
	"testing"
)

const COLLECTION_ADDRESS = "0xB278B62Ea635bEA643aE2418bD52eF38E475Ab92"
const COLLECTION_ADDRESS2 = "0xB278B62Ea635bEA643aE2418bD52eF38E475Ab93"
const COLLECTION_NAME = "NAME"
const COLLECTION_SYMBOL = "SYMBOL"
const COLLECTION_HASH = "0xd22f414eea255a6b362ecae420fa9ae901d9d315f9c52c0fb971b89ccd3c6f3f"

const TOKEN_ID = 1
const TOKEN_URI = "some/uri/"
const TOKEN_OWNER = "0x031c1FAe59dEdCdc243f5C73d5F86D1Dcb02b4c4"

var storage *InMemoryStorage

func getToken(collectionAddress string) *contracts.NftManagerTokenMinted {
	return &contracts.NftManagerTokenMinted{
		Collection: common.HexToAddress(collectionAddress),
		TokenId:    big.NewInt(TOKEN_ID),
		TokenURI:   TOKEN_URI,
		Recipient:  common.HexToAddress(TOKEN_OWNER),
		Raw: types.Log{
			TxHash: common.HexToHash(COLLECTION_HASH),
		},
	}
}

func getCollection(address string) *contracts.NftManagerCollectionCreated {
	return &contracts.NftManagerCollectionCreated{
		Collection: common.HexToAddress(address),
		Name:       COLLECTION_NAME,
		Symbol:     COLLECTION_SYMBOL,
		Raw: types.Log{
			TxHash: common.HexToHash(COLLECTION_HASH),
		},
	}
}

func createStorage() *InMemoryStorage {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	return NewInMemoryStorage(sugar)
}

func setupSuite(t *testing.T) func(t *testing.T) {
	log.Println("setup suite")

	storage = createStorage()

	// Return a function to teardown the test
	return func(t *testing.T) {
		log.Println("teardown suite")
	}
}

func TestInMemoryStorage_GetAllCollections(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	storage.AddCollectionCreatedEvent(getCollection(COLLECTION_ADDRESS))
	collections := storage.GetAllCollections()
	want := models.NFTCollection{
		Address: common.HexToAddress(COLLECTION_ADDRESS),
		Name:    COLLECTION_NAME,
		Symbol:  COLLECTION_SYMBOL,
		TxHash:  common.HexToHash(COLLECTION_HASH),
	}

	if len(collections) != 1 || *collections[0] != want {
		t.Fatalf("Not %v collection", want)
	}
}

func TestInMemoryStorage_GetAllTokens(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	//no collections yet
	tokens := storage.GetAllTokens()
	if len(tokens) != 0 {
		t.Fatalf("Invalid length")
	}

	storage.AddCollectionCreatedEvent(getCollection(COLLECTION_ADDRESS))
	token := getToken(COLLECTION_ADDRESS)
	storage.AddTokenMinted(token)

	tokens = storage.GetAllTokens()
	if len(tokens) != 1 {
		t.Fatalf("Invalid length")
	}

	want := models.NFT{
		Collection: common.HexToAddress(COLLECTION_ADDRESS),
		TokenURI:   TOKEN_URI,
		TokenID:    fmt.Sprintf("%d", TOKEN_ID),
		Owner:      common.HexToAddress(TOKEN_OWNER),
		TxHash:     common.HexToHash(COLLECTION_HASH),
	}
	if *tokens[0] != want {
		t.Fatalf("Expected: %v, got: %v", want, tokens[0])
	}

	token2 := getToken(COLLECTION_ADDRESS2)
	storage.AddTokenMinted(token2)

	tokens = storage.GetAllTokens()
	if len(tokens) != 2 {
		t.Fatalf("Invalid length")
	}

	want2 := models.NFT{
		Collection: common.HexToAddress(COLLECTION_ADDRESS2),
		TokenURI:   TOKEN_URI,
		TokenID:    fmt.Sprintf("%d", TOKEN_ID),
		Owner:      common.HexToAddress(TOKEN_OWNER),
		TxHash:     common.HexToHash(COLLECTION_HASH),
	}
	if *tokens[1] != want2 {
		t.Fatalf("Expected: %v, got: %v", want2, tokens[0])
	}
}

func TestInMemoryStorage_GetTokensByCollection(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	storage.AddCollectionCreatedEvent(getCollection(COLLECTION_ADDRESS))
	storage.AddCollectionCreatedEvent(getCollection(COLLECTION_ADDRESS2))
	token := getToken(COLLECTION_ADDRESS)
	storage.AddTokenMinted(token)

	tokens := storage.GetTokensByCollection(common.HexToAddress(COLLECTION_ADDRESS))
	if len(tokens) != 1 {
		t.Fatalf("Invalid length")
	}

	want := models.NFT{
		Collection: common.HexToAddress(COLLECTION_ADDRESS),
		TokenURI:   TOKEN_URI,
		TokenID:    fmt.Sprintf("%d", TOKEN_ID),
		Owner:      common.HexToAddress(TOKEN_OWNER),
		TxHash:     common.HexToHash(COLLECTION_HASH),
	}
	if *tokens[0] != want {
		t.Fatalf("Expected: %v, got: %v", want, tokens[0])
	}

	tokens2 := storage.GetTokensByCollection(common.HexToAddress(COLLECTION_ADDRESS2))
	if len(tokens2) != 0 {
		t.Fatalf("Invalid length")
	}
}
