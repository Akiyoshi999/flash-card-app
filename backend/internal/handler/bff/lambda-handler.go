package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"backend/internal/repository/deck"
	"backend/internal/service"
	"backend/proto/protobuf"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Request struct {
	Operation string          `json:"operation"`
	Data      json.RawMessage `json:"data"`
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func HandleRequest(ctx context.Context, request Request) (interface{}, error) {
	fmt.Println("request", request)
	// Initialize AWS SDK
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	// Initialize DynamoDB client
	ddbClient := dynamodb.NewFromConfig(cfg)
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	// Initialize repository and service
	repo := deck.NewDeckRepository(ddbClient, tableName)
	deckService := service.NewDeckService(repo)

	// Handle different operations
	switch request.Operation {
	case "CreateDeck":
		var req protobuf.CreateDeckRequest
		if err := json.Unmarshal(request.Data, &req); err != nil {
			return nil, err
		}
		resp, err := deckService.CreateDeck(ctx, &req)
		if err != nil {
			return nil, err
		}
		fmt.Println("resp", resp)
		return resp, nil

	case "GetDeck":
		var req protobuf.GetDeckRequest
		if err := json.Unmarshal(request.Data, &req); err != nil {
			return nil, err
		}
		resp, err := deckService.GetDeck(ctx, &req)
		if err != nil {
			return nil, err
		}
		return resp, nil

	case "ListDecks":
		var req protobuf.ListDecksRequest
		if err := json.Unmarshal(request.Data, &req); err != nil {
			return nil, err
		}
		resp, err := deckService.ListDecks(ctx, &req)
		if err != nil {
			return nil, err
		}
		return resp, nil

	case "UpdateDeck":
		var req protobuf.UpdateDeckRequest
		if err := json.Unmarshal(request.Data, &req); err != nil {
			return nil, err
		}
		resp, err := deckService.UpdateDeck(ctx, &req)
		if err != nil {
			return nil, err
		}
		return resp, nil

	case "DeleteDeck":
		var req protobuf.DeleteDeckRequest
		if err := json.Unmarshal(request.Data, &req); err != nil {
			return nil, err
		}
		resp, err := deckService.DeleteDeck(ctx, &req)
		if err != nil {
			return nil, err
		}
		return resp, nil

	default:
		return nil, errors.New("unknown operation")
	}
}

