package resolvergen

import (
	"fmt"
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/packages"

	"github.com/trevor-scheer/gqlgen/codegen"
	"github.com/trevor-scheer/gqlgen/codegen/config"
)

func TestLayoutSingleFile(t *testing.T) {
	_ = syscall.Unlink("testdata/singlefile/out/resolver.go")

	cfg, err := config.LoadConfig("testdata/singlefile/gqlgen.yml")
	require.NoError(t, err)
	p := Plugin{}

	require.NoError(t, cfg.Init())

	data, err := codegen.BuildData(cfg)
	if err != nil {
		panic(err)
	}

	require.NoError(t, p.GenerateCode(data))
	assertNoErrors(t, "github.com/trevor-scheer/gqlgen/plugin/resolvergen/testdata/singlefile/out")
}

func TestLayoutFollowSchema(t *testing.T) {
	testFollowSchemaPersistence(t, "testdata/followschema")

	b, err := os.ReadFile("testdata/followschema/out/schema.resolvers.go")
	require.NoError(t, err)
	source := string(b)

	require.Contains(t, source, "(_ *customresolver.Resolver, err error)")
	require.Contains(t, source, "// Named return values are supported.")
	require.Contains(t, source, "// CustomerResolverType.Name implementation")
	require.Contains(t, source, "// AUserHelperFunction implementation")
}

func TestLayoutFollowSchemaWithCustomFilename(t *testing.T) {
	testFollowSchemaPersistence(t, "testdata/filetemplate")

	b, err := os.ReadFile("testdata/filetemplate/out/schema.custom.go")
	require.NoError(t, err)
	source := string(b)

	require.Contains(t, source, "// CustomerResolverType.Resolver implementation")
	require.Contains(t, source, "// CustomerResolverType.Name implementation")
	require.Contains(t, source, "// AUserHelperFunction implementation")
}

func TestLayoutInvalidModelPath(t *testing.T) {
	cfg, err := config.LoadConfig("testdata/invalid_model_path/gqlgen.yml")
	require.NoError(t, err)

	require.NoError(t, cfg.Init())

	_, err = codegen.BuildData(cfg)
	require.Error(t, err)
}

func TestOmitTemplateComment(t *testing.T) {
	_ = syscall.Unlink("testdata/omit_template_comment/resolver.go")

	cfg, err := config.LoadConfig("testdata/omit_template_comment/gqlgen.yml")
	require.NoError(t, err)
	p := Plugin{}

	require.NoError(t, cfg.Init())

	data, err := codegen.BuildData(cfg)
	if err != nil {
		panic(err)
	}

	require.NoError(t, p.GenerateCode(data))
	assertNoErrors(t, "github.com/trevor-scheer/gqlgen/plugin/resolvergen/testdata/omit_template_comment/out")
}

func TestResolver_Implementation(t *testing.T) {
	_ = syscall.Unlink("testdata/resolver_implementor/resolver.go")

	cfg, err := config.LoadConfig("testdata/resolver_implementor/gqlgen.yml")
	require.NoError(t, err)
	p := Plugin{}

	require.NoError(t, cfg.Init())

	data, err := codegen.BuildData(cfg, &implementorTest{})
	if err != nil {
		panic(err)
	}

	require.NoError(t, p.GenerateCode(data))
	assertNoErrors(t, "github.com/trevor-scheer/gqlgen/plugin/resolvergen/testdata/resolver_implementor/out")
}

func TestCustomResolverTemplate(t *testing.T) {
	_ = syscall.Unlink("testdata/resolvertemplate/out/resolver.go")
	cfg, err := config.LoadConfig("testdata/resolvertemplate/gqlgen.yml")
	require.NoError(t, err)
	p := Plugin{}

	require.NoError(t, cfg.Init())

	data, err := codegen.BuildData(cfg)
	if err != nil {
		panic(err)
	}

	require.NoError(t, p.GenerateCode(data))
}

func testFollowSchemaPersistence(t *testing.T, dir string) {
	_ = syscall.Unlink(dir + "/out/resolver.go")

	cfg, err := config.LoadConfig(dir + "/gqlgen.yml")
	require.NoError(t, err)
	p := Plugin{}

	require.NoError(t, cfg.Init())

	data, err := codegen.BuildData(cfg)
	if err != nil {
		panic(err)
	}

	require.NoError(t, p.GenerateCode(data))
	assertNoErrors(t, "github.com/trevor-scheer/gqlgen/plugin/resolvergen/"+dir+"/out")
}

func assertNoErrors(t *testing.T, pkg string) {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedTypes |
			packages.NeedTypesSizes,
	}, pkg)
	if err != nil {
		panic(err)
	}

	hasErrors := false
	for _, pkg := range pkgs {
		for _, err := range pkg.Errors {
			hasErrors = true
			fmt.Println(err.Pos + ":" + err.Msg)
		}
	}
	if hasErrors {
		t.Fatal("see compilation errors above")
	}
}

type implementorTest struct{}

func (i *implementorTest) Implement(field *codegen.Field) string {
	return "panic(\"implementor implemented me\")"
}
