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

	sprites := make([]pixel.Picture, 0)
	for _, sprite := range []string{
		"01-Apple.png",
		"02-Banana.png",
		"03-Blueberry.png",
		"04-Orange.png",
		"05-Strawberry.png",
		"06-Watermelon.png",
		"07-Seven.png",
	} {
		var pic pixel.Picture
		if pic, err = loadPicture("Reel-Images/" + sprite); err != nil {
			panic(err)
		}
		sprites = append(sprites, pic)
	}

	// imd := imdraw.New(nil)
	// filledRect(imd, 100, 100, 200, 200, pixel.RGB(0, 1, 0), pixel.RGB(1, 0, 0))

	for !win.Closed() {
		win.Clear(colornames.Whitesmoke)

		i := 0
		p := win.Bounds().Bounds().Norm().Min.Add(sprites[i].Bounds().Size().Scaled(0.5))
		sprite := pixel.NewSprite(sprites[i], sprites[i].Bounds())
		sprite.Draw(win, pixel.IM.Moved(p))

		i++
		p = p.Add(pixel.V(0, sprite.Frame().Max.Y))
		sprite = pixel.NewSprite(sprites[i], sprites[i].Bounds())
		sprite.Draw(win, pixel.IM.Moved(p))

		i++
		p = p.Add(pixel.V(0, sprite.Frame().Max.Y))
		sprite = pixel.NewSprite(sprites[i], sprites[i].Bounds())
		sprite.Draw(win, pixel.IM.Moved(p))

		i++
		p = p.Add(pixel.V(sprite.Frame().Max.X, 0))
		sprite = pixel.NewSprite(sprites[i], sprites[i].Bounds())
		sprite.Draw(win, pixel.IM.Moved(p))

		i++
		p = p.Sub(pixel.V(0, sprite.Frame().Max.Y))
		sprite = pixel.NewSprite(sprites[i], sprites[i].Bounds())
		sprite.Draw(win, pixel.IM.Moved(p))

		i++
		p = p.Sub(pixel.V(0, sprite.Frame().Max.Y))
		sprite = pixel.NewSprite(sprites[i], sprites[i].Bounds())
		sprite.Draw(win, pixel.IM.Moved(p))

		i++
		p = p.Add(pixel.V(sprite.Frame().Max.X, 0))
		sprite = pixel.NewSprite(sprites[i], sprites[i].Bounds())
		sprite.Draw(win, pixel.IM.Moved(p))

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
