# GL Steps

This repo is just a set of example OpenGL applications done in [go-gl](https://github.com/go-gl/gl) as I learn my way around the technology.

There are a couple of useful resources that I've been working with:

- [https://learnopengl.com/Introduction](https://learnopengl.com/Introduction) - a great reference, but written in C, so every single line has to be translated to Go, and that is not super easy. But super useful as a general guide.
- [https://kylewbanks.com/blog/tutorial-opengl-with-golang-part-1-hello-opengl](https://kylewbanks.com/blog/tutorial-opengl-with-golang-part-1-hello-opengl) - This one is specific to `go-gl`, but goes off in the direction of making a Conway's game of life application, so doesn't go too deep into any GL topics beyond making a white square. That said, for getting that far, it was indispensible. Thank you!

But even with those as guides, there are SO many steps involved just with drawing a single triangle to the screen. Very difficult to no where to start. So I started breaking it down to the absolutely bare-bones minimum, making sure I understood that, and adding as little else as I could add in a single unit of understanding. Here's where I'm at so far:

- [Step 1](step01/step01.go) This is the absolute minimum amount of code I could make and still have an OpenGL window appear on the screen and not crash. It initializes `glfw`, creates a window and loops until the window is killed.
- [Step 2](step02/step02.go) I add a bit more safety code here. Just locking the OS thread, as is best practice in `go-gl`, before doing anything, and terminating `glfw` when I'm done.
- [Step 3](step03/step03.go) Here I move a few of the steps into their own functions. And I add quite a bit more to the code that creates the window.
- [Step 4](step04/step04.go) In this one, I defer the `glfw` termination (minor change) and start input processing. The input processing closes the window when the user hits the escape key.
- [Step 5](step05/step05.go) Now I start to add some actual `gl` calls:
    - Initialize OpenGL
    - Set the viewport
    - Set a clear color, clear the screen and swap buffers to make the change take effect
    - Shows a red canvas
- [Step 6](step06/step06.go) Draws a triangle!
    - Create a vertex shader and fragment shader sources as string constants
    - Create shaders from the shader sources and create a program from the shaders
    - Create vertices as a list of float32s
    - Create a vertex buffer object and vertex array object from the vertices
    - Use the program, bind the vertex array and draw the vertices using the shader program
    - Shows a yellow triangle on the red canvas
- [Step 7](step07/step07.go) Loads the shaders and vertices from external files
    - Shaders are defined in two external text files
    - Vertices are defined in an external json files
    - Constants hold the paths to these files
    - Files are loaded and processed appropriately
    - Still draws a yellow triangle on a red canvas
