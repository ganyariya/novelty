package tui

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ganyariya/novelty/internal/application/usecase"
	"github.com/ganyariya/novelty/internal/domain/presentation/valueobject"
	scenarioVO "github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
	"github.com/ganyariya/novelty/internal/infrastructure/tui"
)

type tickMsg time.Time

type GameModel struct {
	gameUseCase  *usecase.GameUseCase
	renderer     *tui.GameRenderer
	lastTick     time.Time
	width        int
	height       int
	ctx          context.Context
	initialized  bool
}

func NewGameModel(gameUseCase *usecase.GameUseCase, width, height int) *GameModel {
	return &GameModel{
		gameUseCase: gameUseCase,
		renderer:    tui.NewGameRenderer(width, height),
		lastTick:    time.Now(),
		width:       width,
		height:      height,
		ctx:         context.Background(),
		initialized: false,
	}
}

func (m *GameModel) Init() tea.Cmd {
	return tea.Batch(
		m.initGame(),
		m.tickCmd(),
	)
}

func (m *GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.renderer = tui.NewGameRenderer(m.width, m.height)
		return m, nil
		
	case tea.KeyMsg:
		return m.handleKeyPress(msg)
		
	case tickMsg:
		return m.handleTick(msg)
		
	default:
		return m, nil
	}
}

func (m *GameModel) View() string {
	if !m.initialized {
		return "Loading..."
	}
	
	m.gameUseCase.ProcessNewMessage()
	return m.renderer.RenderGame(m.gameUseCase.GetDisplayState())
}

func (m *GameModel) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
		
	case "enter", " ":
		m.gameUseCase.AdvanceText()
		return m, nil
		
	case "a":
		m.gameUseCase.ToggleAutoMode()
		return m, nil
		
	case "s":
		m.gameUseCase.ToggleSkipMode()
		return m, nil
		
	case "d":
		m.gameUseCase.ToggleDebugInfo()
		return m, nil
		
	case "1":
		m.gameUseCase.SetDisplayMode(scenarioVO.DisplayModeADV)
		return m, nil
		
	case "2":
		m.gameUseCase.SetDisplayMode(scenarioVO.DisplayModeNVL)
		return m, nil
		
	case "3":
		m.gameUseCase.SetDisplayMode(scenarioVO.DisplayModeHybrid)
		return m, nil
		
	case "ctrl+1":
		m.gameUseCase.SetTextSpeed(valueobject.NewTextSpeed(1))
		return m, nil
		
	case "ctrl+2":
		m.gameUseCase.SetTextSpeed(valueobject.NewTextSpeed(2))
		return m, nil
		
	case "ctrl+3":
		m.gameUseCase.SetTextSpeed(valueobject.NewTextSpeed(3))
		return m, nil
		
	case "ctrl+4":
		m.gameUseCase.SetTextSpeed(valueobject.NewTextSpeed(4))
		return m, nil
	}
	
	return m, nil
}

func (m *GameModel) handleTick(msg tickMsg) (tea.Model, tea.Cmd) {
	now := time.Time(msg)
	elapsed := now.Sub(m.lastTick)
	
	updated := m.gameUseCase.UpdateTyping()
	
	if m.gameUseCase.ShouldAutoAdvance(elapsed) {
		m.gameUseCase.AdvanceText()
		m.lastTick = now
	}
	
	if updated {
		m.lastTick = now
	}
	
	return m, m.tickCmd()
}

func (m *GameModel) tickCmd() tea.Cmd {
	interval := m.gameUseCase.GetTypingInterval()
	return tea.Tick(interval, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m *GameModel) initGame() tea.Cmd {
	return func() tea.Msg {
		if err := m.gameUseCase.InitializeGame(m.ctx); err != nil {
			return tea.Quit()
		}
		m.initialized = true
		return nil
	}
}