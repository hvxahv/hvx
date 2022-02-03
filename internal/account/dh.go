package account

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/hvxahv/hvxahv/api/notify/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/notify"
	"github.com/hvxahv/hvxahv/pkg/cache"
	"github.com/hvxahv/hvxahv/pkg/push"
	"golang.org/x/net/context"
)

type dh struct {
	DeviceID  string
	TO        string
	PublicKey string
	IV        string
}

type send struct {
	DeviceID    string
	DHPublicKey string
	PrivateKey  string
}

func (a *account) DHRequestEncryption(ctx context.Context, in *pb.NewDHRequestEncryption) (*pb.DHEncryptionReply, error) {
	var buf bytes.Buffer
	encode := gob.NewEncoder(&buf)
	if err := encode.Encode(&dh{
		DeviceID:  in.DeviceId,
		TO:        in.To,
		PublicKey: in.DhPublicKey,
		IV:        in.DhIv,
	}); err != nil {
		return nil, err
	}
	if err := cache.SETDHData(in.To, buf.Bytes()); err != nil {
		return nil, err
	}

	client, err := notify.NewNotifyClient()
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(push.NewData(
		"Notify",
		fmt.Sprintf("YOU_HAVE_A_NEW_LOGIN_REQUEST"),
		"https://avatars.githubusercontent.com/u/94792300?s=200&v=4",
		"Authorized"),
	)
	if err != nil {
		return nil, err
	}
	d := &v1alpha1.NewNotifyPush{
		DeviceId: in.To,
		Data:     data,
	}
	reply, err := client.Push(ctx, d)
	if err != nil {
		return nil, err
	}
	return &pb.DHEncryptionReply{Code: reply.Code, Reply: reply.Reply}, nil
}

func (a *account) DHGetPublic(ctx context.Context, in *pb.NewDHGetPublic) (*pb.DHGetPublicReply, error) {
	data, err := cache.GETDHData(in.DeviceId)
	if err != nil {
		return nil, err
	}
	var d dh
	decode := gob.NewDecoder(bytes.NewReader(data))
	if err := decode.Decode(&d); err != nil {
		return nil, err
	}

	return &pb.DHGetPublicReply{
		Code:      "200",
		DeviceId:  d.DeviceID,
		Iv:        d.IV,
		PublicKey: d.PublicKey,
	}, nil
}

func (a *account) DHSendEncryption(ctx context.Context, in *pb.NewDHSendEncryption) (*pb.DHEncryptionReply, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(&send{
		DeviceID:    in.DeviceId,
		DHPublicKey: in.DhPublicKey,
		PrivateKey:  in.PrivateKey,
	}); err != nil {
		return nil, err
	}

	if err := cache.SETDHData(in.DeviceId, buf.Bytes()); err != nil {
		return nil, err
	}
	return &pb.DHEncryptionReply{
		Code:  "200",
		Reply: "ok",
	}, nil
}

func (a *account) DHWaitEncryption(ctx context.Context, in *pb.NewDHWaitEncryption) (*pb.DHWaitEncryptionReply, error) {
	data, err := cache.GETDHData(in.DeviceId)
	if err != nil {
		return nil, err
	}
	var buf send
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&buf); err != nil {
		return nil, err
	}

	return &pb.DHWaitEncryptionReply{
		Code:        "200",
		DhPublicKey: buf.DHPublicKey,
		PrivateKey:  buf.PrivateKey,
	}, nil
}
