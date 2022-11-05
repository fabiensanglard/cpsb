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

var inkscapeBin = "inkscape"
var force = false

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

var mode = "release"
var out = ""

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

func getMode() string {
  var args = os.Args
  if len(args) > 1 {
		return args[1]
  }
  return ""
}

func main() {
	fmt.Println("Building...")

	checkExecutable(inkscapeBin)
	mode = getMode()
	
  var args = os.Args
	if len(args) > 2 {
		force = true
	}

	if mode != "debug" && mode != "release" && mode != "print" {
		fmt.Println("Mode must be either 'debug' or 'release' or 'print'.")
		return
	}

	compileOptions := ""
	if mode == "print" {
		compileOptions = `\def\forprint{}`
		mode = "release"
	}

	outputDirName := "out"
	out = outputDirName + "/" + mode
	os.MkdirAll(out, os.ModePerm)

	makeCover("src/cover/pdf/cover_front.svg", out+"/illu/cover_front.pdf")
	makeCover("src/cover/pdf/cover_back.svg", out+"/illu/cover_back.pdf")

	prepare("illu/img/", prepareImg)
	prepare("illu/d/", prepareDrawing)

	bin := "pdflatex"
	arg0 := "-output-directory"
	arg1 := outputDirName
	arg2 := `\def\base{` + out + `} ` + compileOptions + ` \input{src/book.tex}`
	fmt.Println(bin, arg0, arg1, arg2)

	out, err := exec.Command(bin, arg0, arg1, arg2).CombinedOutput()

	if err != nil {
		fmt.Println("%s %s", string(out), err)
  }

  // Rename
  var src = outputDirName + "/book.pdf"
  var dst =  outputDirName + "/" + currentDir() + "_" + getMode() + ".pdf"
	os.Rename(src, dst)

}
