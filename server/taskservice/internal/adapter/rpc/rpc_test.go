package rpc

// import (
// 	"context"
// 	"log"
// 	"net"
// 	"testing"
// 	"time"

// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/require"
// 	pb "github.com/wralith/aestimatio/server/pb/gen/task"
// 	"github.com/wralith/aestimatio/server/taskservice/internal/core/service"
// 	"github.com/wralith/aestimatio/server/taskservice/internal/repo/inmemory"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// 	"google.golang.org/grpc/metadata"
// 	"google.golang.org/grpc/test/bufconn"
// )

// var h *GRPCHandler

// func Setup(ctx context.Context) (pb.TaskServiceClient, func()) {

// 	mem := inmemory.NewInMemoryTaskRepo()
// 	svc := service.New(mem)
// 	h = NewGRPCHandler(svc, "")

// 	buffer := 101024 * 1024
// 	lis := bufconn.Listen(buffer)
// 	srv := grpc.NewServer()

// 	go func() {
// 		if err := srv.Serve(lis); err != nil {
// 			log.Print("error while serving grpc")
// 		}
// 	}()

// 	pb.RegisterTaskServiceServer(srv, h)

// 	conn, _ := grpc.DialContext(ctx, "",
// 		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) { return lis.Dial() }),
// 		grpc.WithTransportCredentials(insecure.NewCredentials()))

// 	client := pb.NewTaskServiceClient(conn)
// 	return client, srv.Stop
// }

// func TestGRPCHandler(t *testing.T) {
// 	ctx := context.Background()

// 	id := uuid.New()
// 	claims := jwt.MapClaims{}
// 	claims["sub"] = id
// 	claims["exp"] = time.Now().Add(60 * time.Minute).Unix()
// 	claims["iss"] = "aestimatio"
// 	claims["iat"] = time.Now().Unix()

// 	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	str, err := tkn.SignedString([]byte("test-secret"))
// 	require.NoError(t, err)
// 	ctx = metadata.AppendToOutgoingContext(ctx, "jwt", str)

// 	client, stop := Setup(ctx)
// 	defer stop()

// 	in := &pb.CreateTaskRequest{
// 		Title:       "Test",
// 		Description: "Test Description",
// 		DeadlineAt:  time.Now().Add(24 * 5 * time.Hour).Unix(),
// 	}

// 	// TODO: How to pass ctx in the test?
// 	gotMd := metadata.MD{}
// 	send, _ := metadata.FromOutgoingContext(ctx)

// 	require.Equal(t, send, gotMd)
// 	got, err := client.CreateTask(ctx, in)
// 	require.Equal(t, "", gotMd)
// 	require.NoError(t, err)
// 	require.NotNil(t, got)
// }
