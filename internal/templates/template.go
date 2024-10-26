package templates

import (
	"html/template"
	"net/http"
)

// /
// / [Template] struct to hold the parsed HTML template
type Template struct {
	tmpl *template.Template
}

// /
// /
// / [Defines] the data (Files) that we want to pass to the template `
type PageData struct {
	Files []string
}

/// Sample HTML template to display the list of files present in the directory
///
/// [{{range .Files}}] is a Go template directive that iterates over the list of files
///
/// [{{.}}] is a Go template directive that prints the file name

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Snoop</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
        }
        .file-list {
            list-style: none;
            padding: 0;
        }
        .file-item {
            padding: 10px;
            margin: 5px 0;
            background-color: #f5f5f5;
            border-radius: 4px;
        }
        .file-link {
            text-decoration: none;
            color: #2196F3;
        }
        .file-link:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body>
    <h1>Available Files</h1>
    <ul class="file-list">
    {{range .Files}} 
        <li class="file-item">
            <a href="/download/{{.}}" class="file-link">📄 {{.}}</a>
            
        </li>
    {{end}}
    </ul>
</body>

</html>


`

///
///
/// [SpawnTemplate] creates a new instance of the html template with the name `index`

func SpawnTemplate() *Template {
	///template.Must panics if template parsing fails
	///
	/// This is okay during initialization as we want to fail fast if template is invalid
	tmpl := template.Must(template.New("index").Parse(htmlTemplate))
	return &Template{tmpl: tmpl}
}

///
///
/// [RenderFiles] renders the template with the given files data
func (t *Template) RenderFiles(w http.ResponseWriter, files []string) error {
	data := PageData{Files: files}
	return t.tmpl.Execute(w, data)
}
