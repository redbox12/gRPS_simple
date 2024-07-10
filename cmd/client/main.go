package main

import (
	"log"
	"time"

	"github.com/redbox12/gRPS_simple/proto/notification"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// Устанавливаем соединение с сервером gRPC
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	 
	c := notification.NewNotificationServiceClient(conn)

	// Создание контекста с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	response, err := c.Notify(ctx,&notification.NotificationRequest{Message: "Привет, ты обратился ко мне по RPC протоколу"})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("STATUS:", response.Status)
}