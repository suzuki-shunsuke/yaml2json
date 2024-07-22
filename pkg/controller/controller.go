package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/suzuki-shunsuke/go-convmap/convmap"
	"gopkg.in/yaml.v3"
)

type Controller struct {
	stdout io.Writer
}

func New(stdout io.Writer) *Controller {
	return &Controller{
		stdout: stdout,
	}
}

type RunParam struct {
	Path   string
	Indent string
}

func (ctrl *Controller) Run(ctx context.Context, param *RunParam) error {
	f, err := os.Open(param.Path)
	if err != nil {
		return fmt.Errorf("open a file: %w", err)
	}
	defer f.Close()
	var data interface{}
	if err := yaml.NewDecoder(f).Decode(&data); err != nil {
		return fmt.Errorf("parse a file as YAML: %w", err)
	}
	a, err := convmap.Convert(data, nil)
	if err != nil {
		return fmt.Errorf("convert map key from interface{} to string: %w", err)
	}
	encoder := json.NewEncoder(ctrl.stdout)
	encoder.SetIndent("", param.Indent)
	if err := encoder.Encode(a); err != nil {
		return fmt.Errorf("encode data as JSON: %w", err)
	}
	return nil
}
