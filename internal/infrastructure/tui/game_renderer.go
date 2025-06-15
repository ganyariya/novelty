package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/ganyariya/novelty/internal/domain/presentation/entity"
	scenarioVO "github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
)

type GameRenderer struct {
	width       int
	height      int
	baseStyle   lipgloss.Style
	borderStyle lipgloss.Style
}

func NewGameRenderer(width, height int) *GameRenderer {
	return &GameRenderer{
		width:  width,
		height: height,
		baseStyle: lipgloss.NewStyle().
			Width(width).
			Height(height),
		borderStyle: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")),
	}
}

func (r *GameRenderer) RenderGame(displayState *entity.DisplayState) string {
	switch displayState.DisplayMode() {
	case scenarioVO.DisplayModeADV:
		return r.renderADVMode(displayState)
	case scenarioVO.DisplayModeNVL:
		return r.renderNVLMode(displayState)
	case scenarioVO.DisplayModeHybrid:
		return r.renderHybridMode(displayState)
	default:
		return r.renderHybridMode(displayState)
	}
}

func (r *GameRenderer) renderADVMode(displayState *entity.DisplayState) string {
	var content strings.Builder

	backgroundHeight := r.height - 8
	messageWindowHeight := 6

	background := r.createBackground(backgroundHeight)
	messageWindow := r.createMessageWindow(displayState, messageWindowHeight)
	statusBar := r.createStatusBar(displayState)

	content.WriteString(background)
	content.WriteString("\n")
	content.WriteString(messageWindow)
	content.WriteString("\n")
	content.WriteString(statusBar)

	return r.baseStyle.Render(content.String())
}

func (r *GameRenderer) renderNVLMode(displayState *entity.DisplayState) string {
	var content strings.Builder

	textHeight := r.height - 4
	textContent := r.createNVLContent(displayState, textHeight)
	statusBar := r.createStatusBar(displayState)

	content.WriteString(textContent)
	content.WriteString("\n")
	content.WriteString(statusBar)

	return r.baseStyle.Render(content.String())
}

func (r *GameRenderer) renderHybridMode(displayState *entity.DisplayState) string {
	var content strings.Builder

	backlogHeight := r.height - 14
	currentTextHeight := 6

	backlog := r.createBacklog(displayState, backlogHeight)
	currentText := r.createCurrentText(displayState, currentTextHeight)
	statusBar := r.createStatusBar(displayState)

	content.WriteString(backlog)
	content.WriteString("\n")
	content.WriteString(strings.Repeat("─", r.width-4))
	content.WriteString("\n")
	content.WriteString(currentText)
	content.WriteString("\n")
	content.WriteString(statusBar)

	return r.baseStyle.Render(content.String())
}

func (r *GameRenderer) createBackground(height int) string {
	background := lipgloss.NewStyle().
		Width(r.width-4).
		Height(height).
		Align(lipgloss.Center, lipgloss.Center).
		Foreground(lipgloss.Color("240"))

	return background.Render("～ 背景・立ち絵表示領域 ～")
}

func (r *GameRenderer) createMessageWindow(displayState *entity.DisplayState, height int) string {
	messageStyle := lipgloss.NewStyle().
		Width(r.width - 4).
		Height(height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1)

	var content strings.Builder

	if msg := displayState.CurrentMessage(); msg != nil {
		if !msg.IsNarration() {
			speakerStyle := lipgloss.NewStyle().
				Foreground(lipgloss.Color("33")).
				Bold(true)
			content.WriteString(speakerStyle.Render(msg.SpeakerID().Value()))
			content.WriteString(": ")
		}
		content.WriteString(displayState.GetVisibleText())

		if displayState.IsTyping() {
			content.WriteString("█")
		}
	}

	return messageStyle.Render(content.String())
}

func (r *GameRenderer) createNVLContent(displayState *entity.DisplayState, height int) string {
	textStyle := lipgloss.NewStyle().
		Width(r.width - 4).
		Height(height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1)

	var content strings.Builder

	backlog := displayState.Backlog()
	maxMessages := height - 4

	start := 0
	if len(backlog) > maxMessages {
		start = len(backlog) - maxMessages
	}

	for i := start; i < len(backlog); i++ {
		msg := backlog[i]
		if !msg.IsNarration() {
			speakerStyle := lipgloss.NewStyle().
				Foreground(lipgloss.Color("33")).
				Bold(true)
			content.WriteString(speakerStyle.Render(msg.SpeakerID().Value()))
			content.WriteString(": ")
		}
		content.WriteString(msg.Content())
		content.WriteString("\n")
	}

	if msg := displayState.CurrentMessage(); msg != nil {
		if !msg.IsNarration() {
			speakerStyle := lipgloss.NewStyle().
				Foreground(lipgloss.Color("33")).
				Bold(true)
			content.WriteString(speakerStyle.Render(msg.SpeakerID().Value()))
			content.WriteString(": ")
		}
		content.WriteString(displayState.GetVisibleText())

		if displayState.IsTyping() {
			content.WriteString("█")
		}
	}

	return textStyle.Render(content.String())
}

func (r *GameRenderer) createBacklog(displayState *entity.DisplayState, height int) string {
	backlogStyle := lipgloss.NewStyle().
		Width(r.width - 4).
		Height(height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1)

	var content strings.Builder

	backlog := displayState.Backlog()
	maxMessages := height - 4

	start := 0
	if len(backlog) > maxMessages {
		start = len(backlog) - maxMessages
	}

	for i := start; i < len(backlog); i++ {
		msg := backlog[i]
		if !msg.IsNarration() {
			speakerStyle := lipgloss.NewStyle().
				Foreground(lipgloss.Color("240")).
				Bold(false)
			content.WriteString(speakerStyle.Render(msg.SpeakerID().Value()))
			content.WriteString(": ")
		}
		content.WriteString(msg.Content())
		content.WriteString("\n")
	}

	return backlogStyle.Render(content.String())
}

func (r *GameRenderer) createCurrentText(displayState *entity.DisplayState, height int) string {
	currentStyle := lipgloss.NewStyle().
		Width(r.width - 4).
		Height(height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("33")).
		Padding(1)

	var content strings.Builder

	if msg := displayState.CurrentMessage(); msg != nil {
		if !msg.IsNarration() {
			speakerStyle := lipgloss.NewStyle().
				Foreground(lipgloss.Color("33")).
				Bold(true)
			content.WriteString(speakerStyle.Render(msg.SpeakerID().Value()))
			content.WriteString(": ")
		}
		content.WriteString(displayState.GetVisibleText())

		if displayState.IsTyping() {
			content.WriteString("█")
		}
	}

	return currentStyle.Render(content.String())
}

func (r *GameRenderer) createStatusBar(displayState *entity.DisplayState) string {
	statusStyle := lipgloss.NewStyle().
		Width(r.width - 4).
		Height(2).
		Foreground(lipgloss.Color("240"))

	var status []string

	if displayState.AutoMode() {
		status = append(status, "Auto: ON")
	} else {
		status = append(status, "Auto: OFF")
	}

	if displayState.SkipMode() {
		status = append(status, "Skip: ON")
	} else {
		status = append(status, "Skip: OFF")
	}

	status = append(status, "[S]ave", "[L]oad", "[Q]uit")

	if displayState.ShowDebugInfo() {
		status = append(status, fmt.Sprintf("Mode: %s", displayState.DisplayMode().String()))
		status = append(status, fmt.Sprintf("Speed: %s", displayState.TextSpeed().String()))
	}

	return statusStyle.Render(strings.Join(status, " | "))
}
