package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type genfunc struct {
	NumIn  int
	NumOut int
}

var tmpl = template.Must(template.New("genfunc.go").Funcs(
	template.FuncMap{
		"Iterate": func(count int) []int {
			rv := make([]int, count)
			for i := range rv {
				rv[i] = i + 1
			}
			return rv
		},
		"Join": func(count int, prefix string) string {
			var str string
			for i := 0; i < count; i++ {
				str += fmt.Sprintf("%s%d, ", prefix, i+1)
			}
			return str
		},
		"TrimSuffix": func(str string, suffix string) string {
			return strings.TrimSuffix(str, suffix)
		},
	},
).Parse(`package task

import "context"
 
{{- range . }}
{{- $in_out := printf "%d_%d" .NumIn .NumOut }}
{{- $type_I := Join .NumIn "I" }}
{{- $type_O := Join .NumOut "O" }}
{{- $type_I_O := TrimSuffix (printf "%s%s" $type_I $type_O ) ", " }}
{{- $t_in := Join .NumIn "t.in" }}
{{- $t_out := Join .NumOut "t.out" }}

type func{{$in_out}}[{{$type_I_O}} any] struct {
	BaseTask
	{{- range $i := Iterate .NumIn }}	
	in{{$i}} I{{$i}}
	{{- end }}
	{{- range $i := Iterate .NumOut }}
	out{{$i}} O{{$i}}
	{{- end }}
}

func Func{{$in_out}}[{{$type_I_O}} any](f func(context.Context, {{$type_I}}) ({{$type_O}} error)) *func{{$in_out}}[{{$type_I_O}}] {
	t := &func{{$in_out}}[{{$type_I_O}}]{}
	t.BaseTask = *Func(func(ctx context.Context) error {
		var err error
		{{$t_out}} err = f(ctx, {{$t_in}})
		return err
	})
	t.InputFunc = BaseInputFunc(func(in ...any) {
		{{ range $i, $v := Iterate .NumIn -}}
		t.in{{$v}} = in[{{$i}}].(I{{$v}})
		{{ end -}}
	})
	return t
}

func (t *func{{$in_out}}[{{$type_I_O}}]) Output() []any {
	return []any{
		{{- $t_out -}}
	}
}

func (t *func{{$in_out}}[{{$type_I_O}}]) TInput(
	{{- range $i := Iterate .NumIn -}}
	in{{$i}} I{{$i}},
	{{- end -}}
) {
	{{ if gt .NumIn 0 -}}
	{{ TrimSuffix $t_in ", " }} = {{ TrimSuffix (Join .NumIn "in") ", " }}
	{{ end -}}
}

func (t *func{{$in_out}}[{{$type_I_O}}]) TOutput() ({{$type_O}}) {
	return {{ TrimSuffix $t_out ", " }}
}
{{- end }}
`))

func main() {
	in := flag.Int("num_in", 5, "number of input")
	out := flag.Int("num_out", 5, "number of output")
	flag.Parse()
	g := []genfunc{}
	for i := 0; i <= *in; i++ {
		for j := 0; j <= *out; j++ {
			if i == 0 && j == 0 {
				continue
			}
			g = append(g, genfunc{NumIn: i, NumOut: j})
		}
	}
	if err := tmpl.Execute(os.Stdout, g); err != nil {
		panic(err)
	}
}
