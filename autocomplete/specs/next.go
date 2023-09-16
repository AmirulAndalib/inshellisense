// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["next"] = model.Subcommand{
		Name:        []string{"next"},
		Description: `Next.js CLI to start, build and export your application`,
		Options: []model.Option{{
			Name:        []string{"-h", "--help"},
			Description: `Output usage information`,
		}, {
			Name:        []string{"-v", "--version"},
			Description: `Output the version number`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"build"},
			Description: `Create an optimized production build of your application`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFolders},
				Name:        "dir",
				Description: `Represent the directory of the Next.js application`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"--profile"},
				Description: `Enable production profiling`,
			}, {
				Name:        []string{"--debug"},
				Description: `Enable more verbose build output`,
			}},
		}, {
			Name:        []string{"dev"},
			Description: `Start the application in development mode`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFolders},
				Name:        "dir",
				Description: `Represent the directory of the Next.js application`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"-p", "--port"},
				Description: `A port number on which to start the application`,
				Args:        []model.Arg{{}},
			}, {
				Name:        []string{"-H", "--hostname"},
				Description: `Hostname on which to start the application`,
				Args:        []model.Arg{{}},
			}},
		}, {
			Name:        []string{"start"},
			Description: `Start the application in production mode`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFolders},
				Name:        "dir",
				Description: `Represent the directory of the Next.js application`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"-p", "--port"},
				Description: `A port number on which to start the application`,
				Args:        []model.Arg{{}},
			}, {
				Name:        []string{"-H", "--hostname"},
				Description: `Hostname on which to start the application`,
				Args:        []model.Arg{{}},
			}},
		}, {
			Name:        []string{"export"},
			Description: `Exports the application for production deployment`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFolders},
				Name:        "dir",
				Description: `Represent the directory of the Next.js application`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"-s"},
				Description: `Do not print any messages to console`,
			}},
		}, {
			Name:        []string{"telemetry"},
			Description: `Allows you to control Next.js' telemetry collection`,
			Args: []model.Arg{{
				Name:        "status",
				Description: `Turn Next.js' telemetry collection on or off`,
				Suggestions: []model.Suggestion{{
					Name:        []string{`enable`},
					Description: `Enable Next.js' telemetry collection`,
				}, {
					Name:        []string{`disable`},
					Description: `Disable Next.js' telemetry collection`,
				}},
			}},
		}},
	}
}