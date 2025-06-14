package repository

import (
	"context"

	"github.com/ganyariya/novelty/internal/domain/scenario/entity"
	"github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
)

type ScenarioRepository interface {
	LoadScene(ctx context.Context, sceneID valueobject.SceneID) (*entity.Scene, error)
	LoadCharacters(ctx context.Context) (map[valueobject.CharacterID]*entity.Character, error)
	ExecuteSceneFunction(ctx context.Context, scene *entity.Scene, functionName string, gameState interface{}) error
	ReloadScene(ctx context.Context, sceneID valueobject.SceneID) (*entity.Scene, error)
}