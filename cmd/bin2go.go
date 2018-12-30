package main

import (
    "bin2go"

    "bytes"
    "flag"
    "fmt"
    "io"
    "os"
    "strings"
)

var opts struct {
    outfile string
    pkg     string
}

func usage() {
    fmt.Fprintf(os.Stderr,
        "Usage: %s [-h] [-p package] [-o outfile] FILENAME[:VARNAME] [FILENAME2[:VARNAME2]...]\n",
        os.Args[0])
}

func realMain() int {
    flag.Usage = usage
    flag.StringVar(&opts.pkg, "p", "main", "package name (default main)")
    flag.StringVar(&opts.outfile, "o", "", "Output file (default stdout)")
    flag.Parse()

    if len(flag.Args()) < 1 {
        fmt.Fprintln(os.Stderr, "No filename specified")
        usage()
        return 2
    }

    gen, err := bin2go.New(opts.pkg)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to initialize generator: %v\n", err)
        return 1
    }

    for _, f := range flag.Args() {
        sp := strings.SplitN(f, ":", 2)
        if len(sp) == 1 {
            if err := gen.AddFile(f); err != nil {
                fmt.Fprintf(os.Stderr, "Failed to add file %q: %v\n", f, err)
                return 1
            }
        } else {
            if err := gen.AddFileVar(sp[0], sp[1]); err != nil {
                fmt.Fprintf(os.Stderr, "Failed to add file/var %q: %v\n", f, err)
                return 1
            }
        }
    }

    success := false
    useStdout := false
    var out io.Writer

    if opts.outfile != "" {
        out, err = os.Create(opts.outfile)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Failed to open output file: %v\n", err)
            return 1
        }
    } else {
        out = new(bytes.Buffer)
        useStdout = true
    }

    defer func() {
        if !success && !useStdout {
            out.(*os.File).Close()
            if err := os.Remove(opts.outfile); err != nil {
                fmt.Fprintf(os.Stderr, "Warning: failed to delete incomplete output file: %v\n", err)
            }
        } else if success && useStdout {
            os.Stdout.Write(out.(*bytes.Buffer).Bytes())
        }
    }()

    if err := gen.Output(out); err != nil {
        fmt.Fprintf(os.Stderr, "Failed to generate data: %v\n", err)
        return 1
    }

    success = true
    return 0
}

func main() {
    // wrap main since os.Exit doesn't call deferred functions
    os.Exit(realMain())
}
