package step06

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 400
	height = 400

	// shader source
	vertexShaderSource = `
		#version 410
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	`
	fragmentShaderSource = `
		#version 410
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(1.0, 1.0, 0.0, 1.0);
		}
	`
)

// the vertices that will draw the triangle
var vertices = []float32{
	-0.5, -0.5, 0.0,
	0.5, -0.5, 0.0,
	0.0, 0.5, 0.0,
}

// Main draws a triangle.
// There's a whole lot of new code needed for this one!
func Main() {
	runtime.LockOSThread()
	initGlfw()
	window := createWindow("Step 06")
	defer glfw.Terminate()
	initGL()

	// create the shaders, program and vertex array
	program := initShaders()
	vao := initVertices()

	// pass vao and program to app loop
	appLoop(window, vao, program)
}

func initGlfw() {
	if err := glfw.Init(); err != nil {
		log.Fatal(err)
	}
}

func initGL() {
	if err := gl.Init(); err != nil {
		log.Fatal(err)
	}
}

func createWindow(title string) *glfw.Window {
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	window.SetFramebufferSizeCallback(framebufferSizeCallback)
	window.MakeContextCurrent()
	return window
}

func framebufferSizeCallback(window *glfw.Window, w, h int) {
	gl.Viewport(0, 0, int32(w), int32(h))
}

// app loop now gets the vao and program
func appLoop(window *glfw.Window, vao, program uint32) {
	for !window.ShouldClose() {
		processInput(window)
		gl.ClearColor(1.0, 0.0, 0.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// use this program and draw the vertex array with the shader.
		gl.UseProgram(program)
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/3))

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func processInput(window *glfw.Window) {
	if glfw.GetCurrentContext().GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}

// initialize vbo and vao
func initVertices() uint32 {
	// make vertex buffer object
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

	// make vertex array object
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

// initializes the shader programs
func initShaders() uint32 {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		log.Fatal(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		log.Fatal(err)
	}

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)
	return program
}

// compiles shader from shader source
func compileShader(source string, shaderType uint32) (uint32, error) {
	// 1. create shader
	shader := gl.CreateShader(shaderType)

	// 2. add source to shader
	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)

	// 3. free the source
	free()

	// 4. compile
	gl.CompileShader(shader)

	// 5. check compile status
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)

	// 6. if it failed...
	if status == gl.FALSE {
		return 0, fmt.Errorf("failed to compile shader:  %v", source)
	}

	// 7. success!
	return shader, nil
}
