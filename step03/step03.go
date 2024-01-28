package step03

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 400
	height = 400
)

// Main breaks it into functions, more stuff around window creation.
func Main() {
	// 1. Ensure we run in single threaded mode here.
	runtime.LockOSThread()

	// 2. Initialize GLFW.
	initGlfw()

	// 3. Create a window.
	window := createWindow("Step 03")

	// 4. Event loop runs until the window is closed.
	appLoop(window)

	// 5. Clean up.
	glfw.Terminate()
}

func initGlfw() {
	if err := glfw.Init(); err != nil {
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

	window.MakeContextCurrent()

	return window
}

func appLoop(window *glfw.Window) {
	for !window.ShouldClose() {
		glfw.PollEvents()
	}
}
