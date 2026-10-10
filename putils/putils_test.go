package putils_test

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func TestMain(m *testing.M) {
	pterm.DisableOutput()
	m.Run()
}

func equal(t *testing.T, got, want any) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got  %#v\nwant %#v", got, want)
	}
}

type user struct {
	Name    string
	Email   string `pterm:"E-Mail"`
	Age     *int
	Secret  string `pterm:"-"`
	private string
}

func TestTableFromStructs(t *testing.T) {
	age := 36
	got := putils.TableFromStructs([]user{
		{Name: "Ada", Email: "ada@example.com", Age: &age, private: "x"},
		{Name: "Grace"},
	})

	equal(t, got, pterm.TableData{
		{"Name", "E-Mail", "Age"},
		{"Ada", "ada@example.com", "36"},
		{"Grace", "", ""},
	})

	equal(t, putils.TableFromStructs([]*user{nil})[1], []string{"", "", ""})
	equal(t, putils.TableFromStructs([]int{1}), pterm.TableData(nil))
}

func TestTableFromCSV(t *testing.T) {
	got, err := putils.TableFromCSV("name,size\nnpm,\"1,2 GB\"\ngo")
	if err != nil {
		t.Fatal(err)
	}

	equal(t, got, pterm.TableData{{"name", "size"}, {"npm", "1,2 GB"}, {"go"}})

	tsv, err := putils.TableFromSeparated("a\tb", '\t')
	if err != nil {
		t.Fatal(err)
	}

	equal(t, tsv, pterm.TableData{{"a", "b"}})
}

func TestKeyValue(t *testing.T) {
	equal(t, putils.KeyValueFromStruct(&user{Name: "Ada"}), []pterm.KeyValuePair{
		{Key: "Name", Value: "Ada"}, {Key: "E-Mail"}, {Key: "Age"},
	})
	equal(t, putils.KeyValueFromStruct(42), []pterm.KeyValuePair(nil))
	equal(t, putils.KeyValueFromMap(map[string]int{"b": 2, "a": 1}), []pterm.KeyValuePair{
		{Key: "a", Value: "1"}, {Key: "b", Value: "2"},
	})
}

func TestTreeFromMap(t *testing.T) {
	got := putils.TreeFromMap("config", map[string]any{
		"port":  8080,
		"db":    map[string]any{"host": "localhost", "pool": 10},
		"hosts": []string{"a", "b"},
		"nil":   nil,
	})

	equal(t, got, pterm.TreeNode{Text: "config", Children: []pterm.TreeNode{
		{Text: "db", Children: []pterm.TreeNode{
			{Text: "host", Description: "localhost"},
			{Text: "pool", Description: "10"},
		}},
		{Text: "hosts", Children: []pterm.TreeNode{{Text: "a"}, {Text: "b"}}},
		{Text: "nil"},
		{Text: "port", Description: "8080"},
	}})
}

func TestTreeFromIndented(t *testing.T) {
	got := putils.TreeFromIndented(`
src
  main.go
  internal
    db.go

	tabbed
README.md`)

	equal(t, got, pterm.TreeNode{Children: []pterm.TreeNode{
		{Text: "src", Children: []pterm.TreeNode{
			{Text: "main.go"},
			{Text: "internal", Children: []pterm.TreeNode{{Text: "db.go"}}},
			{Text: "tabbed"},
		}},
		{Text: "README.md"},
	}})
}

func TestBullets(t *testing.T) {
	equal(t, putils.BulletsFromSlice([]int{1, 2}), []pterm.BulletListItem{{Text: "1"}, {Text: "2"}})
	equal(t, putils.BulletsFromIndented("a\n  b\n    c\nd"), []pterm.BulletListItem{
		{Text: "a"}, {Level: 1, Text: "b"}, {Level: 2, Text: "c"}, {Text: "d"},
	})
}

func TestRunWithSpinner(t *testing.T) {
	ran := false

	err := putils.RunWithSpinner(t.Context(), "work", func(context.Context, *pterm.SpinnerPrinter) error {
		ran = true
		return nil
	})
	if err != nil || !ran {
		t.Fatalf("err %v, ran %v", err, ran)
	}

	want := errors.New("boom")

	err = putils.RunWithSpinner(t.Context(), "work", func(context.Context, *pterm.SpinnerPrinter) error { return want })
	if !errors.Is(err, want) {
		t.Fatalf("got %v, want %v", err, want)
	}
}

func TestDownloadWithProgress(t *testing.T) {
	body := bytes.Repeat([]byte("pterm"), 10_000)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/missing" {
			http.NotFound(w, r)
			return
		}

		_, _ = w.Write(body)
	}))
	defer srv.Close()

	var buf bytes.Buffer
	if err := putils.DownloadWithProgress(t.Context(), "file", srv.URL, &buf); err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(buf.Bytes(), body) {
		t.Fatalf("downloaded %d bytes, want %d", buf.Len(), len(body))
	}

	if err := putils.DownloadWithProgress(t.Context(), "file", srv.URL+"/missing", &buf); err == nil {
		t.Fatal("want an error for a 404")
	}
}
