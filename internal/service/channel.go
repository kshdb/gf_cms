// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
)

type (
	IChannel interface {
		PcNavigation(ctx context.Context) (out []*model.ChannelPcNavigationListItem, err error)
		BackendApiIndex(ctx context.Context) (out []*model.ChannelBackendApiListItem, err error)
		BackendChannelTree(ctx context.Context, selectedId int) (out []*model.ChannelBackendApiListItem, err error)
		BackendChannelModelTree(ctx context.Context, modelType string, channelId int) (out []*model.ChannelBackendApiListItem, err error)
		BackendApiStatus(ctx context.Context, in *backendApi.ChannelStatusApiReq) (out *backendApi.ChannelStatusApiRes, err error)
		BackendApiDelete(ctx context.Context, in *backendApi.ChannelDeleteApiReq) (out *backendApi.ChannelDeleteApiRes, err error)
		BackendApiAdd(ctx context.Context, in *backendApi.ChannelAddApiReq) (out *backendApi.ChannelAddApiRes, err error)
		BackendApiEdit(ctx context.Context, in *backendApi.ChannelEditApiReq) (out *backendApi.ChannelEditApiRes, err error)
		GetOneById(ctx context.Context, id int) (out *entity.CmsChannel, err error)
		BackendModelMap() map[string]string
		BackendModelCanAddMap() map[string]string
		BackendModelDesc(model string) string
		UpdateRelation(ctx context.Context, originChannelId int) (out interface{}, err error)
		GetChildIds(ctx context.Context, belongChannelId int, andMe bool) (arrAllIds []int, err error)
		PcHomeAboutChannel(ctx context.Context, channelId int) (channel *entity.CmsChannel, err error)
	}
)

var (
	localChannel IChannel
)

func Channel() IChannel {
	if localChannel == nil {
		panic("implement not found for interface IChannel, forgot register?")
	}
	return localChannel
}

func RegisterChannel(i IChannel) {
	localChannel = i
}
