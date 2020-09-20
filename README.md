# example-go-opencv
### Instal (OpenCV on Ubuntu)[https://docs.opencv.org/master/d2/de6/tutorial_py_setup_in_ubuntu.html]
Prerequisites
```
sudo apt-get install cmake gcc g++
sudo apt-get install cmake gcc g++ python3-dev python3-numpy 
sudo apt-get install libavcodec-dev libavformat-dev libswscale-dev libgstreamer-plugins-base1.0-dev libgstreamer1.0-dev
sudo apt-get install libgtk-3-dev

sudo apt-get install libpng-dev libjpeg-dev libopenexr-dev libtiff-dev libwebp-dev
```
Get sources
```
git clone https://github.com/opencv/opencv.git
```

Build
```
cd opencv
mkdir build && cd build
cmake -D OPENCV_GENERATE_PKGCONFIG=YES ..
make
sudo make install
```
NB: Flag `-D OPENCV_GENERATE_PKGCONFIG=YES` is used to avoid following error:
```
# pkg-config --cflags  -- opencv4
Package opencv4 was not found in the pkg-config search path.
Perhaps you should add the directory containing `opencv4.pc'
to the PKG_CONFIG_PATH environment variable
No package 'opencv4' found
```
It generates and installs a file `opencv4.pc` in `/usr/local/lib/pkgconfig/opencv4.pc`. 

Add following line to the `.bashrc`
```
export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/local/lib/pkgconfig/opencv4.pc
```
On error like
```
error while loading shared libraries: libopencv_dnn.so.4.5: cannot open shared object file: No such file or directory
```
Try to find it with a command `sudo find / -name "libopencv_dnn.so.4.5"` and add the path to the found lib to `LD_LIBRARY_PATH` to the `.bashrc`
```
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/lib
```
With fixed `LD_LIBRARY_PATH` go-application can be run with command `go run main.go`, but to run it with JetBrains GoLand - update system cache 
```
sudo ldconfig
```
Some (details)[https://medium.com/@meghamohan/everything-you-need-to-know-about-libraries-in-c-e8ad6138cbb4] about libraries.

Source the updated rc-file
```
source ~/.bashrc
```
