package render

import "testing"

func TestCleanString(t *testing.T) {
	testString := "alt=&#34;IMAGE ALT TEXT HERE&#34; width=&#34;240&#34; height=&#34;180&#34; border=&#34;10&#34; /&gt;&lt;/a&gt;"
	want := `alt="IMAGE ALT TEXT HERE" width="240" height="180" border="10" /></a>`

	have := cleanString(testString)
	if have != want {
		t.Errorf("cleanString:\n\thave %q\n\twant %q", have, want)
	}
}
