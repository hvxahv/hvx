package channel

import (
	"fmt"
	"testing"

	"github.com/hvxahv/hvxahv/api/channel/v1alpha1"
	"golang.org/x/net/context"
)

func TestChannel_IsChannelAdministrator(t *testing.T) {
	c := &channel{}
	administrator, err := c.IsChannelAdministrator(context.Background(), &v1alpha1.IsChannelAdministratorRequest{
		ChannelId: "747232969484730369",
		AdminId:   "",
	})
	if err != nil {
		return
	}
	fmt.Println(administrator.IsAdministrator)
}

func TestChannel_AddAdministrator(t *testing.T) {
	c := &channel{}
	administrator, err := c.AddAdministrator(context.Background(), &v1alpha1.AddAdministratorRequest{
		ChannelId:      "747232969484730369",
		AdminAccountId: "746931987134185473",
		AddAdminId:     "746932029522116609",
		IsOwner:        false,
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(administrator.Reply)
}

func TestChannel_RemoveAdministrator(t *testing.T) {
	c := &channel{}
	administrator, err := c.RemoveAdministrator(context.Background(), &v1alpha1.RemoveAdministratorRequest{
		OwnerId:       "746931987134185473",
		ChannelId:     "747198189416415233",
		RemoveAdminId: "746932029522116609",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(administrator.Reply)
}

func TestChannel_GetAdministrators(t *testing.T) {
	c := &channel{}
	administrators, err := c.GetAdministrators(context.Background(), &v1alpha1.GetAdministratorsRequest{
		AccountId: "746931987134185473",
		ChannelId: "747198189416415233",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(administrators.Administrators)
}
