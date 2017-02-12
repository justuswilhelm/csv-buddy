package lib

import (
	"fmt"
	"github.com/montanaflynn/stats"
)

// Column stores a data frame columns
type Column struct {
	Name    string
	Content []float64
}

// Median TODO
func (c *Column) Median() (float64, error) {
	return stats.Median(c.Content)
}

// Mean TODO
func (c *Column) Mean() (float64, error) {
	return stats.Mean(c.Content)
}

// StdDevP TODO
func (c *Column) StdDevP() (float64, error) {
	return stats.StdDevP(c.Content)
}

// Summary TODO
func (c *Column) Summary() string {
	mean, err := c.Mean()
	if err != nil {
		return fmt.Sprintf("Error in mean: %+v", err)
	}

	median, err := c.Median()
	if err != nil {
		return fmt.Sprintf("Error in median: %+v", err)
	}

	stdDevP, err := c.StdDevP()
	if err != nil {
		return fmt.Sprintf("Error in stdevp: %+v", err)
	}

	return fmt.Sprintf(
		"%s: Mean %f, Median %f, StdDev %f",
		c.Name,
		mean,
		median,
		stdDevP)
}
