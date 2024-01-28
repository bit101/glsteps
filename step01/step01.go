package step01

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

// Main is the simplest possible go-gl OpenGL program.
func Main() {
	// 1. Initialize GLFW.
	if err := glfw.Init(); err != nil {
		log.Fatal(err)
	}

	// 2. Create a window.
	window, err := glfw.CreateWindow(400, 400, "Step 01", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Event loop runs until the window is closed.
	for !window.ShouldClose() {
		glfw.PollEvents()
	}
}
