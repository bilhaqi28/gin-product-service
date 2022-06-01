package servergrpc

import (
	"context"
	"errors"
	"log"
	"time"

	pb "github.com/bilhaqi28/gin-product-service/servergrpc/model/user"
	"google.golang.org/grpc"
)

func getDataUserByToken(client pb.DataUserClient, token string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tokenGrand := pb.TokenGrand{Token: token}
	tokenJwt, err := client.GenerateJwtByToken(ctx, &tokenGrand)
	if err != nil {
		return "", errors.New("Token Is Expired")
	}
	return tokenJwt.Token, nil
}

func ReqAuthServiceToTokenGrpc(tokenGrand string) (string, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":8000", opts...)
	if err != nil {
		log.Fatalln("Error Dial", err.Error())
	}
	defer conn.Close()
	client := pb.NewDataUserClient(conn)
	result, err := getDataUserByToken(client, tokenGrand)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return result, nil
}
