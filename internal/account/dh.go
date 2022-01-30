package account

import (
	"bytes"
	"encoding/gob"
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/cache"
	"golang.org/x/net/context"
)

type req struct {
	DeviceHash string
	PublicKey  string
	IV         string
}

type send struct {
	DeviceHash string
	PublicKey  string
	PrivateKey string
}

func (a *account) RequestEncryption(ctx context.Context, in *pb.NewRequestEncryption) (*pb.EcdhEncryptionReply, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(&req{
		DeviceHash: in.RequestDeviceHash,
		PublicKey:  in.DhPublicKey,
		IV:         in.DhIv,
	}); err != nil {
		return nil, err
	}
	fmt.Println(buf.Bytes())
	//if err := cache.SETDH(req, buf.Bytes()); err != nil {
	//	log.Println(err)
	//	return
	//}

	//d, err := json.Marshal(push.NewData(
	//	"Notify",
	//	fmt.Sprintf("You are preparing to login on another device: %s.", deviceID),
	//	"https://avatars.githubusercontent.com/u/94792300?s=200&v=4",
	//	"Authorized"),
	//)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//if err := notify.NewPush(a.ID, hash.ID, d).Push(); err != nil {
	//	log.Println(err)
	//	return
	//}
	return nil, nil
}

func (a *account) SendEncryption(ctx context.Context, in *pb.NewSendEncryption) (*pb.EcdhEncryptionReply, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(send{
		DeviceHash: in.DeviceHash,
		PublicKey:  in.DhPublicKey,
		PrivateKey: in.PrivateKeyAes,
	}); err != nil {
		return nil, err
	}
	fmt.Println(buf.Bytes())
	//if err := cache.SETDH(in.DeviceHash, buf.Bytes()); err != nil {
	//	return nil, err
	//}
	return nil, nil
}

func (a *account) WaitEncryption(ctx context.Context, in *pb.NewWaitEncryptionDeviceID) (*pb.WaitEncryptionReply, error) {
	data, err := cache.GETDHData(in.DeviceHash)
	if err != nil {
		return nil, err
	}
	var buf send
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&buf); err != nil {
		return nil, err
	}
	fmt.Println(data)
	return nil, nil
}
