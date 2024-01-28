# GL Steps

This repo is just a set of example OpenGL applications done in [go-gl](https://github.com/go-gl/gl) as I learn my way around the technology.

There are a few useful resources that I've been working with:

- [https://kylewbanks.com/blog/tutorial-opengl-with-golang-part-1-hello-opengl](https://kylewbanks.com/blog/tutorial-opengl-with-golang-part-1-hello-opengl) - This one is specific to `go-gl`, but goes off in the direction of making a Conway's game of life application, so doesn't go too deep into any GL topics beyond making a white square. That said, for getting that far, it was indispensible. Thank you!
- [https://learnopengl.com/Introduction](https://learnopengl.com/Introduction) - a great reference, but written in C, so every single line has to be translated to Go, and that is not super easy. But with the help of the above link, I was able to get things up and running and have been using parts of this resource as a guide for each of the steps.
- [https://github.com/go-gl/example/blob/master/gl41core-cube/cube.go](https://github.com/go-gl/example/blob/master/gl41core-cube/cube.go) is another good resource. This is an official example from the `go-gl` repo. I guess this can serve pretty well as a best practices document.

But even with those as guides, there are SO many steps involved just with drawing a single triangle to the screen. Very difficult to know where to start. So I started breaking it down to the absolutely bare-bones minimum, making sure I understood that, and adding as little else as I could add for each step, aiming for a single unit of understanding in each one. 

Although not written expressly as a tutorial, I think it can be used pretty well as a way to learn `go-gl` step by step.

Here's where I'm at so far:

- [Step 1](step01/step01.go) This is the absolute minimum amount of code I could make and still have an OpenGL window appear on the screen and not crash.
    -   Initialize `glfw`
    -   Create a window
    -   Loop until the window is killed
- [Step 2](step02/step02.go) I add a bit more safety code here.
    - Lock the OS thread before doing anything, as is best practice in `go-gl`
    - Terminate `glfw` when the app is done
- [Step 3](step03/step03.go) Orginazation and window stuff.
    - Move a few of the steps into their own functions
    - Add quite a bit more to the code that creates the window
- [Step 4](step04/step04.go) Input processing.
    -  Defer the `glfw` termination (minor change)
    -  Input processing closes the window when the user hits the escape key
- [Step 5](step05/step05.go) Now I start to add some actual `gl` calls:
    - Initialize OpenGL
    - Set the viewport
    - Set a clear color, clear the screen and swap buffers to make the change take effect
    - Shows a red canvas
- [Step 6](step06/step06.go) Draws a triangle! Whole lot of changes here that couldn't really be broken down any further.
    - Create a vertex shader and fragment shader sources as string constants
    - Create shaders from the shader sources and create a program from the shaders
    - Create vertices as a list of float32s
    - Create a vertex buffer object and vertex array object from the vertices
    - Use the program, bind the vertex array and draw the vertices using the shader program
    - Shows a yellow triangle on the red canvas
- [Step 7](step07/step07.go) Loads the shaders and vertices from external files
    - Shaders are defined in two external text files
    - Vertices are defined in an external json file
    - Constants hold the paths to these files
    - Files are loaded and processed appropriately
    - Still draws a yellow triangle on a red canvas
- [Step 8](step08/step08.go) The vertex shader now defines the fragment color
    - Vertex shader creates an out vec4 and assigns a velue to it 
    - Fragment shader reads this value and uses it
- [Step 9](step09/step09.go) Uniforms
    - The fragment shader sets a uniform vec4 to read its color from
    - The main program looks for the location of the uniform variable
    - It sets the value of that uniform
- [Step 10](step10/step10.go) Animated color
    - The main program calculates an ever-changing set of r, g, b values based on time
    - It passes those rgb values to the uniform read by the fragment shader
- [Step 11](step11/step11.go) Some clean up and best practices
    - Lock os thread in init
    - Delete shaders after making program
    - Show gl version
