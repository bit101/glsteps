package step02

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

// Main is minimal, but at least tries to set up and shut down responsibly.
func Main() {
	// 1. Ensure we run in single threaded mode here.
	runtime.LockOSThread()

	// 2. Initialize GLFW.
	if err := glfw.Init(); err != nil {
		log.Fatal(err)
	}

	// 3. Create a window.
	window, err := glfw.CreateWindow(400, 400, "Step 02", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 4. Event loop runs until the window is closed.
	for !window.ShouldClose() {
		glfw.PollEvents()
	}

	// 5. Clean up.
	glfw.Terminate()
}
