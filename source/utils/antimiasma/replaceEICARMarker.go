package antimiasma

import "strings"

func replaceEICARMarker(contents []byte) []byte {

	return []byte(strings.ReplaceAll(string(contents), "{EICAR_MARKER}", decoded_eicar))

}
