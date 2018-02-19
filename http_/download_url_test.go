// see also examples:
// https://gist.github.com/ijt/950790/fca88967337b9371bb6f7155f3304b3ccbf3946f
// https://gist.github.com/jmoiron/e9f72720cef51862b967

package http_

import (
	"testing"
	"net/http"
	"io/ioutil"
)

func TestVerifyDownload(t *testing.T) {
	if response, err := http.Get("http://www.google.com"); err != nil {
		t.Errorf("%s", err)
	} else {
		defer response.Body.Close()
		if contents, err := ioutil.ReadAll(response.Body); err != nil {
			t.Errorf("%s", err)
		} else {
			t.Logf("%s", string(contents))
		}
	}
}
