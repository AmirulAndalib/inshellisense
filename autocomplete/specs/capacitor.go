// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["capacitor"] = model.Subcommand{
		Name:        []string{"capacitor"},
		Description: `The Capacitor command-line interface (CLI) tool is used to develop Capacitor apps`,
		Options: []model.Option{{
			Name:         []string{"--help", "-h"},
			Description:  `Output usage information. Can be used with individual commands too`,
			IsPersistent: true,
		}, {
			Name:        []string{"--version", "-V"},
			Description: `Output the version number`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"add"},
			Description: `Add a native platform project to your app`,
			Args: []model.Arg{{
				Name: "platform",
				Suggestions: []model.Suggestion{{
					Name: []string{`android`},
				}, {
					Name: []string{`ios`},
				}},
			}},
		}, {
			Name:        []string{"copy"},
			Description: `Copy the web app build and Capacitor configuration file into the native platform project. Run this each time you make changes to your web app or change a configuration value`,
			Args: []model.Arg{{
				Name: "platform",
				Suggestions: []model.Suggestion{{
					Name: []string{`android`},
				}, {
					Name: []string{`ios`},
				}},
				IsOptional: true,
			}},
		}, {
			Name:        []string{"ls"},
			Description: `List all installed Cordova and Capacitor plugins`,
			Args: []model.Arg{{
				Name: "platform",
				Suggestions: []model.Suggestion{{
					Name: []string{`android`},
				}, {
					Name: []string{`ios`},
				}},
				IsOptional: true,
			}},
		}, {
			Name:        []string{"open"},
			Description: `Opens the native project workspace in the specified native IDE (Xcode for iOS, Android Studio for Android)`,
			Args: []model.Arg{{
				Name: "platform",
				Suggestions: []model.Suggestion{{
					Name: []string{`android`},
				}, {
					Name: []string{`ios`},
				}},
			}},
		}, {
			Name:        []string{"run"},
			Description: `Opens the native project workspace in the specified native IDE (Xcode for iOS, Android Studio for Android)`,
			Args: []model.Arg{{
				Name: "platform",
				Suggestions: []model.Suggestion{{
					Name: []string{`android`},
				}, {
					Name: []string{`ios`},
				}},
			}},
			Options: []model.Option{{
				Name:        []string{"--list"},
				Description: `Print a list of target devices available to the given platform`,
			}, {
				Name:        []string{"--target"},
				Description: `Run on a specific target device`,
				Args: []model.Arg{{
					Name:      "target",
					Generator: nil, // TODO: port over generator
				}},
			}},
		}, {
			Name:        []string{"sync"},
			Description: `This command runs copy and then update`,
			Args: []model.Arg{{
				Name: "platform",
				Suggestions: []model.Suggestion{{
					Name: []string{`android`},
				}, {
					Name: []string{`ios`},
				}},
				IsOptional: true,
			}},
			Options: []model.Option{{
				Name:        []string{"--deployment"},
				Description: `Podfile.lock won't be deleted and pod install will use --deployment option`,
			}, {
				Name:        []string{"--inline"},
				Description: `After syncing, all JS source maps will be inlined allowing for debugging an Android Web View in Chromium based browsers`,
			}},
		}, {
			Name:        []string{"update"},
			Description: `Updates the native plugins and dependencies referenced in package.json`,
			Args: []model.Arg{{
				Name: "platform",
				Suggestions: []model.Suggestion{{
					Name: []string{`android`},
				}, {
					Name: []string{`ios`},
				}},
				IsOptional: true,
			}},
			Options: []model.Option{{
				Name:        []string{"--deployment"},
				Description: `Podfile.lock won't be deleted and pod install will use --deployment option`,
			}},
		}},
	}
}