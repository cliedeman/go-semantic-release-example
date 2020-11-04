package main

import (
	"github.com/go-semantic-release/changelog-generator-default/pkg/generator"
	"github.com/go-semantic-release/commit-analyzer-cz/pkg/analyzer"
	"github.com/go-semantic-release/condition-gitlab/pkg/condition"
	"github.com/go-semantic-release/provider-gitlab/pkg/provider"
	"github.com/go-semantic-release/semantic-release/v2/pkg/config"
	"github.com/go-semantic-release/semantic-release/v2/pkg/hooks"
	"github.com/go-semantic-release/semantic-release/v2/pkg/runner"
	"log"
)

func main() {
	conf := make(map[string]string)

	provider := provider.GitLabRepository{}
	// TODO
	err := provider.Init(conf)

	if err != nil {
		log.Fatal(err)
	}

	hooksExecutor := hooks.ChainedHooksExecutor{}
	err = hooksExecutor.Init(conf)

	if err != nil {
		log.Fatal(err)
	}

	analyzer := analyzer.DefaultCommitAnalyzer{}

	err = analyzer.Init(conf)

	if err != nil {
		log.Fatal(err)
	}

	generator := generator.DefaultChangelogGenerator{}

	err = generator.Init(conf)

	if err != nil {
		log.Fatal(err)
	}

	semantic := runner.SemanticRelease{
		CI:                 &condition.GitLab{},
		Prov:               &provider,
		HooksExecutor:      &hooksExecutor,
		CommitAnalyzer:     &analyzer,
		ChangelogGenerator: &generator,
		Updater:            nil,
	}

	semantic.Run(&config.Config{
		// TODO: populate fields somehow. possibly urfave or cobra
		Token:                           "",
		ProviderOpts:                    nil,
		CommitAnalyzerOpts:              nil,
		CIConditionOpts:                 nil,
		ChangelogGeneratorOpts:          nil,
		FilesUpdaterOpts:                nil,
		HooksOpts:                       nil,
		Match:                           "",
		VersionFile:                     false,
		Prerelease:                      false,
		Dry:                             false,
		AllowInitialDevelopmentVersions: false,
		AllowNoChanges:                  false,
		MaintainedVersion:               "",
	})
}
