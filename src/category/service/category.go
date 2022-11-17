package service

type CategoryService interface {
	GetCategory() ([]GetCategoryResponse, error)
}

type (
	GetCategoryResponse struct {
		CategoryID            int    `json:"categoryId"`
		CategoryNameEn        string `json:"categoryNameEn"`
		CategoryNameTh        string `json:"categoryNameTh"`
		Type                  int    `json:"type"`
		CategoryType          string `json:"categoryType"`
		DisplayChannel        string `json:"displayChannel"`
		CategoryDescEn        string `json:"categoryDescEn"`
		CategoryDescTh        string `json:"categoryDescTh"`
		CategoryIcon          string `json:"categoryIcon"`
		CategoryImage         string `json:"categoryImage"`
		CategoryURL           string `json:"categoryUrl"`
		MobileIconActive      string `json:"mobileIconActive"`
		MobileIconInactive    string `json:"mobileIconInactive"`
		SerenadeIconActive    string `json:"serenadeIconActive"`
		SerenadeIconInactive  string `json:"serenadeIconInactive"`
		PrivilegeIconActive   string `json:"privilegeIconActive"`
		PrivilegeIconInactive string `json:"privilegeIconInactive"`
		CategoryMsgEn         string `json:"categoryMsgEn"`
		CategoryMsgTh         string `json:"categoryMsgTh"`
	}
)
