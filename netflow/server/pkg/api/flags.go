package api

import (
	"flag"
	"fmt"

	"golang.org/x/exp/slices"
)

type ApiFlags struct {
	Config string
	Type   string
}

func ParseFlags() (*ApiFlags, error) {
	a := &ApiFlags{}
	flag.StringVar(&a.Config, "c", "/app/configs/conf.yaml", "Path to config file. Example: /abs/path/to/config")
	flag.StringVar(&a.Type, "type", "", "Defalut: dev. Defines which of type api will started (prod or dev)")
	flag.Parse()

	if !slices.Contains([]string{"prod", "dev"}, a.Type) {
		return nil, fmt.Errorf("type flag value must be prod/dev")
	}

	return a, nil

}
