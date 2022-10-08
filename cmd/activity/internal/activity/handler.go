package activity

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
)

type Object struct {
	Address string
	Inbox   string
	Err     error
}

func NewObject(address string) *Object {
	object, err := clientv1.New(context.Background(), microsvc.ActorServiceName).GetActorByAddress(address)
	if err != nil {
		return &Object{Err: err}
	}
	return &Object{Address: address, Inbox: object.GetInbox()}
}

func NewObjectId(id uint) *Object {
	object, err := clientv1.New(context.Background(), microsvc.ActorServiceName).GetActor(int64(id))
	if err != nil {
		return &Object{Err: err}
	}
	return &Object{Address: object.Actor.Address, Inbox: object.Actor.Inbox}
}

type Actor struct {
	Id      int64
	Address string
	Inbox   string

	// PrivateKey used to encrypt activitypub delivery requests.
	// It is stored in the account and the corresponding public key is stored in the actor.
	PrivateKey  string
	PublicKeyId string
	Err         error
}

func NewActor(actorId, accountId int64) *Actor {
	ctx := context.Background()
	actors, err := clientv1.New(ctx, microsvc.ActorServiceName).GetActor(actorId)
	if err != nil {
		fmt.Println("actor", err)
		return &Actor{Err: err}
	}
	accounts, err := clientv1.New(ctx, microsvc.AccountServiceName).GetPrivateKey(accountId)
	if err != nil {
		return &Actor{Err: errors.NewFailedToConnect("account")}
	}

	return &Actor{
		Id:          actors.Actor.Id,
		Address:     actors.Actor.GetAddress(),
		Inbox:       actors.Actor.GetInbox(),
		PrivateKey:  accounts.GetPrivateKey(),
		PublicKeyId: fmt.Sprintf("%s#main-key", actors.Actor.Address),
	}
}

type Handler struct {
	Object     Object
	Actor      Actor
	ActivityId string
	Err        error
	Errors     chan error
	Successes  chan string
	Failures   chan string
	Closed     chan bool
}

func NewHandler(object Object, actor Actor) (*Handler, error) {
	if object.Err != nil {
		return nil, errors.NewFailedToConnect("actor")
	}
	if actor.Err != nil {
		return nil, errors.Newf("ACTIVITY_HANDLER", actor.Err)
	}

	return &Handler{
		Object:     object,
		Actor:      actor,
		ActivityId: fmt.Sprintf("%s/%s", actor.Address, uuid.NewString()),
	}, nil
}
