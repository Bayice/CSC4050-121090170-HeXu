package executor

import (
	"fmt"
	"io"
	"text/template"

	"github.com/pkg/errors"
)

type TemplateMgr struct {
	functions *template.FuncMap
	tmpl      *template.Template
	names     []string
}

func (t *TemplateMgr) LoadTemplates(tmpls ...string) error {
	t.names = append(t.names, tmpls...)

	var err error
	t.tmpl = template.New("").Funcs(*t.functions)
	for _, v := range tmpls {
		fmt.Println("###Load内部函数展示v###")
		fmt.Println(v)
		t.tmpl, err = t.tmpl.ParseGlob(v)
		if err != nil {
			return errors.Wrap(err, "parse template dir failed")
		}
	}
	return nil
}

func (t *TemplateMgr) LoadFunc(name string, f any) {
	(*t.functions)[name] = f
}

func (t *TemplateMgr) LoadFuncMap(m *template.FuncMap) {
	t.functions = m
}

func (t *TemplateMgr) Generate(wr io.Writer, name string, data any) error {
	fmt.Printf("\n####开始展示ExecuteTemplate参数###\n")
	// 打印执行的模板名称
	fmt.Println("Executing template:", name)
	fmt.Println("Data to be executed with template:", data)
	fmt.Printf("###参数展示完毕###\n")

	return t.tmpl.ExecuteTemplate(wr, name, data)
}

func (t *TemplateMgr) GetNames() []string {
	names := []string{}
	for _, v := range t.tmpl.Templates() {
		names = append(names, v.Name())
	}

	fmt.Printf("\n ###GetNames### \n")
	fmt.Printf("%v\n", names)
	fmt.Printf("\n ###GetNames Over### \n")
	return names
}
