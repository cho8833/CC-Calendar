package repository

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cho8833/kakao_login_lambda/model"
	"log"
)

type UserRepository struct {
}

func getClient() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	return dynamodb.NewFromConfig(cfg), nil
}

func TableExists(tableName string) (bool, error) {
	exists := true
	client, err := getClient()
	if err != nil {
		return false, err
	}
	_, err = client.DescribeTable(
		context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(tableName)},
	)
	if err != nil {
		var notFoundEx *types.ResourceNotFoundException
		if errors.As(err, &notFoundEx) {
			log.Printf("Table %v does not exist.\n", tableName)
			err = nil
		} else {
			log.Printf("Couldn't determine existence of table %v. Here's why: %v\n", tableName, err)
		}
		exists = false
	}
	return exists, err
}

func (repository UserRepository) GetUser(userId int64) (*model.User, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	user := model.User{Id: userId}
	response, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: user.GetKey(), TableName: aws.String("User"),
	})
	if err != nil {
		log.Printf("Couldn't get info about %v. Here's why: %v\n", userId, err)
		return nil, err
	}

	err = attributevalue.UnmarshalMap(response.Item, &user)
	if err != nil {
		log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		return nil, err
	}
	return &user, nil
}