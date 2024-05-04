package controller

import (
	"bytes"
	"fmt"
	"html/template"
	"kprobe/executor"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/pkg/errors"
)

type Contoller struct {
	Conf *BpfRuntimeConfig
	Opt  *Options

	templMgr *executor.TemplateMgr
}

func InitContorller(opts ...Option) (*Contoller, error) {
	ctl := &Contoller{}
	ctl.Opt = &Options{
		FuncMap:          make(template.FuncMap),
		ExecuteTemplates: make([]string, 0),
	}
	for _, f := range opts {
		f(ctl.Opt)
	}
	fmt.Printf("\n ###Conf point1### \n")
	fmt.Printf("%+v\n", ctl.Opt)

	if ctl.Opt.ConfPath != "" {
		ctl.Conf = &BpfRuntimeConfig{}
		file, err := os.ReadFile(ctl.Opt.ConfPath)

		fmt.Printf("\n ###File Content### \n")
		// fmt.Println(string(file))

		if err != nil {
			return nil, errors.Wrap(err, "open config failed")
		}
		err = toml.Unmarshal(file, ctl.Conf)

		fmt.Printf("\n ###Conf point2### \n")
		fmt.Println(ctl.Conf)
		fmt.Printf("\n ###Conf point2 Over### \n")

		if err != nil {
			return nil, errors.Wrap(err, "decode config failed")
		}
	}
	ctl.templMgr = &executor.TemplateMgr{}

	fmt.Printf("\n ###Load Map发生### \n")
	for name, _ := range ctl.Opt.FuncMap {
		fmt.Println(name)
	}

	fmt.Printf("\n ###参数展示完毕### \n")

	ctl.templMgr.LoadFuncMap(&ctl.Opt.FuncMap) //load function must be call before load templates

	fmt.Printf("\n ###Load模板发生### \n")
	for _, tmpl := range ctl.Opt.Templates {
		fmt.Println(tmpl)
	}
	fmt.Printf("\n ###参数展示完毕### \n")

	err := ctl.templMgr.LoadTemplates(ctl.Opt.Templates...)
	if err != nil {
		return nil, err
	}
	return ctl, nil
}

func (c *Contoller) Run() {
	fmt.Printf("\n ###Conf point3### \n")
	fmt.Println(c.Conf)
	fmt.Printf("\n ###Conf point3 Over### \n")
	for _, name := range c.templMgr.GetNames() {
		// fmt.Printf("\n ###更换 name### \n")
		// fmt.Printf(name)

		for _, v := range c.Opt.ExecuteTemplates {
			// fmt.Printf("\n ###更换 v### \n")
			// fmt.Printf("name: %s, v: %s\n", name, v)
			// fmt.Printf(v)

			filename := path.Base(name)
			// fmt.Println("filename:", filename) // 输出 filename

			if v == filename {
				fmt.Printf("\n ###执行发生### \n")
				fmt.Println("filename:", filename, "v", v)
				fmt.Printf("name: %s, v: %s\n", name, v)

				if c.Opt.OutputStrategy == file {
					file, err := os.OpenFile(path.Join(c.Opt.OutputDir, strings.TrimSuffix(filename, ".gtpl")), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
					if err != nil {
						log.Fatalf("failed to open file: %v", err)
					}
					defer file.Close()

					fmt.Printf("\n ###生成发生### \n")
					fmt.Println("OutputStrategy:", c.Opt.OutputStrategy)
					fmt.Println("file:", file)
					fmt.Println("filename:", filename)
					fmt.Printf("EBPFProgram: %+v\n", c.Conf.EBPFProgram)
					fmt.Printf("\n ###参数展示完毕### \n")

					err = c.templMgr.Generate(file, filename, c.Conf.EBPFProgram)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}

	// if c.Opt.GenerateCCode {
	// 	for _, name := range c.templMgr.GetNames() {
	// 		if strings.HasSuffix(name, "c.gtpl") {
	// 			filename := path.Base(name)
	// 			if c.Opt.OutputStrategy == file {
	// 				file, err := os.OpenFile(path.Join(c.Opt.OutputDir, strings.TrimSuffix(filename, ".gtpl")), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	// 				if err != nil {
	// 					log.Fatalf("failed to open file: %v", err)
	// 				}
	// 				defer file.Close()
	// 				err = c.templMgr.Generate(file, filename, c.Conf.EBPFProgram)
	// 				if err != nil {
	// 					log.Fatal(err)
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	if c.Opt.CompileCCode {
		for _, name := range c.templMgr.GetNames() {
			if strings.HasSuffix(name, "c.gtpl") {
				filename := path.Base(name)
				cmd := exec.Command("go", "run", "github.com/cilium/ebpf/cmd/bpf2go", "-target", "amd64", "-output-dir", c.Opt.OutputDir, "bpf", path.Join(c.Opt.OutputDir, strings.TrimSuffix(filename, ".gtpl")), "--", fmt.Sprintf("-I%s", c.Opt.CHeaders))
				log.Println(cmd.String())
				cmd.Env = append(os.Environ(), "GOPACKAGE=main")
				var stderr bytes.Buffer
				cmd.Stderr = &stderr

				err := cmd.Run()
				if err != nil {
					log.Fatalf("Command execution failed: %v, stderr: %s", err, stderr.String())
				}
			}
		}

	}

	// if c.Opt.GenerateCtlCode {
	// 	for _, name := range c.templMgr.GetNames() {
	// 		if strings.HasSuffix(name, "go.gtpl") {
	// 			filename := path.Base(name)
	// 			if c.Opt.OutputStrategy == file {
	// 				file, err := os.OpenFile(path.Join(c.Opt.OutputDir, strings.TrimSuffix(filename, ".gtpl")), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	// 				if err != nil {
	// 					log.Fatalf("failed to open file: %v", err)
	// 				}
	// 				defer file.Close()
	// 				err = c.templMgr.Generate(file, filename, c.Conf.EBPFProgram)
	// 				if err != nil {
	// 					log.Fatal(err)
	// 				}
	// 			}
	// 		}
	// 	}
	// }
}
