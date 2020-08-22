package videodl

import (
	"fmt"
	"github.com/mattn/go-shellwords"
	"os/exec"
)

func ChangeVideo(reqId string) {
	changeQuery := fmt.Sprintf("./kplay %s", reqId)
	_ = runCmdStr(changeQuery)
}

func DownloadVideo(query string) {
	searchQuery := fmt.Sprintf("youtube-dl -f mp4 ytsearch1:%s --output /mnt/virtual_sd/video.mp4", query)
	_ = runCmdStr(searchQuery)
}

// https://qiita.com/tanksuzuki/items/9205ff70c57c4c03b5e5
func runCmdStr(cmdstr string) error {
	// 文字列をコマンド、オプション単位でスライス化する
	c, err := shellwords.Parse(cmdstr)
	if err != nil {
		return err
	}
	switch len(c) {
	case 0:
		// 空の文字列が渡された場合
		return nil
	case 1:
		// コマンドのみを渡された場合
		err = exec.Command(c[0]).Run()
	default:
		// コマンド+オプションを渡された場合
		// オプションは可変長でexec.Commandに渡す
		err = exec.Command(c[0], c[1:]...).Run()
	}
	if err != nil {
		return err
	}
	return nil
}