package main

import (
	"embed"
	"image"
	_ "image/png"

	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// not a comment, this is a compiler directive meant to indicate to
// ↓ the compiler that it should embed the /assets directory in the binary

//go:embed assets
var assetFS embed.FS

type Sprites struct {
	FlatBedVariants          []*ebiten.Image
	ContainerVariants        []*ebiten.Image
	BinVariants              []*ebiten.Image
	Tags32Wide               []*ebiten.Image
	IndustrialEngineVariants []*ebiten.Image
	TracksBase               *ebiten.Image
	TracksVariants           []*ebiten.Image
	TGVCarriage              *ebiten.Image
	TGVEngine                *ebiten.Image
	TEREngineVariants        []*ebiten.Image
	TERCarriageVariants      []*ebiten.Image
	Tags48Wide               []*ebiten.Image
}

func loadSprites() *Sprites {
	TileSet, _, err := ebitenutil.NewImageFromFileSystem(assetFS, "assets/trains.png")
	if err != nil {
		log.Fatal(err)
	}
	s := Sprites{}
	s.FlatBedVariants = []*ebiten.Image{
		TileSet.SubImage(image.Rect(0, 0, 0+32, 0+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(32, 0, 32+32, 0+32)).(*ebiten.Image),
	}
	s.ContainerVariants = []*ebiten.Image{
		TileSet.SubImage(image.Rect(64, 0, 64+32, 0+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(96, 0, 96+32, 0+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(128, 0, 128+32, 0+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(160, 0, 160+32, 0+32)).(*ebiten.Image),
	}
	s.BinVariants = []*ebiten.Image{
		TileSet.SubImage(image.Rect(0, 32, 0+32, 32+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(32, 32, 32+32, 32+32)).(*ebiten.Image),
	}
	s.Tags32Wide = []*ebiten.Image{
		TileSet.SubImage(image.Rect(64, 32, 64+32, 32+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(96, 32, 96+32, 32+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(128, 32, 128+32, 32+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(160, 32, 160+32, 32+32)).(*ebiten.Image),
	}
	s.IndustrialEngineVariants = []*ebiten.Image{
		TileSet.SubImage(image.Rect(0, 64, 0+48, 64+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(48, 64, 48+48, 64+32)).(*ebiten.Image),
	}
	s.TracksBase = TileSet.SubImage(image.Rect(176, 64, 176+16, 64+32)).(*ebiten.Image)
	s.TracksVariants = []*ebiten.Image{
		TileSet.SubImage(image.Rect(96, 64, 96+16, 64+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(112, 64, 112+16, 64+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(128, 64, 128+16, 64+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(144, 64, 144+16, 64+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(160, 64, 160+16, 64+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(128, 96, 128+16, 96+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(144, 96, 144+16, 96+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(160, 96, 160+16, 96+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(176, 96, 176+16, 96+32)).(*ebiten.Image),
	}
	s.TGVCarriage = TileSet.SubImage(image.Rect(0, 96, 0+64, 96+32)).(*ebiten.Image)
	s.TGVEngine = TileSet.SubImage(image.Rect(64, 96, 64+64, 96+32)).(*ebiten.Image)
	s.TEREngineVariants = []*ebiten.Image{
		TileSet.SubImage(image.Rect(0, 128, 0+48, 128+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(48, 128, 48+48, 128+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(96, 128, 96+48, 128+32)).(*ebiten.Image),
	}
	s.TERCarriageVariants = []*ebiten.Image{
		TileSet.SubImage(image.Rect(0, 160, 0+48, 160+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(48, 160, 48+48, 160+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(96, 160, 96+48, 160+32)).(*ebiten.Image),
	}
	s.Tags48Wide = []*ebiten.Image{
		TileSet.SubImage(image.Rect(144, 128, 144+48, 128+32)).(*ebiten.Image),
		TileSet.SubImage(image.Rect(144, 160, 144+48, 160+32)).(*ebiten.Image),
	}

	return &s
}
