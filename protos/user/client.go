package user_v4

import (
	"google.golang.org/grpc"
	"github.com/nyx-swit/onboarding/config"
	"swit/swit-configs-library-golang/env"
)

// Namespace : version
// Copy 이후 환경 변수명 체크 필.
func getGRPCClientConn() *grpc.ClientConn {

	env.GetString("USER_SVC")
	/*
	a := client.Address{
		IP:        env.GetString("USER_SVC"),
		Namespace: env.GetString("NS_V4"),
		Port:      env.GetString("USER_PORT"),
	}
	conn := client.NewGrpcClientProvider(a)
		*/
	*/
	return conn
}

/*
func UserClientProvider() UserServiceClient {
	return NewUserServiceClient(getGRPCClientConn())
}
*/