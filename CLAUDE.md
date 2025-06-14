# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

**重要：このプロジェクトに関する質問や説明は日本語で回答してください。**

## Project Overview

This is a terminal-based visual novel game engine written in Go. The project follows Domain-Driven Design (DDD) principles and is designed for AI-assisted development with future Unity migration in mind.

## Key Technologies

- **Go**: Primary development language
- **gopher-lua**: Lua scripting engine for scenario scripts
- **bubbletea/lipgloss/bubbles**: TUI libraries for terminal interface
- **DDD Architecture**: Domain-driven design with clear separation of concerns

## Commands

```bash
# Build the game
go build -o bin/game cmd/game/main.go

# Run tests
go test ./...

# Run specific test package
go test ./internal/domain/scenario/...

# Format code
go fmt ./...

# Check for issues
go vet ./...

# Run the game (after building)
./bin/game
```

## Architecture

### DDD Structure
- `internal/domain/`: Core business logic (entities, value objects, repositories, services)
  - `scenario/`: Scenario management, scenes, characters, choices
  - `game/`: Game state, save data, player progress
  - `presentation/`: Display state, text rendering, window layouts
- `internal/infrastructure/`: External dependencies (Lua engine, TUI, persistence, audio)
- `internal/application/`: Use cases orchestrating domain services
- `cmd/`: Application entry points
- `pkg/`: Shared utilities and configuration

### Key Entities
- **Scenario**: Complete game scenario with scenes and characters
- **Scene**: Individual game scenes (mapped to Lua files)
- **Character**: Game characters with display properties
- **GameState**: Current game state including flags, variables, and history
- **Message**: Text messages with speaker and display mode information

### Display Modes
The engine supports three display modes:
- **ADV Mode**: Traditional adventure game style with message windows
- **NVL Mode**: Novel-style full-screen text flow
- **Hybrid Mode**: Combined approach with backlog and current text areas

## Lua Scripting System

### Scenario Structure
```
scripts/scenarios/
├── main.lua              # Entry point
├── chapter01/            # Chapter directories
│   ├── scene01.lua      # Scene files
│   └── scene02.lua
├── common/               # Shared definitions
│   ├── characters.lua   # Character definitions
│   └── system.lua       # System functions
└── endings/              # Ending variations
```

### Core Lua Functions
- `text(character_id, content)`: Display character dialogue
- `narration(content)`: Display narrative text
- `nvl_text(content)`: Display text in NVL mode
- `choice({options})`: Present player choices
- `jump_to_scene(path, function)`: Jump to different scene
- `jump_to_label(function)`: Jump within current scene
- `set_flag(name, value)` / `get_flag(name)`: Flag management
- `set_var(name, value)` / `get_var(name)`: Variable management
- `add_var(name, value)`: Add to variable value
- `set_display_mode(mode)`: Change display mode ("adv", "nvl", "hybrid")
- `register_character(id, config)`: Register character with name, color, voice_dir
- `play_voice(file)`: Play voice file (planned feature)
- `voice_text(character_id, content, voice_file)`: Text with voice (planned)

### Scene Function Pattern
Each Lua scene file contains functions that represent labels/entry points:
```lua
function start()
    -- Scene entry point
end

function custom_label()
    -- Named jump target
end
```

## Game Features

### Text Display System
- **Typewriter effect**: Character-by-character text display with adjustable speed
- **Text control**: Click/Enter to display full line, advance to next text
- **Backlog**: History of previous text messages
- **Skip mode**: Fast-forward through read/unread text
- **Auto mode**: Automatic text advancement with configurable timing

### Save/Load System
- **Multiple save slots**: JSON-based save data (recommended 9 slots)
- **Save content**: Scene position, flags, variables, read history, play time
- **Quick save/load**: Fast save and restore functionality

### Choice System
- **Multiple options**: Keyboard navigation with arrow keys + Enter
- **Conditional choices**: Options based on flags and variables
- **Branching logic**: Story branches based on player decisions

### Character System
- **Character registration**: Name, display name, color theme, voice directory
- **Color coding**: ANSI color support for character differentiation
- **ASCII art icons**: Optional character visual representation

### Audio System (Planned)
- **Voice playback**: Character voice files (WAV/MP3)
- **BGM/SFX**: Background music and sound effects
- **Volume control**: Master, voice, BGM, SFX volume settings
- **Audio sync**: Voice and text synchronization

### Settings System
- **Text speed**: Adjustable typewriter speed
- **Default display mode**: ADV/NVL/Hybrid preference
- **Key bindings**: Customizable keyboard shortcuts
- **Color themes**: UI color customization
- **Window size**: Terminal size adaptation

## Development Guidelines

### File Organization
- Follow DDD layering strictly - domain logic should not depend on infrastructure
- Each domain aggregate gets its own subdirectory
- Interfaces defined in domain layer, implementations in infrastructure
- Use dependency injection for repository and service dependencies

### Lua Integration
- Lua scripts are loaded dynamically from `scripts/scenarios/`
- Go-Lua bridge functions are registered in `infrastructure/lua/`
- All game state changes from Lua go through domain services
- Scene transitions are handled by the jump service

### Testing Strategy
- Unit tests for all domain entities and services
- Integration tests for Lua script execution
- Test files follow `*_test.go` naming convention
- Mock implementations for repositories in tests

### Error Handling
- Domain errors defined in `pkg/errors/domain_errors.go`
- Lua execution errors are wrapped and propagated to UI layer
- Save/load operations include comprehensive error handling

## Key Design Patterns

### Repository Pattern
All data access goes through repository interfaces defined in domain layer with implementations in infrastructure layer.

### Service Layer
Domain services coordinate between entities and handle complex business logic that doesn't belong in a single entity.

### Value Objects
Immutable objects like SceneID, CharacterID, and JumpTarget encapsulate domain concepts and provide type safety.

### Bridge Pattern
Lua-Go integration uses bridge pattern to expose Go functionality to Lua scripts while maintaining clean separation.

## Development Phases

### Phase 1: Core Implementation
1. Domain model design
2. Lua engine integration
3. Basic text display system

### Phase 2: Game Features
1. Choice system implementation
2. Save/load functionality
3. Character management system

### Phase 3: UI/UX Enhancement
1. TUI optimization
2. Animation effects
3. Performance optimization

### Phase 4: Advanced Features
1. Kinematic novel support (NVL/ADV switching)
2. Debug functionality
3. Settings system
4. Audio system preparation

## Success Criteria
- All tests pass with `go test ./...`
- Actual visual novel scenarios run successfully
- Core logic is reusable for Unity migration
- AI developers can easily add features to the codebase

## Technical Notes

### Kinematic Novel Support
- **月姫-style NVL**: Full-screen narrative text with detailed descriptions
- **ADV-style dialogue**: Window-based conversation display
- **Dynamic switching**: Lua script control over display modes
- **Hybrid approach**: Combined backlog and current text areas

### Performance Considerations
- Lua script hot-reloading for development
- Efficient text rendering with minimal redraws
- Memory management for save data and text history
- Cross-platform terminal compatibility

### Debugging Features
- Lua script hot-reload during development
- Debug console for flag/variable inspection
- Scene jump functionality for testing
- Comprehensive error reporting and logging