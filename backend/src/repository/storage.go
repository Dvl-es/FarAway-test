package repository

import (
	"backend/src/contracts"
	"backend/src/models"
	"github.com/ethereum/go-ethereum/common"
)

type Repository interface {
	AddCollectionCreatedEvent(collectionCreatedEvent *contracts.NftManagerCollectionCreated)
	AddTokenMinted(tokenMintedEvent *contracts.NftManagerTokenMinted)

	GetAllCollections() []*models.NFTCollection
	GetAllTokens() []*models.NFT
	GetTokensByCollection(collection common.Address) []*models.NFT
}
