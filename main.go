/*
References:
-----------
https://github.com/gopxl/pixel
*/
package main

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image"
	_ "image/png"
	"os"
)

const (
	WinWidth  = 800
	WinHeight = 600
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Gome",
		Bounds: pixel.R(0, 0, 800, 600),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// load image
	pic, err := loadPicture("res/player.png")
	if err != nil {
		panic(err)
	}

	// mainloop
	sprite := pixel.NewSprite(pic, pic.Bounds())
	mat := pixel.IM
	mat = mat.Moved(win.Bounds().Center())

	win.Clear(colornames.Skyblue)
	sprite.Draw(win, mat)
	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}
