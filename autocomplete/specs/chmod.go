// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["chmod"] = model.Subcommand{
		Name:        []string{"chmod"},
		Description: `Change file modes or Access Control Lists`,
		Args: []model.Arg{{
			Name: "mode",
			Suggestions: []model.Suggestion{{
				Name:        []string{`u+x`},
				Description: `Give execute permission for the user`,
			}, {
				Name:        []string{`a+rx`},
				Description: `Adds read and execute permissions for all classes`,
			}, {
				Name:        []string{`744`},
				Description: `Sets read, write, and execute permissions for user, and sets read permission for Group and Others`,
			}, {
				Name:        []string{`664`},
				Description: `Sets read and write permissions for user and Group, and provides read to Others`,
			}, {
				Name:        []string{`777`},
				Description: `⚠️ allows all actions for all users`,
			}},
		}, {
			Templates: []model.Template{model.TemplateFilepaths},
		}},
		Options: []model.Option{{
			Name:        []string{"-f"},
			Description: `Do not display a diagnostic message if chmod could not modify the mode for file, nor modify the exit status to reflect such failures`,
		}, {
			Name:        []string{"-H"},
			Description: `If the -R option is specified, symbolic links on the command line are followed and hence unaffected by the command.  (Symbolic links encountered during tree traversal are not followed.)`,
		}, {
			Name:        []string{"-h"},
			Description: `If the file is a symbolic link, change the mode of the link itself rather than the file that the link points to`,
		}, {
			Name:        []string{"-L"},
			Description: `If the -R option is specified, all symbolic links are followed`,
		}, {
			Name:        []string{"-P"},
			Description: `If the -R option is specified, no symbolic links are followed. This is the default`,
		}, {
			Name:        []string{"-R"},
			Description: `Change the modes of the file hierarchies rooted in the files, instead of just the files themselves. Beware of unintentionally matching the ""..'' hard link to the parent directory when using wildcards like "".*''`,
		}, {
			Name:        []string{"-v"},
			Description: `Cause chmod to be verbose, showing filenames as the mode is modified. If the -v flag is specified more than once, the old and new modes of the file will also be printed, in both octal and symbolic notation`,
		}},
	}
}