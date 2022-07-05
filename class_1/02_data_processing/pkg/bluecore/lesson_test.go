package bluecore

import (
	"github.com/spf13/afero"
	"os"
	"reflect"
	"testing"
	"time"
)

var testData = `{
  "classification": "testing",
  "created": "2021-10-01 21:52:46.778659",
  "creator": "tj@example.com",
  "id": 12345,
  "journey_id": 999888777,
  "journey_type": "test journey",
  "key": "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
  "library_type": "test_type",
  "name": "Test Data",
  "nodes": {
    "11111": {
      "22222": {
        "33333": null
      }
    }
  },
  "segment_ids": [
    999888777
  ],
  "split_groups": [],
  "status": "test",
  "touch_to_campaign_map": {
    "11111": 11112,
    "22222": 22223,
    "33333": 33334
  }
}`
var testFileName = "mytest.json"

func initTest(t *testing.T, fs afero.Fs) {
	t.Helper()

	file, err := fs.OpenFile(testFileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		t.Fatalf("cannot open test file: %v", err)
	}
	defer file.Close()

	if n, err := file.Write([]byte(testData)); err != nil {
		t.Fatalf("received error writing test data to file: %v", err)
	} else if n != len(testData) {
		t.Fatalf("wrote an unexpected amount of data to file: exp(%d) act(%d)", len(testData), n)
	}
}

func TestLesson(t *testing.T) {
	fs := afero.NewMemMapFs()
	initTest(t, fs)
	data, err := ReadFileData(Config{
		Fs:       fs,
		FilePath: testFileName,
	})
	if err != nil {
		t.Fatalf("read file data error: %v", err)
	}

	jc, err := ProcessJsonJC(data)
	if err != nil {
		t.Fatalf("cannot unmarshal json data: %v", err)
	}

	// Layout is described in the time/format.go file
	// https://cs.opensource.google/go/go/+/refs/tags/go1.18.3:src/time/format.go;l=92
	// And in the example time.Format() provided here: https://pkg.go.dev/time#Time.Format
	expCreated, err := time.Parse(time.RFC3339Nano, "2021-10-01T21:52:46.778659000Z")
	if err != nil {
		t.Fatalf("cannot create expected time: %v", err)
	}

	expJourney := Journey{
		Classification: "testing",
		Creator:        "tj@example.com",
		Id:             12345,
		SegmentIds:     []int{999888777},
		TouchToCampMap: map[string]int{
			"11111": 11112,
			"22222": 22223,
			"33333": 33334,
		},
		Nodes: map[string]Node{
			"11111": Node{
				"22222": Node{
					"33333": nil,
				},
			},
		},
		SomeRandomProp: nil,
		Name:           interface{}("Test Data"),
		Key: Key{
			Name:  "Lorem",
			Other: " ipsum dolor sit amet, consectetur adipiscing elit",
			Len:   len("Lorem ipsum dolor sit amet, consectetur adipiscing elit"),
		},
		Created: CustomTime(expCreated),
	}

	if !reflect.DeepEqual(expJourney, jc) {
		t.Errorf("results mismatch: \nexp(%+v) \nact(%+v)", expJourney, jc)
	}
}
