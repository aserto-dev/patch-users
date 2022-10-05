package main

import (
	"context"
	"log"

	"github.com/aserto-dev/aserto-go/client"
	api "github.com/aserto-dev/go-grpc/aserto/api/v1"
	ds1 "github.com/aserto-dev/go-grpc/aserto/authorizer/directory/v1"
	"github.com/google/uuid"
	flag "github.com/spf13/pflag"
)

var (
	tenantID string
	patch    bool
	authzSvc string
	authzKey string
)

func main() {
	flag.StringVar(&authzSvc, "authz-svc", "authorizer.prod.aserto.com:8443", "authorizer service address")
	flag.StringVar(&authzKey, "authz-key", "", "authorizer API key")
	flag.StringVar(&tenantID, "tenant-id", "", "patch tenant with ID")
	flag.BoolVar(&patch, "patch", false, "patch user")

	flag.Parse()

	ctx := context.Background()

	conn, err := client.NewConnection(ctx,
		client.WithAddr(authzSvc),
		client.WithAPIKeyAuth(authzKey),
		client.WithTenantID(tenantID),
	)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	clnt := ds1.NewDirectoryClient(conn.Conn)

	if err := listUsers(ctx, clnt, tenantID, patch); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func listUsers(ctx context.Context, clnt ds1.DirectoryClient, tenantID string, patch bool) error {

	resp, err := clnt.ListUsers(ctx, &ds1.ListUsersRequest{})
	if err != nil {
		return err
	}

	for _, user := range resp.Results {
		if !isValidID(user.Id) {
			log.Printf("TID:%s UID:%s %s\n", tenantID, user.Id, user.Email)
			if patch {
				if err := patchUser(ctx, clnt, user); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func isValidID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

func patchUser(ctx context.Context, clnt ds1.DirectoryClient, user *api.User) error {
	oldID := user.GetId()

	if _, err := clnt.DeleteUser(ctx, &ds1.DeleteUserRequest{Id: oldID}); err != nil {
		log.Printf("delete user %s failed with error %s", oldID, err)
		return err
	}

	// remove existing ID from user object, before recreating, this will trigger new ID creation.
	user.Id = ""

	resp, err := clnt.CreateUser(ctx, &ds1.CreateUserRequest{User: user})
	if err != nil {
		log.Printf("create user %s failed with error %s", user.Id, err)
		return err
	}

	log.Printf("user %s recreated with id:%s", oldID, resp.Result.Id)

	return nil
}
