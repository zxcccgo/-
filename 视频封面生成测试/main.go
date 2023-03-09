package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	inPath:=".\\public\\testvideo.mp4"
	GetSnapShot(inPath,30)
}

func GetSnapShot(inPath string, frameNum int) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	
	img, err := imaging.Decode(buf)
	if err != nil {
		panic(err)
	}
	name := strings.Split(inPath, "/")
	path := "./" + name[len(name)-1]+".jpeg"
	if err := imaging.Save(img, path); err != nil {
		log.Println("here")
		panic(err)
	}

}
