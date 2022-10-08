/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package internal

import (
	"fmt"
	"github.com/hvxahv/hvx/cfg"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"
	"testing"
)

func init() {
	cfg.Default()
}

func TestDB(t *testing.T) {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Articles{}); err != nil {
		t.Error(errors.NewDatabaseCreate(ArticleTable))
	}
}

func TestArticles_Create(t *testing.T) {
	a := &Articles{}

	create, err := a.Create()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(create)
}

func TestArticles_Get(t *testing.T) {
	get, err := NewArticlesId(801949911405297665).Get()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(get)
}

func TestArticles_GetArticles(t *testing.T) {
	articles, err := NewArticlesActorId(801935105807482881).GetArticles()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(articles)
}
