package game

import (
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	started bool
	counter int
}

func (g *Game) Update() error {
	if !g.started {
		g.started = true

		go func() {
			for {
				time.Sleep(1 * time.Second)
				// Surely this needs to be sync'd.... right??
				g.counter++
			}
		}()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World! - "+strconv.Itoa(g.counter))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
