/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package internal

import (
	"github.com/hvxahv/hvx/cfg"
	"testing"
)

func init() {
	cfg.Default()
}

func TestArticles_Get(t *testing.T) {
	g, err := NewArticlesId(787516945347018753).Get(785518573776797697)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("ok", g)

	g2, err := NewArticlesId(787516945347018753).Get(785747557033967617)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("ok", g2)

	g3, err := NewArticlesId(787516945347018753).Get(787507052643319809)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("err", g3)
}
