# go 语言的 template 包使用

- [go 语言的 template 包使用](#go-语言的-template-包使用)
  - [1. 说明](#1-说明)
    - [1.1 参考链接](#11-参考链接)
    - [1.2 Go语言的模板引擎](#12-go语言的模板引擎)
    - [1.3 模板引擎的作用](#13-模板引擎的作用)
  - [2. 模板引擎的使用](#2-模板引擎的使用)
    - [2.1 定义模板文件](#21-定义模板文件)
    - [2.2 解析模板文件](#22-解析模板文件)
    - [2.3 模板渲染](#23-模板渲染)
  - [3. 代码使用示例](#3-代码使用示例)
  - [4. 模板语法](#4-模板语法)
    - [4.1 {{.}}](#41-)
    - [4.2 注释](#42-注释)
    - [4.3 管道(pipeline)](#43-管道pipeline)
    - [4.4 变量](#44-变量)
    - [4.5 移除空格](#45-移除空格)
    - [4.6 条件判断: if-else-end](#46-条件判断-if-else-end)
    - [4.7 遍历: range-end](#47-遍历-range-end)
    - [4.7 作用域: with-end](#47-作用域-with-end)
  - [5. 模板引擎函数](#5-模板引擎函数)
    - [5.1 比较函数](#51-比较函数)
    - [5.2 预定义的全局函数](#52-预定义的全局函数)
    - [5.3 自定义函数](#53-自定义函数)
  - [6. 模板嵌套](#6-模板嵌套)
    - [6.1 嵌套模板文件](#61-嵌套模板文件)
    - [6.2 嵌套 `define` 定义的模板](#62-嵌套-define-定义的模板)
  - [7. 模板继承](#7-模板继承)
  - [8. 修改默认的标识符](#8-修改默认的标识符)
  - [9. text/template 与 html/tempalte 的区别](#9-texttemplate-与-htmltempalte-的区别)
    - [9.1 演示demo](#91-演示demo)

## 1. 说明

### 1.1 参考链接

- https://www.liwenzhou.com/posts/Go/template/

### 1.2 Go语言的模板引擎

go 语言内置了文本模板引擎 text/template 和用于HTML文档的 html/template。

- html/template 实现了数据驱动模板，用于生成可防止代码注入的安全的HTML内容
- html/templat 提供了和 text/template 包相同的接口
- Go语言中输出HTML的场景都应使用html/template这个包

### 1.3 模板引擎的作用

- 模板文件通常定义为.tmpl和.tpl为后缀（也可以使用其他的后缀），必须使用UTF8编码
- 模板文件中使用{{和}}包裹和标识需要传入的数据
- 传给模板这样的数据就可以通过点号（.）来访问，如果数据是复杂类型的数据，可以通过{ { .FieldName }}来访问它的字段
- 除{{和}}包裹的内容外，其他内容均不做修改原样输出

## 2. 模板引擎的使用

### 2.1 定义模板文件

定义模板文件时需要我们按照相关语法规则去编写即可，一般定义模板后缀为 .tmpl 或 .tpl

### 2.2 解析模板文件

解析定义好的模板，得到模板对象

```go
func (t *Template) Parse(src string) (*Template, error)
func ParseFiles(filenames ...string) (*Template, error)
func ParseGlob(pattern string) (*Template, error)
```

当然，你也可以使用func New(name string) *Template函数创建一个名为name的模板，然后对其调用上面的方法去解析模板字符串或模板文件。

### 2.3 模板渲染

渲染模板简单来说就是使用数据去填充模板

```go
func (t *Template) Execute(wr io.Writer, data interface{}) error
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
```

## 3. 代码使用示例

- 定义模板

```go
<h1>{{.title}}</h1>
<hr>
<h2>user struct</h2>
<p>ID : {{.user.ID}}</p>
<p>姓名 : {{.user.Name}}</p>
<p>性别 : {{.user.Gender}}</p>
<hr>
<h2>map[string]any</h2>
<p>ID : {{.mp.id}}</p>
<p>姓名 : {{.mp.name}}</p>
<p>性别 : {{.mp.gender}}</p>
```

- 解析并渲染模板

```go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ErrorPage struct {
	Code   int64
	ErrMsg string
}

func NewError(code int64, err error) *ErrorPage {
	return &ErrorPage{
		Code:   code,
		ErrMsg: err.Error(),
	}
}

func PageError(w http.ResponseWriter, inErr error, code ...int64) {
	// 解析模板
	tpl, err := template.ParseFiles("./tpl/error.tpl")
	if err != nil {
		fmt.Fprintln(w, "解析错误文件失败")
		return
	}
	// 错误代码处理
	var errCode int64
	if len(code) > 0 {
		errCode = code[0]
	} else {
		errCode = 500
	}
	// 渲染模板
	err = tpl.Execute(w, NewError(errCode, inErr))
	if err != nil {
		fmt.Fprintln(w, "渲染模板失败")
		return
	}
}

// 要在模板中渲染的结构体字段必须大写
type User struct {
	ID     int64
	Gender string
	Name   string
}

func PageTemplateMutil(w http.ResponseWriter, r *http.Request) {
	// 数据准备
	user := User{
		ID:     88,
		Name:   "bing",
		Gender: "female",
	}
	mp := map[string]any{
		"id":     88,
		"name":   "bing",
		"gender": "male",
	}
	title := "Mutil Title"

	// 解析模板
	tpl, err := template.ParseFiles("./tpl/mutil.tpl")
	if err != nil {
		PageError(w, err)
		return
	}
	// 渲染模板
	err = tpl.Execute(w, map[string]any{
		"user":  user,
		"mp":    mp,
		"title": title,
	})
	if err != nil {
		PageError(w, err)
		return
	}
}

func main() {
	// 注册路由函数 PrintMsg
	http.HandleFunc("/struct", PageTemplateMutil)

	// 启动服务
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Printf("Http serve failed, err: %+v", err)
	}
}
```

## 4. 模板语法

### 4.1 {{.}}

模板语法都包含在`{{`和`}}`中间，其中`{{.}}`中的点表示当前对象。当我们传入一个结构体对象时，我们可以根据.来访问结构体的对应字段

- 渲染

```go
user := User{
	ID:     88,
	Name:   "bing",
	Gender: "female",
}
err = tpl.Execute(w, user)
```

- 模板文件

```html
<p>User : {{.}}</p>
<p>ID : {{.ID}}</p>
<p>姓名 : {{.Name}}</p>
<p>性别 : {{.Gender}}</p>
```

### 4.2 注释

注释，执行时会忽略。可以多行。注释不能嵌套，并且必须紧贴分界符始止。

```html
{{/* a comment */}}
```

### 4.3 管道(pipeline)

go 的模板语法中，pipeline 的概念是传递数据，只要能产生数据的，都是pipeline. eg: {{.}}, {{.Name}}

### 4.4 变量

变量用 `$` 开始定义

```html
{{ $id:=.ID }}
<p>custom id: {{ $id }}</p>
{{ $str:="test name" }}
<p>custom var: {{ $str }}</p>
```

### 4.5 移除空格

`{{-`语法去除模板内容左侧的所有空白符号， 使用`-}}`去除模板内容右侧的所有空白符号

```html
{{- .Name -}}
```

> Note: -要紧挨{{和}}，同时与模板值之间需要使用空格分隔

### 4.6 条件判断: if-else-end

```html
{{if pipeline}} T1 {{end}}
{{if pipeline}} T1 {{else}} T0 {{end}}
{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
```

### 4.7 遍历: range-end

Go的模板语法中使用range关键字进行遍历，其中pipeline的值必须是数组、切片、字典或者通道

```html
<!-- 如果pipeline的值其长度为0，不会有任何输出 -->
{{range pipeline}} T1 {{end}}

<!-- 如果pipeline的值其长度为0，则会执行T0 -->
{{range pipeline}} T1 {{else}} T0 {{end}}
```

```html
<h2>map range</h2>
{{range $key,$value := .mp}}
<p>key: {{$key}}, value: {{$value}}</p>
{{else}}
<p>空数据</p>
{{end}}
```

### 4.7 作用域: with-end

```html
<!-- 如果pipeline为empty不产生输出，否则将 `.` 设为pipeline的值并执行T1。不修改外面的dot -->
{{with pipeline}} T1 {{end}}

<!-- 如果pipeline为empty，不改变dot并执行T0，否则dot设为pipeline的值并执行T1 -->
{{with pipeline}} T1 {{else}} T0 {{end}}
```

```html
{{ with .mp}}
<p>ID : {{.id}}</p>
<p>姓名 : {{.name}}</p>
<p>性别 : {{.gender}}</p>
{{end}}
```

## 5. 模板引擎函数

### 5.1 比较函数

- 布尔函数会将任何类型的零值视为假，其余视为真
- 为了简化多参数相等检测，eq（只有eq）可以接受2个或更多个参数，它会将第一个参数和其余参数依次比较: `{{eq arg1 arg2 arg3}}`
- 比较函数只适用于基本类型（或重定义的基本类型，如”type Celsius float32”）
- 整数和浮点数不能互相比较

| 函数 | 结果描述                 |
| :--- | :----------------------- |
| eq   | 如果arg1 == arg2则返回真 |
| ne   | 如果arg1 != arg2则返回真 |
| lt   | 如果arg1 < arg2则返回真  |
| le   | 如果arg1 <= arg2则返回真 |
| gt   | 如果arg1 > arg2则返回真  |
| ge   | 如果arg1 >= arg2则返回真 |

### 5.2 预定义的全局函数

- and: 函数返回它的第一个empty参数或者最后一个参数；就是说"and x y"等价于"if x then y else x"；所有参数都会执行
- or: 返回第一个非empty参数或者最后一个参数；亦即"or x y"等价于"if x then x else y"；所有参数都会执行
- not: 返回它的单个参数的布尔值的否定
- len: 返回它的参数的整数类型长度, eg: `{{len .fieldName}}`
- index: 执行结果为第一个参数以剩下的参数为索引/键指向的值；如"index x 1 2 3"返回x[1][2][3]的值；每个被索引的主体必须是数组、切片或者字典。 eg: `{{index .mp "gender"}}`
- print: 即fmt.Sprint
- printf: 即fmt.Sprintf
- println: 即fmt.Sprintln
- html: 返回与其参数的文本表示形式等效的转义HTML。这个函数在html/template中不可用。
- urlquery: 以适合嵌入到网址查询中的形式返回其参数的文本表示的转义值。这个函数在html/template中不可用。
- js: 返回与其参数的文本表示形式等效的转义JavaScript。
- call: 执行结果是调用第一个参数的返回值，该参数必须是函数类型，其余参数作为调用该函数的参数；如"`call .X.Y 1 2`"等价于go语言里的 .X.Y(1, 2)；
  - 该函数类型值必须有1到2个返回值，如果有2个则后一个必须是error接口类型；
  - 如果有2个返回值的方法返回的error非nil，模板执行会中断并返回给调用模板执行者该错误；
  - NOTE: 在使用这个函数调用时, `call` 标记不需要, 直接 `.X.Y 1 2` => `.X.Y(1, 2)`

```html
{{/* 函数调用 */}}
<p>Call : {{.user.Say}}</p>
<p>index : {{index .mp "gender"}}</p>
```

### 5.3 自定义函数

自定义函数的注册必须在解析模板之前(ParseXxx 之前),否则模板中不能正确解析到

> go

```go
	// template.New 创建一个模板, 名字需要和模板的名字对应
	tpl := template.New("mutil.tpl")
	// tpl.Funcs 为该模板注册自定义函数
	tpl.Funcs(template.FuncMap{
		"myfunc1": func(name string) string {
			return "myfunc1: hello" + name
		},
		"myfunc2": func(name string) (string, error) {
			return "myfunc2: hello" + name, nil
		},
	})
	// 解析模板
	_, err := tpl.ParseFiles("./tpl/mutil.tpl")
	if err != nil {
		PageError(w, err)
		return
	}
	// 渲染模板
	err = tpl.Execute(w, map[string]any{})
```

> html

```html
<p>myfunc1: {{myfunc1 "world"}}</p>
<p>myfunc2: {{myfunc1 "xxx"}}</p>
```

## 6. 模板嵌套

- 嵌套模板文件
- 嵌套使用 `define` 定义的模板

### 6.1 嵌套模板文件

嵌套模板文件需要在模板解析时将嵌套的模板一并解析,并且内嵌的模板必须在主模板之前

- 嵌套模板 ul.tpl

```html
<ul>
    <li>吃饭</li>
    <li>睡觉</li>
    <li>打豆豆</li>
</ul>
```

- go 模板解析

```go
// template.New 创建一个模板, 名字需要和模板的名字对应
tpl := template.New("template.tpl")
// 解析模板
_, err := tpl.ParseFiles("./tpl/ul.tpl", "./tpl/template.tpl")
// 渲染模板
err = tpl.Execute(w, map[string]any{})
if err != nil {
	PageError(w, err)
	return
}
```

- 主模板文件

```html
<h1>嵌套模板文件</h1>
<hr>
{{ template "ul.tpl" }}
```

### 6.2 嵌套 `define` 定义的模板

- 主模板文件

```html
<h1>嵌套 define 定义的模板</h1>
<hr>
{{ template "ol.tpl" }}

{{ define "ol.tpl"}}
<ol>
    <li>吃饭</li>
    <li>睡觉</li>
    <li>打豆豆</li>
</ol>
{{end}}
```

## 7. 模板继承

- 定义模板语法: `{{define "name"}} T1 {{end}}`
- 执行模板语法: `{{template "name" pipeline}}`
- 模板继承可以用于作布局, 语法: `{{block "name" pipeline}} T1 {{end}}`; block 可看作定义模板`{{define "name"}} T1 {{end}}`和执行`{{template "name" pipeline}}`的缩写

典型的用法是定义一组根模板，然后通过在其中重新定义块模板进行自定义

- 定义基础模板

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .title }}</title>
    <style>
        *{
            margin: 0;
        }
        .nav{
            height: 50px;
            width: 100%;
            position: fixed;
            top: 0;
            background-color: beige;
        }
        .main{
            margin-top: 50px;
        }
        .menu{
            width: 20%;
            height: 100%;
            position: fixed;
            left: 0;
            background-color: antiquewhite;
        }
        .content{
            margin-left: 20%;
            width: 80%;
            height: 100%;
            position: fixed;
            background-color: azure;
        }
    </style>
</head>
<body>
    <div class="nav">nav</div>
    <div class="main">
        <div class="menu">menu</div>
        <div class="content">
            {{block "content" .}} {{end}}
        </div>
    </div>
</body>
</html>
```

- 子模板

```html
{{/* 继承根模板 */}}
{{template "layout.tpl" .}}

{{/* 重新定义内容 */}}
{{define "content"}}
{{.content}}
{{end}}
```

- 指定并渲染模板

```go
// 解析模板
tpl, err := template.ParseFiles("./tpl/layout.tpl", "./tpl/con2.tpl")
// 渲染模板
err = tpl.ExecuteTemplate(w, "con2.tpl", map[string]any{
	"title":   "Content2",
	"content": "这个时内容2",
})
if err != nil {
	PageError(w, err)
	return
}
```

## 8. 修改默认的标识符

go 标准库的模板引擎使用的花括号 `{{` 和 `}}` 作为标识，许多前端框架（如Vue和 AngularJS）也使用{{和}}作为标识符，所以当我们同时使用Go语言模板引擎和以上前端框架时就会出现冲突，这个时候我们需要修改标识符，修改前端的或者修改Go语言的。

```go
// 修改 go 标准库模板引擎的默认花括号
template.New("test").Delims("{[", "]}").ParseFiles("./t.tmpl")
```

## 9. text/template 与 html/tempalte 的区别

1. html/tempalte 针对的是需要返回HTML内容的场景, 渲染过程中会对一些有风险的内容进行转义，以此来防范跨站脚本攻击(js,css,特殊字符等)
2. 可自定义一个 safe 函数对认为安全的数据不进行转义

### 9.1 演示demo

- html xss 测试模板

```html
<p>{{ .script }}</p>
<h2 id="test"></h2>
{{ safe .safeScript}}
<h2 id="test2"></h2>
{{ .safeScript2 | safe}}
```

- go 模型解析和渲染

```go
// 定义一个模板
// template.New 创建一个模板, 名字需要和模板的名字对应
tpl := template.New("xss.tpl")
// tpl.Funcs 为该模板注册自定义函数
tpl.Funcs(template.FuncMap{
    // 自定义函数对认为安全的数据或脚本不进行转义
	"safe": func(s string) template.HTML {
	return template.HTML(s)
	},
})
// 解析模板
tpl, err := tpl.ParseFiles("./tpl/xss.tpl")
// 渲染模板
err = tpl.Execute(w, map[string]any{
	"script":      "<script>alert(123);</script>",
	"safeScript":  "<script>document.getElementById('test').innerHTML='This is a test safe script';</script>",
	"safeScript2": "<script>document.getElementById('test2').innerHTML='This is a test2 safe script';</script>",
})
if err != nil {
	PageError(w, err)
	return
}
```