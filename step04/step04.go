package step04

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 400
	height = 400
)

// Main adds input processing - close window by hitting escape.
func Main() {
	runtime.LockOSThread()

	initGlfw()
	window := createWindow("Step 04")

	// defer cleanup
	defer glfw.Terminate()

	appLoop(window)
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

		// process input events
		processInput(window)

		glfw.PollEvents()
	}
}

// if escape key pressed, set window should close
func processInput(window *glfw.Window) {
	if glfw.GetCurrentContext().GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}
