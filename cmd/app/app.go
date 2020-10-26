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

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type Game struct {
	app.Compo
	name  string
	board Board
}

func pageLoad() app.UI {
	// Check for ongoing game

	// Otherwise load new game
	return newGame()
}

func newGame() app.UI {
	return &Game{
		board: newBoard(),
	}
}

// Render ...
func (g *Game) Render() app.UI {
	return app.Div().Body(
		app.Main().Body(
			g.board.Render(),
		),
	)
}

// OnInputChange ...
func (g *Game) OnInputChange(ctx app.Context, e app.Event) {
	g.name = ctx.JSSrc.Get("value").String()
	g.Update()
}

func main() {
	app.Route("/", pageLoad())
	app.Run()
}
