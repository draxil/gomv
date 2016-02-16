package gomv

import (
	"os"
	"testing"
	"io/ioutil"
	"path/filepath"
//	"log"
)

func TestBasicMove (t *testing.T) {
	a := tempfile()
	b := tempfile()
	f, err := os.Create(a)
	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte("fish"))
	if err != nil {
		panic(err)
	}
	f.Close()
	mve := MoveFile( a, b)
	
	if mve != nil {
		t.Fatal( "error on move : " + mve.Error() )
	}

	_, err = os.Stat(a)
	if err == nil  {
		t.Fatal("Old file still exists, or something else went wrong")
	}
	_, err = os.Stat(b)
	if err != nil {
		t.Fatal("Cant stat target")
	}
	f, err = os.Open(b)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll( f )
	if err != nil {
		panic(err)
	}
	
	if string(bytes) != "fish" {
		t.Fatal("File content didn't make it:" + string(bytes))
	}
	os.Remove(a)
	os.Remove(b)
}

func TestCrossVolumeUnix(t *testing.T){
	if os.Getenv("GOMV_TEST_CROSS_VOLUME_UNIX_TARGET") == "" {
		t.Skip("GOMV_TEST_CROSS_VOLUME_UNIX_TARGET not set skipping that test")
	}
	a := tempfile()
	_, fn := filepath.Split( a );
	b := os.Getenv("GOMV_TEST_CROSS_VOLUME_UNIX_TARGET") + string(os.PathSeparator) + fn

	f, err := os.Create(a)
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte("fish"))
	if err != nil {
		panic(err)
	}
	f.Close()

	mve := MoveFile( a, b)
	
	if mve != nil {
		t.Fatal( "error on move : " + mve.Error() )
	}
	
	_, err = os.Stat(a)
	if err == nil  {
		t.Fatal("Old file still exists, or something else went wrong")
	}
	_, err = os.Stat(b )
	if err != nil {
		t.Fatal("Cant stat target")
	}

	f2, err := os.Open(b)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll( f2 )
	if err != nil {
		panic(err)
	}
	
	if string(bytes) != "fish" {
		t.Fatal("File content didn't make it:" + string(bytes))
	}
	os.Remove(a)
	os.Remove(b)

}

func tempfile()(string){
	f, err := ioutil.TempFile("", "gomv") 
	if err != nil {
		panic( err )
	}
	return f.Name()
}
