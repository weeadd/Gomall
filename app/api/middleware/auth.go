package middleware

import (
	"context"

	apiutils "Gomall/app/api/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type SessionUserIdKey string

const SessionUserId SessionUserIdKey = "user_id"

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		ctx = context.WithValue(ctx, apiutils.SessionUserId, session.Get("user_id"))
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			c.JSON(401, "Unauthorized")
			return
		}
		c.Next(ctx)
	}
}
