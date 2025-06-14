package repository

import (
	"context"

	"github.com/ganyariya/novelty/internal/domain/game/entity"
	"github.com/ganyariya/novelty/internal/domain/game/valueobject"
)

type SaveRepository interface {
	Save(ctx context.Context, saveData *entity.SaveData) error
	Load(ctx context.Context, slot valueobject.SaveSlot) (*entity.SaveData, error)
	ListSaves(ctx context.Context) ([]valueobject.SaveSlot, error)
	Delete(ctx context.Context, slot valueobject.SaveSlot) error
	Exists(ctx context.Context, slot valueobject.SaveSlot) (bool, error)
}