package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/imijanur/islamteachhelps.com/internal/handlers"
	"github.com/imijanur/islamteachhelps.com/internal/models"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	distDir := "dist"

	if err := os.RemoveAll(distDir); err != nil {
		return fmt.Errorf("remove dist: %w", err)
	}
	if err := os.MkdirAll(distDir, 0o755); err != nil {
		return fmt.Errorf("create dist: %w", err)
	}

	pagesWritten := 0
	for _, r := range handlers.Routes {
		name := strings.TrimPrefix(r.Path, "/")
		if name == "" {
			name = "index"
		}
		outPath := filepath.Join(distDir, name+".html")
		if err := renderToFile(outPath, r.Template, r.Page); err != nil {
			return fmt.Errorf("render %s: %w", r.Path, err)
		}
		pagesWritten++
	}

	if err := renderToFile(filepath.Join(distDir, "404.html"), handlers.NotFound.Template, handlers.NotFound.Page); err != nil {
		return fmt.Errorf("render 404: %w", err)
	}
	pagesWritten++

	filesCopied, err := copyDir("static", filepath.Join(distDir, "static"))
	if err != nil {
		return fmt.Errorf("copy static: %w", err)
	}

	headers := "/static/*\n  Cache-Control: public, max-age=2592000\n"
	if err := os.WriteFile(filepath.Join(distDir, "_headers"), []byte(headers), 0o644); err != nil {
		return fmt.Errorf("write _headers: %w", err)
	}

	fmt.Printf("Exported %d pages, copied %d static files to %s/\n", pagesWritten, filesCopied, distDir)
	return nil
}

func renderToFile(outPath, templateName string, page models.Page) error {
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := handlers.RenderPage(f, templateName, page); err != nil {
		return err
	}
	return f.Close()
}

func copyDir(src, dst string) (int, error) {
	count := 0
	err := filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		if err := copyFile(path, target); err != nil {
			return err
		}
		count++
		return nil
	})
	return count, err
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}
