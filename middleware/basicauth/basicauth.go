package basicauth

import (
	"encoding/base64"
	"strconv"

	"github.com/kataras/iris"
	"github.com/kataras/iris/config"
)

type (
	encodedUser struct {
		HeaderValue string
		Username    string
	}
	encodedUsers []encodedUser

	basicAuthMiddleware struct {
		config config.BasicAuth
		// these are filled from the config.Users map at the startup
		auth             encodedUsers
		realmHeaderValue string
	}
)

//

// New takes one parameter, the config.BasicAuth returns a HandlerFunc
// use: iris.UseFunc(New(...)), iris.Get(...,New(...),...)
func New(c config.BasicAuth) iris.HandlerFunc {
	return NewHandler(c).Serve
}

// NewHandler takes one parameter, the config.BasicAuth returns a Handler
// use: iris.Use(NewHandler(...)), iris.Get(...,iris.HandlerFunc(NewHandler(...)),...)
func NewHandler(c config.BasicAuth) iris.Handler {
	b := &basicAuthMiddleware{config: config.DefaultBasicAuth().MergeSingle(c)}
	b.init()
	return b
}

// Default takes one parameter, the users returns a HandlerFunc
// use: iris.UseFunc(Default(...)), iris.Get(...,Default(...),...)
func Default(users map[string]string) iris.HandlerFunc {
	return DefaultHandler(users).Serve
}

// DefaultHandler takes one parameter, the users returns a Handler
// use: iris.Use(DefaultHandler(...)), iris.Get(...,iris.HandlerFunc(Default(...)),...)
func DefaultHandler(users map[string]string) iris.Handler {
	c := config.DefaultBasicAuth()
	c.Users = users
	return NewHandler(c)
}

//

func (b *basicAuthMiddleware) init() {
	// pass the encoded users from the user's config's Users value
	b.auth = make(encodedUsers, 0, len(b.config.Users))

	for k, v := range b.config.Users {
		fullUser := k + ":" + v
		header := "Basic " + base64.StdEncoding.EncodeToString([]byte(fullUser))
		b.auth = append(b.auth, encodedUser{HeaderValue: header, Username: k})
	}

	// set the auth realm header's value
	b.realmHeaderValue = "Basic realm=" + strconv.Quote(b.config.Realm)
}

func (b *basicAuthMiddleware) findUsername(headerValue string) (username string, found bool) {
	if len(headerValue) == 0 {
		return
	}

	for _, user := range b.auth {
		if user.HeaderValue == headerValue {
			username = user.Username
			found = true
			break
		}
	}

	return
}

// Serve the actual middleware
func (b *basicAuthMiddleware) Serve(ctx *iris.Context) {

	if username, found := b.findUsername(ctx.RequestHeader("Authorization")); !found {
		ctx.SetHeader("WWW-Authenticate", b.realmHeaderValue)
		ctx.SetStatusCode(iris.StatusUnauthorized)
		// don't continue to the next handler
	} else {
		// all ok set the context's value in order to be getable from the next handler
		ctx.Set(b.config.ContextKey, username)
		ctx.Next() // continue
	}

}
