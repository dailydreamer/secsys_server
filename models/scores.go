package models

// Score type represents score
type Score struct {
  ID string `json:"id"`
	// TODO
}

// CreateScore create score and return id
func CreateScore(score Score) (string, error) {
  // TODO
}

// GetScoreByID get score by id
func GetScoreByID(scoreID string) (Score, error) {
  // TODO
}

// GetScores return Score list
func GetScores() ([]Score, error) {
  // TODO
}

// UpdateScore update score with whole score entity
func UpdateScore(score Score) (error) {

}

// DeleteScoreByID delete score by id
func DeleteScoreByID(scoreID string) (error) {

}
