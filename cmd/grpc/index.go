
package grpcclient

// import (
// 	"log"
// 	"sync"

// 	pb "github.com/NithishNithi/GoTask/proto"
// 	"google.golang.org/grpc"
// )

// var once sync.Once

// type GrpcClient pb.GoTaskServiceClient

// var (
// 	instance GrpcClient
// )

// func GetGrpcClientInstance() (GrpcClient, *grpc.ClientConn) {
// 	var conn *grpc.ClientConn
// 	once.Do(func() { // <-- atomic, does not allow repeating
// 		conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
// 		if err != nil {
// 			log.Fatalf
// 			("Failed to connect: %v", err)
// 		}
// 		//defer conn.Close()

// 		instance = pb.NewGoTaskServiceClient(conn)
// 	})

// 	return instance, conn
// }
