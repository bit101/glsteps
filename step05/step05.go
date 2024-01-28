package step05

import (
	"log"
	"runtime"

	// imports gl
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 400
	height = 400
)

// Main pulls in actual gl calls. Initializes GL, sets window clear color, viewport size
func Main() {
	runtime.LockOSThread()
	initGlfw()
	window := createWindow("Step 05")

	// init gl
	initGL()

	defer glfw.Terminate()
	appLoop(window)
}

func initGlfw() {
	if err := glfw.Init(); err != nil {
		log.Fatal(err)
	}
}

// initializes opengl
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

	// set viewport size
	window.SetFramebufferSizeCallback(framebufferSizeCallback)
	window.MakeContextCurrent()

	return window
}

// sets viewport size
func framebufferSizeCallback(window *glfw.Window, w, h int) {
	gl.Viewport(0, 0, int32(w), int32(h))
}

func appLoop(window *glfw.Window) {
	for !window.ShouldClose() {
		processInput(window)

		// clear the window to red
		gl.ClearColor(1.0, 0.0, 0.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// swap buffers to update
		window.SwapBuffers()

		glfw.PollEvents()
	}
}

func processInput(window *glfw.Window) {
	if glfw.GetCurrentContext().GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}
