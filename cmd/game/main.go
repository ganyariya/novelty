package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ganyariya/novelty/internal/application/usecase"
	"github.com/ganyariya/novelty/internal/infrastructure/persistence"
	"github.com/ganyariya/novelty/internal/presentation/tui"
	"github.com/ganyariya/novelty/pkg/config"
)

func main() {
	cfg := config.DefaultConfig()

	scenarioRepo := persistence.NewFileScenarioRepository(cfg.Game.ScriptDir)

	characterRepo := persistence.NewFileCharacterRepository()

	gameUseCase := usecase.NewGameUseCase(scenarioRepo, characterRepo, cfg)

	initialWidth := cfg.Display.WindowWidth
	initialHeight := cfg.Display.WindowHeight

	model := tui.NewGameModel(gameUseCase, initialWidth, initialHeight)

	p := tea.NewProgram(model, tea.WithAltScreen())

	// fmt.Println("Starting Terminal Novel Game Engine...")
	// fmt.Println("Controls:")
	// fmt.Println("  Enter/Space: Advance text")
	// fmt.Println("  A: Toggle Auto mode")
	// fmt.Println("  S: Toggle Skip mode")
	// fmt.Println("  D: Toggle Debug info")
	// fmt.Println("  1/2/3: Switch display mode (ADV/NVL/Hybrid)")
	// fmt.Println("  Ctrl+1-4: Change text speed")
	// fmt.Println("  Q: Quit")
	// fmt.Println()
	// fmt.Println("Press any key to start...")

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
