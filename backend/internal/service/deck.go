package service

import (
	"context"
	"time"

	"backend/proto/protobuf"

	"backend/internal/repository/deck"

	"github.com/google/uuid"
)

type DeckService struct {
	repo *deck.DynamoRepository
}

func NewDeckService(repo *deck.DynamoRepository) *DeckService {
	return &DeckService{
		repo: repo,
	}
}

func (s *DeckService) CreateDeck(ctx context.Context, req *protobuf.CreateDeckRequest) (*protobuf.Deck, error) {
	deck := &protobuf.Deck{
		DeckId:    uuid.New().String(),
		Name:      req.Name,
		Owner:     req.Owner,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	if err := s.repo.Create(ctx, deck); err != nil {
		return nil, err
	}

	return deck, nil
}

func (s *DeckService) GetDeck(ctx context.Context, req *protobuf.GetDeckRequest) (*protobuf.GetDeckResponse, error) {
	deck, err := s.repo.Get(ctx, req.DeckId, req.Owner)
	if err != nil {
		return nil, err
	}

	return &protobuf.GetDeckResponse{
		Deck: deck,
	}, nil
}

func (s *DeckService) ListDecks(ctx context.Context, req *protobuf.ListDecksRequest) (*protobuf.ListDecksResponse, error) {
	decks, err := s.repo.List(ctx, req.Owner)
	if err != nil {
		return nil, err
	}

	return &protobuf.ListDecksResponse{
		Decks: decks,
	}, nil
}

func (s *DeckService) UpdateDeck(ctx context.Context, req *protobuf.UpdateDeckRequest) (*protobuf.UpdateDeckResponse, error) {
	deck, err := s.repo.Get(ctx, req.DeckId, req.Owner)
	if err != nil {
		return nil, err
	}

	deck.Name = req.Name
	if err := s.repo.Update(ctx, deck); err != nil {
		return nil, err
	}

	return &protobuf.UpdateDeckResponse{
		Deck: deck,
	}, nil
}

func (s *DeckService) DeleteDeck(ctx context.Context, req *protobuf.DeleteDeckRequest) (*protobuf.DeleteDeckResponse, error) {
	if err := s.repo.Delete(ctx, req.DeckId, req.Owner); err != nil {
		return nil, err
	}

	return &protobuf.DeleteDeckResponse{
		Success: true,
	}, nil
} 
