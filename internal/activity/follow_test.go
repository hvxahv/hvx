package activity

import (
	"fmt"
	"testing"
)

func TestFollowRequests_Create(t *testing.T) {
	IniTestConfig(t)

	nf := NewFollowRequests("https://mas.to/c467d206-5aeb-4d44-8d03-98e93175fe04", 710444110081654785, 698619813575491585)
	if err := nf.Create(); err != nil {
		fmt.Println(err)
	}
}

func TestFollowAccepts_Create(t *testing.T) {
	IniTestConfig(t)

	nf := NewFollowAccepts("https://mas.to/users/hvturingga#accepts/follows/126745", 710444110081654785, 698619813575491585, "https://394c-2408-832f-20b2-6fb0-1d07-725-5183-8fa1.ngrok.io/e988acf2-0188-4839-a7b6-39e7203d8fc7")
	if err := nf.Create(); err != nil {
		fmt.Println(err)
	}
}

func TestFollows_Create(t *testing.T) {
	IniTestConfig(t)

	f := NewFollows(710444110081654785,698619813575491585)
	if err := f.Create(); err != nil {
		fmt.Println(err)
	}

	f2 := NewFollows(698619813575491585,710444110081654785)
	if err := f2.Create(); err != nil {
		fmt.Println(err)
	}
}

func TestFollows_GetFollowers(t *testing.T) {
	IniTestConfig(t)

	f, err := NewByActorID(698619813575491585).GetFollowers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
}

func TestFollows_GetFollowing(t *testing.T) {
	IniTestConfig(t)

	f, err := NewByActorID(698619813575491585).GetFollowing()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
}

func TestFollowRequests_Delete(t *testing.T) {
	IniTestConfig(t)

	if err := NewFollowRequestsActivityID("https://mas.to/2d7db601-f78e-40f8-b9cb-5fc3e285883f").Delete(); err != nil {
		fmt.Println(err)
		return 
	}
}

func TestFollows_Remove(t *testing.T) {
	IniTestConfig(t)

	if err := NewFollows(698619814874251265, 710444110081654785).Remove(); err != nil {
		fmt.Println(err)
	}
}

func TestFollowRequests_GetDetailsByID(t *testing.T) {
	IniTestConfig(t)

	r, err := NewFollowRequestsByID(714081785916784641).GetDetailsByID()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}