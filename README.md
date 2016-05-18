Iris Web Framework
===========================
[![Project Status](https://img.shields.io/badge/version-3.0.0_alpha5-blue.svg?style=float-square)](HISTORY.md)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kataras/iris?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![Build Status](https://travis-ci.org/kataras/iris.svg?branch=master&style=flat-square)](https://travis-ci.org/kataras/iris)
[![Go Report Card](https://goreportcard.com/badge/github.com/kataras/iris)](https://goreportcard.com/report/github.com/kataras/iris)
[![GoDoc](https://godoc.org/github.com/kataras/iris?status.svg)](https://godoc.org/github.com/kataras/iris)


A Community-driven, mini web application framework for Go Programming Language. Comes with the [highest](#benchmarks) performance achieved so far.

**[Easy to learn](https://kataras.gitbooks.io/iris/content/)** while providing robust set of features for building **modern web applications**.


```html
<!-- ./templates/hi.html -->
<h1> Hi {{ .Name }} </h1>
```

```go
// ./main.go
import  "github.com/kataras/iris"

func main() {
	iris.Get("/hi_html", func(ctx *iris.Context){
	   page := map[string]interface{}{"Name": "Iris"}
	   ctx.HTML(iris.StatusOK, "hi.html", page)
	})

	iris.Get("/hi_markdown", func(ctx *iris.Context){
	   ctx.Markdown(iris.StatusOK, "# Hi Iris")
	})

	iris.Listen(":8080")
}

```

> Learn about [configuration](https://kataras.gitbooks.io/iris/content/configuration.html) and [render](https://kataras.gitbooks.io/iris/content/render.html).



Install
------------
 The only requirement is Go 1.6

`$ go get -u github.com/kataras/iris`

 >If you are connected to the Internet through China [click here](https://kataras.gitbooks.io/iris/content/install.html)

How to use
------------
<a href="https://www.gitbook.com/book/kataras/iris/details"><img align="right" width="185" src="https://raw.githubusercontent.com/kataras/iris/gh-pages/assets/book/cover_1.png"></a>

- Read the [book](https://www.gitbook.com/book/kataras/iris/details) or [wiki](https://github.com/kataras/iris/wiki)

- Take a look at the [examples](https://github.com/iris-contrib/examples)




If you'd like to discuss this package, or ask questions about it, feel free to

* Post an issue or  idea [here](https://github.com/kataras/iris/issues)
* [Chat]( https://gitter.im/kataras/iris) with us

Open debates

 - [E-book Cover - Which one you suggest?](https://github.com/kataras/iris/issues/67)



Benchmarks
------------


Benchmarks results taken [from external source](https://github.com/smallnest/go-web-framework-benchmark), created by [@smallnest](https://github.com/smallnest).

This is the most realistic benchmark suite than you will find for Go Web Frameworks. Give attention to its readme.md.

May 12 2016


![Benchmark Wizzard Processing Time](http://kataras.github.io/iris/assets/benchmark_11_05_2016_different_processing_time.png)

[click here to view detailed tables of different benchmarks](https://github.com/smallnest/go-web-framework-benchmark)


Versioning
------------

[Current](HISTORY.md): **v3.0.0-alpha.5**
>  Iris is a project with active development


Read more about Semantic Versioning 2.0.0

 - http://semver.org/
 - https://en.wikipedia.org/wiki/Software_versioning
 - https://wiki.debian.org/UpstreamGuide#Releases_and_Versions


Third party packages
------------

- [Iris is build on top of fasthttp](https://github.com/valyala/fasthttp)
- [pongo2 as one of the build'n template engines](https://github.com/flosch/pongo2)
- [blackfriday markdown as one of the build'n template engines](https://github.com/russross/blackfriday)
- [mergo for merge configs](https://github.com/imdario/mergo)
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
Nowadays I spend all my time in the construction of Iris, therefore I have no income value.
I cannot support this project alone.

If you,

- think that any information you obtained here is worth some money
- believe that Iris worths to remains a highly active project

feel free to send any amount through paypal

[![](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=makis%40ideopod%2ecom&lc=GR&item_name=Iris%20web%20framework&item_number=iriswebframeworkdonationid2016&amount=2%2e00&currency_code=EUR&bn=PP%2dDonationsBF%3abtn_donateCC_LG%2egif%3aNonHosted)


License
------------

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
License can be found [here](https://github.com/kataras/iris/blob/master/LICENSE).

