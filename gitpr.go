package main

import (
	"fmt"
	"strings"
)

type Prefix struct {
   prefix string
   describe string
}

func main() {
    var prefixList[]Prefix
    prefixList = append(prefixList, Prefix{"fix", "バグを修正"})
    prefixList = append(prefixList, Prefix{"feat", "新規機能・新規ファイル追加"})
    prefixList = append(prefixList, Prefix{"update", "バグではない機能修正"})
    prefixList = append(prefixList, Prefix{"refactor", "リファクタリング"})
    prefixList = append(prefixList, Prefix{"disable", "コメントアウト等で無効化した"})
    prefixList = append(prefixList, Prefix{"delete", "ファイルを削除・コードの一部を取り除いた"})
    prefixList = append(prefixList, Prefix{"rename", "ファイル名を変更した"})
    prefixList = append(prefixList, Prefix{"move", "ファイルを移動した"})
    prefixList = append(prefixList, Prefix{"docs", "ドキュメントのみを修正した"})
    prefixList = append(prefixList, Prefix{"style", "空白・セミコロン・行・コーディングフォーマットなどの修正"})
    prefixList = append(prefixList, Prefix{"pref", "性能向上する修正をした"})
    prefixList = append(prefixList, Prefix{"test", "テスト追加や間違っていたテストの修正"})
    prefixList = append(prefixList, Prefix{"chore", "ビルドツールやライブラリで自動生成されたものをコミットした"})

    // prefixList = append(prefixList, Prefix{"hotfix", "クリティカルなバグを修正"})
    // prefixList = append(prefixList, Prefix{"add", "新規機能・新規ファイルを追加"})
    // prefixList = append(prefixList, Prefix{"change", "仕様変更による機能修正"})
    // prefixList = append(prefixList, Prefix{"clean", "リファクタリング"})
    // prefixList = append(prefixList, Prefix{"improve", "リファクタリング"})
    // prefixList = append(prefixList, Prefix{"remove", "ファイルを削除・コードの一部を取り除いた"})
    // prefixList = append(prefixList, Prefix{"upgrade", "バージョンアップした"})
    // prefixList = append(prefixList, Prefix{"revert", "修正を取り消した"})

    maxLen := 0
    for _, t := range prefixList {
        prefixLen := len(t.prefix)
        if (maxLen < prefixLen) {
            maxLen = prefixLen
        }
    }

    for _, t := range prefixList {
        show(t, maxLen)
    }
}
func show(p Prefix, maxLen int) {
    var line[]string
    line = append(line, fmt.Sprintf("\033[48;5;43m%s", p.prefix)) // 43が色
    line = append(line, fmt.Sprintf("\033[49m"))

    prefixLen := len(p.prefix)
    for i := 0; i < (maxLen - prefixLen); i++ {
        line = append(line, " ")
    }
    line = append(line, fmt.Sprintf("\033[49m: %s", p.describe))

    fmt.Println(strings.Join(line, ""))
}
