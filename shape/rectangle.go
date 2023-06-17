package shape

type Rectangle struct{
	Width float32
	Length float32
}

func (rec *Rectangle) Area() float32{
	return rec.Width * rec.Length
}

func (rec *Rectangle) Perimeter() float32{
	return 2 * (rec.Width + rec.Length)
}
