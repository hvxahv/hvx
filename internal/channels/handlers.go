package channels

//New Implementation of the method of new channels.
//func (s *server) New(ctx context.Context, in *pb.NewChannelData) (*pb.NewChannelReply, error) {
//	nc := NewChannels(in.Name, in.Id, in.Avatar, in.Bio, in.Owner, in.IsPrivate)
//	code, res, id, err := nc.New()
//	if err != nil {
//		log.Println(err)
//		return &pb.NewChannelReply{Code: int32(code), Message: res, Id: id}, err
//	}
//	return &pb.NewChannelReply{Code: int32(code), Message: res, Id: id}, nil
//}
//
//func (s *server) NewAdmin(ctx context.Context, in *pb.NewAdminData) (*pb.ChannelReply, error) {
//	nca, err := NewAddAdmins(in.Owner, in.Id, in.Admin)
//	if err != nil {
//		return &pb.ChannelReply{Code: 500, Message: err.Error()}, nil
//	}
//
//	code, res, err := nca.AddAdmin()
//	if err != nil {
//		return &pb.ChannelReply{Code: int32(code), Message: res}, nil
//	}
//
//
//	return &pb.ChannelReply{Code: int32(code), Message: res}, nil
//}
//
//func (s *server) NewSubscriber(ctx context.Context, in *pb.NewSubscriberData) (*pb.ChannelReply, error) {
//	ns, err := NewSubscriber(in.Id, in.Name)
//	if err != nil {
//		return &pb.ChannelReply{Code: int32(500), Message: err.Error()}, nil
//	}
//
//	code, res, err := ns.New()
//	if err != nil {
//		return &pb.ChannelReply{Code: int32(code), Message: res}, nil
//	}
//	return &pb.ChannelReply{Code: int32(code), Message: res}, nil
//}