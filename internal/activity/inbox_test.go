package activity

import (
	"fmt"
	"testing"
)

func TestFollows_Create(t *testing.T) {
	fo := []uint{746932029260431361, 747232968277196801, 747524801241219073}
	for _, i := range fo {
		if err := NewFollows(746931986864701441, i).Create(); err != nil {
			t.Error(err)
			return
		}
	}
}

func TestFollows_Delete(t *testing.T) {
	if err := NewFollows(747524801241219073, 747232968277196801).Delete(); err != nil {
		t.Error(err)
		return
	}
}

func TestFollows_GetFollowers(t *testing.T) {
	x, err := NewFollower(746931986864701441).GetFollowers()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(x)
}

func TestFollows_GetFollowing(t *testing.T) {
	x, err := NewFollowing(747524801241219073).GetFollowings()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(x)
}
