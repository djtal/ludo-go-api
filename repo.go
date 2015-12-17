package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func GetDb() *sql.DB {
	var connectStr = fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"))
	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
func RepoFindGames() []Game {
	var games []Game
	var db = GetDb()

	var game struct {
		ID          int            `json:"id"`
		Name        string         `json:"name"`
		Description sql.NullString `json:"description"`
		Difficulty  int            `json:"difficulty"`
		MinPlayer   int            `json:"min-player"`
		MaxPlayer   int            `json:"max-player"`
	}

	rows, err := db.Query("select id, name, description, difficulty, min_player, max_player from games")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		g := Game{}
		err := rows.Scan(&game.ID, &game.Name, &game.Description, &game.Difficulty, &game.MinPlayer, &game.MaxPlayer)
		if err != nil {
			log.Fatal(err)
		}
		g.ID = game.ID
		if game.Description.Valid {
			// use s.String
			g.Description = game.Description.String
		} else {
			g.Description = ""
		}
		g.Difficulty = game.Difficulty
		g.MaxPlayer = game.MaxPlayer
		g.MinPlayer = game.MinPlayer
		g.Name = game.Name
		games = append(games, g)
	}

	defer db.Close()
	return games
}

func RepoFindGame(id int) Game {
	var game struct {
		ID          int            `json:"id"`
		Name        string         `json:"name"`
		Description sql.NullString `json:"description"`
		Difficulty  int            `json:"difficulty"`
		MinPlayer   int            `json:"min-player"`
		MaxPlayer   int            `json:"max-player"`
	}

	var db = GetDb()
	db.QueryRow("select id, name, description, difficulty, min_player, max_player from games where id = ?", id).Scan(&game.ID, &game.Name, &game.Description, &game.Difficulty, &game.MinPlayer, &game.MaxPlayer)
	g := Game{}
	g.ID = game.ID
	if game.Description.Valid {
		// use s.String
		g.Description = game.Description.String
	} else {
		g.Description = ""
	}
	g.Difficulty = game.Difficulty
	g.MaxPlayer = game.MaxPlayer
	g.MinPlayer = game.MinPlayer
	g.Name = game.Name

	return g
}
