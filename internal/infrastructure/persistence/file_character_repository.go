package persistence

import (
	"context"

	scenarioEntity "github.com/ganyariya/novelty/internal/domain/scenario/entity"
	"github.com/ganyariya/novelty/internal/domain/scenario/repository"
	"github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
	"github.com/ganyariya/novelty/pkg/errors"
)

type FileCharacterRepository struct {
	characters map[valueobject.CharacterID]*scenarioEntity.Character
}

func NewFileCharacterRepository() repository.CharacterRepository {
	return &FileCharacterRepository{
		characters: make(map[valueobject.CharacterID]*scenarioEntity.Character),
	}
}

func (r *FileCharacterRepository) GetCharacter(ctx context.Context, characterID valueobject.CharacterID) (*scenarioEntity.Character, error) {
	character, exists := r.characters[characterID]
	if !exists {
		return nil, errors.WrapCharacterNotFound(characterID.Value(), nil)
	}
	return character, nil
}

func (r *FileCharacterRepository) GetAllCharacters(ctx context.Context) (map[valueobject.CharacterID]*scenarioEntity.Character, error) {
	result := make(map[valueobject.CharacterID]*scenarioEntity.Character)
	for k, v := range r.characters {
		result[k] = v
	}
	return result, nil
}

func (r *FileCharacterRepository) SaveCharacter(ctx context.Context, character *scenarioEntity.Character) error {
	r.characters[character.ID()] = character
	return nil
}

func (r *FileCharacterRepository) RegisterCharacter(ctx context.Context, character *scenarioEntity.Character) error {
	r.characters[character.ID()] = character
	return nil
}