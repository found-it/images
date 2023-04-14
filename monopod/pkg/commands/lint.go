package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/chainguard-images/images/monopod/pkg/commands/options"
	"github.com/chainguard-images/images/monopod/pkg/images"
)

func Lint() *cobra.Command {
	mo := &options.MatrixOptions{}
	cmd := &cobra.Command{
		Use:   "lint",
		Short: "Apply linting rules to images.",
		Example: `
monopod lint
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			impl := &lintImpl{
				ModifiedFiles: mo.ModifiedFiles,
				MelangeMode:   mo.MelangeMode,
				UniqueImages:  mo.UniqueImages,
			}
			return impl.Do()
		},
	}
	mo.AddFlags(cmd)
	return cmd
}

type lintImpl struct {
	ModifiedFiles string
	MelangeMode   string
	UniqueImages  bool
}

type imageTags struct {
	prod []string
	Dev  []string
}

func (i *lintImpl) Do() error {
	allImages, err := images.ListAll()
	if err != nil {
		return err
	}

	imagesMap := map[string][]string{}

	for _, image := range allImages {
        imagesMap[image.ApkoBaseTag]= append(imagesMap[image.ApkoBaseTag], image.ApkoTargetTag)
	}

	for k, v := range imagesMap {
		containsDev := false
        seen := map[string]struct{}
		for _, tag := range v {
			if strings.Contains(tag, "-dev") {

			}
		}
		if !containsDev {
			fmt.Printf("%s:%v\n", k, v)
		}
	}

	return nil
}
