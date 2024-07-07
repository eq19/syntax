package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/ubuntu/wsl/wsl-builder/common"
	"golang.org/x/exp/slices"
)

type columns struct {
	rootfsArch                                 *string
	appId, fullName, launcher, codeName, short *bool
}

func writeReleaseInfo(csvPath string, distros []string, cols columns) error {
	releases, err := selectReleases(csvPath, distros)
	if err != nil {
		return err
	}

	return writeTable(releases, cols)
}

// writeTable writes the information in a table into stdout.
func writeTable(releases []common.WslReleaseInfo, cols columns) error {
	w := tabwriter.NewWriter(os.Stdout, 1, 2, 2, ' ', 0)
	defer w.Flush()

	if err := writeHeader(w, cols); err != nil {
		return err
	}

	for _, r := range releases {
		if err := writeRow(w, r, cols); err != nil {
			return err
		}
	}

	return nil
}

func writeHeader(w io.Writer, f columns) error {
	if *f.short {
		return nil
	}

	text := []string{}
	text = append(text, "WslID")
	if *f.appId {
		text = append(text, "AppID")
	}
	if *f.fullName {
		text = append(text, "FullName")
	}
	if *f.launcher {
		text = append(text, "Launcher")
	}
	if *f.codeName {
		text = append(text, "CodeName")
	}
	if len(*f.rootfsArch) != 0 {
		text = append(text, "Rootfs")
	}

	if _, err := fmt.Fprintln(w, strings.Join(text, "\t")); err != nil {
		return fmt.Errorf("got an error writing header: %v", err)
	}
	return nil
}

func writeRow(w io.Writer, r common.WslReleaseInfo, f columns) error {
	var text []string
	text = append(text, r.WslID)
	if *f.appId {
		text = append(text, r.AppID)
	}
	if *f.fullName {
		text = append(text, r.FullName)
	}
	if *f.codeName {
		text = append(text, r.CodeName)
	}
	if *f.launcher {
		text = append(text, r.LauncherName+".exe")
	}
	if len(*f.rootfsArch) != 0 {
		text = append(text, r.RootfsURL(*f.rootfsArch))
	}

	if _, err := fmt.Fprintln(w, strings.Join(text, "\t")); err != nil {
		return fmt.Errorf("got an error writing row for %q: %v", r.WslID, err)
	}
	return nil
}

func selectReleases(csvPath string, targets []string) (releases []common.WslReleaseInfo, err error) {
	releases, err = common.ReleasesInfo(csvPath)
	if err != nil {
		return releases, fmt.Errorf("could not read release info: %v", err)
	}

	end := len(releases)
	if len(targets) != 0 {
		end = partition(releases, func(r common.WslReleaseInfo) bool {
			idx := slices.Index(targets, r.WslID)
			if idx == -1 {
				return false
			}
			// Found targets are blanked so we're left with a
			// slice of missing targets.
			targets[idx] = ""
			return true
		})
	}

	if end == 0 {
		return releases, errors.New("no releases were found")
	}

	if end < len(targets) {
		end := partition(targets, func(t string) bool { return len(t) != 0 })
		return nil, fmt.Errorf("the following releases were not found: %s", strings.Join(targets[:end], ", "))
	}
	return releases[:end], nil
}

// partition returns a partition index "p" and rearranges the slice such that
// all elements that fulfil the predicate have index strictly less than "p".
// Stable only for items that fulfil the predicate.
func partition[T any](slice []T, predicate func(t T) bool) (p int) {
	if len(slice) == 0 {
		return 0
	}
	for i := range slice {
		if !predicate(slice[i]) {
			continue
		}
		if i == p {
			p++
			continue
		}
		slice[i], slice[p] = slice[p], slice[i]
		p++
	}
	return p
}
