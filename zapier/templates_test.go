package zapier

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"testing"
)

func TestTemplatesService_List_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zap-templates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	templates, _, err := client.Templates.List(context.Background(), 0)
	if err != nil {
		t.Errorf("Templates.List returned error: %v", err)
	}

	want := []*Template{{Id: 1}, {Id: 2}}
	if !reflect.DeepEqual(templates, want) {
		t.Errorf("Templates.List returned %+v, want %+v", templates, want)
	}
}

func TestTemplatesService_List_limit(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zap-templates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		q := r.URL.Query()
		limit, err := strconv.Atoi(string(q.Get("limit")[0]))
		if err != nil {
			t.Error(err)
		}
		resp := "["
		for i := 0; i < limit; i += 1 {
			resp += `{"id":` + strconv.Itoa(i) + `}`
			if i < limit-1 {
				resp += ","
			}
		}
		resp += "]"
		fmt.Fprint(w, resp)
	})

	limit := 3
	templates, _, err := client.Templates.List(context.Background(), limit)
	if err != nil {
		t.Errorf("Templates.List returned error: %v", err)
	}

	if len(templates) != limit {
		t.Errorf("Templates.List returned %+v, want %+v", len(templates), limit)
	}
}
