package step12

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 400
	height = 400

	// these are altered to get per-vertex colors
	vertexShaderSource   = "step12/vertex_shader.gl"
	fragmentShaderSource = "step12/fragment_shader.gl"

	// this now contains color data per vertex
	verticesSource = "step12/vertices.json"
)

var (
	vertices []float32
)

func init() {
	runtime.LockOSThread()
}

// Main - We have colors mixed in with coordinates in the vertex data.
// we have to parse those in initVertices
func Main() {
	initGlfw()
	window := createWindow("Step 12")
	defer glfw.Terminate()

	initGL()
	program := initShaders()
	vao := initVertices()

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

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)
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

func appLoop(window *glfw.Window, vao, program uint32) {
	for !window.ShouldClose() {
		processInput(window)

		// changed clear color to black to better see the triangle colors
		gl.ClearColor(0.0, 0.0, 0.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// removed dynamic changing color
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

func initVertices() uint32 {
	verticesData, err := os.ReadFile(verticesSource)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(verticesData, &vertices)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	// this gets the position values from the vertex array.
	// it's the first three values in each row.
	stride := int32(6 * 4) // 6 ints, 4 bytes per int
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride, nil)
	gl.EnableVertexAttribArray(0)

	// this gets the color values, the second three values in each row
	// pointer points to where the color values start: 3 floats in * 4 byes per float.
	var pointer uintptr = 3 * 4
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, stride, gl.Ptr(pointer))
	gl.EnableVertexAttribArray(1)

	return vao
}

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

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	return program
}

func compileShader(path string, shaderType uint32) (uint32, error) {
	sourceBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	source := string(sourceBytes)

	shader := gl.CreateShader(shaderType)
	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		return 0, fmt.Errorf("failed to compile shader:  %v", source)
	}

	return shader, nil
}
