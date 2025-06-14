package valueobject

import (
	"fmt"
	"strings"
)

type SceneID struct {
	value string
}

func NewSceneID(chapter, scene string) SceneID {
	return SceneID{value: fmt.Sprintf("%s/%s", chapter, scene)}
}

func NewSceneIDFromPath(path string) SceneID {
	return SceneID{value: path}
}

func (s SceneID) String() string {
	return s.value
}

func (s SceneID) Value() string {
	return s.value
}

func (s SceneID) Chapter() string {
	parts := strings.Split(s.value, "/")
	if len(parts) >= 1 {
		return parts[0]
	}
	return ""
}

func (s SceneID) Scene() string {
	parts := strings.Split(s.value, "/")
	if len(parts) >= 2 {
		return parts[1]
	}
	return ""
}

func (s SceneID) IsEmpty() bool {
	return s.value == ""
}

func (s SceneID) Equals(other SceneID) bool {
	return s.value == other.value
}