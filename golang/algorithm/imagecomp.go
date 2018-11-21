package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	IMAGECOMP_MAX_URL_COUNT = 1000

	IMAGECOMP_ORDER_BY_ROW = 0
	IMAGECOMP_ORDER_BY_COL = 1
)

const (
	H_ALIGN_LEFT   = "left"
	H_ALIGN_RIGHT  = "right"
	H_ALIGN_CENTER = "center"
	V_ALIGN_TOP    = "top"
	V_ALIGN_BOTTOM = "bottom"
	V_ALIGN_MIDDLE = "middle"
)

/*

imagecomp
/bucket/<string>
/format/<string> 	optional, default jpg
/rows/<int>			optional, default 1
/cols/<int>			optional, default 1
/halign/<string> 	optional, default left
/valign/<string> 	optional, default top
/order/<int>		optional, default 1
/alpha/<int> 		optional, default 0
/bgcolor/<string>	optional, default gray
/margin/<int>		optional, default 0
/url/<string>
/url/<string>

*/

func maxInt(vals ...int) int {
	var max int
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	var files string
	var destFile string
	var format string
	var halign string
	var valign string
	var rows int
	var cols int
	var order int
	var alpha int
	var margin int
	var bgColorStr string
	flag.StringVar(&files, "files", "", "file path joined by comma(,)")
	flag.StringVar(&destFile, "dest", "", "dest image file path")
	flag.StringVar(&format, "format", "", "dest image format, png, jpg or jpeg")
	flag.StringVar(&halign, "halign", "", "horizontal align, left, right or center")
	flag.StringVar(&valign, "valign", "", "vertical align, top, bottom or middle")
	flag.IntVar(&rows, "rows", 0, "dest image layout rows")
	flag.IntVar(&cols, "cols", 0, "dest image layout cols")
	flag.IntVar(&order, "order", 0, "dest sub image layout order, 0 or 1")
	flag.IntVar(&alpha, "alpha", 0, "dest image alpha, 0 ~ 255")
	flag.IntVar(&margin, "margin", 0, "dest sub image margin")
	flag.StringVar(&bgColorStr, "bgcolor", "", "background color, like #000000")
	flag.Parse()

	//format
	if format == "" {
		format = "jpg"
	}

	//halign
	if halign == "" {
		halign = H_ALIGN_LEFT
	}

	//valign

	if valign == "" {
		valign = V_ALIGN_TOP
	}

	if format == "png" {
		alpha = 0
	}

	if alpha < 0 || alpha > 255 {
		fmt.Println("invalid `alhpa`, should between [0,255]")
		return
	}

	//bgcolor, default white
	bgColor := color.RGBA{0xFF, 0xFF, 0xFF, uint8(alpha)}

	if bgColorStr != "" {
		colorPattern := `^#[a-fA-F0-9]{6}$`
		if matched, _ := regexp.Match(colorPattern, []byte(bgColorStr)); !matched {
			fmt.Println("invalid `bgcolor`, should in format '#FFFFFF'")
			return
		}

		bgColorStr = bgColorStr[1:]

		redPart := bgColorStr[0:2]
		greenPart := bgColorStr[2:4]
		bluePart := bgColorStr[4:6]

		redInt, _ := strconv.ParseInt(redPart, 16, 64)
		greenInt, _ := strconv.ParseInt(greenPart, 16, 64)
		blueInt, _ := strconv.ParseInt(bluePart, 16, 64)

		bgColor = color.RGBA{
			uint8(redInt),
			uint8(greenInt),
			uint8(blueInt),
			uint8(alpha),
		}
	}

	//check rows and cols valid or not
	fileItems := strings.Split(files, ",")
	fileCount := len(fileItems)

	if rows == 0 && cols == 0 {
		cols = 1
		rows = fileCount / cols
	} else if rows == 0 && cols != 0 {
		if cols > fileCount {
			fmt.Println("cols larger than url count error")
			return
		}
		if fileCount%cols == 0 {
			rows = fileCount / cols
		} else {
			rows = fileCount/cols + 1
		}
	} else if rows != 0 && cols == 0 {
		if rows > fileCount {
			fmt.Println("rows larger than url count error")
			return
		}
		if fileCount%rows == 0 {
			cols = fileCount / rows
		} else {
			cols = fileCount/rows + 1
		}
	} else {
		if fileCount > rows*cols {
			fmt.Println("url count larger than rows*cols error")
			return
		}

		if fileCount < rows*cols {
			switch order {
			case IMAGECOMP_ORDER_BY_ROW:
				if fileCount < (rows-1)*cols+1 {
					fmt.Println("url count less than (rows-1)*cols+1 error")
					return
				}
			case IMAGECOMP_ORDER_BY_COL:
				if fileCount < rows*(cols-1)+1 {
					fmt.Println("url count less than rows*(cols-1)+1 error")
					return
				}
			}
		}
	}

	//layout the images
	localImgFps := make([]*os.File, 0, fileCount)

	var localImgObjs [][]image.Image = make([][]image.Image, rows*cols)

	for index := 0; index < rows; index++ {
		localImgObjs[index] = make([]image.Image, cols)
	}

	var rowIndex int = 0
	var colIndex int = 0

	for _, file := range fileItems {

		imgFp, openErr := os.Open(file)
		if openErr != nil {
			fmt.Println("open local image of failed", openErr.Error())
			return
		}

		localImgFps = append(localImgFps, imgFp)

		var imgObj image.Image
		var dErr error

		if strings.HasSuffix(file, ".png") {
			imgObj, dErr = png.Decode(imgFp)
			if dErr != nil {
				fmt.Println("decode png image failed", file, dErr.Error())
				return
			}
		} else if strings.HasSuffix(file, ".jpg") || strings.HasSuffix(file, ".jpeg") {
			imgObj, dErr = jpeg.Decode(imgFp)
			if dErr != nil {
				fmt.Println("decode jpeg image failed", file, dErr.Error())
				return
			}
		} else {
			fmt.Println("src image must have suffix", file)
			return
		}

		localImgObjs[rowIndex][colIndex] = imgObj

		//update index
		switch order {
		case IMAGECOMP_ORDER_BY_ROW:
			if colIndex < cols-1 {
				colIndex += 1
			} else {
				colIndex = 0
				rowIndex += 1
			}

		case IMAGECOMP_ORDER_BY_COL:
			if rowIndex < rows-1 {
				rowIndex += 1
			} else {
				rowIndex = 0
				colIndex += 1
			}
		}
	}

	//close file handlers
	defer func() {
		for _, fp := range localImgFps {
			fp.Close()
		}
	}()

	//calc the dst image size
	dstImageWidth := 0
	dstImageHeight := 0

	rowImageMaxWidths := make([]int, 0)
	rowImageMaxHeights := make([]int, 0)

	for _, rowSlice := range localImgObjs {
		if len(rowSlice) == 0 {
			continue
		}

		rowImageColWidths := make([]int, 0)
		rowImageColHeights := make([]int, 0)

		for _, imgObj := range rowSlice {
			if imgObj != nil {
				bounds := imgObj.Bounds()
				rowImageColWidths = append(rowImageColWidths, bounds.Dx())
				rowImageColHeights = append(rowImageColHeights, bounds.Dy())
			}
		}

		rowImageColMaxWidth := maxInt(rowImageColWidths...)
		rowImageColMaxHeight := maxInt(rowImageColHeights...)

		rowImageMaxWidths = append(rowImageMaxWidths, rowImageColMaxWidth)
		rowImageMaxHeights = append(rowImageMaxHeights, rowImageColMaxHeight)
	}

	blockWidth := maxInt(rowImageMaxWidths...)
	blockHeight := maxInt(rowImageMaxHeights...)

	//dest image width & height with margin
	dstImageWidth = blockWidth*cols + (cols+1)*margin
	dstImageHeight = blockHeight*rows + (rows+1)*margin

	//compose the dst image
	dstRect := image.Rect(0, 0, dstImageWidth, dstImageHeight)
	dstImage := image.NewRGBA(dstRect)

	draw.Draw(dstImage, dstImage.Bounds(), image.NewUniform(bgColor), image.ZP, draw.Src)

	for rowIndex, rowSlice := range localImgObjs {
		for colIndex := 0; colIndex < len(rowSlice); colIndex++ {
			imgObj := rowSlice[colIndex]

			//check nil
			if imgObj == nil {
				continue
			}

			imgWidth := imgObj.Bounds().Max.X - imgObj.Bounds().Min.X
			imgHeight := imgObj.Bounds().Max.Y - imgObj.Bounds().Min.Y

			//calc the draw rect start point
			p1 := image.Point{
				colIndex*blockWidth + (colIndex+1)*margin,
				rowIndex*blockHeight + (rowIndex+1)*margin,
			}

			//check halign and valign
			//default is left and top
			switch halign {
			case H_ALIGN_CENTER:
				offset := (blockWidth - imgWidth) / 2
				p1.X += offset
			case H_ALIGN_RIGHT:
				offset := (blockWidth - imgWidth)
				p1.X += offset
			}

			switch valign {
			case V_ALIGN_MIDDLE:
				offset := (blockHeight - imgHeight) / 2
				p1.Y += offset
			case V_ALIGN_BOTTOM:
				offset := (blockHeight - imgHeight)
				p1.Y += offset
			}

			//calc the draw rect end point
			p2 := image.Point{}
			p2.X = p1.X + blockWidth
			p2.Y = p1.Y + blockHeight

			drawRect := image.Rect(p1.X, p1.Y, p2.X, p2.Y)

			//draw
			draw.Draw(dstImage, drawRect, imgObj, imgObj.Bounds().Min, draw.Src)
		}
	}

	//write result
	var buffer = bytes.NewBuffer(nil)
	switch format {
	case "png":
		eErr := png.Encode(buffer, dstImage)
		if eErr != nil {
			fmt.Println("create dst png image failed", eErr)
			return
		}

	case "jpg", "jpeg":
		eErr := jpeg.Encode(buffer, dstImage, &jpeg.Options{
			Quality: 75,
		})
		if eErr != nil {
			fmt.Println("create dst jpeg image failed", eErr)
			return
		}
	}

	ioutil.WriteFile(destFile, buffer.Bytes(), 0644)
}
