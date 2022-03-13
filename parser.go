package esbuild

import (
	"errors"

	"github.com/evanw/esbuild/internal/config"
	"github.com/evanw/esbuild/internal/js_ast"
	"github.com/evanw/esbuild/internal/js_parser"
	"github.com/evanw/esbuild/internal/logger"
)

func ParseTsx(contents string) (js_ast.AST, error) {
	log := logger.NewDeferLog(logger.DeferLogNoVerboseOrDebug)

	source := logger.Source{
		Index:          0,
		KeyPath:        logger.Path{Text: "<stdin>"},
		PrettyPath:     "<stdin>",
		Contents:       contents,
		IdentifierName: "stdin",
	}

	tree, ok := js_parser.Parse(log, source, js_parser.OptionsFromConfig(&config.Options{
		TS: config.TSOptions{
			Parse: true,
		},
		JSX: config.JSXOptions{
			Parse: true,
		},
	}))

	if !ok {
		return tree, errors.New("failed to parse tsx")
	}

	return tree, nil
}
