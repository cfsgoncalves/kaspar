package api

import (
	"context"
	api "kaspar/api/grpcprotos"
	"kaspar/configuration"
	"kaspar/entities"
	usecase "kaspar/usecase/implementation"
	"time"
)

type GrpcServer struct {
	api.UnimplementedStockHandleServer
	StockRedditApi usecase.StockRedditApi
}

func (g *GrpcServer) GetStockByNameAndOptionalDate(ctx context.Context, in *api.StockRequest) (*api.StockResponse, error) {
	dateParam := in.GetDate()
	stockName := in.GetName()

	DATE_FORMAT := configuration.GetEnvAsString("DATE_FORMAT", "2006-01-02")
	date := time.Now().UTC().Format(DATE_FORMAT)

	//Validate date paramenter
	if dateParam != "" {
		_, err := time.Parse(DATE_FORMAT, dateParam)
		if err != nil {
			return &api.StockResponse{}, err
		}
		date = dateParam
	}

	redditStruct, err := g.StockRedditApi.GetStockByName(date, stockName)

	switch err {
	case nil:
		return &api.StockResponse{
			NoOfComments:   redditStruct.(entities.RedditStock).NoOfComments,
			Sentiment:      redditStruct.(entities.RedditStock).Sentiment,
			SentimentScore: redditStruct.(entities.RedditStock).SentimentScore,
			Ticker:         redditStruct.(entities.RedditStock).Ticker,
		}, nil
	default:
		return nil, err
	}
}
