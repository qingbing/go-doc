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

func PagePrintMsg(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "RequestURI:%s", r.RequestURI)
}

func PageTemplateError(w http.ResponseWriter, r *http.Request) {
	PageError(w, fmt.Errorf("测试一个错误"))
}

func PageTemplateMessage(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	tpl, err := template.ParseFiles("./tpl/message.tpl")
	if err != nil {
		PageError(w, err)
		return
	}
	// 渲染模板
	err = tpl.Execute(w, "传递的单个消息")
	if err != nil {
		PageError(w, err)
		return
	}
}

// 要在模板中渲染的结构体字段必须大写
type User struct {
	ID     int64
	Gender string
	Name   string
}

func (u User) Say() string {
	return fmt.Sprintf("ID:%d, Name:%s, Gender:%s.", u.ID, u.Name, u.Gender)
}

func CallSay() string {
	return "CallSay func"
}

func PageTemplateStruct(w http.ResponseWriter, r *http.Request) {
	// 数据准备
	u := User{
		ID:     88,
		Name:   "bing",
		Gender: "male",
	}
	// 解析模板
	tpl, err := template.ParseFiles("./tpl/struct.tpl")
	if err != nil {
		PageError(w, err)
		return
	}
	// 渲染模板
	err = tpl.Execute(w, u)
	if err != nil {
		PageError(w, err)
		return
	}
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

func PageTemplateInsert(w http.ResponseWriter, r *http.Request) {
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
}

func PageTemplateContent1(w http.ResponseWriter, r *http.Request) {
	// template.New 创建一个模板, 名字需要和模板的名字对应
	tpl := template.New("layout.tpl")
	// 解析模板
	_, err := tpl.ParseFiles("./tpl/layout.tpl", "./tpl/con1.tpl")
	// 渲染模板
	err = tpl.ExecuteTemplate(w, "con1.tpl", map[string]any{
		"title":   "Content1",
		"content": "这个时内容1",
	})
	if err != nil {
		PageError(w, err)
		return
	}
}
func PageTemplateContent2(w http.ResponseWriter, r *http.Request) {
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
}

func PageTemplateNewFlag(w http.ResponseWriter, r *http.Request) {
	// 定义一个模板
	tpl := template.New("new_flag.tpl").Delims("{[", "]}")
	// 解析模板
	tpl, err := tpl.ParseFiles("./tpl/new_flag.tpl")
	// 渲染模板
	err = tpl.Execute(w, map[string]any{
		"title": "This is a page title",
	})
	if err != nil {
		PageError(w, err)
		return
	}
}

func PageTemplateXss(w http.ResponseWriter, r *http.Request) {
	// 定义一个模板
	// template.New 创建一个模板, 名字需要和模板的名字对应
	tpl := template.New("xss.tpl")
	// tpl.Funcs 为该模板注册自定义函数
	tpl.Funcs(template.FuncMap{
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
}

func main() {
	// 注册路由函数 PrintMsg
	http.HandleFunc("/", PagePrintMsg)
	http.HandleFunc("/test_error", PageTemplateError)
	http.HandleFunc("/message", PageTemplateMessage)
	http.HandleFunc("/struct", PageTemplateStruct)
	http.HandleFunc("/mutil", PageTemplateMutil)
	http.HandleFunc("/in", PageTemplateInsert)
	http.HandleFunc("/con1", PageTemplateContent1)
	http.HandleFunc("/con2", PageTemplateContent2)
	http.HandleFunc("/new_flag", PageTemplateNewFlag)
	http.HandleFunc("/xss", PageTemplateXss)

	// 启动服务
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Printf("Http serve failed, err: %+v", err)
	}
}
