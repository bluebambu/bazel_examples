package test_filegroup

import (
	"fmt"
	// "encoding/json"
	"path"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/label"
	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/repo"
	"github.com/bazelbuild/bazel-gazelle/resolve"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

const testFilegroupName = "test_filegroup"

type testFilegroupLang struct {
	language.BaseLang

	sawDone bool
}

func NewLanguage() language.Language {
	return &testFilegroupLang{}
}

func (*testFilegroupLang) Name() string { return testFilegroupName }

func (*testFilegroupLang) Loads() []rule.LoadInfo { 
	fmt.Println("++++++++++++++  Loads")
	return nil 
}

func (*testFilegroupLang) Kinds() map[string]rule.KindInfo {
	fmt.Println("++++++++++++++  kinds")
	return kinds
}

var kinds = map[string]rule.KindInfo{
	"filegroup": {
		NonEmptyAttrs:  map[string]bool{"srcs": true, "deps": true},
		MergeableAttrs: map[string]bool{"srcs": true},
	},
	"tconf_binary": {
		NonEmptyAttrs:  map[string]bool{"srcs": true, "deps": true},
		MergeableAttrs: map[string]bool{"srcs": true},
	},
}

func (l *testFilegroupLang) GenerateRules(args language.GenerateArgs) language.GenerateResult {
	fmt.Println("++++++++++++++  GenerateRules: " + args.Rel)
	// j,_ := json.MarshalIndent(args, "", "  ")
	// fmt.Printf("%s", string(j))
	// fmt.Println("++++++++++++++")

	if l.sawDone {
		panic("GenerateRules must not be called after DoneGeneratingRules")
	}

	r := rule.NewRule("filegroup", "all_files")
	srcs := make([]string, 0, len(args.Subdirs)+len(args.RegularFiles))
	srcs = append(srcs, args.RegularFiles...)
	for _, f := range args.Subdirs {
		pkg := path.Join(args.Rel, f)
		srcs = append(srcs, "//"+pkg+":all_files")
	}
	r.SetAttr("srcs", srcs)
	r.SetAttr("testonly", true)
	if args.File == nil || !args.File.HasDefaultVisibility() {
		r.SetAttr("visibility", []string{"//visibility:public"})
	}


	t := rule.NewRule("tconf_binary", "tconf_bin_" + args.Rel)
	srcs = make([]string, 0, len(args.Subdirs)+len(args.RegularFiles))
	srcs = append(srcs, args.RegularFiles...)
	for _, f := range args.Subdirs {
		pkg := path.Join(args.Rel, f)
		srcs = append(srcs, "//"+pkg+":all_files")
	}
	t.SetAttr("srcs", srcs)
	t.SetAttr("custom_xiang", 2324)
	if args.File == nil || !args.File.HasDefaultVisibility() {
		t.SetAttr("visibility", []string{"//visibility:public"})
	}

	return language.GenerateResult{
		Gen:     []*rule.Rule{r},
		Imports: []interface{}{nil},
	}
}

func (l *testFilegroupLang) DoneGeneratingRules() {
	fmt.Println("++++++++++++++  DoneGeneratingRules")
	l.sawDone = true
}

func (l *testFilegroupLang) Resolve(c *config.Config, ix *resolve.RuleIndex, rc *repo.RemoteCache, r *rule.Rule, imports interface{}, from label.Label) {
	// j, _ := json.MarshalIndent(c, "", " ")
	fmt.Println("++++++++++++++  Resolve") 
	if !l.sawDone {
		panic("Expected a call to DoneGeneratingRules before Resolve")
	}
}