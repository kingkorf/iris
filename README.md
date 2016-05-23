[![Iris Logo](http://iris-go.com/assets/iris_full_logo_2.png)](http://iris-go.com)

Fast, unopinionated, minimalist web framework for [Go Programming Language](https://github.com/golang/go).

[![Build Status](https://travis-ci.org/kataras/iris.svg?branch=master&style=flat-square)](https://travis-ci.org/kataras/iris)
[![Go Report Card](https://goreportcard.com/badge/github.com/kataras/iris)](https://goreportcard.com/report/github.com/kataras/iris)
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kataras/iris?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![Project Status](https://img.shields.io/badge/version-3.0.0_alpha7-blue.svg?style=float-square)](HISTORY.md)
[![License](https://img.shields.io/badge/license-BSD-blue.svg?style=flat-square)](LICENSE)

[![Benchmark Wizzard Processing Time](http://iris-go.com/assets/benchmark_11_05_2016_different_processing_time.png)](#benchmarks)

```sh
$ cat main.go
```
```go
package main

import  "github.com/kataras/iris"

func main() {
	iris.Get("/hi_json", func(c *iris.Context) {
		c.JSON(200, iris.Map{
			"Name": "Iris",
			"Age":  2,
		})
	})
	iris.Listen(":8080")
}
```

> Learn about [configuration](https://kataras.gitbooks.io/iris/content/configuration.html) and [render](https://kataras.gitbooks.io/iris/content/render.html).



Installation
------------
 The only requirement is Go 1.6

`$ go get -u github.com/kataras/iris`

 >If you are connected to the Internet through China [click here](https://kataras.gitbooks.io/iris/content/install.html)

Features
------------
- Focus on high performance
- Robust routing & subdomains
- View system supporting [5+](https://kataras.gitbooks.io/iris/content/render_templates.html) template engines
- Highly scalable Websocket & Sessions API
- Middlewares & Plugins were never be easier
- Full REST API
- Custom HTTP Errors
- Typescript compiler + Browser editor
- Content negotiation & streaming
- Transport Layer Security

Docs & Community
------------

- Read the [book](https://www.gitbook.com/book/kataras/iris/details) or [wiki](https://github.com/kataras/iris/wiki)

- Take a look at the [examples](https://github.com/iris-contrib/examples)




If you'd like to discuss this package, or ask questions about it, feel free to

* Post an issue or  idea [here](https://github.com/kataras/iris/issues)
* [Chat]( https://gitter.im/kataras/iris) with us

Open debates

 - [E-book Cover - Which one you suggest?](https://github.com/kataras/iris/issues/67)

**TIP** Be sure to read the [history](HISTORY.md) for Migrating from 2.x to 3.x.

#### Table of Contents
<a href="https://www.gitbook.com/book/kataras/iris/details"><img align="right" width="185" src="http://iris-go.com/assets/book/cover_1.png"></a>

* [Hi](https://kataras.gitbooks.io/iris/content/hi.html)
* [Transport Layer Security](https://kataras.gitbooks.io/iris/content/tls.html)
* [Handlers](https://kataras.gitbooks.io/iris/content/handlers.html)
   * [Using Handlers](https://kataras.gitbooks.io/iris/content/using-handlers.html)
   * [Using HandlerFuncs](https://kataras.gitbooks.io/iris/content/using-handlerfuncs.html)
   * [Using native http.Handler](https://kataras.gitbooks.io/iris/content/using-native-httphandler.html)
       * [Using native http.Handler via iris.ToHandlerFunc()](https://kataras.gitbooks.io/iris/content/using-native-httphandler-via-tohandlerfunc.html)
* [Middlewares](https://kataras.gitbooks.io/iris/content/middlewares.html)
* [API](https://kataras.gitbooks.io/iris/content/api.html)
* [Declaration](https://kataras.gitbooks.io/iris/content/declaration.html)
* [Configuration](https://kataras.gitbooks.io/iris/content/configuration.html)
* [Party](https://kataras.gitbooks.io/iris/content/party.html)
* [Subdomains](https://kataras.gitbooks.io/iris/content/subdomains.html)
* [Named Parameters](https://kataras.gitbooks.io/iris/content/named-parameters.html)
* [Static files](https://kataras.gitbooks.io/iris/content/static-files.html)
* [Send files](https://kataras.gitbooks.io/iris/content/send-files.html)
* [Render](https://kataras.gitbooks.io/iris/content/render.html)
   * [REST](https://kataras.gitbooks.io/iris/content/render_rest.html)
   * [Templates](https://kataras.gitbooks.io/iris/content/render_templates.html)
* [Gzip](https://kataras.gitbooks.io/iris/content/gzip.html)
* [Streaming](https://kataras.gitbooks.io/iris/content/streaming.html)
* [Cookies](https://kataras.gitbooks.io/iris/content/cookies.html)
* [Flash messages](https://kataras.gitbooks.io/iris/content/flashmessages.html)
* [Body binder](https://kataras.gitbooks.io/iris/content/request-body-bind.html)
* [Custom HTTP Errors](https://kataras.gitbooks.io/iris/content/custom-http-errors.html)
* [Context](https://kataras.gitbooks.io/iris/content/context.html)
* [Logger](https://kataras.gitbooks.io/iris/content/logger.html)
* [HTTP access control](https://kataras.gitbooks.io/iris/content/middleware-cors.html)
* [Secure](https://kataras.gitbooks.io/iris/content/middleware-secure.html)
* [Sessions](https://kataras.gitbooks.io/iris/content/package-sessions.html)
* [Websockets](https://kataras.gitbooks.io/iris/content/package-websocket.html)
* [Graceful](https://kataras.gitbooks.io/iris/content/package-graceful.html)
* [Recovery](https://kataras.gitbooks.io/iris/content/middleware-recovery.html)
* [Plugins](https://kataras.gitbooks.io/iris/content/plugins.html)
* [Internationalization and Localization](https://kataras.gitbooks.io/iris/content/middleware-internationalization-and-localization.html)
* [Easy Typescript](https://kataras.gitbooks.io/iris/content/plugin-typescript.html)
* [Browser based Editor](https://kataras.gitbooks.io/iris/content/plugin-editor.html)
* [Routes info](https://kataras.gitbooks.io/iris/content/plugin-routesinfo.html)
* [Control panel](https://kataras.gitbooks.io/iris/content/plugin-iriscontrol.html)
* [Examples](https://github.com/iris-contrib/examples)


Philosophy
------------

The Iris philosophy is to provide robust tooling for HTTP, making it a great solution for single page applications, web sites, hybrids, or public HTTP APIs.

Iris does not force you to use any specific ORM or template engine. With support for the most used template engines, you can quickly craft the perfect application.

Benchmarks
------------

[This Benchmark suite]((https://github.com/smallnest/go-web-framework-benchmark)) aims to compare the whole HTTP request processing between Go web frameworks.

![Benchmark Wizzard Processing Time](http://iris-go.com/assets/benchmark_11_05_2016_different_processing_time.png)

[Click here to view all benchmarks.](https://github.com/smallnest/go-web-framework-benchmark)

Testing
------------

Iris suggests you to use [this, new,](https://github.com/gavv/httpexpect) suite to test your API.
[Httpexpect](https://github.com/gavv/httpexpect) supports fasthttp & Iris after [recommandation](https://github.com/gavv/httpexpect/issues/2). Its author is active and a very good developer so I believe its a promising library. You can see an Iris example [here](https://github.com/gavv/httpexpect/blob/master/example/iris_test.go) .

Versioning
------------

[Current](HISTORY.md): **v3.0.0-alpha.7**
>  Iris is an active project


Read more about Semantic Versioning 2.0.0

 - http://semver.org/
 - https://en.wikipedia.org/wiki/Software_versioning
 - https://wiki.debian.org/UpstreamGuide#Releases_and_Versions


Todo
------------
> for the next release 'v3'

- [ ] Create server & client side (js) library for .on('event', func action(...)) / .emit('event')... (like socket.io but supports only [websocket](https://github.com/kataras/iris/tree/master/websocket)).
- [x] Find and provide support for the most stable template engine and be able to change it via the configuration, keep html/templates  support.
- [ ] Extend, test and publish to the public the Iris' cmd.


People
------------
The author of Iris is [kataras](https://twitter.com/MakisMaropoulos)

[List of all contributors](https://github.com/kataras/iris/graphs/contributors).


I spend all my time in the construction of Iris, therefore I have no income value.
I cannot support this project alone for a long period without your support.

If you,

- think that any information you obtained here is worth some money
- believe that Iris worths to remains a highly active project

feel free to send any amount through paypal

[![](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=makis%40ideopod%2ecom&lc=GR&item_name=Iris%20web%20framework&item_number=iriswebframeworkdonationid2016&amount=2%2e00&currency_code=EUR&bn=PP%2dDonationsBF%3abtn_donateCC_LG%2egif%3aNonHosted)


### More about your donations

**Thank you**!

I'm  grateful for all the generous donations. Iris is fully funded by these donations.

#### Donors

- [Ryan Brooks](https://github.com/ryanbyyc) donated 50 EUR at May 11

> The name of the donator added after his/her permission.

#### Report, so far

- 13 EUR for the domain, [iris-go.com](https://iris-go.com)


**Available**: VAT(50)-13 = 47.5-13 = 34.5 EUR




Third party packages
------------

- [Iris is build on top of fasthttp](https://github.com/valyala/fasthttp)
- [pongo2 is one of the supporting template engines](https://github.com/flosch/pongo2)
- [amber is one of the supporting template engines](https://github.com/eknkc/amber)
- [jade is one of the supporting template engines](https://github.com/Joker/jade)
- [blackfriday is one of the supporting template engines](https://github.com/russross/blackfriday)
- [klauspost/gzip for faster compression](https://github.com/klauspost/compress/gzip)
- [mergo for merge configs](https://github.com/imdario/mergo)
- [formam as form binder](https://github.com/monoculum/formam)
- [i18n for internalization](https://github.com/Unknwon/i18n)

License
------------

This project is licensed under the BSD 3-Clause License.
License can be found [here](https://github.com/kataras/iris/blob/master/LICENSE).
