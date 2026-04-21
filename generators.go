package main

import (
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

type Wagon struct {
	Sprite *ebiten.Image
	Tag    *ebiten.Image
	Tagged bool
	Flip   bool
	Length int
}

func _as_wagon(sprite *ebiten.Image) *Wagon {
	return &Wagon{
		Sprite: sprite,
		Tagged: false,
		Length: sprite.Bounds().Dx(),
		Flip:   false,
	}
}

type Train struct {
	Wagons []*Wagon
	X      int
	Length int
	Speed  int
}

func (t *Train) Append(t2 *Train) {
	t.Wagons = append(t.Wagons, t2.Wagons...)
	t.Length += t2.Length
	t.X += t2.X
	t.Speed = min(t.Speed, t2.Speed)
}

func _choose_one(bag []*ebiten.Image) *ebiten.Image {
	return bag[rand.Int()%len(bag)]
}

func GenerateIndustrialTrain(s *Sprites) *Train {
	typeDiceThrow := rand.Int() % 100
	wagons := []*Wagon{
		_as_wagon(_choose_one(s.IndustrialEngineVariants)),
	}
	if typeDiceThrow < 30 {
		// container and flatbeds, 5-7 wagons, 30% chance of tag
		for range (rand.Int() % 3) + 5 {
			if rand.Int()%2 == 0 {
				wagons = append(wagons, _as_wagon(_choose_one(s.FlatBedVariants)))
			} else {
				container := _as_wagon(_choose_one(s.ContainerVariants))
				if rand.Int()%3 == 0 {
					container.Tag = _choose_one(s.Tags32Wide)
					container.Tagged = true
				}
				wagons = append(wagons, container)
			}
		}
	} else if typeDiceThrow < 60 {
		// bins and flatbed, 3-6 wagons, 30-70 partition, 10% chance of tag on bins
		for range (rand.Int() % 4) + 3 {
			if rand.Int()%10 < 4 {
				bin := _as_wagon(_choose_one(s.BinVariants))
				if rand.Int()%10 == 0 {
					bin.Tag = _choose_one(s.Tags32Wide)
					bin.Tagged = true
				}
				wagons = append(wagons, bin)
			} else {
				wagons = append(wagons, _as_wagon(_choose_one(s.FlatBedVariants)))
			}
		}
	} else if typeDiceThrow < 90 {
		// container only, 3-7 wagons, 60% chance of tag
		for range (rand.Int() % 5) + 3 {
			container := _as_wagon(_choose_one(s.ContainerVariants))
			if rand.Int()%10 < 6 {
				container.Tag = _choose_one(s.Tags32Wide)
				container.Tagged = true
			}
			wagons = append(wagons, container)
		}
	} else {
		// 12-21 bins, 10% chance of tag
		for range (rand.Int() % 9) + 12 {
			bin := _as_wagon(_choose_one(s.BinVariants))
			if rand.Int()%10 == 0 {
				bin.Tag = _choose_one(s.Tags32Wide)
				bin.Tagged = true
			}
			wagons = append(wagons, bin)
		}
	}
	length := 0
	for _, w := range wagons {
		length += w.Length
	}
	return &Train{
		Wagons: wagons,
		Length: length,
		Speed:  1,
	}
}

func GenerateTER(s *Sprites) *Train {
	typeDiceThrow := rand.Int() % 100
	if typeDiceThrow < 50 {
		typeDiceThrow = 0
	} else if typeDiceThrow < 90 {
		typeDiceThrow = 1
	} else {
		typeDiceThrow = 2
	}
	if rand.Int()%2 == 0 {
		return _generateOneTER(s, typeDiceThrow)
	} else {
		train := _generateOneTER(s, typeDiceThrow)
		train.Append(_generateOneTER(s, typeDiceThrow))
		return train
	}
}

func _generateOneTER(s *Sprites, typeDiceThrow int) *Train {
	wagons := []*Wagon{}
	wagons = append(wagons, _as_wagon(s.TEREngineVariants[typeDiceThrow]))

	for range (rand.Int() % 3) + 3 {
		carriage := _as_wagon(s.TERCarriageVariants[typeDiceThrow])
		if rand.Int()%100 < 15 {
			carriage.Tag = _choose_one(s.Tags48Wide)
			carriage.Tagged = true
		}
		wagons = append(wagons, carriage)
	}
	flipped_engine := _as_wagon(s.TEREngineVariants[typeDiceThrow])
	flipped_engine.Flip = true
	wagons = append(wagons, flipped_engine)

	length := 0
	for _, w := range wagons {
		length += w.Length
	}
	return &Train{
		Wagons: wagons,
		Length: length,
		Speed:  1,
	}
}

func GenerateTGV(s *Sprites) *Train {
	wagons := []*Wagon{}
	wagons = append(wagons, _as_wagon(s.TGVEngine))

	for range (rand.Int() % 4) + 5 {
		wagons = append(wagons, _as_wagon(s.TGVCarriage))
	}

	flipped_engine := _as_wagon(s.TGVEngine)
	flipped_engine.Flip = true
	wagons = append(wagons, flipped_engine)

	length := 0
	for _, w := range wagons {
		length += w.Length
	}
	return &Train{
		Wagons: wagons,
		Length: length,
		Speed:  2,
	}
}

func GenerateTrain(s *Sprites) *Train {
	typeDiceThrow := rand.Int() % 100
	if typeDiceThrow < 45 {
		return GenerateTER(s)
	} else if typeDiceThrow < 85 {
		return GenerateIndustrialTrain(s)
	} else {
		return GenerateTGV(s)
	}
}

func GenerateTracks(s *Sprites, screenWidth int, onlyCleanTracks bool) []*ebiten.Image {
	// 40% chance of mutation for each track tile
	tracks := []*ebiten.Image{}

	tile_count := screenWidth / 16
	if onlyCleanTracks {
		for range tile_count {
			tracks = append(tracks, s.TracksBase)
		}
	} else {
		for range tile_count {
			if rand.Int()%10 < 6 {
				tracks = append(tracks, s.TracksBase)
			} else {
				tracks = append(tracks, _choose_one(s.TracksVariants))
			}
		}
	}
	return tracks
}
