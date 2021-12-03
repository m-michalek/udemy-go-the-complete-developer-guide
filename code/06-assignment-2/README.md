# Here are the directions:

1. Write a program that creates two custom struct types called 'triangle' and 'square'
2. The 'square' type should be a struct with a field called 'sideLength' of type float64
3. The 'triangle' type should be a struct with a field called 'height' of type float64 and a field of type 'base' of type float64
4. Both types should have function called 'getArea' that returns the calculated area of the square or triangle
5. Area of a triangle = 0.5 * base * height. </br>
Area of a square = sideLength * sideLength
6. Add a 'shape' interface that defines a function called 'getArea'.  This function should calculate the area of the given shape and print it out to the terminal Design the interface so that the 'printArea' function can be called with either a triangle or a square.
