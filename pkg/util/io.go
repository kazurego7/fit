package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/kazurego7/fit/pkg/global"
)

var (
	queryCache = make(map[string]gitQueryResult)
)

type gitQueryResult struct {
	out      []byte
	exitCode int
	err      error
}

func GitCommand(globalFlag global.GlobalFlag, args []string) int {
	// 更新時にクエリ結果のキャッシュをクリア
	queryCache = make(map[string]gitQueryResult)

	// デバッグモードの場合は標準エラー出力にコマンドを表示
	extArgs := append([]string{"-c", "core.quotepath=false"}, args...)
	if global.RootFlag.Debug {
		if globalFlag.Dryrun {
			fmt.Fprintln(os.Stderr, "dry-run: git "+strings.Join(extArgs, " "))
		} else {
			fmt.Fprintln(os.Stderr, "command: git "+strings.Join(extArgs, " "))
		}
	}

	// ドライランの場合は実行せずに終了コード 0 を返す
	if globalFlag.Dryrun {
		return 0
	}

	// 実行して終了コードを返す
	cmd := exec.Command("git", extArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	return cmd.ProcessState.ExitCode()
}

func GitQuery(globalFlag global.GlobalFlag, args []string) ([]byte, int, error) {
	// デバッグモードの場合は標準エラー出力にクエリを表示
	extArgs := append([]string{"-c", "core.quotepath=false"}, args...)
	extArgsString := strings.Join(extArgs, " ")
	cache, cacheExists := queryCache[extArgsString]
	if global.RootFlag.Debug {
		queryResult := "query(execute): git " + extArgsString
		if cacheExists {
			queryResult = "query(cache): git " + extArgsString
		}
		fmt.Fprintln(os.Stderr, queryResult)
	}

	// クエリ結果のキャッシュがあればそれを返す
	if cacheExists {
		return cache.out, cache.exitCode, cache.err
	}

	// クエリ結果のキャッシュがなければクエリを実行してキャッシュに保存
	cmd := exec.Command("git", extArgs...)
	out, err := cmd.Output()
	exitCode := cmd.ProcessState.ExitCode()
	queryCache[extArgsString] = gitQueryResult{out, exitCode, err}
	return out, exitCode, err
}

func InputTextLn() (string, error) {
	var ans string
	_, err := fmt.Scanf("%s\n", &ans)
	if err != nil {
		return "", err
	}
	return ans, nil
}

func InputYesOrNo(allwaysYes bool) (bool, error) {
	if allwaysYes {
		return true, nil
	}
	for {
		ans, err := InputTextLn()
		if err != nil {
			return false, err
		}
		switch ans {
		case "Yes", "Y", "yes", "y":
			return true, nil
		case "No", "N", "no", "n":
			return false, nil
		default:
			fmt.Println(`"yes" か "no" で入力してください`)
			continue
		}
	}
}
