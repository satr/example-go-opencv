package main

import (
	"gocv.io/x/gocv" //to install - run: go get -u -d gocv.io/x/gocv
)

func main() {
	webcam, _ := gocv.OpenVideoCapture(0) //0 - first webcam, 1 - second webcam (if exists), etc.
	defer webcam.Close()
	window := gocv.NewWindow("Some title")
	defer window.Close()
	img := gocv.NewMat()
	defer img.Close()

	//Loop to show webcam captured output
	for {
		webcam.Read(&img)
		window.IMShow(img)
		key := window.WaitKey(1)
		//exit on hit ESC
		if key == 27 {
			break
		}
	}
}
