package arg

import (
	"testing"
)

func TestWord(t *testing.T) {
	expected := `\033[1;37mWord\033[0m`
	result := wrapText("Word", C_WHITE)

	if result != expected {
		t.Errorf("Failed to wrap word %s - %s", expected, result)
	}
}

func TestText(t *testing.T) {
	expected := `\033[1;37mSome more complex
text\033[0m`
	result := wrapText("Some more complex\ntext", C_WHITE)

	if result != expected {
		t.Errorf("Failed to wrap text %s - %s", expected, result)
	}
}
