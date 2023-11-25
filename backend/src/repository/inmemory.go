package repository

import (
	"backend/src/contracts"
	"backend/src/models"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"sync"
)

type InMemoryStorage struct {
	logger      *zap.SugaredLogger
	tokensMx    sync.Mutex
	tokens      map[common.Address][]*models.NFT        // collection address => token
	collections map[common.Address]models.NFTCollection // collection address => collection
}

func NewInMemoryStorage(logger *zap.SugaredLogger) *InMemoryStorage {
	return &InMemoryStorage{
		collections: make(map[common.Address]models.NFTCollection),
		tokens:      make(map[common.Address][]*models.NFT),
		logger:      logger,
	}
}

func (s *InMemoryStorage) AddCollectionCreatedEvent(collectionCreatedEvent *contracts.NftManagerCollectionCreated) {
	collection := models.NFTCollection{
		Address: collectionCreatedEvent.Collection,
		Name:    collectionCreatedEvent.Name,
		Symbol:  collectionCreatedEvent.Symbol,
		TxHash:  collectionCreatedEvent.Raw.TxHash,
	}

	s.collections[collection.Address] = collection
}

func (s *InMemoryStorage) AddTokenMinted(tokenMintedEvent *contracts.NftManagerTokenMinted) {
	nft := models.NFT{
		Collection: tokenMintedEvent.Collection,
		TokenURI:   tokenMintedEvent.TokenURI,
		TokenID:    tokenMintedEvent.TokenId.String(),
		Owner:      tokenMintedEvent.Recipient,
		TxHash:     tokenMintedEvent.Raw.TxHash,
	}

	s.tokensMx.Lock()
	defer s.tokensMx.Unlock()
	s.tokens[nft.Collection] = append(s.tokens[nft.Collection], &nft)
}

func (s *InMemoryStorage) GetAllCollections() []*models.NFTCollection {
	collections := make([]*models.NFTCollection, 0)
	for _, c := range s.collections {
		collections = append(collections, &c)
	}
	return collections
}

func (s *InMemoryStorage) GetAllTokens() []*models.NFT {
	nfts := make([]*models.NFT, 0)
	for _, collection := range s.tokens {
		nfts = append(nfts, collection...)
	}
	return nfts
}

func (s *InMemoryStorage) GetTokensByCollection(collection common.Address) []*models.NFT {
	results, ok := s.tokens[collection]
	if !ok {
		return make([]*models.NFT, 0)
	} else {
		return results
	}
}
