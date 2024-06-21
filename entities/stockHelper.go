package entities

type RedditStock struct {
	NoOfComments   string `json:"no_of_comments"`
	Sentiment      string `json:"sentiment"`
	SentimentScore string `json:"sentiment_score"`
	Ticker         string `json:"ticker"`
}
