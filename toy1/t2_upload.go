package toy1

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// 模拟客户端表单上传文件
// postFile <filename> to <targetUrl>
func postFile(filename string, targetURL string) error {
	bodyBuf := &bytes.Buffer{}

	bodyWriter := multipart.NewWriter(bodyBuf)
	defer bodyWriter.Close()
	formFile, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		return err
	}

	// open file and copy to formFile
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		return err
	}

	// contentType
	contentType := bodyWriter.FormDataContentType()

	fmt.Println("ContentType: " + contentType)
	fmt.Println(bodyBuf.String())

	// HTTP POST
	resp, err := http.Post(targetURL, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read Response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(respBody))

	return nil
}
