package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"simple-msd/core"
)

func main() {
	app := &cli.App{
		Name:  "simple-msd",
		Usage: "A simple multi-source downloader",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Usage:    "URLs to download",
				Required: true,
			},
			&cli.Int64Flag{
				Name:        "split",
				Aliases:     []string{"s"},
				Usage:       "Download a file using N number of chunks. This will also determine the number of concurrent workers.",
				DefaultText: "5",
			},
			&cli.Int64Flag{
				Name:        "min-split-size",
				Aliases:     []string{"k"},
				Usage:       "Minimum size of a single chunk in bytes. If the size of the file is less than this value, it will be downloaded in a single chunk.",
				DefaultText: "20MB",
			},
			&cli.BoolFlag{
				Name:        "single-bar",
				Aliases:     []string{"b"},
				Usage:       "Enable single progress bar for downloading",
				DefaultText: "Multiple Progress Bar",
			},
		},
		Action: func(context *cli.Context) error {
			splitSize := context.Int64("split")
			if splitSize <= 0 {
				log.Println("Split value must be greater than 0, using default value of 5")
				splitSize = 5
			}

			minSplitSize := context.Int64("min-split-size")
			if minSplitSize <= 0 {
				log.Println("Minimum value must be greater than 0, using default value of 20MB")
				minSplitSize = 1024 * 1024 * 20
			}

			singleProgressBar := context.Bool("single-bar")

			urls := core.NewUrlCollector(context.StringSlice("url")...)
			f, err := core.NewFileDetails(*urls)
			if err != nil {
				log.Fatalln(err)
			}

			d := core.Downloader{
				File:              *f,
				ChunkSize:         minSplitSize,
				NumberOfChunks:    splitSize,
				SingleProgressBar: singleProgressBar}

			d.Execute()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
