package channel

import (
	pb "github.com/disism/hvxahv/api/hvxahv/v1alpha1"
	"golang.org/x/net/context"
	"log"
)

// NewChannel Implementation of the method of new channel.
func (s *server) NewChannel(ctx context.Context, in *pb.NewChannelData) (*pb.NewChannelReply, error) {
	nc := NewChannels(in.Name, in.Id, in.Avatar, in.Bio, in.Owner, in.IsPrivate)
	code, res, id, err := nc.New()
	if err != nil {
		log.Println(err)
		return &pb.NewChannelReply{Code: int32(code), Message: res, Id: id}, err
	}
	return &pb.NewChannelReply{Code: int32(code), Message: res, Id: id}, nil
}

func (s *server) NewChannelAdmin(ctx context.Context, in *pb.NewChanAdmData) (*pb.NewChanAdmReply, error) {
	nca, err := NewChanAdmins(in.Owner, in.Id, in.Admin)
	if err != nil {
		return &pb.NewChanAdmReply{Code: 500, Message: err.Error()}, nil
	}

	code, res, err := nca.AddAdmin()
	if err != nil {
		return &pb.NewChanAdmReply{Code: 500, Message: res}, nil
	}
	

	return &pb.NewChanAdmReply{Code: int32(code), Message: res}, nil
}