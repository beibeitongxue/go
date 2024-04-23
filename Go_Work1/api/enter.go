package api

import (
	"Go_Work/api/node_api"
	"Go_Work/api/settings_api"
	"Go_Work/api/user_api"
)

//import "Go_Work/api/user_api"

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	NodeApi     node_api.NodeApi
	//ImagesApi   images_api.ImagesApi
	//AdvertApi   advert_api.AdvertApi
	//MenuApi     menu_api.MenuApi
	UserApi user_api.UserApi
	//TagApi      tag_api.TagApi
	//MessageApi  message_api.MessageApi
	//ArticleApi  article_api.ArticleApi
	//CommentApi  comment_api.CommentApi
	//NewsApi     new_api.NewApi
	//ChatApi     chat_api.ChatApi
	//LogApi      log_api.LogApi
	//DataApi     data_api.DataApi
}

var ApiGroupApp = new(ApiGroup)
