// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ICronClient interface {
		Start(ctx context.Context)
		Stop(ctx context.Context)
		IsLogin() bool
		OnCronDelete(ctx context.Context, args ...interface{})
		OnCronEdit(ctx context.Context, args ...interface{})
		OnCronStatus(ctx context.Context, args ...interface{})
		OnCronOnlineExec(ctx context.Context, args ...interface{})
	}
	IAuthClient interface {
		Start(ctx context.Context)
		Stop(ctx context.Context)
		IsLogin() bool
		OnResponseAuthSummary(ctx context.Context, args ...interface{})
	}
)

var (
	localAuthClient IAuthClient
	localCronClient ICronClient
)

func AuthClient() IAuthClient {
	if localAuthClient == nil {
		panic("implement not found for interface IAuthClient, forgot register?")
	}
	return localAuthClient
}

func RegisterAuthClient(i IAuthClient) {
	localAuthClient = i
}

func CronClient() ICronClient {
	if localCronClient == nil {
		panic("implement not found for interface ICronClient, forgot register?")
	}
	return localCronClient
}

func RegisterCronClient(i ICronClient) {
	localCronClient = i
}
