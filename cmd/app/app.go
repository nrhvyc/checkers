// return app.Div().Body(
// 	app.Main().Body(
// 		app.H1().Body(
// 			app.Text("Hello, "),
// 			app.If(g.name != "",
// 				app.Text(g.name),
// 			).Else(
// 				app.Text("World"),
// 			),
// 		),
// 		app.Input().
// 			Value(g.name).
// 			Placeholder("What is your name?").
// 			AutoFocus(true).
// 			OnChange(g.OnInputChange),
// 	),
// )
package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/nrhvyc/checkers/internal/game"
)

// Game ...
type Game struct {
	app.Compo
	Board game.Board
}

func pageLoad() app.UI {
	// Check for ongoing game

	// Otherwise load new game
	return newGame()
}

func newGame() app.UI {
	return &Game{
		Board: game.NewBoard(),
	}
}

// Render ...
func (g *Game) Render() app.UI {
	return app.Div().Body(
		app.Main().Body(
			g.Board.Render(),
		),
	)
}

// OnInputChange ...
// func (g *Game) OnInputChange(ctx app.Context, e app.Event) {
// 	// g.name = ctx.JSSrc.Get("value").String()
// 	g.Update()
// }

func main() {
	app.Route("/", pageLoad())
	app.Run()
}

// Initialize the encoder and decoder.  Normally enc and dec would be
// bound to network connections and the encoder and decoder would
// run in different processes.
// var network bytes.Buffer        // Stand-in for a network connection
// enc := gob.NewEncoder(&network) // Will write to network.
// dec := gob.NewDecoder(&network) // Will read from network.
// // Encode (send) the value.
// err := enc.Encode(P{3, 4, 5, "Pythagoras"})
// if err != nil {
// 	log.Fatal("encode error:", err)
// }
// // Decode (receive) the value.
// var q Q
// err = dec.Decode(&q)
// if err != nil {
// 	log.Fatal("decode error:", err)
// }
// fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)
