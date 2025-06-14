package entity

import (
	"github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
)

type MessageID struct {
	value string
}

func NewMessageID(value string) MessageID {
	return MessageID{value: value}
}

func (m MessageID) String() string {
	return m.value
}

type Message struct {
	id          MessageID
	speakerID   valueobject.CharacterID
	content     string
	displayMode valueobject.DisplayMode
	voiceFile   string
}

func NewMessage(
	id MessageID,
	speakerID valueobject.CharacterID,
	content string,
	displayMode valueobject.DisplayMode,
	voiceFile string,
) *Message {
	return &Message{
		id:          id,
		speakerID:   speakerID,
		content:     content,
		displayMode: displayMode,
		voiceFile:   voiceFile,
	}
}

func NewNarrationMessage(id MessageID, content string, displayMode valueobject.DisplayMode) *Message {
	return &Message{
		id:          id,
		speakerID:   valueobject.NewCharacterID(""),
		content:     content,
		displayMode: displayMode,
		voiceFile:   "",
	}
}

func (m *Message) ID() MessageID {
	return m.id
}

func (m *Message) SpeakerID() valueobject.CharacterID {
	return m.speakerID
}

func (m *Message) Content() string {
	return m.content
}

func (m *Message) DisplayMode() valueobject.DisplayMode {
	return m.displayMode
}

func (m *Message) VoiceFile() string {
	return m.voiceFile
}

func (m *Message) IsNarration() bool {
	return m.speakerID.IsEmpty()
}

func (m *Message) HasVoice() bool {
	return m.voiceFile != ""
}