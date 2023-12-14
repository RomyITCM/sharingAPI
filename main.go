package main

import (
	"context"
	"database/sql"
	"net/http"
	"smile-service/database"
	"smile-service/method"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

var db *sql.DB
var logger *zap.Logger

func init() {
	log, _ := zap.NewProduction()
	logger = log
	dbConnection, err := database.GetConnection()

	if err != nil {
		logger.Error("Error connection to database", zap.Error(err))
		panic(err)
	}

	dbConnection.Ping()
	if err != nil {
		logger.Error("Error ping to database", zap.Error(err))
		panic(err)
	}

	db = dbConnection
}

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	var response *events.APIGatewayProxyResponse

	if event.HTTPMethod == http.MethodGet {
		response = method.GET(ctx, db, event, logger)
	} else {
		response = method.POST(ctx, db, event, logger)
	}

	return response, nil
}

func main() {
	lambda.Start(Handler)
}
