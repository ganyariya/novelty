package service

import (
	"time"

	"github.com/ganyariya/novelty/internal/domain/presentation/entity"
	scenarioEntity "github.com/ganyariya/novelty/internal/domain/scenario/entity"
)

type TextDisplayService struct{}

func NewTextDisplayService() *TextDisplayService {
	return &TextDisplayService{}
}

func (s *TextDisplayService) UpdateTyping(displayState *entity.DisplayState) bool {
	if !displayState.IsTyping() {
		return false
	}
	
	if displayState.SkipMode() {
		displayState.FinishTyping()
		return true
	}
	
	charsPerTick := 1
	if displayState.TextSpeed().Value() >= 3 {
		charsPerTick = 2
	}
	
	displayState.AdvanceTyping(charsPerTick)
	return true
}

func (s *TextDisplayService) GetTypingInterval(displayState *entity.DisplayState) time.Duration {
	if displayState.SkipMode() {
		return 1 * time.Millisecond
	}
	
	return displayState.TextSpeed().Duration()
}

func (s *TextDisplayService) ShouldAutoAdvance(displayState *entity.DisplayState, elapsed time.Duration) bool {
	if !displayState.AutoMode() || displayState.IsTyping() {
		return false
	}
	
	autoDelay := 3 * time.Second
	return elapsed >= autoDelay
}

func (s *TextDisplayService) ShowMessage(displayState *entity.DisplayState, message *scenarioEntity.Message) {
	displayState.SetCurrentMessage(message)
}

func (s *TextDisplayService) AdvanceToNextMessage(displayState *entity.DisplayState) {
	if displayState.IsTyping() {
		displayState.FinishTyping()
	} else {
		displayState.SetCurrentMessage(nil)
	}
}