package Upload

import (
	"fmt"
	"io/ioutil"
	"net/http"

	l "github.com/acky666/ackyLog"
)

func HandleUploadedFile(w http.ResponseWriter, r *http.Request) {

	l.INFO("File Uploaded")

	file, handler, err := r.FormFile("portsFile")
	if err != nil {
		l.WARNING("Error Retrieving the File: %v", err)
		return
	}

	defer file.Close()
	l.INFO("Uploaded File: %+v Size: %+v \n", handler.Filename, handler.Size)

	tempFile, err := ioutil.TempFile("/tmp", "port-*.json")
	if err != nil {
		l.ERROR("Error opening Temp File %v", err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		l.ERROR("Error Saving Temp File %v", err)
	}

	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "Successfully Uploaded File\n")

	go processor(tempFile.Name())

}
