package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/message"
)

type server struct {
	pb.MessagesServer
}

const serviceName = "message"
