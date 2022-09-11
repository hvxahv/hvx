package clientv1

import pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"

type Activity interface {
	Inbox(name string, body []byte) (*pb.InboxResponse, error)
}

func (svc *Svc) Inbox(name string, body []byte) (*pb.InboxResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	i, err := pb.NewInboxClient(c.Conn).Inbox(svc.ctx, &pb.InboxRequest{
		Name: name,
		Data: body,
	})
	if err != nil {
		return nil, err
	}
	return i, nil
}
