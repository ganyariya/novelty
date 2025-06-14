package entity

import (
	"github.com/ganyariya/novelty/internal/domain/presentation/valueobject"
	scenarioEntity "github.com/ganyariya/novelty/internal/domain/scenario/entity"
	scenarioVO "github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
)

type DisplayState struct {
	currentMessage   *scenarioEntity.Message
	displayMode      scenarioVO.DisplayMode
	textSpeed        valueobject.TextSpeed
	colorTheme       valueobject.ColorTheme
	backlog          []*scenarioEntity.Message
	isTyping         bool
	typingPosition   int
	autoMode         bool
	skipMode         bool
	showDebugInfo    bool
}

func NewDisplayState(
	displayMode scenarioVO.DisplayMode,
	textSpeed valueobject.TextSpeed,
	colorTheme valueobject.ColorTheme,
) *DisplayState {
	return &DisplayState{
		displayMode:    displayMode,
		textSpeed:      textSpeed,
		colorTheme:     colorTheme,
		backlog:        make([]*scenarioEntity.Message, 0),
		isTyping:       false,
		typingPosition: 0,
		autoMode:       false,
		skipMode:       false,
		showDebugInfo:  false,
	}
}

func (d *DisplayState) CurrentMessage() *scenarioEntity.Message {
	return d.currentMessage
}

func (d *DisplayState) SetCurrentMessage(message *scenarioEntity.Message) {
	if d.currentMessage != nil {
		d.backlog = append(d.backlog, d.currentMessage)
	}
	d.currentMessage = message
	d.isTyping = true
	d.typingPosition = 0
}

func (d *DisplayState) DisplayMode() scenarioVO.DisplayMode {
	return d.displayMode
}

func (d *DisplayState) SetDisplayMode(mode scenarioVO.DisplayMode) {
	d.displayMode = mode
}

func (d *DisplayState) TextSpeed() valueobject.TextSpeed {
	return d.textSpeed
}

func (d *DisplayState) SetTextSpeed(speed valueobject.TextSpeed) {
	d.textSpeed = speed
}

func (d *DisplayState) ColorTheme() valueobject.ColorTheme {
	return d.colorTheme
}

func (d *DisplayState) SetColorTheme(theme valueobject.ColorTheme) {
	d.colorTheme = theme
}

func (d *DisplayState) Backlog() []*scenarioEntity.Message {
	result := make([]*scenarioEntity.Message, len(d.backlog))
	copy(result, d.backlog)
	return result
}

func (d *DisplayState) ClearBacklog() {
	d.backlog = make([]*scenarioEntity.Message, 0)
}

func (d *DisplayState) IsTyping() bool {
	return d.isTyping
}

func (d *DisplayState) TypingPosition() int {
	return d.typingPosition
}

func (d *DisplayState) AdvanceTyping(chars int) {
	if d.currentMessage == nil {
		return
	}
	
	d.typingPosition += chars
	if d.typingPosition >= len([]rune(d.currentMessage.Content())) {
		d.typingPosition = len([]rune(d.currentMessage.Content()))
		d.isTyping = false
	}
}

func (d *DisplayState) FinishTyping() {
	if d.currentMessage != nil {
		d.typingPosition = len([]rune(d.currentMessage.Content()))
	}
	d.isTyping = false
}

func (d *DisplayState) AutoMode() bool {
	return d.autoMode
}

func (d *DisplayState) SetAutoMode(auto bool) {
	d.autoMode = auto
}

func (d *DisplayState) SkipMode() bool {
	return d.skipMode
}

func (d *DisplayState) SetSkipMode(skip bool) {
	d.skipMode = skip
}

func (d *DisplayState) ShowDebugInfo() bool {
	return d.showDebugInfo
}

func (d *DisplayState) SetShowDebugInfo(show bool) {
	d.showDebugInfo = show
}

func (d *DisplayState) GetVisibleText() string {
	if d.currentMessage == nil {
		return ""
	}
	
	content := []rune(d.currentMessage.Content())
	if d.typingPosition >= len(content) {
		return string(content)
	}
	
	return string(content[:d.typingPosition])
}