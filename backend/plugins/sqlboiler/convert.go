// +build ignore

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	gbgen "github.com/web-ridge/gqlgen-sqlboiler/v2"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	root, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to get current directory", err.Error())
		os.Exit(-1)
	}

	// supports root or sub directories
	output := gbgen.Config{
		Directory:   filepath.Join(root, "generated/helpers"),
		PackageName: "helpers",
	}
	backend := gbgen.Config{
		Directory:   filepath.Join(root, "generated/db/models"),
		PackageName: "models",
	}
	frontend := gbgen.Config{
		Directory:   filepath.Join(root, "generated/graphql/models"),
		PackageName: "graphql",
	}

	err = api.Generate(cfg,
		api.AddPlugin(gbgen.NewConvertPlugin(
			output,   // directory where convert.go, convert_input.go and preload.go should live
			backend,  // directory where sqlboiler files are put
			frontend, // directory where gqlgen models live
		)),
		api.AddPlugin(gbgen.NewResolverPlugin(
			output,
			backend,
			frontend,
			"github.com/web-ridge/yourapp/yourauth", // leave empty if you don't have auth
		)),
	)
	if err != nil {
		fmt.Println("error!!")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
