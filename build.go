package main

import "bytes"
import "fmt"
import "os"
import "os/exec"
import "io/ioutil"
import "io"
import "strings"
import "golang.org/x/image/draw"
import "image/jpeg"
import "image/png"
import "image"
import "path"
import "regexp"

var inkscapeBin = "inkscape"
var force = false
var mode = ""
var out = ""
var lang = "en" // in english by default

func isOlder(src string, than string) bool {
	if force {
		return false
	}
	stat, err := os.Stat(src)

	// If the file does not exist, it is not older.
	if os.IsNotExist(err) {
		return false
	}

	statThan, _ := os.Stat(than)

	return stat.ModTime().After(statThan.ModTime())
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func rescale(ssrc string, sdst string) {

	input, _ := os.Open(ssrc)
	defer input.Close()

	output, err := os.Create(sdst)
	defer output.Close()
	if err != nil {
		panic(fmt.Sprintf("Error creating %s", sdst))
	}

	// Decode the image (from PNG to image.Image):
	var src image.Image
	if strings.Contains(ssrc, ".png") {
		src, err = png.Decode(input)
	} else {
		src, err = jpeg.Decode(input)
	}

	if err != nil {
		panic(fmt.Sprintf("Error loading %s", ssrc))

	}

	// Set the expected size that you want:
	dst := image.NewRGBA(image.Rect(0, 0, 100, 100.0*src.Bounds().Max.Y/src.Bounds().Max.X))

	// Resize:
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	// Encode to `output`:
	png.Encode(output, dst)
}

func prepareImg(src string, dst string) {
	// Compare last modified and skip if necessary
	if isOlder(dst, src) {
		return
	}

	if mode == "release" {
		// Copy
		fmt.Println("Copy:", src, "->", dst)
		copy(src, dst)
	} else {
		fmt.Println("Scale:", src, "->", dst)
		rescale(src, dst)
	}
}

func prepareDrawing(src string, dst string) {
	dst = strings.ReplaceAll(dst, ".svg", ".png")

	// Compare last modified and skip if necessary
	if isOlder(dst, src) {
		return
	}

	bin := inkscapeBin
	arg0 := "--export-filename=" + dst
	var arg1 = "--export-dpi=300"
	if mode == "debug" {
		arg1 = "--export-dpi=100"
	}
	arg2 := src

	fmt.Println(inkscapeBin, arg0, arg1, arg2)

	out, err := exec.Command(bin, arg0, arg1, arg2).CombinedOutput()

	if err != nil {
		fmt.Println("%s %s", string(out), err)
		return
	}
}

func prepare(folder string, f func(string, string)) {
	os.MkdirAll(cwd()+out+"/"+folder, os.ModePerm)
	items, _ := ioutil.ReadDir(folder)
	for _, item := range items {
		if item.IsDir() {
			os.MkdirAll(cwd()+out+"/"+folder+item.Name()+"/", os.ModePerm)
			prepare(folder+item.Name()+"/", f)
		}
		var src = cwd() + folder + item.Name()
		var dst = cwd() + out + "/" + folder + item.Name()
		f(src, dst)
	}
}

func toString(fs []string) string {
	var b bytes.Buffer
	for _, s := range fs {
		fmt.Fprintf(&b, "%s ", s)
	}
	return b.String()
}

func makeCover(src string, dst string) {
	src = cwd() + src
	dst = cwd() + dst

	if isOlder(dst, src) {
		return
	}
	bin := inkscapeBin
	args := make([]string, 0)
	args = append(args, "--export-filename="+dst)

	if mode == "debug" {
		args = append(args, "--export-dpi=100")
	} else {
		args = append(args, "--export-dpi=300")
	}
	args = append(args, "--export-type=pdf")
	args = append(args, "--export-text-to-path")
	// arg3 := "--verb"
	// arg4 := "ObjectToPath"
	// arg3 := ""
	args = append(args, src)

	fmt.Printf("%s %s\n", inkscapeBin, toString(args))

	out, err := exec.Command(bin, args...).CombinedOutput()
	if err != nil {
		fmt.Println("%s %s", string(out), err)
		return
	}
}

func cwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	cwd += "/"
	return cwd
}

func currentDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	return path.Base(cwd)
}

func checkExecutable(bin string) {
	path, err := exec.LookPath(bin)
	if err != nil {
		fmt.Println(fmt.Sprintf("Could not find executable '%s'", bin))
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("Found '%s' -> '%s'", bin, path))
}

// parse argument (debug|release|print -f fr|en)
func parseArgs() {
	var args = os.Args
	for _, arg := range args {
		if arg == "-f" {
			force = true
			continue
		}

		// lang = fr
		if arg == "fr" {
			lang = "fr"
			continue
		}

		// lang = en (default)
		if arg == "en" {
			lang = "en"
			continue
		}

		if arg == "debug" {
			mode = "debug"
			continue
		}

		if arg == "print" {
			mode = "print"
			continue
		}

		if arg == "release" {
			mode = "release"
			continue
		}
	}

	if mode == "" {
		fmt.Println("Mode must be either 'debug', 'release' or 'print'.")
		os.Exit(1)
	}
}

// hack book.tex to insert correct translation (if the file exists)
func insertLanguageIntoBookTex() string {

	// open book.tex
	input, err := ioutil.ReadFile("src/book.tex")
	if err != nil {
		fmt.Println("Could not open src/book.tex file",err)
		os.Exit(2)
	}

	// hack the \subfile if translation exists
	lines := strings.Split(string(input), "\n")
	langSufix := "-"+lang
	if lang == "en" { // no sufix in filename for english
		langSufix = ""
	}
	latexSubfileExpression := regexp.MustCompile(`\\subfile\{src\/([^-]+)(-.*)?\}`)
	for i,_ := range lines {
		if matches := latexSubfileExpression.FindStringSubmatch(lines[i]) ; matches != nil { // find subfile
			if matches[2] == langSufix { // same language, no need to translate
				continue
			}

			if fileExists,_ := os.Stat("src/"+matches[1]+langSufix+".tex") ; fileExists != nil { // translation file exists
				fmt.Println("Replace subfile",matches[1]+matches[2],"with",matches[1]+langSufix)
				lines[i] = "\\subfile{src/"+matches[1]+langSufix+"}"
			}
		}
	}

	// save the "hacked" book.tex
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("src/book.tex", []byte(output), 0644)
	if err != nil {
		fmt.Println("Could not save src/book.tex file",err)
		os.Exit(3)
	}

	return langSufix
}

func main() {
	
	checkExecutable(inkscapeBin)
	parseArgs()

	if lang == "en" {
		fmt.Println("Building in english in", mode, "mode...")
	} else if lang == "fr" {
		fmt.Println("Construction du livre en français en mode", mode, "...")
	}

	langSufix := insertLanguageIntoBookTex()
	//fmt.Printf("langSufix="+langSufix)

	compileOptions := ""
	if mode == "print" {
		compileOptions = `\def\forprint{}`
		mode = "release"
	}

	outputDirName := "out"
	out = outputDirName + "/" + mode
	os.MkdirAll(out, os.ModePerm)

	makeCover("src/cover/pdf/cover_front"+langSufix+".svg", out+"/illu/cover_front.pdf")
	makeCover("src/cover/pdf/cover_back"+langSufix+".svg", out+"/illu/cover_back.pdf")

	prepare("illu/img/", prepareImg)
	prepare("illu/d/", prepareDrawing)

	bin := "pdflatex"
	arg0 := "-output-directory"
	arg1 := outputDirName
	arg2 := `\def\base{` + out + `} ` + compileOptions + ` \input{src/book.tex}`
	
    var err error
	var out []byte
	draftMode := "-draftmode"

	// Compile in draft mode to generate only necessary files for Table of Contents
	if mode != "debug" {
		fmt.Println(bin, draftMode, arg0, arg1, arg2)
		out, err = exec.Command(bin, draftMode, arg0, arg1, arg2).CombinedOutput()
	}

	// Full Compile to generate PDF
	if err == nil {
		fmt.Println(bin, arg0, arg1, arg2)
		out, err = exec.Command(bin, arg0, arg1, arg2).CombinedOutput()
	}

	if err != nil {
		fmt.Println("%s %s", string(out), err)
    }

	// Rename
	var src = outputDirName + "/book.pdf"
	var dst =  outputDirName + "/" + currentDir() + "_" + mode + langSufix + ".pdf"
	os.Rename(src, dst)
}
