package api

import "time"

// Mostly Auto-Generated from https://mholt.github.io/json-to-go/
type CtfdChallListResponse struct {
	Success bool        `json:"success"`
	Data    []CtfdChall `json:"data"`
}

type Requirements struct {
}

type CtfdChall struct {
	ID             int          `json:"id"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	ConnectionInfo string       `json:"connection_info"`
	NextID         int          `json:"next_id"`
	MaxAttempts    int          `json:"max_attempts"`
	Value          int          `json:"value"`
	Category       string       `json:"category"`
	Type           string       `json:"type"`
	State          string       `json:"state"`
	Requirements   Requirements `json:"requirements"`
	Solves         int          `json:"solves"`
	SolvedByMe     bool         `json:"solved_by_me"`
}

type CtfdHintResponse struct {
	Success bool `json:"success"`
	Data    []struct {
		Content      string `json:"content"`
		Cost         int    `json:"cost"`
		Challenge    int    `json:"challenge"`
		Type         string `json:"type"`
		Requirements struct {
			Prerequisites []any `json:"prerequisites"`
		} `json:"requirements"`
		ChallengeID int `json:"challenge_id"`
		ID          int `json:"id"`
	} `json:"data"`
}

type CtfdChallResponse struct {
	Success bool `json:"success"`
	Data    struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Value          int    `json:"value"`
		Description    string `json:"description"`
		ConnectionInfo any    `json:"connection_info"`
		NextID         any    `json:"next_id"`
		Category       string `json:"category"`
		State          string `json:"state"`
		MaxAttempts    int    `json:"max_attempts"`
		Type           string `json:"type"`
		TypeData       struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			Templates struct {
				Create string `json:"create"`
				Update string `json:"update"`
				View   string `json:"view"`
			} `json:"templates"`
			Scripts struct {
				Create string `json:"create"`
				Update string `json:"update"`
				View   string `json:"view"`
			} `json:"scripts"`
		} `json:"type_data"`
		Solves     int      `json:"solves"`
		SolvedByMe bool     `json:"solved_by_me"`
		Attempts   int      `json:"attempts"`
		Files      []string `json:"files"`
		Tags       []any    `json:"tags"`
		Hints      []any    `json:"hints"`
		View       string   `json:"view"`
	} `json:"data"`
}

type CtfdTeamResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Email       any    `json:"email"`
		Members     []int  `json:"members"`
		BracketID   any    `json:"bracket_id"`
		Fields      []any  `json:"fields"`
		Affiliation any    `json:"affiliation"`
		Website     any    `json:"website"`
		Country     any    `json:"country"`
		ID          int    `json:"id"`
		CaptainID   int    `json:"captain_id"`
		Name        string `json:"name"`
		OauthID     any    `json:"oauth_id"`
		Place       string `json:"place"`
		Score       int    `json:"score"`
	} `json:"data"`
}

type CtfdScoreBoardResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Num1 struct {
			ID          int    `json:"id"`
			AccountURL  string `json:"account_url"`
			Name        string `json:"name"`
			Score       int    `json:"score"`
			BracketID   any    `json:"bracket_id"`
			BracketName any    `json:"bracket_name"`
			Solves      []struct {
				ChallengeID int       `json:"challenge_id"`
				AccountID   int       `json:"account_id"`
				TeamID      int       `json:"team_id"`
				UserID      int       `json:"user_id"`
				Value       int       `json:"value"`
				Date        time.Time `json:"date"`
			} `json:"solves"`
		} `json:"1"`
		Num2 struct {
			ID          int    `json:"id"`
			AccountURL  string `json:"account_url"`
			Name        string `json:"name"`
			Score       int    `json:"score"`
			BracketID   any    `json:"bracket_id"`
			BracketName any    `json:"bracket_name"`
			Solves      []struct {
				ChallengeID int       `json:"challenge_id"`
				AccountID   int       `json:"account_id"`
				TeamID      int       `json:"team_id"`
				UserID      int       `json:"user_id"`
				Value       int       `json:"value"`
				Date        time.Time `json:"date"`
			} `json:"solves"`
		} `json:"2"`
		Num3 struct {
			ID          int    `json:"id"`
			AccountURL  string `json:"account_url"`
			Name        string `json:"name"`
			Score       int    `json:"score"`
			BracketID   any    `json:"bracket_id"`
			BracketName any    `json:"bracket_name"`
			Solves      []struct {
				ChallengeID int       `json:"challenge_id"`
				AccountID   int       `json:"account_id"`
				TeamID      int       `json:"team_id"`
				UserID      int       `json:"user_id"`
				Value       int       `json:"value"`
				Date        time.Time `json:"date"`
			} `json:"solves"`
		} `json:"3"`
	} `json:"data"`
}
