// Package gomv: move files without caring if we're moving to another device etc.
package gomv

import (
	"os"
	"errors"
	"io"
)

const invalidCrossDevice = "invalid cross-device link"
const crossDevice = "cross-device link"

var (
	/* ErrDestDir : where destination is a directory */
	ErrDestDir = errors.New("Desination file can't be a directory")
	/* ErrSourceDir : where source is a directory */
	ErrSourceDir = errors.New("Source file can't be a directory")
)

/* MoveFile : move source file to dest file. Will try to rename the file, if that fails in a way we recognise as an OS restriction we'll copy the file and then remove the original.
*/
func MoveFile(source string, dest string) (error) {

	di, err := os.Stat(dest)
	
	if err != nil && ! os.IsNotExist(err) {
		return errors.New("Error checking dest file: " + err.Error())
	} 
	
	if err == nil && di.IsDir() {
		return ErrDestDir
	}

	si, err := os.Stat(source)
	if err != nil {
		return errors.New("Cannot stat source file: " + err.Error())
	}
	if err == nil && si.IsDir() {
		return ErrSourceDir
	}

	err = os.Rename(source, dest)

	if err == nil {
		// job done
		return nil
	}
	
	 li, ok := err.(*os.LinkError)
	
	if ! ok {
		return err
	}

	switch li.Err.Error() {
	case invalidCrossDevice, crossDevice: 
		return cpmv( source, dest, si )
	}
	
	return err
}

func cpmv(source string, dest string, si os.FileInfo) (error) {
	cpe := copyfile( source, dest, si )
	if cpe != nil {
		return errors.New( "Could not copymove file, copy failed: " + cpe.Error())
	}
	rme := os.Remove( source )
	if rme != nil {
		return rme
	}
	return nil
}

func copyfile(source string, dest string, si os.FileInfo ) ( err error ){
	in, err := os.Open(source)

	if err != nil {
		return
	}
	
	defer in.Close()
	
	out, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, si.Mode())
	
	if err != nil {
		return
	}
	
	defer func() {
		if e := out.Close(); e != nil {
			err = e
	
		}
	}()
	
	_, err = io.Copy(out, in)
	
	return
}
