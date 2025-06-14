# ターミナルノベルゲームエンジン要件仕様書

## 概要
ターミナル環境で動作するビジュアルノベルゲームエンジンを Go で実装する。
AIによる開発を前提とし、後にUnityなどのグラフィカル環境への移植を想定している。

## 技術スタック

### 開発言語・環境
- **Go**: 静的型付け、後方互換性、AIでの開発しやすさ
- **gopher-lua**: シナリオスクリプトをLuaで記述
- **TUIライブラリ**: 
  - `github.com/charmbracelet/bubbletea` (推奨)
  - `github.com/charmbracelet/lipgloss` (スタイリング)
  - `github.com/charmbracelet/bubbles` (コンポーネント)

### アーキテクチャ
- **DDD (Domain Driven Design)**: ドメインロジックを明確に分離
- **テスト駆動開発**: `go test` で容易にイテレーション
- **Go ベストプラクティス準拠**

## 機能要件

### 1. コアゲームシステム

#### 1.1 テキスト表示システム
- **文字送り機能**
  - 1文字ずつのタイプライター効果
  - 速度調整可能（設定で変更）
  - スキップ機能（既読・未読問わず）

- **テキスト表示制御**
  - クリック（Enter）で1行分を即座に表示
  - 表示完了後のクリックで次テキストへ進行
  - バックログ機能（過去のテキスト履歴）

#### 1.2 選択肢システム
- 複数選択肢の表示
- キーボードナビゲーション（矢印キー + Enter）
- 選択結果による分岐処理
- 選択肢の条件分岐（フラグ依存）

#### 1.3 キャラクター識別システム
- **キャラクター名表示**
  - 話者名を明確に表示
  - キャラクター別の色分け（ANSIカラー）
  - キャラクター設定（名前、色、属性）

- **ターミナル向けキャラ表現**
  - ASCII Art アイコン（オプション）
  - 文字装飾（太字、斜体、色）
  - キャラ別のテキストスタイル

#### 1.4 画面表示形式システム
- **ハイブリッド方式（推奨）**
  ```
  ┌─────────────────────────────────────────────────────┐
  │ [バックログ領域] - 過去のテキストがスクロール表示    │
  │ 彼女は振り返ると、微笑みかけた。                    │
  │ 「こんにちは」                                      │
  │ その声は、春の風のように優しかった。                │
  │                                                     │
  ├─────────────────────────────────────────────────────┤
  │ [現在表示中テキスト] - 文字送り中                   │
  │ 主人公: 「君の名前は？」█                           │
  │                                                     │
  ├─────────────────────────────────────────────────────┤
  │ [ステータス] Auto: OFF | Skip: OFF | [S]ave [L]oad │
  └─────────────────────────────────────────────────────┘
  ```

- **NVL形式（Novel形式）**
  - 画面全体にテキストが流れるように表示
  - 長い地の文、心理描写に最適
  - 小説的な没入感を提供
  - 画面いっぱいになったら自動クリア
  - 代表例：月姫、ひぐらしのなく頃に

- **ADV形式（Adventure形式）**
  - メッセージウィンドウでセリフを1つずつ表示
  - 会話シーンに最適
  - キャラクター名の明確な表示
  - クリック毎に次のセリフに進行
  - 代表例：一般的なギャルゲー、RPG会話

- **表示形式の切り替え**
  - Luaスクリプトから表示形式を指定可能
  - ゲーム設定での既定形式選択
  - シーンに応じた自動切り替え対応

#### 1.5 セーブ・ロードシステム
- **セーブデータ**
  - JSON形式でのデータ保存
  - 複数スロット対応（推奨：9スロット）
  - セーブ日時の記録
  - プレイ時間の記録

- **セーブ内容**
  - 現在のシーン位置
  - フラグ・変数の状態
  - 既読テキストの記録
  - キャラクター好感度等
  - 現在の表示形式状態

### 2. 高度なノベルゲーム機能

#### 2.1 キネマティックノベル対応
- **月姫スタイル（NVL形式）**
  - 地の文と会話の明確な区別
  - 状況描写の詳細表示
  - 心理描写の表現
  - 画面全体を使った文章表示
  
- **ADV形式対応**
  - ウィンドウベースの会話表示
  - キャラクター中心の対話シーン
  - 選択肢との自然な連携

- **形式間の切り替え**
  - Luaスクリプトでの動的切り替え
  - シーンの性質に応じた最適な表示
  - ユーザー設定での既定形式選択

#### 2.2 表示形式別ウィンドウシステム
- **ハイブリッド形式レイアウト**
  - 上部：バックログ領域（スクロール表示）
  - 中部：現在テキスト表示領域（文字送り）
  - 下部：ステータス・操作情報
  
- **NVL形式レイアウト**
  - 画面全体がテキスト表示領域
  - 上から下へ順次テキスト追加
  - 画面満杯時の自動クリア機能
  
- **ADV形式レイアウト**
  - 上部：背景・状況表示領域
  - 下部：固定メッセージウィンドウ
  - キャラクター名とセリフの明確な区分

- **共通UI要素**
  - ボーダー装飾（形式に応じて調整）
  - テキストの自動改行
  - スクロール対応
  - 色分けによる視認性向上

#### 2.3 スキップ・オート機能
- **スキップモード**
  - 既読テキストの高速スキップ
  - 未読テキストのスキップ（オプション）
  - 選択肢での一時停止

- **オートモード**
  - 自動テキスト送り
  - 待機時間の設定可能

### 3. システム機能

#### 3.1 設定システム
- テキスト速度調整
- 音量設定（将来的な音響対応）
- ウィンドウサイズ対応
- キーバインド設定
- **既定表示形式設定**（NVL/ADV/ハイブリッド）
- 色テーマ設定

#### 3.2 音声システム（設計のみ）
- **音声再生機能**
  - WAV/MP3ファイルの再生対応
  - キャラクターボイス再生
  - BGM・効果音の再生
  - クロスプラットフォーム対応設計

- **音声制御**
  - 音量調整（マスター・ボイス・BGM・SE）
  - 音声スキップ対応
  - 音声とテキストの同期制御
  - 音声ファイルの非同期読み込み

- **実装優先度**
  - Phase 1: 設計のみ（音声なし実装）
  - Phase 2: システム効果音
  - Phase 3: 音声ファイル再生
  - Phase 4: 高度な音声制御

#### 3.3 デバッグ・開発支援
- Luaスクリプトのホットリロード
- デバッグコンソール
- フラグ・変数の確認機能
- シーン間ジャンプのデバッグ機能

## 技術仕様

### 1. DDD設計によるプロジェクト構造

```
cmd/
  game/
    main.go
internal/
  domain/
    scenario/
      entity/
        scenario.go          # Scenario Entity
        scene.go             # Scene Entity  
        character.go         # Character Entity
        choice.go            # Choice Entity
        message.go           # Message Entity
      valueobject/
        scene_id.go          # SceneID Value Object
        character_id.go      # CharacterID Value Object
        display_mode.go      # DisplayMode Value Object
        jump_target.go       # JumpTarget Value Object
      repository/
        scenario_repository.go    # ScenarioRepository Interface
        character_repository.go   # CharacterRepository Interface
      service/
        scenario_service.go       # ScenarioService
        jump_service.go          # JumpService
        choice_service.go        # ChoiceService
    game/
      entity/
        game_state.go        # GameState Entity
        save_data.go         # SaveData Entity
        player_progress.go   # PlayerProgress Entity
      valueobject/
        save_slot.go         # SaveSlot Value Object
        game_flags.go        # GameFlags Value Object
        variables.go         # Variables Value Object
      repository/
        save_repository.go   # SaveRepository Interface
        progress_repository.go # ProgressRepository Interface
      service/
        game_service.go      # GameService
        save_service.go      # SaveService
    presentation/
      entity/
        display_state.go     # DisplayState Entity
        text_renderer.go     # TextRenderer Entity
        window_layout.go     # WindowLayout Entity
      valueobject/
        text_speed.go        # TextSpeed Value Object
        color_theme.go       # ColorTheme Value Object
      service/
        text_display_service.go  # TextDisplayService
        input_handler_service.go # InputHandlerService
  infrastructure/
    lua/
      scenario_engine.go    # Lua実行エンジン
      function_registry.go  # Lua関数登録
      script_loader.go      # スクリプトローダー
    tui/
      game_renderer.go      # TUIレンダラー
      input_manager.go      # 入力管理
      layout_manager.go     # レイアウト管理
    persistence/
      json_save_repository.go    # JSONセーブリポジトリ実装
      file_scenario_repository.go # ファイルシナリオリポジトリ実装
    audio/
      audio_engine.go       # 音声エンジン（設計のみ）
  application/
    usecase/
      game_usecase.go       # ゲーム進行ユースケース
      save_usecase.go       # セーブ・ロードユースケース
      scenario_usecase.go   # シナリオ実行ユースケース
      choice_usecase.go     # 選択肢処理ユースケース
  presentation/
    tui/
      game_screen.go        # ゲーム画面
      menu_screen.go        # メニュー画面
      save_screen.go        # セーブ・ロード画面
      choice_screen.go      # 選択肢画面
pkg/
  config/
    config.go
  errors/
    domain_errors.go
scripts/
  scenarios/
    main.lua              # エントリーポイント
    common/
      characters.lua      # キャラクター定義
      system.lua          # システム関数
    chapter01/
      scene01.lua
      scene02.lua
    chapter02/
      scene01.lua
test/
  domain/
    scenario/
      entity/
        scenario_test.go
        scene_test.go
      service/
        scenario_service_test.go
        jump_service_test.go
  integration/
    scenario_execution_test.go
    save_load_test.go
  lua/
    script_execution_test.go
```

### 2. シナリオ構造・ファイル管理

#### 2.1 シナリオファイル構造
```
scenarios/
├── main.lua              # メインシナリオ（エントリーポイント）
├── chapter01/            # チャプター1
│   ├── scene01.lua      # オープニングシーン
│   ├── scene02.lua      # 学校シーン  
│   └── scene03.lua      # 放課後シーン
├── chapter02/            # チャプター2
│   ├── scene01.lua      # 次の日
│   └── scene02.lua      # メインイベント
├── common/               # 共通定義
│   ├── characters.lua   # キャラクター定義
│   ├── system.lua       # システム関数
│   └── config.lua       # ゲーム設定
└── endings/              # エンディング分岐
    ├── good_end.lua
    ├── bad_end.lua
    └── true_end.lua
```

#### 2.2 ノベルゲーム用語体系
- **プロジェクト**: ゲーム全体
- **チャプター**: 大きな章分け（ディレクトリ単位）
- **シーン**: チャプター内の場面（ファイル単位）
- **ラベル**: シーン内の特定位置（@マーク）
- **メッセージ**: 個々のテキスト・コマンド
- **メッセージウィンドウ**: テキスト表示領域
- **バックログ**: 過去のテキスト履歴
- **システムメニュー**: セーブ・ロード・設定画面
- **選択肢ウィンドウ**: 分岐選択UI

#### 2.2 DDD Entity設計例

```go
// ===== Domain Entities =====

// Scenario Entity - シナリオ全体を表現
type Scenario struct {
    id          ScenarioID
    title       string
    scenes      map[SceneID]*Scene
    characters  map[CharacterID]*Character
    startScene  SceneID
}

// Scene Entity - 個別シーンを表現  
type Scene struct {
    id          SceneID
    filePath    string
    functions   map[string]lua.LFunction  // Lua関数のマッピング
    loaded      bool
}

// Character Entity - キャラクターを表現
type Character struct {
    id          CharacterID
    name        string
    displayName string
    colorTheme  ColorTheme
    voiceDir    string
}

// Choice Entity - 選択肢を表現
type Choice struct {
    id          ChoiceID
    text        string
    condition   string
    callback    lua.LFunction
    jumpTarget  *JumpTarget
}

// Message Entity - 表示メッセージを表現
type Message struct {
    id          MessageID
    speakerID   CharacterID
    content     string
    displayMode DisplayMode
    voiceFile   string
}

// GameState Entity - ゲーム状態を表現
type GameState struct {
    currentScene    SceneID
    currentFunction string
    flags          GameFlags
    variables      Variables
    history        []Message
    saveSlots      map[SaveSlot]*SaveData
}

// ===== Value Objects =====

// SceneID Value Object
type SceneID struct {
    value string
}

func NewSceneID(chapter, scene string) SceneID {
    return SceneID{value: fmt.Sprintf("%s/%s", chapter, scene)}
}

// JumpTarget Value Object
type JumpTarget struct {
    sceneID      SceneID
    functionName string
}

func NewJumpTarget(scenePath, funcName string) *JumpTarget {
    return &JumpTarget{
        sceneID:      SceneID{value: scenePath},
        functionName: funcName,
    }
}

// DisplayMode Value Object
type DisplayMode int

const (
    DisplayModeADV DisplayMode = iota
    DisplayModeNVL
    DisplayModeHybrid
)

// ===== Repository Interfaces =====

type ScenarioRepository interface {
    LoadScene(sceneID SceneID) (*Scene, error)
    LoadCharacters() (map[CharacterID]*Character, error)
    ExecuteSceneFunction(scene *Scene, functionName string, state *GameState) error
}

type SaveRepository interface {
    Save(slot SaveSlot, data *SaveData) error
    Load(slot SaveSlot) (*SaveData, error)
    ListSaves() ([]SaveSlot, error)
}

// ===== Domain Services =====

type ScenarioService struct {
    repo ScenarioRepository
}

func (s *ScenarioService) ExecuteJump(target *JumpTarget, state *GameState) error {
    scene, err := s.repo.LoadScene(target.sceneID)
    if err != nil {
        return err
    }
    
    state.currentScene = target.sceneID
    state.currentFunction = target.functionName
    
    return s.repo.ExecuteSceneFunction(scene, target.functionName, state)
}

type JumpService struct {
    scenarioService *ScenarioService
}

func (j *JumpService) JumpToScene(scenePath, functionName string, state *GameState) error {
    target := NewJumpTarget(scenePath, functionName)
    return j.scenarioService.ExecuteJump(target, state)
}

func (j *JumpService) JumpToLabel(functionName string, state *GameState) error {
    target := &JumpTarget{
        sceneID:      state.currentScene,
        functionName: functionName,
    }
    return j.scenarioService.ExecuteJump(target, state)
}
```

#### 2.3 Lua-Go連携の実装パターン

```go
// Lua関数をGoから呼び出すためのブリッジ
type LuaBridge struct {
    state         *lua.LState
    gameState     *GameState
    jumpService   *JumpService
    choiceService *ChoiceService
}

func (b *LuaBridge) RegisterFunctions() {
    // シナリオ制御関数
    b.state.SetGlobal("jump_to_scene", b.state.NewFunction(b.jumpToScene))
    b.state.SetGlobal("jump_to_label", b.state.NewFunction(b.jumpToLabel))
    b.state.SetGlobal("text", b.state.NewFunction(b.displayText))
    b.state.SetGlobal("choice", b.state.NewFunction(b.showChoice))
    
    // ゲーム状態管理関数
    b.state.SetGlobal("set_flag", b.state.NewFunction(b.setFlag))
    b.state.SetGlobal("get_flag", b.state.NewFunction(b.getFlag))
    b.state.SetGlobal("set_var", b.state.NewFunction(b.setVariable))
    b.state.SetGlobal("get_var", b.state.NewFunction(b.getVariable))
}

func (b *LuaBridge) jumpToScene(L *lua.LState) int {
    scenePath := L.CheckString(1)
    functionName := L.CheckString(2)
    
    err := b.jumpService.JumpToScene(scenePath, functionName, b.gameState)
    if err != nil {
        L.Push(lua.LBool(false))
        L.Push(lua.LString(err.Error()))
        return 2
    }
    
    L.Push(lua.LBool(true))
    return 1
}
```

### 2. Luaスクリプト仕様

#### 2.1 基本構文・ジャンプシステム（gopher-lua対応）
```lua
-- ===== common/characters.lua =====
-- キャラクター定義
register_character("hero", {
    name = "主人公", 
    color = "blue", 
    voice_dir = "hero/"
})
register_character("heroine", {
    name = "ヒロイン", 
    color = "red", 
    voice_dir = "heroine/"
})

-- ===== chapter01/scene01.lua =====
-- ファイル内でのラベル定義（Luaの関数として）
function start()
    set_display_mode("adv")
    text("hero", "今日は学校に行こう")
    
    choice({
        {text = "教室に行く", callback = function()
            jump_to_scene("chapter01/scene02.lua", "classroom")
        end},
        {text = "屋上に行く", callback = function()
            jump_to_scene("chapter01/scene02.lua", "rooftop")
        end}
    })
end

function bad_route()
    set_display_mode("nvl")
    nvl_text("運命の分かれ道だった...")
    jump_to_scene("endings/bad_end.lua", "start")
end

-- ===== chapter01/scene02.lua =====
function classroom()
    text("hero", "教室に到着した")
    set_var("location", "classroom")
    jump_to_label("after_school")
end

function rooftop()
    text("hero", "屋上は気持ちいい")
    set_var("location", "rooftop")
    add_var("heroine_affection", 5)
    jump_to_label("after_school")
end

function after_school()
    narration("放課後になった。")
    
    if get_var("heroine_affection") >= 50 then
        jump_to_scene("chapter02/scene01.lua", "good_path")
    else
        jump_to_scene("chapter02/scene01.lua", "normal_path")
    end
end

-- ===== システム関数例 =====
-- 条件付きジャンプ
function conditional_jump()
    if get_flag("talked_to_heroine") and get_var("day") >= 2 then
        jump_to_scene("chapter02/scene01.lua", "morning")
    elseif get_var("heroine_affection") < 30 then
        jump_to_label("bad_route")
    else
        jump_to_label("normal_continue")
    end
end

-- 音声付きテキスト（設計のみ）
function voice_text(character_id, text_content, voice_file)
    if voice_file then
        play_voice(voice_file)
    end
    text(character_id, text_content)
end

-- 複雑な選択肢（条件付き）
function complex_choice()
    local choices = {}
    
    -- 条件に応じて選択肢を動的生成
    if get_var("courage") >= 3 then
        table.insert(choices, {
            text = "告白する", 
            callback = function()
                jump_to_scene("endings/good_end.lua", "confession")
            end
        })
    end
    
    table.insert(choices, {
        text = "普通に話す",
        callback = function()
            jump_to_label("normal_talk")
        end
    })
    
    if get_flag("has_gift") then
        table.insert(choices, {
            text = "プレゼントを渡す",
            callback = function()
                set_flag("gave_gift", true)
                add_var("heroine_affection", 10)
                jump_to_label("gift_scene")
            end
        })
    end
    
    choice(choices)
end
```

### 3. テスト戦略

#### 3.1 単体テスト
- ドメインロジックの徹底的なテスト
- Luaエンジンの動作テスト
- テキスト処理ロジックのテスト

#### 3.2 統合テスト
- シナリオ実行の端到端テスト
- セーブ・ロード機能のテスト
- TUIコンポーネントの動作テスト

#### 3.3 テストファイル例
```go
// internal/domain/game/service/game_test.go
func TestGameService_ProcessText(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected TextDisplayInfo
    }{
        {
            name:  "通常テキスト",
            input: "こんにちは",
            expected: TextDisplayInfo{
                Text: "こんにちは",
                DisplayDuration: 5 * time.Second,
            },
        },
    }
    // テスト実装
}
```

## 開発フェーズ

### Phase 1: コア実装
1. ドメインモデル設計
2. Luaエンジン統合
3. 基本的なテキスト表示

### Phase 2: ゲーム機能
1. 選択肢システム
2. セーブ・ロード
3. キャラクター管理

### Phase 3: UI/UX改善
1. TUI向け最適化
2. アニメーション効果
3. パフォーマンス最適化

### Phase 4: 高度な機能
1. キネマティックノベル対応
2. デバッグ機能
3. 設定システム

## 成功指標
- `go test ./...` で全テストが通る
- 実際のノベルゲームシナリオが動作する
- Unityへの移植時にコアロジックが再利用できる
- AI開発者が容易に機能追加できる設計
