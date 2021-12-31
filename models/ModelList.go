package models

import "github.com/Amazeful/Amazeful-Backend/util"

type ModelList struct {
	List []util.BaseModel

	R util.Repository
}
