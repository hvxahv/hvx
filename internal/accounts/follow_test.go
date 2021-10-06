package accounts

import (
	"fmt"
	"log"
	"testing"
)

func TestNewFollowers(t *testing.T) {
	TestInitDB(t)

	nf := NewFollows(698619813575491585, 699445624156061697)
	if err := nf.New(); err != nil {
		log.Println(err)
		return
	}

	nf2 := NewFollows(699445624156061697, 698619813575491585)
	if err := nf2.New(); err != nil {
		log.Println(err)
		return
	}
}

func TestFollows_FetchFollowing(t *testing.T) {
	TestInitDB(t)

	na := NewFetchByID(698619813575491585)
	n := na.FetchFollowing()
	fmt.Println(n)
}

func TestFollows_FetchFollowers(t *testing.T) {
	TestInitDB(t)

	nf := NewFetchByID(698619813575491585)
	n := nf.FetchFollowers()
	fmt.Println(n)
}