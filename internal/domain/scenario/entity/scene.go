package entity

import (
	"github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
)

type Scene struct {
	id       valueobject.SceneID
	filePath string
	loaded   bool
}

func NewScene(id valueobject.SceneID, filePath string) *Scene {
	return &Scene{
		id:       id,
		filePath: filePath,
		loaded:   false,
	}
}

func (s *Scene) ID() valueobject.SceneID {
	return s.id
}

func (s *Scene) FilePath() string {
	return s.filePath
}

func (s *Scene) IsLoaded() bool {
	return s.loaded
}

func (s *Scene) SetLoaded(loaded bool) {
	s.loaded = loaded
}

func (s *Scene) SetFilePath(filePath string) {
	s.filePath = filePath
}