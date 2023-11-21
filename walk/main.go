package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	ext     string
	size    int64
	list    bool
	del     bool
	wLog    io.Writer
	archive string
}

func run(root string, out io.Writer, cfg config) error {
	delLogger := log.New(cfg.wLog, "DELETED FILE", log.LstdFlags)

	return filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filterOut(path, cfg.ext, cfg.size, info) {
			return nil
		}

		if cfg.list {
			return listFile(path, out)
		}

		if cfg.archive != "" {
			if err := archiveFile(cfg.archive, root, path); err != nil {
				return err
			}
		}

		if cfg.del {
			return delFile(path, delLogger)
		}

		return listFile(path, out)
	})
}

func main() {
	var (
		f   = os.Stdout
		err error
	)

	list := flag.Bool("list", false, "List the files")
	del := flag.Bool("del", false, "Delete the filtered files")
	archive := flag.String("archive", "", "Archive directory")
	root := flag.String("root", "", "Root directory to start")
	ext := flag.String("ext", "", "Extension name of files to filter")
	size := flag.Int64("size", 0, "Minimum file size filter")
	logFile := flag.String("log", "", "Log deletes to this file")
	flag.Parse()

	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}

	c := config{ext: *ext, list: *list, size: *size, del: *del, wLog: f, archive: *archive}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

}
