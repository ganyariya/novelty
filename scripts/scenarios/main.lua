-- メインシナリオエントリーポイント

function start()
    set_display_mode("hybrid")
    
    narration("これは、ターミナルで動作するビジュアルノベルゲームエンジンのデモです。")
    
    text("システム", "ようこそ、Terminal Novel Game Engineへ！")
    
    narration("このエンジンは以下の機能をサポートしています：")
    narration("・ADV、NVL、Hybridの3つの表示モード")
    narration("・タイプライター効果によるテキスト表示")
    narration("・Luaスクリプトによるシナリオ制御")
    narration("・フラグと変数による分岐処理")
    
    text("システム", "では、簡単なデモを始めましょう。")
    
    set_flag("demo_started", true)
    set_var("demo_step", 1)
    
    jump_to_label("demo_adv_mode")
end

function demo_adv_mode()
    set_display_mode("adv")
    
    text("システム", "これはADVモードです。")
    text("システム", "ゲームによくある、下部にメッセージウィンドウが表示される形式ですね。")
    
    narration("上部には背景や立ち絵が表示される予定です。")
    
    add_var("demo_step", 1)
    
    jump_to_label("demo_nvl_mode")
end

function demo_nvl_mode()
    set_display_mode("nvl")
    
    nvl_text("これはNVLモードです。")
    nvl_text("画面全体にテキストが表示され、小説を読んでいるような感覚になります。")
    nvl_text("長い地の文や心理描写に適した表示形式です。")
    
    text("システム", "月姫やひぐらしのなく頃になどで使われている形式ですね。")
    
    add_var("demo_step", 1)
    
    jump_to_label("demo_hybrid_mode")
end

function demo_hybrid_mode()
    set_display_mode("hybrid")
    
    text("システム", "そして、これがHybridモードです。")
    text("システム", "上部にバックログ、下部に現在のテキストが表示されます。")
    
    narration("過去のテキストを確認しながら、現在の内容も見やすい形式です。")
    
    add_var("demo_step", 1)
    
    jump_to_label("demo_controls")
end

function demo_controls()
    text("システム", "操作方法を説明します。")
    text("システム", "Enterキーまたはスペースキーで次のテキストに進みます。")
    text("システム", "Aキーでオートモード、Sキーでスキップモードを切り替えられます。")
    text("システム", "1、2、3キーで表示モードを変更できます。")
    text("システム", "Ctrl+1-4でテキスト速度を変更できます。")
    text("システム", "Dキーでデバッグ情報の表示を切り替えられます。")
    
    jump_to_label("demo_flags_vars")
end

function demo_flags_vars()
    text("システム", "フラグと変数のデモです。")
    
    local step = get_var("demo_step")
    text("システム", "現在のデモステップ: " .. step)
    
    local started = get_flag("demo_started")
    if started then
        text("システム", "デモ開始フラグが立っています。")
    else
        text("システム", "デモ開始フラグが立っていません。")
    end
    
    set_var("demo_complete", 1)
    
    jump_to_label("demo_end")
end

function demo_end()
    text("システム", "デモは以上です。")
    text("システム", "このエンジンはまだ開発中ですが、基本的な機能は動作しています。")
    
    narration("今後、選択肢システム、セーブ・ロード機能、キャラクター管理システムなどが追加される予定です。")
    
    text("システム", "Qキーで終了できます。お疲れ様でした！")
    
    jump_to_label("loop_end")
end

function loop_end()
    text("システム", "デモループです。Enterで繰り返し、Qで終了してください。")
    -- 無限ループを避けるために、ここでは何もしない
end