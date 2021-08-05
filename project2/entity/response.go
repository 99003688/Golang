package entity

type Response struct {
	Date             int64  `json:"date"`
	State            string `json:"state"`
	Positive         string `json:"positive"`
	ProbableCases    int64  `json:"probableCases"`
	Negative         int64  `json:"negative"`
	TotalTestResults int64  `json:"totalTestResults"`
	Recovered        int64  `josn:"recovered"`
	LastUpdateEt     string `json:"lastUpdateEt"`
	DateModified     string `json:"dateModified"`
	Death            int64  `josn:"death"`
	Hospitalized     int64  `json:"hospitalized"`
	DeathConfirmed   int64  `json:"deathConfirmed"`
}
