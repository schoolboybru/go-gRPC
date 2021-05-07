package chat

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	RegisterChatServiceServer(s, &Server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}
func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestSayHello(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := NewChatServiceClient(conn)

	message := Message {
		Body: "Hello from the client",
	}
	resp, err := client.SayHello(ctx, &message)

	if err != nil {
		t.Fatalf("SayHello failed %v", err)
	}

	log.Printf("Response: %+v", resp)
}

func TestLoadMessages(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	client := NewChatServiceClient(conn)

	//message := Message {
	//	Body: "First message",
	//}
	//
	//secondMessage := Message {
	//	Body: "Second message",
	//}
	//
	//messages := []*Message {
	//	&message,
	//	&secondMessage,
	//}

	//messageResponse := MessageResponse{
	//	Messages: messages,
	//}

	itemQuery := ItemQuery{
		Id: 1,
	}

	resp, err := client.LoadMessages(ctx, &itemQuery)

	if err != nil {
		t.Fatalf("failed to load messages: %v", err)
	}

	log.Printf("Response: %v", resp)
}
