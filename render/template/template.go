package template

import (
	"fmt"

	"github.com/kataras/iris/config"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/render/template/engine/html"
	"github.com/kataras/iris/render/template/engine/pongo"
)

type (
	Engine interface {
		BuildTemplates() error
		Execute(ctx context.IContext, name string, binding interface{}, layout string) error
		ExecuteGzip(ctx context.IContext, name string, binding interface{}, layout string) error
	}

	Template struct {
		Engine Engine

		IsDevelopment bool
		Gzip          bool
		ContentType   string
		Layout        string
	}
)

// New creates and returns a Template instance which keeps the Template Engine and helps with render
func New(c config.Template) *Template {

	var e Engine
	// [ENGINE-2]
	switch c.Engine {
	case config.HTMLEngine:
		e = html.New(c) //  HTMLTemplate
	case config.PongoEngine:
		e = pongo.New(c) // Pongo2
	default: // it's the config.NoEngine
		return nil
	}

	if err := e.BuildTemplates(); err != nil { // first build the templates, if error then panic because this is called before server's run
		panic(err)
	}

	compiledContentType := c.ContentType + "; charset=" + c.Charset

	return &Template{
		Engine:        e,
		IsDevelopment: c.IsDevelopment,
		Gzip:          c.Gzip,
		ContentType:   compiledContentType,
		Layout:        c.Layout,
	}

}

func (t *Template) Render(ctx context.IContext, name string, bindings interface{}, layout ...string) error {

	if t == nil { // No engine was given but .Render was called
		ctx.WriteHTML(403, "<b> Iris </b> <br/> Templates are disabled via config.NoEngine, check your iris' configuration please.")
		return fmt.Errorf("[IRIS TEMPLATES] Templates are disabled via config.NoEngine, check your iris' configuration please.\n")
	}

	// build templates again on each render if IsDevelopment.
	if t.IsDevelopment {
		if err := t.Engine.BuildTemplates(); err != nil {
			return err
		}
	}
	ctx.GetRequestCtx().Response.Header.Set("Content-Type", t.ContentType)
	// I don't like this, something feels wrong
	_layout := ""
	if len(layout) > 0 {
		_layout = layout[0]
	}
	if _layout == "" {
		_layout = t.Layout
	}

	//

	if t.Gzip {
		return t.Engine.ExecuteGzip(ctx, name, bindings, _layout)
	}

	return t.Engine.Execute(ctx, name, bindings, _layout)

}
