package graphics

import (
	"fmt"
	"image/color"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func PlotGraph(x, y []float64, text string) error {
	p := plot.New()

	f, err := os.Create(fmt.Sprintf("%s.png", text))
	if err != nil {
		return fmt.Errorf("error creating png: %v", err)
	}
	defer f.Close()

	pxys := make(plotter.XYs, len(x))
	for i := range x {
		pxys[i].X = x[i]
		pxys[i].Y = y[i]
	}

	s, err := plotter.NewScatter(pxys)
	if err != nil {
		return fmt.Errorf("error creating scatter: %v", err)
	}
	s.Color = color.RGBA{R: 255, A: 255}

	p.Add(s)

	l, err := plotter.NewLine(pxys)
	if err != nil {
		return fmt.Errorf("error creating lines: %v", err)
	}
	l.Color = color.RGBA{G: 255, A: 255}

	p.Add(l)

	wt, err := p.WriterTo(512, 512, "png")
	if err != nil {
		return fmt.Errorf("error init plot writer: %v", err)
	}

	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("error writting plot: %v", err)
	}

	return nil
}
