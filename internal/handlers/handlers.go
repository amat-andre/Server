package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/amat-andre/Server/internal/service"
)


func MainHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "server not support %s requests. Use %s method", r.Method, http.MethodGet)
        return
    }

	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type",  "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func UploadHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "server not support %s requests. Use %s method", r.Method, http.MethodPost)
        return
    }

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	file, handler, err := r.FormFile("myFile")
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := service.DefinitionAndConversion(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(handler.Filename)
	fileName := fmt.Sprint(time.Now().UTC().String(), ext)

	newFile, err := os.Create(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.WriteString(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(result))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
