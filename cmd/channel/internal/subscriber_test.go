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
	if err := NewSubscribe(807240711882047489, 801979262380376065).AddSubscriber(801935106069495809); err != nil {
		t.Error(err)
		return
	}
}
