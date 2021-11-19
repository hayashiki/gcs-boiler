package gcsboiler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func Test_WriteReadDelete(t *testing.T) {
	ctx := context.Background()
	os.Setenv("STORAGE_EMULATOR_HOST", "http://localhost:4443")
	bucket := "sample"
	fromName := "dog1.jpg"
	objName := "dog2.jpg"
	f, err := os.Open(fmt.Sprintf("testdata/%s/%s", bucket, fromName))
	defer func() {
		cerr := f.Close()
		if cerr != nil {
			return
		}
		err = fmt.Errorf("failed to close: %v", err)
	}()
	if err != nil {
		t.Error(err)
	}
	gcs := New(bucket)
	if err := gcs.Write(ctx, objName, f); err != nil {
		t.Error(err)
	}
	_, err = gcs.Read(ctx, objName)

	//buf := bytes.Buffer{}
	//size, err := io.Copy(io.Discard, dog)
	//if err != nil {
	//	t.Error(err)
	//}
	//if got := size; got == 0 {
	//	t.Errorf("size more then 0, got: %v:", got)
	//}

	resp, err := http.Get(fmt.Sprintf("http://localhost:4443/%s/%s", bucket, objName))
	if err != nil {
		t.Error(err)
	}
	if want, got := http.StatusOK, resp.StatusCode; want != got {
		t.Errorf("want %v: got: %v", want, got)
	}

	if err := gcs.Delete(ctx, objName); err != nil {
		t.Error(err)
	}

	resp, err = http.Get(fmt.Sprintf("http://localhost:4443/%s/%s", bucket, objName))
	if err != nil {
		t.Error(err)
	}
	if want, got := http.StatusNotFound, resp.StatusCode; want != got {
		t.Errorf("want %v: got: %v", want, got)
	}
}
