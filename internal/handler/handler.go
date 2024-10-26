package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"deepraj02/snoop/internal/templates"
)

///
///
/// [FileHandler] `struct` stores the directory path where the CLI is started.
/// It also stores a reference to the template object to render the files data into HTML.

type FileHandler struct {
	dir      string
	template *templates.Template
}

///
///
/// [New] function creates a new FileHandler object with the given directory path and initializes them.

func New(dir string) *FileHandler {
	return &FileHandler{
		dir:      dir,
		template: templates.SpawnTemplate(),
	}
}

/// [HandleIndex] wrapper around [getFiles]
///
///Handles HTTP GET requests to the root path (/).
///
/// Retrieves a list of files from the shared directory using the getFiles function.
///
///Renders an HTML template (index.html) to display the list of files.

func (h *FileHandler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files, err := h.getFiles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.template.RenderFiles(w, files)
}

/// [HandleDownload] Sets the Content-Disposition header to prompt the browser to download the file.
///
/// Serves the file using http.ServeFile.

func (h *FileHandler) HandleDownload(w http.ResponseWriter, r *http.Request) {
	filename := filepath.Base(r.URL.Path)
	filePath := filepath.Join(h.dir, filename)

	absDir, err := filepath.Abs(h.dir)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if !strings.HasPrefix(absPath, absDir) {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	http.ServeFile(w, r, filePath)
}

/// [getFiles] reads the directory and returns a list of files.
///
///

func (h *FileHandler) getFiles() ([]string, error) {
	var files []string
	entries, err := os.ReadDir(h.dir)
	if err != nil {
		return nil, err
	}

	/// It ignores directories and returns only file names.
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}
