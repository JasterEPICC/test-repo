package service

import (
	"fmt"
	"privilege-api-myais/lib"
	"privilege-api-myais/src/category/repository"
)

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(connect repository.CategoryRepository) CategoryService {
	return categoryService{connect}
}

func (s categoryService) GetCategory() (res []GetCategoryResponse, errCategory error) {
	category, errCategory := s.categoryRepository.GetCategory()
	if errCategory != nil {
		lib.LogInfoSQL(fmt.Sprintf("GetCategory: %v", errCategory))
		return nil, lib.NewError(50001, "user not found")
	}

	for _, detail := range category {
		res = append(res, GetCategoryResponse{
			CategoryID:            int(detail.CategoryID.Int64),
			CategoryNameEn:        detail.Category.String,
			CategoryNameTh:        detail.CategoryTh.String,
			Type:                  int(detail.Type.Int64),
			CategoryType:          detail.CategoryType.String,
			DisplayChannel:        detail.DisplayChannel.String,
			CategoryDescEn:        detail.CategoryEnDesc.String,
			CategoryDescTh:        detail.CategoryThDesc.String,
			CategoryIcon:          detail.CategoryIcon.String,
			CategoryImage:         detail.CategoryImage.String,
			CategoryURL:           detail.CategoryURL.String,
			MobileIconActive:      detail.MobileIconActive.String,
			MobileIconInactive:    detail.MobileIconInactive.String,
			SerenadeIconActive:    detail.SerenadeIconActive.String,
			SerenadeIconInactive:  detail.SerenadeIconInactive.String,
			PrivilegeIconActive:   detail.PrivilegeIconActive.String,
			PrivilegeIconInactive: detail.PrivilegeIconInactive.String,
			CategoryMsgEn:         detail.CategoryMsgEn.String,
			CategoryMsgTh:         detail.CategoryMsgTh.String,
		})
	}

	return res, nil
}
