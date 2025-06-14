package config

import (
	"github.com/ganyariya/novelty/internal/domain/presentation/valueobject"
	scenarioVO "github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
)

type Config struct {
	Game        GameConfig        `json:"game"`
	Display     DisplayConfig     `json:"display"`
	Audio       AudioConfig       `json:"audio"`
	Development DevelopmentConfig `json:"development"`
}

type GameConfig struct {
	ScriptDir     string `json:"script_dir"`
	SaveDir       string `json:"save_dir"`
	MaxSaveSlots  int    `json:"max_save_slots"`
	AutoSaveSlot  int    `json:"auto_save_slot"`
	StartScene    string `json:"start_scene"`
	StartFunction string `json:"start_function"`
}

type DisplayConfig struct {
	DefaultMode    string `json:"default_mode"`
	TextSpeed      int    `json:"text_speed"`
	AutoModeDelay  int    `json:"auto_mode_delay"`
	WindowWidth    int    `json:"window_width"`
	WindowHeight   int    `json:"window_height"`
	ColorTheme     string `json:"color_theme"`
	ShowLineNumbers bool  `json:"show_line_numbers"`
}

type AudioConfig struct {
	Enabled          bool    `json:"enabled"`
	MasterVolume     float64 `json:"master_volume"`
	VoiceVolume      float64 `json:"voice_volume"`
	BGMVolume        float64 `json:"bgm_volume"`
	SFXVolume        float64 `json:"sfx_volume"`
	VoiceDir         string  `json:"voice_dir"`
	BGMDir           string  `json:"bgm_dir"`
	SFXDir           string  `json:"sfx_dir"`
}

type DevelopmentConfig struct {
	HotReload     bool   `json:"hot_reload"`
	DebugMode     bool   `json:"debug_mode"`
	LogLevel      string `json:"log_level"`
	ShowDebugInfo bool   `json:"show_debug_info"`
}

func DefaultConfig() *Config {
	return &Config{
		Game: GameConfig{
			ScriptDir:     "scripts/scenarios",
			SaveDir:       "saves",
			MaxSaveSlots:  9,
			AutoSaveSlot:  0,
			StartScene:    "main.lua",
			StartFunction: "start",
		},
		Display: DisplayConfig{
			DefaultMode:     "hybrid",
			TextSpeed:       2,
			AutoModeDelay:   3000,
			WindowWidth:     80,
			WindowHeight:    24,
			ColorTheme:      "blue",
			ShowLineNumbers: false,
		},
		Audio: AudioConfig{
			Enabled:      false,
			MasterVolume: 1.0,
			VoiceVolume:  1.0,
			BGMVolume:    0.8,
			SFXVolume:    0.8,
			VoiceDir:     "audio/voice",
			BGMDir:       "audio/bgm",
			SFXDir:       "audio/sfx",
		},
		Development: DevelopmentConfig{
			HotReload:     true,
			DebugMode:     false,
			LogLevel:      "info",
			ShowDebugInfo: false,
		},
	}
}

func (c *Config) GetDisplayMode() scenarioVO.DisplayMode {
	return scenarioVO.NewDisplayModeFromString(c.Display.DefaultMode)
}

func (c *Config) GetTextSpeed() valueobject.TextSpeed {
	return valueobject.NewTextSpeed(c.Display.TextSpeed)
}

func (c *Config) GetColorTheme() valueobject.ColorTheme {
	themes := valueobject.DefaultColorThemes()
	if theme, exists := themes[c.Display.ColorTheme]; exists {
		return theme
	}
	return themes["blue"]
}