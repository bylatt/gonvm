package utils

import "github.com/mholt/archiver/v3"

// Unarchive extract source into destination
func Unarchive(src, dst string) error {
	return archiver.Unarchive(src, dst)
}
