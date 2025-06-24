package image

type Image struct{
	image string
}

func NewImage(image string)*Image{
	return &Image{image: image}
}

func (i*Image)getImage()(string,error){
	return i.image,nil
}