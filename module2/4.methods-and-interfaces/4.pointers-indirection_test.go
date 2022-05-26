package main

// functions with a pointer argument must take a pointer
// methods with pointer receivers take either a value or a pointer as the receiver when they are called

// type Vertex struct {
// 	X, Y float64
// }

// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func ScaleFunc(v *Vertex, f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	v.Scale(2)
// 	ScaleFunc(&v, 10)

// 	p := &Vertex{4, 3}
// 	p.Scale(3)
// 	ScaleFunc(p, 8)
// }

// -----

// in reverse, functions with a value argument must take a value of that specific type
// and methods with value receivers take either a value or a pointer as the receiver when they are called

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func AbsFunc(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func main() {
// 	v := Vertex{3, 4}
// 	v.Abs()
// 	AbsFunc(v)

// 	p := &Vertex{4, 3}
// 	p.Abs()
// 	AbsFunc(*p)
// }
