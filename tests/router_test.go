package tests

import (
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/gavv/httpexpect/fasthttpexpect"
	"github.com/kataras/iris"
	"github.com/kataras/iris/config"
)

type route struct {
	Method   string
	Path     string
	Body     string
	Status   int
	Register bool
}

var routes = []route{
	// FOUND - registed
	route{"GET", "/test_get", "hello, get!", 200, true},
	route{"POST", "/test_post", "hello, post!", 200, true},
	route{"PUT", "/test_put", "hello, put!", 200, true},
	route{"DELETE", "/test_delete", "hello, delete!", 200, true},
	route{"HEAD", "/test_head", "hello, head!", 200, true},
	route{"OPTIONS", "/test_options", "hello, options!", 200, true},
	route{"CONNECT", "/test_connect", "hello, connect!", 200, true},
	route{"PATCH", "/test_patch", "hello, patch!", 200, true},
	route{"TRACE", "/test_trace", "hello, trace!", 200, true},
	// NOT FOUND - not registed
	route{"GET", "/test_get_nofound", "Not Found", 404, false},
	route{"POST", "/test_post_nofound", "Not Found", 404, false},
	route{"PUT", "/test_put_nofound", "Not Found", 404, false},
	route{"DELETE", "/test_delete_nofound", "Not Found", 404, false},
	route{"HEAD", "/test_head_nofound", "Not Found", 404, false},
	route{"OPTIONS", "/test_options_nofound", "Not Found", 404, false},
	route{"CONNECT", "/test_connect_nofound", "Not Found", 404, false},
	route{"PATCH", "/test_patch_nofound", "Not Found", 404, false},
	route{"TRACE", "/test_trace_nofound", "Not Found", 404, false},
}

func TestRouter(t *testing.T) {
	api := iris.New()
	// first register the routes needed
	for idx, _ := range routes {
		r := routes[idx]
		if r.Register {
			api.Handle(r.Method, r.Path, iris.HandlerFunc(func(ctx *iris.Context) {
				ctx.SetStatusCode(r.Status)
				ctx.SetBodyString(r.Body)
			}))
		}
	}
	api.PreListen(config.Server{ListeningAddr: ""})

	// create httpexpect instance that will call fasthtpp.RequestHandler directly
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client:   fasthttpexpect.NewBinder(api.ServeRequest),
	})

	// run the tests
	for idx, _ := range routes {
		r := routes[idx]
		e.Request(r.Method, r.Path).
			Expect().
			Status(r.Status).Body().Equal(r.Body)
	}

}
