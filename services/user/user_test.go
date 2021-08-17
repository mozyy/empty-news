package user

import (
	"context"
	"log"
	"net"
	"reflect"
	"testing"

	"github.com/mozyy/empty-news/proto/pbuser"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pbuser.RegisterUserServer(server, New())

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func Test_user_Register(t *testing.T) {
	tests := []struct {
		name    string
		req     *pbuser.RegisterRequest
		want    *pbuser.LoginResponse
		wantErr bool
	}{
		{"success", &pbuser.RegisterRequest{Mobile: "18381335182", Password: "yyue"}, &pbuser.LoginResponse{}, false},
		{"success", &pbuser.RegisterRequest{Mobile: "18381335183", Password: "yyue"}, &pbuser.LoginResponse{}, false},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pbuser.NewUserClient(conn)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.Register(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("user.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}
