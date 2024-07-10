package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/redbox12/gRPS_simple/proto/notification"
	"google.golang.org/grpc"
)

type server struct {
	//обеспечивает реализацию всех методов интерфейса NotificationServiceServer, которые вы не реализовали.
	// Это позволяет избежать ошибки, связанной с отсутствием метода 
	notification.UnimplementedNotificationServiceServer
}

// Описываем метод Notify ответа
func (s *server) Notify(ctx context.Context, n *notification.NotificationRequest) (*notification.NotificationResponse, error) {
	fmt.Println("RECEIVED NOTIFICATION:", n.Message)
	return &notification.NotificationResponse{Status: "OK"}, nil
}

func main() {
	//Создаем tcp сервер
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	//регистрируем сервер, который создали выше
	notification.RegisterNotificationServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
