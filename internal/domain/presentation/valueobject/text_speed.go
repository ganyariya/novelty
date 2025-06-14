package valueobject

import "time"

type TextSpeed struct {
	value int
}

const (
	TextSpeedSlow   = 1
	TextSpeedMedium = 2
	TextSpeedFast   = 3
	TextSpeedMax    = 4
)

func NewTextSpeed(speed int) TextSpeed {
	if speed < TextSpeedSlow {
		speed = TextSpeedSlow
	}
	if speed > TextSpeedMax {
		speed = TextSpeedMax
	}
	return TextSpeed{value: speed}
}

func (t TextSpeed) Value() int {
	return t.value
}

func (t TextSpeed) Duration() time.Duration {
	switch t.value {
	case TextSpeedSlow:
		return 150 * time.Millisecond
	case TextSpeedMedium:
		return 100 * time.Millisecond
	case TextSpeedFast:
		return 50 * time.Millisecond
	case TextSpeedMax:
		return 10 * time.Millisecond
	default:
		return 100 * time.Millisecond
	}
}

func (t TextSpeed) String() string {
	switch t.value {
	case TextSpeedSlow:
		return "slow"
	case TextSpeedMedium:
		return "medium"
	case TextSpeedFast:
		return "fast"
	case TextSpeedMax:
		return "max"
	default:
		return "medium"
	}
}