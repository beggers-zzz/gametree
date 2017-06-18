package main

// Game presents an interface for a two-player game.
type Game interface {

    // Score() tells us how the game is going--a positive number indicates
    // player 1 has the advantage, and a negative number indicates player 2 has
    // the advantage.
    Score() float64

    // NextPositions returns possible positions for a game to be in after one
    // move of whichever player's turn it is.
    NextPositions() *[]Game

    // Move makes a move in the Game.
    Move(*Game) error

    // PlayerOneWin returns true if the game is over and player 1 has won,
    // else false.
    PlayerOneWin() bool

    // PlayerTwoWin returns true if the game is over and player 2 has won,
    // else false.
    PlayerTwoWin() bool

    // Draw() returns true if the game is over and is a draw, else false.
    Draw() bool
}
