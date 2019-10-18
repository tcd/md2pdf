package github

import (
	"fmt"
	"image"
	_ "image/gif" // gif, jpeg, & png imported for initialization side-effects. See: https://golang.org/pkg/image/#example__decodeConfig
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/lib"
)

// Image writes a JPEG, PNG or GIF to a gofpdf.Fpdf.
func Image(pdf *gofpdf.Fpdf, src, link string) {
	isRemote, err := regexp.MatchString(`^(https://|http://|www)`, strings.ToLower(src))
	if err != nil {
		log.Println(err)
	}

	if isRemote {
		width, height, format, err := remoteImageInfo(src)
		if err != nil {
			log.Println(err)
			return
		}
		remoteImage(pdf, src, width, height, format, "")
	} else {
		width, height, format, err := localImageInfo(src)
		if err != nil {
			log.Println(err)
			return
		}
		localImage(pdf, src, width, height, format, "")
	}
	pdf.Ln(9)
}

func remoteImage(pdf *gofpdf.Fpdf, src string, width, height int, format, link string) {
	res, err := http.Get(src)
	if err != nil {
		log.Println("Unable to fetch image: ", src)
		return
	}
	defer res.Body.Close()

	opt, err := newImageOpts(format)
	if err != nil {
		res.Body.Close()
		log.Println(err)
		return
	}

	_ = pdf.RegisterImageOptionsReader(src, opt, res.Body)
	res.Body.Close()

	x, y := pdf.GetXY()

	imgHeight := lib.PxToMm(height)
	imgWidth := lib.PxToMm(width)
	cbWidth := lib.ContentBoxWidth(pdf)
	if imgWidth > cbWidth {
		imgHeight = imgHeight * (cbWidth / imgWidth)
		imgWidth = cbWidth
	}

	pdf.ImageOptions(
		src,       // path to image
		x,         // x
		y,         // y
		imgWidth,  // width
		imgHeight, // height
		true,      // flow;  If flow is true, the current y value is advanced after placing the image and a page break may be made if necessary
		opt,       // options
		0,         // link
		"",        // linkStr
	)
}

func localImage(pdf *gofpdf.Fpdf, src string, width, height int, format, link string) {
	file, err := os.Open(src)
	if err != nil {
		log.Println(err)
		return
	}

	opt, err := newImageOpts(format)
	if err != nil {
		log.Println(err)
		return
	}

	_ = pdf.RegisterImageOptionsReader(src, opt, file)

	x, y := pdf.GetXY()

	imgHeight := lib.PxToMm(height)
	imgWidth := lib.PxToMm(width)
	cbWidth := lib.ContentBoxWidth(pdf)
	if imgWidth > cbWidth {
		imgHeight = imgHeight * (cbWidth / imgWidth)
		imgWidth = cbWidth
	}

	pdf.ImageOptions(
		src,       // path to image
		x,         // x
		y,         // y
		imgWidth,  // width
		imgHeight, // height
		true,      // flow;
		opt,       // options
		0,         // link
		"",        // linkStr
	)
}

func remoteImageInfo(imageURL string) (width, height int, format string, err error) {
	res, err := http.Get(imageURL)
	if err != nil {
		return 0, 0, "", err
	}
	defer res.Body.Close()

	image, format, err := image.DecodeConfig(res.Body)
	if err != nil {
		res.Body.Close()
		return 0, 0, "", err
	}

	res.Body.Close()
	return image.Width, image.Height, format, nil
}

func localImageInfo(imagePath string) (width, height int, format string, err error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return 0, 0, "", err
	}

	image, format, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, "", err
	}

	return image.Width, image.Height, format, nil
}

func newImageOpts(format string) (gofpdf.ImageOptions, error) {
	var opt gofpdf.ImageOptions
	var err error

	switch format {
	case "png":
		opt.ImageType = "png"
		opt.ReadDpi = true
	case "jpg":
		opt.ImageType = "jpg"
		opt.ReadDpi = false
	case "jpeg":
		opt.ImageType = "jpeg"
		opt.ReadDpi = false
	case "gif":
		opt.ImageType = "gif"
		opt.ReadDpi = false
	default:
		err = fmt.Errorf("Only jpg, png, & gif formats are supported 😔")
	}

	return opt, err
}
