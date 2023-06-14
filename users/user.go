package users

import (
	"fmt"
	"github.com/elcorderista/godesdecero/modelos"
	"time"
)

func AltaUsuario() {
	//Create a new object type User
	u := new(modelos.User)
	u.AddUser(10, "Camila", time.Now(), true)
	fmt.Println(u)
}
