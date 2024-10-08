package main

import (
	"image"
	"os"

	_ "image/png"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"golang.org/x/image/colornames"
)

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

func run() {
	cfg := opengl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := opengl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	// imd := imdraw.New(nil)

	// filledRect(imd, 100, 100, 200, 200, pixel.RGB(0, 1, 0), pixel.RGB(1, 0, 0))

	spritesheet, err := loadPicture("spritesheet.png")
	if err != nil {
		panic(err)
	}

	size := float64(312)
	sprites := []*pixel.Sprite{
		pixel.NewSprite(spritesheet, pixel.R(8, 400+size+22, 8+size, 42+3*size+22)),
		pixel.NewSprite(spritesheet, pixel.R(8, 400, 8+size, 42+2*size)),
		pixel.NewSprite(spritesheet, pixel.R(8, 42, 8+size, 42+size)),
	}

	for !win.Closed() {
		win.Clear(colornames.Whitesmoke)

		// delta := float64(280)
		p := win.Bounds().Center()
		sprites[0].Draw(win, pixel.IM.Moved(p))
		// p = p.Add(pixel.V(0, -delta))
		sprites[1].Draw(win, pixel.IM.Moved(p))
		// p = p.Add(pixel.V(0, delta*2))
		sprites[2].Draw(win, pixel.IM.Moved(p))
		// imd.Draw(win)

		win.Update()
	}
}

// func filledRect(imd *imdraw.IMDraw, x, y, w, h float64, fill_color, stroke_color pixel.RGBA) {
// 	imd.Color = fill_color
// 	imd.Push(pixel.V(x, y))
// 	imd.Push(pixel.V(x+w, y+h))
// 	imd.Rectangle(0)

// 	imd.Color = stroke_color
// 	imd.Push(pixel.V(x, y))
// 	imd.Push(pixel.V(x+w, y+h))
// 	imd.Rectangle(2)
// }

func main() {
	opengl.Run(run)
}
