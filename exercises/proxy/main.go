package main

/*
Let’s say we want to represent a proxy that controls access to a real image
object (like loading an image only when it’s needed). We'll have a RealImage
object that loads an image, and a ProxyImage object that controls when the
real image is loaded.
*/

type Image interface {
	Display()
}

type RealImage struct {
	fileName string
}

func (r *RealImage) Display() {
	println("Displaying", r.fileName)
}

func NewRealImage(fileName string) *RealImage {
	return &RealImage{fileName: fileName}
}

type ProxyImage struct {
	realImage *RealImage
	name      string
}

func NewProxyImage(fileName string) *ProxyImage {
	return &ProxyImage{name: fileName}
}

func (p *ProxyImage) Display() {
	if p.realImage == nil {
		p.realImage = NewRealImage(p.name)
	}
	p.realImage.Display()
}

func main() {
	var image1 Image = NewProxyImage("image1.jpg")
	var image2 Image = NewProxyImage("image2.png")

	image1.Display()

	image1.Display()

	image2.Display()
}
