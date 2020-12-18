package presure

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func TestSingleHandler(c *gin.Context) {
	fmt.Printf("presure :  %v \n", time.Now())
}




