package usecase

import (
	"context"
	"fmt"
	"time"

	gameEntity "github.com/ganyariya/novelty/internal/domain/game/entity"
	presentationEntity "github.com/ganyariya/novelty/internal/domain/presentation/entity"
	"github.com/ganyariya/novelty/internal/domain/presentation/service"
	"github.com/ganyariya/novelty/internal/domain/presentation/valueobject"
	"github.com/ganyariya/novelty/internal/domain/scenario/repository"
	scenarioService "github.com/ganyariya/novelty/internal/domain/scenario/service"
	scenarioVO "github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
	"github.com/ganyariya/novelty/internal/infrastructure/persistence"
	"github.com/ganyariya/novelty/pkg/config"
)

type GameUseCase struct {
	scenarioRepo     repository.ScenarioRepository
	scenarioService  *scenarioService.ScenarioService
	jumpService      *scenarioService.JumpService
	textDisplayService *service.TextDisplayService
	gameState        *gameEntity.GameState
	displayState     *presentationEntity.DisplayState
	config           *config.Config
}

func NewGameUseCase(
	scenarioRepo repository.ScenarioRepository,
	characterRepo repository.CharacterRepository,
	cfg *config.Config,
) *GameUseCase {
	scenarioSvc := scenarioService.NewScenarioService(scenarioRepo, characterRepo)
	jumpSvc := scenarioService.NewJumpService(scenarioSvc)
	textDisplaySvc := service.NewTextDisplayService()
	
	startSceneID := scenarioVO.NewSceneIDFromPath(cfg.Game.StartScene)
	gameState := gameEntity.NewGameState(startSceneID)
	
	displayState := presentationEntity.NewDisplayState(
		cfg.GetDisplayMode(),
		cfg.GetTextSpeed(),
		cfg.GetColorTheme(),
	)
	
	return &GameUseCase{
		scenarioRepo:       scenarioRepo,
		scenarioService:    scenarioSvc,
		jumpService:        jumpSvc,
		textDisplayService: textDisplaySvc,
		gameState:          gameState,
		displayState:       displayState,
		config:             cfg,
	}
}

func (u *GameUseCase) InitializeGame(ctx context.Context) error {
	fmt.Printf("[DEBUG] Initializing game with scene: %s, function: %s\n", u.config.Game.StartScene, u.config.Game.StartFunction)
	err := u.jumpService.JumpToScene(
		ctx,
		u.config.Game.StartScene,
		u.config.Game.StartFunction,
		u.gameState,
	)
	if err != nil {
		fmt.Printf("[DEBUG] Game initialization failed: %v\n", err)
	} else {
		fmt.Printf("[DEBUG] Game initialized successfully. Messages in history: %d\n", len(u.gameState.History()))
	}
	return err
}

func (u *GameUseCase) UpdateTyping() bool {
	return u.textDisplayService.UpdateTyping(u.displayState)
}

func (u *GameUseCase) GetTypingInterval() time.Duration {
	return u.textDisplayService.GetTypingInterval(u.displayState)
}

func (u *GameUseCase) ShouldAutoAdvance(elapsed time.Duration) bool {
	return u.textDisplayService.ShouldAutoAdvance(u.displayState, elapsed)
}

func (u *GameUseCase) AdvanceText() {
	if u.displayState.IsTyping() {
		// タイピング中の場合はタイピングを完了
		u.textDisplayService.AdvanceToNextMessage(u.displayState)
	} else {
		// タイピング完了後の場合は次のメッセージを表示
		if scenarioRepo, ok := u.scenarioRepo.(*persistence.FileScenarioRepository); ok {
			engine := scenarioRepo.GetEngine()
			executionState := engine.GetExecutionState()
			
			if nextMessage := executionState.GetNextMessage(); nextMessage != nil {
				fmt.Printf("[DEBUG] AdvanceText: Showing next message from queue: %s\n", nextMessage.Content())
				
				// メッセージを履歴に追加
				u.gameState.AddToHistory(nextMessage)
				
				// メッセージを表示
				u.textDisplayService.ShowMessage(u.displayState, nextMessage)
			} else {
				fmt.Printf("[DEBUG] No more messages in queue\n")
			}
		}
	}
}

func (u *GameUseCase) ToggleAutoMode() {
	u.displayState.SetAutoMode(!u.displayState.AutoMode())
}

func (u *GameUseCase) ToggleSkipMode() {
	u.displayState.SetSkipMode(!u.displayState.SkipMode())
}

func (u *GameUseCase) ToggleDebugInfo() {
	u.displayState.SetShowDebugInfo(!u.displayState.ShowDebugInfo())
}

func (u *GameUseCase) SetDisplayMode(mode scenarioVO.DisplayMode) {
	u.displayState.SetDisplayMode(mode)
}

func (u *GameUseCase) SetTextSpeed(speed valueobject.TextSpeed) {
	u.displayState.SetTextSpeed(speed)
}

func (u *GameUseCase) GetGameState() *gameEntity.GameState {
	return u.gameState
}

func (u *GameUseCase) GetDisplayState() *presentationEntity.DisplayState {
	return u.displayState
}

func (u *GameUseCase) ProcessNewMessage() {
	// メッセージキューから新しいメッセージがあるかチェック
	if scenarioRepo, ok := u.scenarioRepo.(*persistence.FileScenarioRepository); ok {
		engine := scenarioRepo.GetEngine()
		executionState := engine.GetExecutionState()
		
		// 現在表示中のメッセージがない場合、キューから次のメッセージを取得
		if u.displayState.CurrentMessage() == nil && executionState.HasPendingMessages() {
			if nextMessage := executionState.GetNextMessage(); nextMessage != nil {
				fmt.Printf("[DEBUG] ProcessNewMessage: Showing message from queue: %s\n", nextMessage.Content())
				
				// メッセージを履歴に追加
				u.gameState.AddToHistory(nextMessage)
				
				// メッセージを表示
				u.textDisplayService.ShowMessage(u.displayState, nextMessage)
				return
			}
		}
	}
	
	// 従来の履歴ベースの処理をフォールバックとして保持
	history := u.gameState.History()
	if len(history) > 0 {
		lastMessage := history[len(history)-1]
		if u.displayState.CurrentMessage() == nil || 
		   u.displayState.CurrentMessage().ID().String() != lastMessage.ID().String() {
			u.textDisplayService.ShowMessage(u.displayState, lastMessage)
		}
	}
}

