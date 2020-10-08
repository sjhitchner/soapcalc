package django

import (
	"path/filepath"
	// "syscall"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin"
)

const (
	pluginName = "django"
	tmplDir    = "/Users/steve/src/github.com/sjhitchner/soapcalc/backend/plugins/django/*.gotpl"
)

// Maxlength of char fields

// New create new plugin
func New() plugin.Plugin {
	// func New(filename string, typename string) plugin.Plugin {
	//	return &Plugin{filename: filename, typeName: typename}
	return &Plugin{
		filename: "model.py",
	}
}

// Plugin plugin struct
type Plugin struct {
	filename string
	typeName string
}

var _ plugin.CodeGenerator = &Plugin{}
var _ plugin.ConfigMutator = &Plugin{}

// Name plugin name
func (m *Plugin) Name() string {
	return pluginName
}

// MutateConfig mutate config
func (t *Plugin) MutateConfig(cfg *config.Config) error {
	// _ = syscall.Unlink(m.filename)
	fmt.Println("MutateConfig", cfg)
	return nil
}

// GenerateCode generate code
func (t *Plugin) GenerateCode(data *codegen.Data) error {
	/*
		abs, err := filepath.Abs(t.filename)
		if err != nil {
			return err
		}
	*/

	models, err := t.GenerateModels(data)
	if err != nil {
		return err
	}

	funcMap := template.FuncMap{
		"snake": strings.ToLower,
		"djangoField": func(field Field) string {

			if field.Name == "id" {
				return models.PrimaryKey()
			}

			switch field.Type {
			case "string":
				return "models.CharField(max_length=100)"
			case "int":
				return "models.IntegerField()"
			case "float32", "float64":
				return "models.FloatField()"
			default:
				return "models.ForeignKey()"
			}
		},
	}

	tmpl := template.Must(
		template.New("models.gotpl").
			Funcs(funcMap).
			ParseGlob(tmplDir))

	if err != nil {
		return err
	}

	err = tmpl.Execute(os.Stdout, &Resolver{
		Models: models,
	})
	if err != nil {
		return err
	}

	return nil
}

func (t *Plugin) GenerateModels(data *codegen.Data) ([]Model, error) {
	models := make([]Model, 0, 10)

	for _, obj := range data.Objects {

		if obj.Type == nil {
			continue
		}

		kind := filepath.Base(obj.Type.String())
		name := filepath.Ext(obj.Type.String())[1:]

		if !strings.HasPrefix(kind, "model") {
			continue
		}

		model := Model{
			Name: name,
		}

		for _, field := range obj.Fields {

			model.Fields = append(model.Fields, Field{
				Name:     field.GoFieldName,
				Type:     field.TypeReference.GO.String(),
				Required: field.TypeReference.IsNilable(),
			})
		}

		models = append(models, model)
	}

	return models, nil
}

// ResolverBuild build resolver
type Resolver struct {
	*codegen.Data

	Models []Model
}

type Model struct {
	Name       string
	Implements string
	Fields     []Field
}

type Field struct {
	Name     string
	Type     string
	Required bool
	Args     map[string]string
}
