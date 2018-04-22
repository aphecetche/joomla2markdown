package wsub

import (
	"testing"
)

func TestRemoveStyle(t *testing.T) {
	s := `<h1>titre1</h1>
	<p class="retain" style="background:yellow">para1</p>
	<p>para2</p>
	<img src="toto.png">`

	expected := `<h1>titre1</h1>
	<p class="retain">para1</p>
	<p>para2</p>
	<img src="toto.png"/>`

	r := removeStyle(s)
	if r != expected {
		t.Errorf("Expected:\n>%s<\nbut got:\n>%s<\n", expected, r)
	}
}
