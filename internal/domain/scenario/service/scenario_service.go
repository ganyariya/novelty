package service

import (
	"context"

	"github.com/ganyariya/novelty/internal/domain/game/entity"
	"github.com/ganyariya/novelty/internal/domain/scenario/repository"
	"github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
)

type ScenarioService struct {
	scenarioRepo  repository.ScenarioRepository
	characterRepo repository.CharacterRepository
}

func NewScenarioService(
	scenarioRepo repository.ScenarioRepository,
	characterRepo repository.CharacterRepository,
) *ScenarioService {
	return &ScenarioService{
		scenarioRepo:  scenarioRepo,
		characterRepo: characterRepo,
	}
}

func (s *ScenarioService) ExecuteJump(
	ctx context.Context,
	target *valueobject.JumpTarget,
	gameState *entity.GameState,
) error {
	scene, err := s.scenarioRepo.LoadScene(ctx, target.SceneID())
	if err != nil {
		return err
	}

	gameState.SetCurrentScene(target.SceneID(), target.FunctionName())

	return s.scenarioRepo.ExecuteSceneFunction(ctx, scene, target.FunctionName(), gameState)
}

func (s *ScenarioService) LoadScene(ctx context.Context, sceneID valueobject.SceneID) error {
	_, err := s.scenarioRepo.LoadScene(ctx, sceneID)
	return err
}

func (s *ScenarioService) ReloadScene(ctx context.Context, sceneID valueobject.SceneID) error {
	_, err := s.scenarioRepo.ReloadScene(ctx, sceneID)
	return err
}