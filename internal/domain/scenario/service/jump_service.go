package service

import (
	"context"

	"github.com/ganyariya/novelty/internal/domain/game/entity"
	"github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
)

type JumpService struct {
	scenarioService *ScenarioService
}

func NewJumpService(scenarioService *ScenarioService) *JumpService {
	return &JumpService{
		scenarioService: scenarioService,
	}
}

func (j *JumpService) JumpToScene(
	ctx context.Context,
	scenePath, functionName string,
	gameState *entity.GameState,
) error {
	target := valueobject.NewJumpTarget(scenePath, functionName)
	return j.scenarioService.ExecuteJump(ctx, target, gameState)
}

func (j *JumpService) JumpToLabel(
	ctx context.Context,
	functionName string,
	gameState *entity.GameState,
) error {
	target := valueobject.NewJumpTargetFromSceneID(gameState.CurrentScene(), functionName)
	return j.scenarioService.ExecuteJump(ctx, target, gameState)
}