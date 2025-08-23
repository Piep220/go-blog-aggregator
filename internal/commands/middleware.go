package commands

import (
	"context"
)

func MiddlewareLoggedIn(next handlerLoggedIn) handlerFn {
	return func(s *State, cmd Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
        if err != nil {
            return err
        }
        return next(s, cmd, user)
	}
}