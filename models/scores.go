package models

import (
	"time"
	"secsys/db"
)

// Score type represents score
type Score struct {
  ID string `json:"id" db:"id"`
  UserID string `json:"userID" db:"user_id"`
  ComName string `json:"comName" db:"com_name"`
  Year string `json:"year" db:"year"`
  Standard string `json:"standard" db:"standard"`
  ScoreNo string `json:"scoreNo" db:"score_no"`
  ScoreType string `json:"scoreType" db:"score_type"`
  Satisfied string `json:"satisfied" db:"satisfied"`
  Score float64 `json:"score" db:"score"`
  Reason string `json:"reason" db:"reason"`
  Created time.Time `db:"created"`
  Modified time.Time `db:"modified"`
}

// CreateScore create score and return id
func CreateScore(score Score) (string, error) {
  var id string
  createScoreSQL := `INSERT INTO scores (
    user_id,
    com_name,
    year,
    standard,
    score_no,
    score_type,
    satisfied,
    score,
    reason
  ) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
    RETURNING id;`
  err := db.Pool.Get(&id, createScoreSQL,
    score.UserID,
    score.ComName,
    score.Year,
    score.Standard,
    score.ScoreNo,
    score.ScoreType,
    score.Satisfied,
    score.Score,
    score.Reason)
  return id, err
}

// GetScoreByID get score by id
func GetScoreByID(scoreID string) (Score, error) {
  var score Score
  getScoreSQL := `SELECT * FROM scores WHERE id=$1`
  err := db.Pool.Get(&score, getScoreSQL, scoreID)
  return score, err
}

// GetScores return Score list
func GetScores() ([]Score, error) {
  scores := []Score{}
  getScoresSQL := `SELECT * FROM scores`
  err := db.Pool.Select(&scores, getScoresSQL)
  return scores, err
}

// UpdateScore update score with whole score entity
func UpdateScore(score Score) (error) {
  updateScoreSQL := `UPDATE scores SET
    user_id=$1,
    com_name=$2,
    year=$3,
    standard=$4,
    score_no=$5,
    score_type=$6,
    satisfied=$7,
    score=$8,
    reason=$9
  WHERE id=$10`
  _, err := db.Pool.Exec(updateScoreSQL,
    score.UserID,
    score.ComName,
    score.Year,
    score.Standard,
    score.ScoreNo,
    score.ScoreType,
    score.Satisfied,
    score.Score,
    score.Reason,
    score.ID)
  return err
}

// DeleteScoreByID delete score by id
func DeleteScoreByID(scoreID string) (error) {
  deleteScoreSQL := `DELETE FROM scores WHERE id=$1`
  _, err := db.Pool.Exec(deleteScoreSQL, scoreID)
  return err
}
