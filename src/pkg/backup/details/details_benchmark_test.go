package details

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
)

var (
	details512k = detailsBuilder(512000)
	details256k = detailsBuilder(256000)
	details102k = detailsBuilder(102000)
)

func detailsBuilder(fileCount int) *Details {
	alpha := make([]string, 26)

	for i := 0; i < 26; i++ {
		alpha[i] = string(rune(97 + i))
	}

	alphaPath := basePath.Append(alpha...)
	fes := FolderEntriesForPath(alphaPath, alphaPath)

	// populate items
	for i := 0; i < fileCount; i++ {
		var (
			ri       = rand.Int31n(26)
			itemName = "item_" + uuid.NewString()
			itemP    = basePath.Append(alpha[:ri]...).Append(itemName)
		)

		info := ItemInfo{
			OneDrive: &OneDriveInfo{
				ItemName: itemName,
			},
		}

		ent := folderEntry{
			RepoRef:   itemP.String(),
			ShortRef:  itemP.ShortRef(),
			ParentRef: itemP.Dir().String(),
			Info:      info,
		}

		fes = append(fes, ent)
	}

	// populate 13 empty folders with dirmeta items
	for i := 0; i < 13; i++ {
		var (
			ri       = rand.Int31n(26)
			fldName  = "empty_" + uuid.NewString()
			fldP     = basePath.Append(alpha[:ri]...).Append(fldName)
			itemName = fldName + ".dirmeta"
			itemP    = fldP.Append(itemName)
		)

		// folder
		info := ItemInfo{
			Folder: &FolderInfo{
				DisplayName: fldName,
			},
		}

		ent := folderEntry{
			RepoRef:   fldP.String(),
			ShortRef:  fldP.ShortRef(),
			ParentRef: fldP.Dir().String(),
			Info:      info,
		}

		fes = append(fes, ent)

		// dirmeta
		info = ItemInfo{
			OneDrive: &OneDriveInfo{
				ItemName: itemName,
			},
		}

		ent = folderEntry{
			RepoRef:   itemP.String(),
			ShortRef:  itemP.ShortRef(),
			ParentRef: itemP.Dir().String(),
			Info:      info,
		}

		fes = append(fes, ent)
	}

	// shuffle the array, to avoid miscalculations due to ordering
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(fes), func(i, j int) {
		fes[i], fes[j] = fes[j], fes[i]
	})

	// return it as a details struct
	return toDetails(fes)
}

var result *Details

func BenchmarkDetailsFiltering_512_both(b *testing.B)      { benchmarkBoth(details512k, b) }
func BenchmarkDetailsFiltering_256_both(b *testing.B)      { benchmarkBoth(details256k, b) }
func BenchmarkDetailsFiltering_102_both(b *testing.B)      { benchmarkBoth(details102k, b) }
func BenchmarkDetailsFiltering_512_meta(b *testing.B)      { benchmarkMeta(details512k, b) }
func BenchmarkDetailsFiltering_256_meta(b *testing.B)      { benchmarkMeta(details256k, b) }
func BenchmarkDetailsFiltering_102_meta(b *testing.B)      { benchmarkMeta(details102k, b) }
func BenchmarkDetailsFiltering_512_container(b *testing.B) { benchmarkContainer(details512k, b) }
func BenchmarkDetailsFiltering_256_container(b *testing.B) { benchmarkContainer(details256k, b) }
func BenchmarkDetailsFiltering_102_container(b *testing.B) { benchmarkContainer(details102k, b) }

func benchmarkBoth(d *Details, b *testing.B) {
	var d2 *Details

	for n := 0; n < b.N; n++ {
		d2 = d.FilterMetaFiles().FilterEmptyContainers()
	}

	result = d2
}

func benchmarkMeta(d *Details, b *testing.B) {
	var d2 *Details

	for n := 0; n < b.N; n++ {
		d2 = d.FilterMetaFiles()
	}

	result = d2
}

func benchmarkContainer(d *Details, b *testing.B) {
	var d2 *Details

	for n := 0; n < b.N; n++ {
		d2 = d.FilterEmptyContainers()
	}

	result = d2
}
