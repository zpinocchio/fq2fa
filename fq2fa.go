package main

import (

	"flag"
	"github.com/shenwei356/bio/seq"
	"github.com/shenwei356/bio/seqio/fastx"
	"github.com/shenwei356/xopen"
	"io"
	"os"
	)

func main() {

	var alphabet *seq.Alphabet
	idRegexp := "^(\\S+)\\s?"
	LineWidth := 0
	//byname := true
	var infileR1, infileR2, outfile string
	var h bool
	flag.StringVar(&infileR1, "R1", "clean_R1.fastq.gz", "input R1 file,support gz or fq")
	flag.StringVar(&infileR2, "R2", "clean_R2.fastq.gz", "input R2 file,support gz or fq")
	flag.BoolVar(&h, "h", false, "same like idba/fq2fa but support gz\nconvert R1 R2 fastq to [R1R2]'fasta")
	flag.StringVar(&outfile, "o", ".", "output file")

	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(1)
	}

	var recordR1 *fastx.Record
	var recordR2 *fastx.Record
	var fastxReaderR1 *fastx.Reader
	var fastxReaderR2 *fastx.Reader

	out, _ := xopen.Wopen(outfile)
	defer out.Close()
	fastxReaderR1, err := fastx.NewReader(alphabet, infileR1, idRegexp)
	fastxReaderR2, err = fastx.NewReader(alphabet, infileR2, idRegexp)

	if err != nil {
		panic(err)
	}
	for {
		recordR1, err = fastxReaderR1.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			break
		}
		recordR2, err = fastxReaderR2.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			break
		}
		recordR1.Seq.Qual = []byte{}
		recordR2.Seq.Qual = []byte{}
		recordR1.FormatToWriter(out, LineWidth)
		recordR2.FormatToWriter(out, LineWidth)

	}
}
