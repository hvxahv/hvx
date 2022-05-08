/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package client

import (
	acct "github.com/hvxahv/hvx/api/grpc/proto/account/v1alpha1"
	activity "github.com/hvxahv/hvx/api/grpc/proto/activity/v1alpha1"
	article "github.com/hvxahv/hvx/api/grpc/proto/article/v1alpha1"
	channel "github.com/hvxahv/hvx/api/grpc/proto/channel/v1alpha1"
	device "github.com/hvxahv/hvx/api/grpc/proto/device/v1alpha1"
	message "github.com/hvxahv/hvx/api/grpc/proto/message/v1alpha1"
	public "github.com/hvxahv/hvx/api/grpc/proto/public/v1alpha1"
	saved "github.com/hvxahv/hvx/api/grpc/proto/saved/v1alpha1"
	"google.golang.org/grpc"
)

type Account interface {
	acct.AccountsClient
	acct.ActorClient
	acct.AuthClient
}

type account struct {
	acct.AccountsClient
	acct.ActorClient
	acct.AuthClient
}

func NewAccount(conn *grpc.ClientConn) Account {
	return &account{
		AccountsClient: acct.NewAccountsClient(conn),
		ActorClient:    acct.NewActorClient(conn),
		AuthClient:     acct.NewAuthClient(conn),
	}
}

type Activity interface {
	activity.ActivityClient
	activity.FollowClient
}

type activities struct {
	activity.ActivityClient
	activity.FollowClient
}

func NewActivity(conn *grpc.ClientConn) Activity {
	return &activities{
		ActivityClient: activity.NewActivityClient(conn),
		FollowClient:   activity.NewFollowClient(conn),
	}
}

type Article interface {
	article.ArticleClient
}

type art struct {
	article.ArticleClient
}

func NewArticle(conn *grpc.ClientConn) Article {
	return &art{
		ArticleClient: article.NewArticleClient(conn),
	}
}

type Channel interface {
	channel.AdministrativeClient
	channel.BroadcastClient
	channel.ChannelClient
	channel.SubscriberClient
}

type channels struct {
	channel.AdministrativeClient
	channel.BroadcastClient
	channel.ChannelClient
	channel.SubscriberClient
}

func NewChannel(conn *grpc.ClientConn) Channel {
	return &channels{
		AdministrativeClient: channel.NewAdministrativeClient(conn),
		BroadcastClient:      channel.NewBroadcastClient(conn),
		ChannelClient:        channel.NewChannelClient(conn),
		SubscriberClient:     channel.NewSubscriberClient(conn),
	}
}

type Device interface {
	device.DevicesClient
}

type devices struct {
	device.DevicesClient
}

func NewDevice(conn *grpc.ClientConn) Device {
	return &devices{
		DevicesClient: device.NewDevicesClient(conn),
	}
}

type Message interface {
	message.MessagesClient
}

type messages struct {
	message.MessagesClient
}

func NewMessage(conn *grpc.ClientConn) Message {
	return &messages{message.NewMessagesClient(conn)}
}

type Public interface {
	public.PublicClient
}

type pub struct {
	public.PublicClient
}

func NewPublic(conn *grpc.ClientConn) Public {
	return &pub{
		PublicClient: public.NewPublicClient(conn),
	}
}

type Saved interface {
	saved.SavedClient
}

type saves struct {
	saved.SavedClient
}

func NewSaved(conn *grpc.ClientConn) Saved {
	return &saves{saved.NewSavedClient(conn)}
}
