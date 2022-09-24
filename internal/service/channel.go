// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
)

type IChannel interface {
	BackendIndex(ctx context.Context) (out []*model.ChannelBackendApiListItem, err error)
	BackendChannelTree(ctx context.Context) (out []*model.ChannelBackendApiListItem, err error)
	BackendApiStatus(ctx context.Context, in *backendApi.ChannelStatusApiReq) (out *backendApi.ChannelStatusApiRes, err error)
	BackendApiDelete(ctx context.Context, in *backendApi.ChannelDeleteApiReq) (out *backendApi.ChannelDeleteApiRes, err error)
}

var localChannel IChannel

func Channel() IChannel {
	if localChannel == nil {
		panic("implement not found for interface IChannel, forgot register?")
	}
	return localChannel
}

func RegisterChannel(i IChannel) {
	localChannel = i
}
