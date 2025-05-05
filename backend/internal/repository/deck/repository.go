package deck

import (
	"context"
)

type Deck struct {
	DeckId    string `json:"deck_id"`
	Name      string `json:"name"`
	Owner     string `json:"owner"`
	CreatedAt string `json:"created_at"`
}

// Deck の CRUD 操作を行う Repository インターフェース
type Repository interface {
	Create(ctx context.Context, deck *Deck) error
	Get(ctx context.Context, deckId string) (*Deck, error)
	Update(ctx context.Context, deck *Deck) error
	Delete(ctx context.Context, deckId string) error
}
