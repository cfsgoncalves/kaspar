package entities

type RedditStock struct {
	NoOfComments   int32   `json:"no_of_comments"`
	Sentiment      string  `json:"sentiment"`
	SentimentScore float64 `json:"sentiment_score"`
	Ticker         string  `json:"ticker"`
}
