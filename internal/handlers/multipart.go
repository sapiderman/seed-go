package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"gopkg.in/yaml.v2"
)

func (h *Handlers) RequestDump(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(requestDump))
	var name string
	var password string
	ct := r.Header.Get("content-type")
	switch {
	case strings.HasPrefix(ct, "text/plain"):
		fmt.Println(r.Body)
	case strings.HasPrefix(ct, "multipart/form-data"):
		name = r.PostFormValue("name")
		password = r.PostFormValue("password")
	case ct == "application/x-www-form-urlencoded":
		name = r.FormValue("name")
		password = r.FormValue("password")
	case strings.HasPrefix(ct, "application/json"):
		type body struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		}
		var b body
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&b)
		name = b.Name
		password = b.Password
	case strings.HasPrefix(ct, "text/yaml"):
		type body struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		}
		var b body
		decoder := yaml.NewDecoder(r.Body)
		decoder.Decode(&b)
		name = b.Name
		password = b.Password
	case strings.HasPrefix(ct, "application/xml"):
		type body struct {
			Name     string `xml:"name"`
			Password string `xml:"password"`
		}
		var b body
		decoder := xml.NewDecoder(r.Body)
		decoder.Decode(&b)
		name = b.Name
		password = b.Password
	}
	fmt.Printf("%s\n%s\n%s\n", ct, name, password)
	w.WriteHeader(200)
	w.Write(requestDump)

}
