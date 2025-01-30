package utils

import (
	"context"
	"fmt"
	"github.com/RA341/multipacman/models"
)

const DefaultFilePerm = 0o775

func GetUserContext(ctx context.Context) (*models.User, error) {
	userVal := ctx.Value("user")
	if userVal == nil {
		return nil, fmt.Errorf("could not find user in context")
	}
	user, ok := userVal.(*models.User)
	if !ok {
		return nil, fmt.Errorf("invalid user type in context")
	}

	return user, nil
}
