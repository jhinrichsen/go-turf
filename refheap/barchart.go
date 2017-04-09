package main

import (
	"flag"
	"github.com/ajstarks/svgo"
	"github.com/vdobler/chart"
	"github.com/vdobler/chart/imgg"
	"github.com/vdobler/chart/svgg"
	"github.com/vdobler/chart/txtg"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
//	"math"
	"os"
//	"time"
)

// -------------------------------------------------------------------------
// Dumper

// Dumper helps saving plots of size WxH in a NxM grid layout
// in several formats
type Dumper struct {
	N, M, W, H, Cnt           int
	S                         *svg.SVG
	I                         *image.RGBA
	svgFile, imgFile, txtFile *os.File
}

func NewDumper(name string, n, m, w, h int) *Dumper {
	var err error
	dumper := Dumper{N: n, M: m, W: w, H: h}

	dumper.svgFile, err = os.Create(name + ".svg")
	if err != nil {
		panic(err)
	}
	dumper.S = svg.New(dumper.svgFile)
	dumper.S.Start(n*w, m*h)
	dumper.S.Title(name)
	dumper.S.Rect(0, 0, n*w, m*h, "fill: #ffffff")

	dumper.imgFile, err = os.Create(name + ".png")
	if err != nil {
		panic(err)
	}
	dumper.I = image.NewRGBA(image.Rect(0, 0, n*w, m*h))
	bg := image.NewUniform(color.RGBA{0xff, 0xff, 0xff, 0xff})
	draw.Draw(dumper.I, dumper.I.Bounds(), bg, image.ZP, draw.Src)

	dumper.txtFile, err = os.Create(name + ".txt")
	if err != nil {
		panic(err)
	}

	return &dumper
}
func (d *Dumper) Close() {
	png.Encode(d.imgFile, d.I)
	d.imgFile.Close()

	d.S.End()
	d.svgFile.Close()

	d.txtFile.Close()
}

func (d *Dumper) Plot(c chart.Chart) {
	row, col := d.Cnt/d.N, d.Cnt%d.N

	igr := imgg.AddTo(d.I, col*d.W, row*d.H, d.W, d.H, color.RGBA{0xff, 0xff, 0xff, 0xff}, nil, nil)
	c.Plot(igr)

	sgr := svgg.AddTo(d.S, col*d.W, row*d.H, d.W, d.H, "", 12, color.RGBA{0xff, 0xff, 0xff, 0xff})
	c.Plot(sgr)

	tgr := txtg.New(100, 30)
	c.Plot(tgr)
	d.txtFile.Write([]byte(tgr.String() + "\n\n\n"))

	d.Cnt++

}

//
// Bar Charts
//
func barChart() {
	dumper := NewDumper("xbar1", 3, 2, 400, 300)
	defer dumper.Close()

	green := chart.Style{Symbol: '#', LineColor: color.NRGBA{0x00, 0xcc, 0x00, 0xff},
		FillColor: color.NRGBA{0x80, 0xff, 0x80, 0xff},
		LineStyle: chart.SolidLine, LineWidth: 2}

	barc := chart.BarChart{Title: "Connection pool"}
	// barc.Key.Hide = true
	// barc.XRange.Fixed(0, 60, 10)

	barc.XRange.Time = true
	// All x and all y values
	barc.AddDataPair("free", []float64{10, 20, 30, 40}, []float64{40, 130, 15, 100}, green)

	dumper.Plot(&barc)
}

//
//  Main
//
func main() {
	var debugging = flag.Bool("debug", false, "output debug information to stderr")

	var bar *bool = flag.Bool("bar", false, "show bar charts")

	flag.Parse()
	if *debugging {
		chart.DebugLogger = log.New(os.Stdout, "", log.LstdFlags)
	}

	if *bar {
		barChart()
	}
}
