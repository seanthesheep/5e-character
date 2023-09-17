package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// App struct
type App struct {
	ctx context.Context
}

type Ability struct {
	name     string
	modifier int
}

type Attribute struct {
	score     int
	modifier  int
	save      int
	abilities []Ability
}

type Character struct {
	id           int
	name         string
	level        int
	initiative   int
	hp           int
	hitDice      string
	ac           int
	proficiency  int
	strength     Attribute
	intelligence Attribute
	dexterity    Attribute
	wisdom       Attribute
	constitution Attribute
	charisma     Attribute
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
