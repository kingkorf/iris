Iris Web Framework
===========================
 [![Build Status](https://travis-ci.org/kataras/iris.svg?branch=master&style=flat-square)](https://travis-ci.org/kataras/iris)
[![Go Report Card](https://goreportcard.com/badge/github.com/kataras/iris?style=flat-square)](https://goreportcard.com/report/github.com/kataras/iris)
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kataras/iris?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![GoDoc](https://godoc.org/github.com/kataras/iris?status.svg)](https://godoc.org/github.com/kataras/iris)
[![License](https://img.shields.io/badge/license-BSD3-blue.svg?style=flat-square)](LICENSE)


Iris: Community-driven, mini web application framework for Go Programming Language. Comes with the [highest](#benchmarks) performance so far.

**[Easy to learn](https://kataras.gitbooks.io/iris/content/)** while providing robust set of features for building **modern web applications**.

<a href="https://www.gitbook.com/book/kataras/iris/details"><img align="left" width="185" src="https://raw.githubusercontent.com/kataras/iris/gh-pages/assets/book/cover_1.png"></a>

![Hi Iris GIF](http://kataras.github.io/iris/assets/hi_iris_may.gif)



```html
<!-- ./templates/hi.html -->
<html><head> <title> Hi Django-syntax </title> </head>
  <body>
    <h1> Hi {{ Name }}
  </body>
</html>

```

```go
// ./main.go
import  "github.com/kataras/iris"

func main() {
    iris.Config().Render.Template.Engine = iris.PongoEngine
	iris.Get("/hi", hi)
	iris.Listen(":8080")
}

func hi(ctx *iris.Context){
   ctx.Render("hi.html", map[string]interface{}{"Name": "iris"})
}
```

> Learn about [configuration](https://kataras.gitbooks.io/iris/content/configuration.html) and [render](https://kataras.gitbooks.io/iris/content/render.html).



# Features

* **Switch between template engines**: Select the way you like to parse your html files, switchable via one-line-configuration, [read more](https://kataras.gitbooks.io/iris/content/render.html)
* **Typescript**: Auto-compile & Watch your client side code via the [typescript plugin](https://kataras.gitbooks.io/iris/content/plugin-typescript.html)
* **Online IDE**: Edit & Compile your client side code when you are not home via the [editor plugin](https://kataras.gitbooks.io/iris/content/plugin-editor.html)
* **Iris Online Control**: Web-based interface to control the basics functionalities of your server via the [iriscontrol plugin](https://kataras.gitbooks.io/iris/content/plugin-iriscontrol.html). Note that Iris control is still young
* **Subdomains**: Easy way to express your api via custom and dynamic subdomains[*](https://kataras.gitbooks.io/iris/content/subdomains.html)
* **Named Path Parameters**: Probably you already know what that means. If not, [It's easy to learn about](https://kataras.gitbooks.io/iris/content/named-parameters.html)
* **Custom HTTP Errors**: Define your own html templates or plain messages when http errors occurs[*](https://kataras.gitbooks.io/iris/content/custom-http-errors.html)
* **Internationalization**: [i18n](https://kataras.gitbooks.io/iris/content/middleware-internationalization-and-localization.html)
* **Bindings**: Need a fast way to convert data from body or form into an object? Take a look [here](https://kataras.gitbooks.io/iris/content/request-body-bind.html)
* **Streaming**: You have only one option when streaming comes in game[*](https://kataras.gitbooks.io/iris/content/streaming.html)
* **Middlewares**: Create and/or use global or per route middlewares with the Iris' simplicity[*](https://kataras.gitbooks.io/iris/content/middlewares.html)
* **Sessions**:  Sessions provides a secure way to authenticate your clients/users, access them via Context or the "Low-level" api [*](https://kataras.gitbooks.io/iris/content/package-sessions.html)
* **Realtime**: Realtime is fun when you use websockets[*](https://kataras.gitbooks.io/iris/content/package-websocket.html)
* **Context**: [Context](https://kataras.gitbooks.io/iris/content/context.html) is used for storing route params, storing handlers, sharing variables between middlewares, render rich content, send file and much more[*](https://kataras.gitbooks.io/iris/content/context.html)
* **Plugins**: You can build your own plugins to  inject the Iris framework[*](https://kataras.gitbooks.io/iris/content/plugins.html)
* **Full API**: All http methods are supported[*](https://kataras.gitbooks.io/iris/content/api.html)
* **Party**:  Group routes when sharing the same resources or middlewares. You can organise a party with domains too! [*](https://kataras.gitbooks.io/iris/content/party.html)
* **Transport Layer Security**: Provide privacy and data integrity between your server and the client[*](https://kataras.gitbooks.io/iris/content/tls.html)
* **Multi server instances**: Besides the fact that Iris has a default main server. You can declare as many as you need[*](https://kataras.gitbooks.io/iris/content/declaration.html)
* **Zero configuration**:  No need to configure anything, unless you're forced to. Default configurations everywhere, which you can change with ease
*  **Zero allocations**: Iris generates zero garbage


----

Install
------------
 The only requirement is Go 1.6

`$ go get -u github.com/kataras/iris`

 >If you are connected to the Internet through China [click here](https://kataras.gitbooks.io/iris/content/install.html)

How to use
------------

- Read the [book](https://www.gitbook.com/book/kataras/iris/details) or [wiki](https://github.com/kataras/iris/wiki)

- Take a look at the [examples](https://github.com/iris-contrib/examples)




If you'd like to discuss this package, or ask questions about it, feel free to

* Post an issue or  idea [here](https://github.com/kataras/iris/issues)
* [Chat]( https://gitter.im/kataras/iris) with us

Open debates

 - [E-book Cover - Which one you suggest?](https://github.com/kataras/iris/issues/67)



# Benchmarks
------------


Benchmarks results taken [from external source](https://github.com/smallnest/go-web-framework-benchmark), created by [@smallnest](https://github.com/smallnest).

This is the most realistic benchmark suite than you will find for Go Web Frameworks. Give attention to its readme.md.

May 12 2016


![Benchmark Wizzard Processing Time](http://kataras.github.io/iris/assets/benchmark_11_05_2016_different_processing_time.png)

[click here to view detailed tables of different benchmarks](https://github.com/smallnest/go-web-framework-benchmark)


Versioning
------------

Current: **v3.0.0-alpha.3**


Read more about Semantic Versioning 2.0.0

 - http://semver.org/
 - https://en.wikipedia.org/wiki/Software_versioning
 - https://wiki.debian.org/UpstreamGuide#Releases_and_Versions


Third party packages
------------

- [Iris is build on top of fasthttp](https://github.com/valyala/fasthttp)
- [pongo2 as one of the build'n template engines](https://github.com/flosch/pongo2)
- [mergo as for merge configs](https://github.com/imdario/mergo)
- [formam as form binder](https://github.com/monoculum/formam)
- [i18n for internalization](https://github.com/Unknwon/i18n)

Contributors
------------

Thanks goes to the people who have contributed code to this package, see the

- [GitHub Contributors page](https://github.com/kataras/iris/graphs/contributors).


Todo
------------
> for the next release 'v3'

- [ ] Implement a middleware or plugin for easy & secure user authentication, stored in (no)database redis/mysql and make use of [sessions](https://github.com/kataras/iris/tree/master/sessions).
- [ ] Create server & client side (js) library for .on('event', func action(...)) / .emit('event')... (like socket.io but supports only [websocket](https://github.com/kataras/iris/tree/master/websocket)).
- [x] Find and provide support for the most stable template engine and be able to change it via the configuration, keep html/templates  support.
- [ ] Extend, test and publish to the public the Iris' cmd.



I am a student at the [University of Central Macedonia](http://teiser.gr/index.php?lang=en).
I spend all my time in the construction of Iris, therefore I have no income value.

If you think that any information you obtained here is worth some money, feel free to send any amount through paypal, thanks.

[![](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=makis%40ideopod%2ecom&lc=GR&item_name=Iris%20web%20framework&item_number=iriswebframeworkdonationid2016&amount=2%2e00&currency_code=EUR&bn=PP%2dDonationsBF%3abtn_donateCC_LG%2egif%3aNonHosted)


License
------------

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
License can be found [here](https://github.com/kataras/iris/blob/master/LICENSE).
