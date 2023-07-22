package main

import (
	"github.com/gin-gonic/gin"
)

type structA struct {
	FieldA string `form:"field_a"`
}

type structB struct {
	NestedStruct structA
	FieldB       string `form:"field_b"`
}

type structC struct {
	NestedStructPointer *structA
	FieldC              string `form:"field_c"`
}

type structD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func main() {
	r := gin.Default()

	r.GET("/b", func(ctx *gin.Context) {
		b := structB{}
		ctx.Bind(&b)
		ctx.JSON(200, gin.H{
			"a": b.NestedStruct.FieldA,
			"b": b.FieldB,
		})
	})

	r.GET("/c", func(ctx *gin.Context) {
		c := structC{}
		ctx.Bind(&c)
		ctx.JSON(200, gin.H{
			"a": c.NestedStructPointer,
			"c": c.FieldC,
		})
	})

	r.GET("/d", func(ctx *gin.Context) {
		d := structD{}
		ctx.Bind(&d)
		ctx.JSON(200, gin.H{
			"d": d.FieldD,
			"x": d.NestedAnonyStruct,
		})
	})

	r.Run()
}
