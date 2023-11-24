package graphics

import (
	"fmt"
	"image/color"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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

func PlotHist(values []int, binCount int, name string, min, max float64) error {
	p := plot.New()
	//	p.X.Min = min
	//	p.X.Max = max

	vals := make(plotter.Values, len(values))
	for i := range vals {
		vals[i] = float64(values[i])
	}

	hist, err := plotter.NewHist(vals, binCount)
	if err != nil {
		return fmt.Errorf("new hist: %w", err)
	}
	hist.Width = float64(vg.Points(50))
	hist.Color = color.RGBA{G: 255, A: 255}

	p.Add(hist)

	err = p.Save(500, 500, fmt.Sprintf("%s.png", name))
	if err != nil {
		return fmt.Errorf("save: %w", err)
	}

	return nil
}
