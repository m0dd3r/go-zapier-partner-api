package tests

import "testing"

func TestListTemplates(t *testing.T) {
	templates, _, err := client.Templates.List(ctx, 0)
	if err != nil {
		t.Error(err)
	}

	if len(templates) != 10 {
		t.Error("expected 10 templates but got ", len(templates))
	}

}
