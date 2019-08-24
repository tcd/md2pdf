package render

import "testing"

func TestCleanString(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{
			in:   "alt=&#34;IMAGE ALT TEXT HERE&#34; width=&#34;240&#34; height=&#34;180&#34; border=&#34;10&#34; /&gt;&lt;/a&gt;",
			want: `alt="IMAGE ALT TEXT HERE" width="240" height="180" border="10" /></a>`,
		},
		{
			in:   "Chroma — A general purpose syntax highlighter",
			want: "Chroma -- A general purpose syntax highlighter",
		},
	}

	for _, c := range cases {
		got := cleanString(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
			t.Errorf("cleanString:\n\thave %q\n\twant %q", got, c.want)
		}
	}
}
