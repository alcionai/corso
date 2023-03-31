package benchmark

import (
	"compress/gzip"
	"io"
	"os"
	"testing"

	"github.com/alcionai/corso/src/cmd/jsondebug/common"
	"github.com/alcionai/corso/src/cmd/jsondebug/decoder"
)

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

func Benchmark_Jsonparser(b *testing.B) {
	d := decoder.JsonParser{}
	runBenchmark(b, d)
}

func Benchmark_Stdlib(b *testing.B) {
	d := decoder.Stdlib{}
	runBenchmark(b, d)
}

func Benchmark_Array(b *testing.B) {
	d := decoder.Array{}
	runBenchmark(b, d)
}

func Benchmark_ArrayFull(b *testing.B) {
	d := decoder.ArrayFull{}
	runBenchmark(b, d)
}

func Benchmark_Map(b *testing.B) {
	d := decoder.Map{}
	runBenchmark(b, d)
}
