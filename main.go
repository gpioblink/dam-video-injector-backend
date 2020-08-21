package main

import (
	"dam-video-injector-backend/cli"
	"fmt"
	"log"
	"os"
)



func main() {
	fmt.Println("dam-video-injector-backend")

	//apicall.ConnectDevice("0a33d3302039dc346d41174d1c315f30")
	//apicall.SendRequest("0a33d3302039dc346d41174d1c315f30","373085", 0)
	//apicall.SendRemocon("0a33d3302039dc346d41174d1c315f30", "F1")
	//apicall.SearchSongs("デート・ア・ライブ")

	//videodl.DownloadVideo("君のせい")

	app := cli.GenerateClientApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}