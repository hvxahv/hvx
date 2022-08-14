package internal

import "testing"

func TestSubscribes_IsSubscriber(t *testing.T) {
	isSub, err := NewSubscribe(787709956263870465, 785747724097224705).IsSubscriber()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(isSub)
}

func TestSubscribes_AddSubscriber(t *testing.T) {
	if err := NewSubscribe(786088233592553473, 785747724097224705).AddSubscriber(785518573776797697); err != nil {
		t.Error(err)
		return
	}
}
