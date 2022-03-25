package main
import "fmt"
import "github.com/gin-gonic/gin"

type Person struct{
   Name string  `form:"field_a"`
   Email string `form:"field_b"`
}


type MyPerson struct{
   Name string  `uri:"name" binding:"required"`
   Email string `uri:"email" binding:"required"`
}

func getPerson(c *gin.Context){
   var p Person
   c.Bind(&p)
   c.JSON(200,gin.H{"name":p.Name,"email":p.Email})
}

func main(){
  router:=gin.Default()
  router.GET("/example",func(c *gin.Context){
      fmt.Println("Inside the GET request")
      c.JSON(200,gin.H{"message":"hello world"})
  })
  router.GET("/test",func(c *gin.Context){
    data:=map[string]interface{}{"language":"Hindi","language2":"English"}
    c.JSON(200,data)
  })
  router.GET("/person",getPerson)
  router.GET("/:name/:email",func(c *gin.Context){
    var person MyPerson
    if err:=c.ShouldBindUri(&person);err!=nil{
       c.JSON(400,gin.H{"msg":err})
       return
    }
    c.JSON(200,gin.H{"name":person.Name,"email":person.Email})
  })
  router.Run()
}

