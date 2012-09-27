package main

import (
	"flag"
	"reflect"
	"github.com/robfig/revel"
	
  "github.com/jamesward/hellorevel/app/controllers"
  
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	rev.INFO.Println("Running revel server")
	flag.Parse()
	rev.Init(*runMode, *importPath, *srcPath)
	
	rev.RegisterController((*controllers.Application)(nil),
		[]*rev.MethodType{
			&rev.MethodType{
				Name: "Index",
				Args: []*rev.MethodArg{ 
			  },
				RenderArgNames: map[int][]string{ 
					13: []string{ 
						"message",
					},
				},
			},
			
		})
	
	rev.Run(*port)
}
