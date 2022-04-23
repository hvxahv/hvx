/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package device

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/device/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/x"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

const serverName = "device"

type device struct {
	pb.DevicesServer
	*Devices
}

// Run starts the server. It will block until the server is shutdown. If the server fails to start, it will return an error.
func Run() error {
	log.Printf("App %s Started at %s\n", serverName, time.Now())

	// Create a new server instance.
	s := grpc.NewServer()

	pb.RegisterDevicesServer(s, &device{})

	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", x.NewService("device").GetPort()))
	if err != nil {
		return err
	}

	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Println(err)
			return
		}
	}()

	return nil
}
