package step10

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width                = 400
	height               = 400
	vertexShaderSource   = "step10/vertex_shader.gl"
	fragmentShaderSource = "step10/fragment_shader.gl"
	verticesSource       = "step10/vertices.json"
)

var (
	vertices []float32
)

// Main - first pass at animation. The uniform color passed into the
// fragment shader is now changing over time.
func Main() {
	runtime.LockOSThread()
	initGlfw()
	window := createWindow("Step 10")
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
		gl.ClearColor(1.0, 0.0, 0.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		vertexColorLocation := gl.GetUniformLocation(program, gl.Str("uni_color\x00"))
		gl.UseProgram(program)

		// get a color from the method
		r, g, b := getColor()
		gl.Uniform4f(vertexColorLocation, r, g, b, 1.0)

		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/3))
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

// calculate an ever changing r, g, b value based on time.
func getColor() (float32, float32, float32) {
	timeValue := glfw.GetTime()
	red := math.Sin(timeValue*1.9)/2 + 0.5
	green := math.Sin(timeValue*1.3)/2 + 0.5
	blue := math.Sin(timeValue*0.7)/2 + 0.5
	return float32(red), float32(green), float32(blue)
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
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

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
	return program
}

func compileShader(path string, shaderType uint32) (uint32, error) {
	sourceBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
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
