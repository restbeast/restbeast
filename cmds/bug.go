// Copyright 2016 The Go Authors. All rights reserved.

package cmds

import (
	"bytes"
	. "fmt"
	"github.com/restbeast/restbeast/lib"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	urlpkg "net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
)

var cmdBug = &cobra.Command{
	Use:   "bug",
	Short: "start a bug report",
	Run:   runBug,
	Long: `
Bug opens the default browser and starts a new bug report.
The report includes useful system information.
	`,
}

func init() {
	rootCmd.AddCommand(cmdBug)
}

func runBug(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		Printf("go bug: bug takes no arguments")
	}
	var buf bytes.Buffer
	buf.WriteString(bugHeader)
	printGoVersion(&buf)
	printRestbeastVersion(&buf)
	buf.WriteString("### Does this issue reproduce with the latest release?\n\n\n")
	printEnvDetails(&buf)
	buf.WriteString(bugFooter)

	body := buf.String()
	url := "https://github.com/restbeast/restbeast/issues/new?body=" + urlpkg.QueryEscape(body)
	if !lib.OpenBrowser(url) {
		Print("Please file a new issue at https://github.com/restbeast/restbeast/issues/new using this template:\n\n")
		Print(body)
	}
}

const bugHeader = `<!-- Please answer these questions before submitting your issue. Thanks! -->

`
const bugFooter = `### What did you do?

<!--
If possible, provide a recipe for reproducing the error.
A complete runnable program is good.
A link on play.golang.org is best.
-->



### What did you expect to see?



### What did you see instead?

`

func printGoVersion(w io.Writer) {
	Fprintf(w, "### What version of Go are you using (`go version`)?\n\n")
	Fprintf(w, "<pre>\n")
	Fprintf(w, "$ go version\n")
	printCmdOut(w, "", "go", "version")
	Fprintf(w, "</pre>\n")
	Fprintf(w, "\n")
}

func printRestbeastVersion(w io.Writer) {
	Fprintf(w, "### What version of Restbeast are you using (`restbeast version`)?\n\n")
	Fprintf(w, "<pre>\n")
	Fprintf(w, "$ restbeast version\n")
	printCmdOut(w, "", "restbeast", "version")
	Fprintf(w, "</pre>\n")
	Fprintf(w, "\n")
}

func printEnvDetails(w io.Writer) {
	Fprintf(w, "### What operating system and processor architecture are you using (`go env`)?\n\n")
	Fprintf(w, "<details><summary><code>go env</code> Output</summary><br><pre>\n")
	Fprintf(w, "$ go env\n")
	printCmdOut(w, "", "go", "env")
	printGoDetails(w)
	printOSDetails(w)
	printCDetails(w)
	Fprintf(w, "</pre></details>\n\n")
}

func printGoDetails(w io.Writer) {
	printCmdOut(w, "GOROOT/bin/go version: ", filepath.Join(runtime.GOROOT(), "bin/go"), "version")
	printCmdOut(w, "GOROOT/bin/go tool compile -V: ", filepath.Join(runtime.GOROOT(), "bin/go"), "tool", "compile", "-V")
}

func printOSDetails(w io.Writer) {
	switch runtime.GOOS {
	case "darwin", "ios":
		printCmdOut(w, "uname -v: ", "uname", "-v")
		printCmdOut(w, "", "sw_vers")
	case "linux":
		printCmdOut(w, "uname -sr: ", "uname", "-sr")
		printCmdOut(w, "", "lsb_release", "-a")
		printGlibcVersion(w)
	case "openbsd", "netbsd", "freebsd", "dragonfly":
		printCmdOut(w, "uname -v: ", "uname", "-v")
	case "illumos", "solaris":
		// Be sure to use the OS-supplied uname, in "/usr/bin":
		printCmdOut(w, "uname -srv: ", "/usr/bin/uname", "-srv")
		out, err := ioutil.ReadFile("/etc/release")
		if err == nil {
			Fprintf(w, "/etc/release: %s\n", out)
		}
	}
}

func printCDetails(w io.Writer) {
	printCmdOut(w, "lldb --version: ", "lldb", "--version")
	cmd := exec.Command("gdb", "--version")
	out, err := cmd.Output()
	if err == nil {
		// There's apparently no combination of command line flags
		// to get gdb to spit out its version without the license and warranty.
		// Print up to the first newline.
		Fprintf(w, "gdb --version: %s\n", firstLine(out))
	}
}

// printCmdOut prints the output of running the given command.
// It ignores failures; 'go bug' is best effort.
func printCmdOut(w io.Writer, prefix, path string, args ...string) {
	cmd := exec.Command(path, args...)
	out, err := cmd.Output()
	if err != nil {
		return
	}
	Fprintf(w, "%s%s\n", prefix, bytes.TrimSpace(out))
}

// firstLine returns the first line of a given byte slice.
func firstLine(buf []byte) []byte {
	idx := bytes.IndexByte(buf, '\n')
	if idx > 0 {
		buf = buf[:idx]
	}
	return bytes.TrimSpace(buf)
}

// printGlibcVersion prints information about the glibc version.
// It ignores failures.
func printGlibcVersion(w io.Writer) {
	tempdir := os.TempDir()
	if tempdir == "" {
		return
	}
	src := []byte(`int main() {}`)
	srcfile := filepath.Join(tempdir, "go-bug.c")
	outfile := filepath.Join(tempdir, "go-bug")
	err := ioutil.WriteFile(srcfile, src, 0644)
	if err != nil {
		return
	}
	defer os.Remove(srcfile)
	cmd := exec.Command("gcc", "-o", outfile, srcfile)
	if _, err = cmd.CombinedOutput(); err != nil {
		return
	}
	defer os.Remove(outfile)

	cmd = exec.Command("ldd", outfile)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	re := regexp.MustCompile(`libc\.so[^ ]* => ([^ ]+)`)
	m := re.FindStringSubmatch(string(out))
	if m == nil {
		return
	}
	cmd = exec.Command(m[1])
	out, err = cmd.Output()
	if err != nil {
		return
	}
	Fprintf(w, "%s: %s\n", m[1], firstLine(out))

	// print another line (the one containing version string) in case of musl libc
	if idx := bytes.IndexByte(out, '\n'); bytes.Index(out, []byte("musl")) != -1 && idx > -1 {
		Fprintf(w, "%s\n", firstLine(out[idx+1:]))
	}
}
