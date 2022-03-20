/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package article

import (
	"context"
	"fmt"
	"github.com/hvxahv/hvxahv/api/article/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func init() {
	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Search configs in home directory with name ".hvxahv" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".hvxahv")

	viper.AutomaticEnv()

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}

	// Initialize the database.
	n := cockroach.NewDBAddr()
	if err := n.InitDB(); err != nil {
		fmt.Println(err)
		return
	}

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println(err)
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
		return
	}
}

func TestArticle_CreateArticle(t *testing.T) {
	aid := "746166817947975681"
	title := "Test Article2"
	summary := "Test Summary"
	articles := "This is a test article2."
	tags := []string{"Arts", "Sports"}
	AttachmentType := "image/jpeg"
	attachments := []string{"http://49.233.26.52:9001/api/v1/buckets/avatar/objects/download?preview=true&prefix=N2I3N2QyYjgtZWNjZi00ODg2LWFjMGYtMTgwNmFiNzQ3MjYxLeWwj+ael+eUseS+nS5qcGc=&version_id=null"}
	to := []string{"https://mas.to/users/hvturingga/inbox"}
	cc := []string{""}
	state := false
	nsfw := false
	visibility := "0"

	data := &v1alpha1.CreateArticleRequest{
		AccountId:      aid,
		Title:          title,
		Summary:        summary,
		Article:        articles,
		Tags:           tags,
		AttachmentType: AttachmentType,
		Attachments:    attachments,
		To:             to,
		Cc:             cc,
		State:          state,
		Nsfw:           nsfw,
		Visibility:     visibility,
	}
	s := article{}
	c, err := s.CreateArticle(context.Background(), data)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(c.Code, c.Reply)
}

func TestArticle_GetArticle(t *testing.T) {
	data := &v1alpha1.GetArticleRequest{
		Id: "746142390030860289",
	}
	s := article{}
	c, err := s.GetArticle(context.Background(), data)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(c.Article)
}

func TestArticle_GetArticlesByAccountID(t *testing.T) {
	data := &v1alpha1.GetArticlesByAccountIDRequest{
		AccountId: "737973421798785025",
	}
	s := article{}
	c, err := s.GetArticlesByAccountID(context.Background(), data)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(c.Articles)
}

func TestArticle_UpdateArticle(t *testing.T) {
	id := "746151348954857473"
	aid := "737973421798785025"
	title := "Test Article10"
	summary := "Test Summary20"
	articles := "This is a test article21."
	tags := []string{"Arts"}
	AttachmentType := "image/jpeg"
	attachments := []string{"http://49.233.26.52:9001/api/v1/buckets/avatar/objects/download?preview=true&prefix=N2I3N2QyYjgtZWNjZi00ODg2LWFjMGYtMTgwNmFiNzQ3MjYxLeWwj+ael+eUseS+nS5qcGc=&version_id=null"}
	nsfw := false
	visibility := "0"

	data := &v1alpha1.UpdateArticleRequest{
		Id:             id,
		AccountId:      aid,
		Title:          title,
		Summary:        summary,
		Article:        articles,
		Tags:           tags,
		AttachmentType: AttachmentType,
		Attachments:    attachments,
		Nsfw:           nsfw,
		Visibility:     visibility,
	}
	s := article{}
	c, err := s.UpdateArticle(context.Background(), data)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(c.Code, c.Reply)
}

func TestArticle_DeleteArticle(t *testing.T) {
	id := "746151788033015809"
	aid := "737973421798785025"
	data := &v1alpha1.DeleteArticleRequest{
		Id:        id,
		AccountId: aid,
	}
	s := article{}
	c, err := s.DeleteArticle(context.Background(), data)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(c.Code, c.Reply)
}
