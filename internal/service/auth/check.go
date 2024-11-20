package auth

import (
	"context"
	"fmt"
	"slices"
)

var endpoitsByRole = map[string][]string{
	"USER": {
		"/user_v1.UserV1/Get",
		"/user_v1.UserV1/Delete",
	},
	"ADMIN": {
		"/user_v1.UserV1/Update",
		"/user_v1.UserV1/Get",
		"/user_v1.UserV1/Delete",
	},
}

func (s *authServ) Check(ctx context.Context, role, endpoint string) error {
	fmt.Println(role)
	endpointsList, ok := endpoitsByRole[role]
	if !ok {
		return fmt.Errorf("role doesnt found: %s", role)
	}

	if !slices.Contains(endpointsList, endpoint) {
		return fmt.Errorf("for role: %s access denied to: %s", role, endpoint)
	}

	return nil
}
