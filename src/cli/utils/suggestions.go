package utils

// The code in this file is mostly lifted out of cobra itself. It as
// of now has  some issue with how it handles completions for
// subcommands and only autocompletes for top level commands.
// https://github.com/spf13/cobra/issues/981#issuecomment-547003669

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// SubcommandsRequiredWithSuggestions will ensure we have a subcommand provided by the user and augments it with
// suggestion for commands, alias and help on root command.
func SubcommandsRequiredWithSuggestions(cmd *cobra.Command, args []string) error {
	requireMsg := "%s requires a valid subcommand"
	// This will be triggered if cobra didn't find any subcommands.
	// Find some suggestions.
	var suggestions []string

	if len(args) != 0 && !cmd.DisableSuggestions {
		typedName := args[0]

		if cmd.SuggestionsMinimumDistance <= 0 {
			cmd.SuggestionsMinimumDistance = 2
		}
		// subcommand suggestions
		suggestions = append(cmd.SuggestionsFor(args[0]))

		// subcommand alias suggestions (with distance, not exact)
		for _, c := range cmd.Commands() {
			if c.IsAvailableCommand() {
				for _, alias := range c.Aliases {
					levenshteinDistance := levenshteinDistance(typedName, alias, true)
					suggestByLevenshtein := levenshteinDistance <= cmd.SuggestionsMinimumDistance
					suggestByPrefix := strings.HasPrefix(strings.ToLower(alias), strings.ToLower(typedName))

					if suggestByLevenshtein || suggestByPrefix {
						suggestions = append(suggestions, alias)
					}
				}
			}
		}

		// help for root command
		if !cmd.HasParent() {
			help := "help"
			levenshteinDistance := levenshteinDistance(typedName, help, true)
			suggestByLevenshtein := levenshteinDistance <= cmd.SuggestionsMinimumDistance
			suggestByPrefix := strings.HasPrefix(strings.ToLower(help), strings.ToLower(typedName))

			if suggestByLevenshtein || suggestByPrefix {
				suggestions = append(suggestions, help)
			}
		}
	}

	var suggestionsMsg string
	if len(suggestions) > 0 {
		suggestionsMsg += "\n\nDid you mean this?\n"
		for _, s := range suggestions {
			suggestionsMsg += fmt.Sprintf("\t%v\n", s)
		}
	}

	if len(suggestionsMsg) > 0 {
		requireMsg = fmt.Sprintf("%s. %s", requireMsg, suggestionsMsg)
	}

	return fmt.Errorf(requireMsg, cmd.Name())
}

// levenshteinDistance compares two strings and returns the levenshtein distance between them.
func levenshteinDistance(s, t string, ignoreCase bool) int {
	if ignoreCase {
		s = strings.ToLower(s)
		t = strings.ToLower(t)
	}

	d := make([][]int, len(s)+1)
	for i := range d {
		d[i] = make([]int, len(t)+1)
	}

	for i := range d {
		d[i][0] = i
	}

	for j := range d[0] {
		d[0][j] = j
	}

	for j := 1; j <= len(t); j++ {
		for i := 1; i <= len(s); i++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				min := d[i-1][j]
				if d[i][j-1] < min {
					min = d[i][j-1]
				}

				if d[i-1][j-1] < min {
					min = d[i-1][j-1]
				}

				d[i][j] = min + 1
			}
		}
	}

	return d[len(s)][len(t)]
}
