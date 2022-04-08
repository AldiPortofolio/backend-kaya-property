package tag

import (
	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/utils/constants"
	
)

func (svc *TagService) Filter(req request.TagBlog, res *models.Response) {
	data, err := svc.tagRepo.Filter(req)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	res.Data = data
	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "success"
	res.Data = data
}
