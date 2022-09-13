package friendship

import (
	"fmt"
	"github.com/hvxahv/hvx/cfg"
	"testing"
)

func init() {
	cfg.DefaultLocalConfig()
}
func TestFollows_CreateFollower(t *testing.T) {
	f := NewFollower(795104263552663553, 795165717415231489)
	if err := f.Follow(); err != nil {
		t.Error(err)
		return
	}
}

func TestFollows_CreateFollowing(t *testing.T) {
	f := NewFollowing(795104263552663553, 795165717415231489)
	if err := f.Follow(); err != nil {
		t.Error(err)
		return
	}
}

func TestFollows_UNFollower(t *testing.T) {
	//f := NewFollower(746931986864701441, 746932029260431361)
	//if err := f.UNFollow(); err != nil {
	//	t.Error(err)
	//	return
	//}
}

func TestFollows_UNFollowing(t *testing.T) {
	f := NewFollowing(746931986864701441, 746932029260431361)
	if err := f.UNFollow(); err != nil {
		t.Error(err)
		return
	}
}

func TestFollows_GetFollowsForFollower(t *testing.T) {
	f := NewFollows(746931986864701441, Follower)
	followers, err := f.Get()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(followers)
}

func TestFollows_GetFollowsForFollowing(t *testing.T) {
	//f := NewGetFollows(746931986864701441, "following")
	//followers, err := f.GetFollows()
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//fmt.Println(followers)
}

func TestFollows_GetFollowsForFriend(t *testing.T) {
	//f := NewGetFollows(746931986864701441, "friend")
	//followers, err := f.GetFollows()
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//fmt.Println(followers)
}
