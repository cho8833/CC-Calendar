package api

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/cho8833/CC-Calendar/internal/user/service"
	"strconv"
)

type UserAPI struct {
	svc *service.UserService
}

func NewUserAPI(svc *service.UserService) *UserAPI {
	return &UserAPI{svc: svc}
}

func (api UserAPI) GetUser(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userId, _ := strconv.ParseInt(event.QueryStringParameters["userId"], 10, 64)

	user, _ := api.svc.GetUser(userId)

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("%#v", user),
	}

	return response, nil
}