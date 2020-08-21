package cli

import (
	"dam-video-injector-backend/apicall"
	"dam-video-injector-backend/videodl"
	"errors"
"github.com/urfave/cli/v2"
	"strconv"
)

var MissingArgumentError = errors.New("not enough argument error")

func GenerateClientApp() *cli.App {
	return &cli.App{
		// マニュアル https://github.com/urfave/cli/blob/master/docs/v2/manual.md
		Name:      "dam-photo-injector",
		Usage:     "manipulate DAM API from your PC and use youtube-dl to download videos",
		HideHelp:  true,

		Commands: []*cli.Command{
			{
				Name:    "search",
				Usage:   "search request codes",
				ArgsUsage: "SONG_TITLE",
				Action:  func(c *cli.Context) error {
					if c.NArg() != 1 {
						_ = cli.ShowAppHelp(c)
						return MissingArgumentError
					}
					apicall.SearchSongs(c.Args().Get(0))
					return nil
				},
			},
			{
				Name:    "connect",
				Usage:   "connect to DAM with qr code",
				Action:  func(c *cli.Context) error {
					if c.NArg() != 1 {
						_ = cli.ShowAppHelp(c)
						return MissingArgumentError
					}
					apicall.ConnectDevice(c.Args().Get(0))
					return nil
				},
			},
			{
				Name: "send",
				Usage: "send Song request",
				ArgsUsage: "QR_CODE REQUEST_NUM INTERRUPT",
				Action:  func(c *cli.Context) error {
					if c.NArg() != 3 {
						_ = cli.ShowAppHelp(c)
						return MissingArgumentError
					}
					i, _ := strconv.Atoi(c.Args().Get(2))
					apicall.SendRequest(c.Args().Get(0), c.Args().Get(1), i)
					return nil
				},
			},
			{
				Name: "remote",
				Usage: "send Remote Ids",
				ArgsUsage: "QR_CODE REMOTE_ID",
				Action:  func(c *cli.Context) error {
					if c.NArg() != 2 {
						_ = cli.ShowAppHelp(c)
						return MissingArgumentError
					}
					apicall.SendRemocon(c.Args().Get(0), c.Args().Get(1))
					return nil
				},
			},
			{
				Name: "download",
				Usage: "download video",
				ArgsUsage: "QUERY",
				Action:  func(c *cli.Context) error {
					if c.NArg() != 1 {
						_ = cli.ShowAppHelp(c)
						return MissingArgumentError
					}
					videodl.DownloadVideo(c.Args().Get(0))
					return nil
				},
			},
		},
	}
}

