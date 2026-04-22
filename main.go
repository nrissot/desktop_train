package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Mouse_x     int
	Mouse_y     int
	SpriteSet   *Sprites
	ScreenWidth int
	Train       *Train
	Tracks      []*ebiten.Image
	LogStats    bool
	GameStats   *Stats
}

func (g *Game) Update() error {
	g.Mouse_x, g.Mouse_y = ebiten.CursorPosition()

	if g.Train.X <= -g.Train.Length {
		if g.LogStats {
			g.GameStats.Train_Counts += 1
			g.GameStats.Save()
		}
		newTrain := GenerateTrain(g.SpriteSet)
		newTrain.X = g.ScreenWidth + newTrain.Length
		g.Train = newTrain

	} else {
		g.Train.X -= g.Train.Speed
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.Tracks != nil {
		for i, tile := range g.Tracks {
			opts := ebiten.DrawImageOptions{}
			if g.Mouse_y >= 0 && g.Mouse_y <= 32 {
				opts.ColorScale.ScaleAlpha(.25)
			}
			opts.GeoM.Translate(float64(i*16), 0)
			screen.DrawImage(tile, &opts)
		}
	}

	currentOffset := 0
	for _, wagon := range g.Train.Wagons {
		opts := ebiten.DrawImageOptions{}
		// confusing, the sprites are flipped by default, so wagon.Flip actually mean "DO NOT FLIP"
		// TODO switch it around ?
		if wagon.Flip {
			opts.GeoM.Scale(1, 1)
			opts.GeoM.Translate(-float64(wagon.Length), 0)
		} else {
			opts.GeoM.Scale(-1, 1)
		}

		opts.GeoM.Translate(float64(g.Train.X+(currentOffset)+wagon.Length), 0)
		currentOffset += wagon.Length

		if g.Mouse_y >= 0 && g.Mouse_y <= 32 {
			opts.ColorScale.ScaleAlpha(.25)
		}

		screen.DrawImage(wagon.Sprite, &opts)
		if wagon.Tagged {
			screen.DrawImage(wagon.Tag, &opts)
		}
	}
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("tps:%v fps:%v", ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	var config Config
	config.Load()

	monitor_width, monitor_height := ebiten.Monitor().Size()

	window_width := monitor_width
	window_height := 32

	ebiten.SetWindowMousePassthrough(true)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)

	ebiten.SetWindowPosition(0, monitor_height-window_height-config.Window_y_offset)

	s := loadSprites()
	train := GenerateTER(s)
	train.X = window_width + train.Length

	ebiten.SetWindowSize(window_width, window_height)
	ebiten.SetWindowTitle("Tchoo tchooo !")
	icon_engine := s.IndustrialEngineVariants[0]

	// TODO, multiple sizes could be nice
	ebiten.SetWindowIcon([]image.Image{icon_engine.SubImage(icon_engine.Bounds())})

	var stats *Stats = &Stats{}

	if config.Log_Statistics {
		stats.Load()
	}

	g := Game{
		SpriteSet:   s,
		Train:       train,
		ScreenWidth: window_width,
		LogStats:    config.Log_Statistics,
		GameStats:   stats,
	}

	if config.Tracks_Visible {
		g.Tracks = GenerateTracks(s, window_width, config.Only_Clean_Tracks)
	} else {
		g.Tracks = nil
	}

	if err := ebiten.RunGameWithOptions(&g, &ebiten.RunGameOptions{ScreenTransparent: true}); err != nil {
		log.Fatal(err)
	}
}
