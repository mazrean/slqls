package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mazrean/slqls/ast"
	"github.com/mazrean/slqls/ast/astutil"
	"github.com/mazrean/slqls/internal/database"
	"github.com/mazrean/slqls/internal/lsp"
	"github.com/mazrean/slqls/parser"
	"github.com/mazrean/slqls/parser/parseutil"
	"github.com/mazrean/slqls/token"
	"github.com/sourcegraph/jsonrpc2"
)

func (s *Server) handleDefinition(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (result interface{}, err error) {
	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	var params lsp.DefinitionParams
	if err := json.Unmarshal(*req.Params, &params); err != nil {
		return nil, err
	}

	f, ok := s.files[params.TextDocument.URI]
	if !ok {
		return nil, fmt.Errorf("document not found: %s", params.TextDocument.URI)
	}

	return definition(params.TextDocument.URI, f.Text, params, s.worker.Cache())
}

func definition(url, text string, params lsp.DefinitionParams, dbCache *database.DBCache) (lsp.Definition, error) {
	pos := token.Pos{
		Line: params.Position.Line,
		Col:  params.Position.Character + 1,
	}
	parsed, err := parser.Parse(text)
	if err != nil {
		return nil, err
	}

	nodeWalker := parseutil.NewNodeWalker(parsed, pos)
	m := astutil.NodeMatcher{
		NodeTypes: []ast.NodeType{ast.TypeIdentifer},
	}
	currentVariable := nodeWalker.CurNodeButtomMatched(m)
	if currentVariable == nil {
		return nil, nil
	}

	aliases := parseutil.ExtractAliased(parsed)
	if len(aliases) == 0 {
		return nil, nil
	}

	var define ast.Node
	for _, v := range aliases {
		alias, _ := v.(*ast.Aliased)
		if alias.AliasedName.String() == currentVariable.String() {
			define = alias.AliasedName
			break
		}
	}

	if define == nil {
		return nil, nil
	}

	res := []lsp.Location{
		{
			URI: url,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      define.Pos().Line,
					Character: define.Pos().Col,
				},
				End: lsp.Position{
					Line:      define.End().Line,
					Character: define.End().Col,
				},
			},
		},
	}

	return res, nil
}
