package repository

import (
	"context"

	"github.com/ganyariya/novelty/internal/domain/scenario/entity"
	"github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
)

type CharacterRepository interface {
	GetCharacter(ctx context.Context, characterID valueobject.CharacterID) (*entity.Character, error)
	GetAllCharacters(ctx context.Context) (map[valueobject.CharacterID]*entity.Character, error)
	SaveCharacter(ctx context.Context, character *entity.Character) error
	RegisterCharacter(ctx context.Context, character *entity.Character) error
}