package admin

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	imgWebRoot   = "/upload/images"
	imgDiskRoot  = "webroot"
	imgOutputDir = imgDiskRoot + imgWebRoot
)

// ProcessImage copies image from file reader and processes it, returning its public web url
func ProcessImage(fileName string, entityType string, entityID string, file io.Reader) (webPath string, err error) {
	fileName = randomString(6) + strings.Split(filepath.Ext(fileName), "?")[0]
	fileDir := imgOutputDir + "/" + entityType + "/" + entityID
	filePath := fileDir + "/" + fileName
	webPath = imgWebRoot + "/" + entityType + "/" + entityID + "/" + fileName

	if _, err = os.Stat(fileDir); os.IsNotExist(err) {
		os.MkdirAll(fileDir, 0777)
	} else {
		if err != nil {
			return
		}
	}

	dst, err := os.Create(filePath)
	defer dst.Close()
	if err != nil {
		return
	}

	n, err := io.Copy(dst, file)
	if err != nil || n == 0 {
		os.Remove(filePath)
		log.Printf("%s failed copying to disk, %d bytes copied!", filePath, n)
		return
	}

	cmd := exec.Command("convert", filePath, "-resize", "500x500>", "-virtual-pixel", "white", "-set", "option:distort:viewport",
		"%[fx:max(w,h)]x%[fx:max(w,h)]-%[fx:max((h-w)/2,0)]-%[fx:max((w-h)/2,0)]",
		"-filter", "point", "-distort", "SRT", "0", "+repage", "-strip", filePath)
	gc := 0
resize:
	err = cmd.Start()
	if err != nil {
		if gc > 10 { // 10 retries fail :)
			os.Remove(filePath)
			log.Printf("%s resizing NOT STARTED: %s!", filePath, err)
			return
		}
		runtime.GC()
		gc++
		goto resize
	}
	err = cmd.Wait()
	if err != nil {
		os.Remove(filePath)
		log.Printf("%s resizing FAILED: %s!", filePath, err)
		return
	}
	return
}

func randomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func uploadImages(c *gin.Context, entityType string) (uploadedImages []string) {
	id := c.Param("id")
	uploadedImages = make([]string, 0, 1)
	reader, err := c.Request.MultipartReader()
	check(err)
	//copy each file to destination.
	for {
		file, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if file.FileName() == "" {
			continue
		}

		webPath, err := ProcessImage(file.FileName(), entityType, id, file)
		check(err)
		// success
		uploadedImages = append(uploadedImages, webPath)
	}
	return
}

func removeImage(c *gin.Context, images []string) []string {
	src := c.Param("src")
	src = src[1:] // cut first slash (double slash *src problem)

	for i, isrc := range images {
		if isrc == src {
			os.Remove(imgDiskRoot + src)
			images = append(images[:i], images[i+1:]...)
			break
		}
	}
	return images
}
