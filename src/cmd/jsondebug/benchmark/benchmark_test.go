package benchmark

import (
	"compress/gzip"
	"io"
	"os"
	"testing"

	"github.com/alcionai/corso/src/cmd/jsondebug/common"
	"github.com/alcionai/corso/src/cmd/jsondebug/decoder"
)

func runBenchmarkByteInput(b *testing.B, d common.ByteManifestDecoder) {
	for i := 0; i < b.N; i++ {
		fn := common.ManifestFileName

		f, err := os.Open(fn)
		if err != nil {
			b.Logf("Error opening input file: %v", err)
			b.FailNow()
		}

		data, err := io.ReadAll(f)
		if err != nil {
			b.Logf("Error reading input data: %v", err)
			b.FailNow()
		}

		f.Close()
		b.ResetTimer()

		err = d.DecodeBytes(data, false)
		if err != nil {
			b.Logf("Error decoding json: %v", err)
			b.FailNow()
		}
	}
}

func runBenchmark(b *testing.B, d common.ManifestDecoder) {
	for _, unzip := range []string{"NotZipped", "Zipped"} {
		b.Run(unzip, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn := common.ManifestFileName
				if unzip == "Zipped" {
					fn += common.GzipSuffix
				}

				f, err := os.Open(fn)
				if err != nil {
					b.Logf("Error opening input file: %v", err)
					b.FailNow()
				}

				defer f.Close()

				var r io.ReadCloser = f

				if unzip == "Zipped" {
					r, err = gzip.NewReader(f)
					if err != nil {
						b.Logf("Error getting gzip reader: %v", err)
						b.FailNow()
					}

					defer r.Close()
				}

				b.ResetTimer()

				err = d.Decode(r, false)
				if err != nil {
					b.Logf("Error decoding json: %v", err)
					b.FailNow()
				}
			}
		})
	}
}

type benchmarkInfo struct {
	name string
	dec  common.Decoder
}

var decoderTable = []benchmarkInfo{
	{
		name: "Stdlib",
		dec:  decoder.Stdlib{},
	},
	{
		name: "JsonParser",
		dec:  decoder.JsonParser{},
	},
	{
		name: "Array",
		dec:  decoder.Array{},
	},
	{
		name: "ArrayFull",
		dec:  decoder.ArrayFull{},
	},
	{
		name: "Map",
		dec:  decoder.Map{},
	},
}

func Benchmark_FromFile(b *testing.B) {
	for _, benchmark := range decoderTable {
		b.Run(benchmark.name, func(b *testing.B) {
			runBenchmark(b, benchmark.dec)
		})
	}
}

func Benchmark_FromBytes(b *testing.B) {
	for _, benchmark := range decoderTable {
		b.Run(benchmark.name, func(b *testing.B) {
			runBenchmarkByteInput(b, benchmark.dec)
		})
	}
}
