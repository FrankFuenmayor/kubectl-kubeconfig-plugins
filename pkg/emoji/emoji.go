package emoji

import "strconv"

type Emoji string

func (e Emoji) String() string {
	r, _ := strconv.ParseInt(string(e), 16, 32)
	return string(r)
}
