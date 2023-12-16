package tests

import (
	"github.com/stretchr/testify/assert"
	c "lmth/convert"
	"testing"
)

func TestFindClosingTagIndex(t *testing.T) {
	assert.Equal(t, c.FindClosingTagIndex("<lmth>", "html"), -1)
	assert.Equal(t, c.FindClosingTagIndex("<html><lmth>", "html"), 6)
	assert.Equal(t, c.FindClosingTagIndex("<html><html><lmth><lmth>", "html"), 18)
	assert.Equal(t, c.FindClosingTagIndex("<html><lmth><html><lmth>", "html"), 6)
	assert.Equal(t, c.FindClosingTagIndex("<html><html><html><lmth><lmth><lmth>", "html"), 30)
	assert.Equal(t, c.FindClosingTagIndex("<p><q>", "p"), 3)
	assert.Equal(t, c.FindClosingTagIndex("<p><div><vid><q>", "p"), 13)
	assert.Equal(t, c.FindClosingTagIndex("<p><div><q>", "p"), 8)
}

var testCases = []string{
	"<html></html>",
	"<html><div></div></html>",
	"<html>aaa<div>bbb</div>ccc</html>",
	"<p>hello</p>",
	`<span class="abc">hello</span>`,
	`<span class="abc"><div class="def">hello</div></span>`,
	`<span class="<div>">hello</span>`,
	`<a class="main-heading" id="title">Contact Us</a>`,
	`<video autoplay muted src="https://google.com" />`,
	`<script defer>console.log(abc)<script>`,
	`<script src="<div>.com">"</script>"</script>`,
}

func TestEncodeDecode(t *testing.T) {
	for _, s := range testCases {
		e, err := c.EncodeHTML(s)
		assert.Nil(t, err)

		d, err := c.DecodeLMHT(e)
		assert.Nil(t, err)

		assert.Equal(t, s, d)
	}
}
