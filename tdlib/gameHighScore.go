// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
	"fmt"
)

// GameHighScore Contains one row of the game high score table
type GameHighScore struct {
	tdCommon
	Position int32 `json:"position"` // Position in the high score table
	UserID   int32 `json:"user_id"`  // User identifier
	Score    int32 `json:"score"`    // User score
}

// MessageType return the string telegram-type of GameHighScore
func (gameHighScore *GameHighScore) MessageType() string {
	return "gameHighScore"
}

// NewGameHighScore creates a new GameHighScore
//
// @param position Position in the high score table
// @param userID User identifier
// @param score User score
func NewGameHighScore(position int32, userID int32, score int32) *GameHighScore {
	gameHighScoreTemp := GameHighScore{
		tdCommon: tdCommon{Type: "gameHighScore"},
		Position: position,
		UserID:   userID,
		Score:    score,
	}

	return &gameHighScoreTemp
}