package handlers

import (
	"github.com/bugagych84/go1fl-sprint6-final-tpl/internal/service"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "parse form: "+err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "get file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "read file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	result, wasMorse, err := service.DetectAndConvert(string(data))
	if err != nil {
		http.Error(w, "convert: "+err.Error(), http.StatusInternalServerError)
		return
	}

	timestamp := time.Now().UTC().Format("20060102T150405Z0700")
	ext := filepath.Ext(header.Filename)
	prefix := "from-text"
	if wasMorse {
		prefix = "from-morse"
	}
	outName := prefix + "_" + timestamp + ext

	dst, err := os.Create(outName)
	if err != nil {
		http.Error(w, "create result file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := dst.WriteString(result); err != nil {
		http.Error(w, "write result file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(result))
}
