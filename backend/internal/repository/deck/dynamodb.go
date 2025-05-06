package deck

import (
	"context"
	"errors"
	"fmt"

	"backend/proto/protobuf"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoRepository struct {
	client    *dynamodb.Client
	tableName string
}

// NewDeckRepository は DynamoRepository を初期化します。
func NewDeckRepository(client *dynamodb.Client, tableName string) *DynamoRepository {
	return &DynamoRepository{
		client:    client,
		tableName: tableName,
	}
}

// Create は新しい Deck を作成します。
func (r *DynamoRepository) Create(ctx context.Context, deck *protobuf.Deck) error {
	_, err := r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "USER#" + deck.Owner},
			"sk": &types.AttributeValueMemberS{Value: "DECK#" + deck.DeckId},
			"name": &types.AttributeValueMemberS{Value: deck.Name},
			"owner": &types.AttributeValueMemberS{Value: deck.Owner},
			"created_at": &types.AttributeValueMemberS{Value: deck.CreatedAt},
		},
	})
	return err
}

// Get は指定された Deck を取得します。
func (r *DynamoRepository) Get(ctx context.Context, id, owner string) (*protobuf.Deck, error) {
	fmt.Printf("id: %s, owner: %s\n", id, owner)
	result, err := r.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "USER#" + owner},
			"sk": &types.AttributeValueMemberS{Value: "DECK#" + id},
		},
	})
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, errors.New("deck not found")
	}

	return &protobuf.Deck{
		DeckId:    result.Item["sk"].(*types.AttributeValueMemberS).Value[5:], // Remove "DECK#" prefix
		Name:      result.Item["name"].(*types.AttributeValueMemberS).Value,
		Owner:     result.Item["owner"].(*types.AttributeValueMemberS).Value,
		CreatedAt: result.Item["created_at"].(*types.AttributeValueMemberS).Value,
	}, nil
}

// List は指定されたユーザーの Deck を一覧表示します。
func (r *DynamoRepository) List(ctx context.Context, owner string) ([]*protobuf.Deck, error) {
	result, err := r.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(r.tableName),
		KeyConditionExpression: aws.String("pk = :pk AND begins_with(sk, :sk)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "USER#" + owner},
			":sk": &types.AttributeValueMemberS{Value: "DECK#"},
		},
	})
	if err != nil {
		return nil, err
	}

	decks := make([]*protobuf.Deck, 0, len(result.Items))
	for _, item := range result.Items {
		decks = append(decks, &protobuf.Deck{
			DeckId:    item["sk"].(*types.AttributeValueMemberS).Value[5:], // Remove "DECK#" prefix
			Name:      item["name"].(*types.AttributeValueMemberS).Value,
			Owner:     item["owner"].(*types.AttributeValueMemberS).Value,
			CreatedAt: item["created_at"].(*types.AttributeValueMemberS).Value,
		})
	}

	return decks, nil
}

// Update は指定された Deck を更新します。
func (r *DynamoRepository) Update(ctx context.Context, deck *protobuf.Deck) error {
	_, err := r.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "USER#" + deck.Owner},
			"sk": &types.AttributeValueMemberS{Value: "DECK#" + deck.DeckId},
		},
		UpdateExpression: aws.String("SET #n = :name"),
		ExpressionAttributeNames: map[string]string{
			"#n": "name",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name": &types.AttributeValueMemberS{Value: deck.Name},
		},
	})
	return err
}

// Delete は指定された Deck を削除します。
func (r *DynamoRepository) Delete(ctx context.Context,id, owner string) error {
	result, err := r.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "USER#" + owner},
			"sk": &types.AttributeValueMemberS{Value: "DECK#" + id},
		},
		ReturnValues: types.ReturnValueAllOld,
	})
	if err != nil {
	return err
	}

	fmt.Println("result", result)

	if len(result.Attributes) > 0 {
		// 削除成功
		return nil
	} else {
		// 削除失敗
		return errors.New("deck not found")
	}
} 
