package form

import (
    "github.com/robfig/revel"
)

type roomForm struct {
    Name string
    Desc string
}

func (rf *roomForm) Validation(v *revel.Validation) {
    v.Required(rf.Name)
    v.Required(rf.Desc)
}
