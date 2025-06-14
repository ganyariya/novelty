package valueobject

type JumpTarget struct {
	sceneID      SceneID
	functionName string
}

func NewJumpTarget(scenePath, funcName string) *JumpTarget {
	return &JumpTarget{
		sceneID:      NewSceneIDFromPath(scenePath),
		functionName: funcName,
	}
}

func NewJumpTargetFromSceneID(sceneID SceneID, funcName string) *JumpTarget {
	return &JumpTarget{
		sceneID:      sceneID,
		functionName: funcName,
	}
}

func (j *JumpTarget) SceneID() SceneID {
	return j.sceneID
}

func (j *JumpTarget) FunctionName() string {
	return j.functionName
}

func (j *JumpTarget) IsValid() bool {
	return !j.sceneID.IsEmpty() && j.functionName != ""
}

func (j *JumpTarget) String() string {
	return j.sceneID.String() + "#" + j.functionName
}