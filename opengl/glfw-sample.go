package main

import (
       "log"
       "github.com/jteeuwen/glfw"
)

func main() {
	err := glfw.Init()
	if err != nil {
	    log.Fatal("Failed to init GLFW: " + err.Error())
	}
	
	err = glfw.OpenWindow(1024, 768, 0,0,0,0, 32,0, glfw.Windowed)
	if err != nil {
	    log.Fatal("Failed to open GLFW window: " + err.Error())
	}
	
	if gl.Init() != 0 {
	    log.Fatal("Failed to init GL")
	}
	
	gl.ClearColor(0.0, 0.0, 0.3, 0.0)
	
	// create vertexbuffer
	gVertexBufferData := []float32{-1.0,-1.0,0.0, 1.0,-1.0,0.0, 0.0,1.0,0.0}
	vertexBuffer := gl.GenBuffer()
	vertexBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(gVertexBufferData), gVertexBufferData, gl.STATIC_DRAW)
	
	for {
	    // clear screen
	    gl.Clear(gl.COLOR_BUFFER_BIT)
	
	    // first attribute buffer: vertices
	    var vertexAttrib gl.AttribLocation = 0
	    vertexAttrib.EnableArray()
	    vertexBuffer.Bind(gl.ARRAY_BUFFER)
	    var f float32 = 0.0
	    vertexAttrib.AttribPointer(
	        3,     // size
	        false, // normalized?
	        0,     // stride
	        &f) // array buffer offset
	
	    // draw the triangle
	    gl.DrawArrays(gl.TRIANGLES, 0, 3)
	
	    vertexAttrib.DisableArray()
	
	    glfw.SwapBuffers()
	}
}