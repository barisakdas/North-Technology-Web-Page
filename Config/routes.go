package Config

import (
	admin "NorthTechWebPage/Admin/controllers"
	"NorthTechWebPage/Site/controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Routes() *httprouter.Router {
	r:= httprouter.New()

	/******************************* <ADMIN> *******************************/
	// About Operations
	r.GET("/admin/about-us",admin.AboutController{}.Index)
	r.GET("/admin/edit-about/:id",admin.AboutController{}.EditAbout)
	r.POST("/admin/update-about/:id",admin.AboutController{}.UpdateAbout)

	// Contact Operations
	r.GET("/admin/contacts",admin.ContactController{}.Index)
	r.GET("/admin/reply-message/:id",admin.ContactController{}.ReplyMessage)
	r.GET("/admin/old-contacts",admin.ContactController{}.OldContacts)
	r.GET("/admin/unreply-message/:id",admin.ContactController{}.UnReplyMessage)

	// Menu Operations
	r.GET("/admin/site-menus", admin.MenuController{}.Index)
	r.GET("/admin/add-new-menu",admin.MenuController{}.AddNewMenu)
	r.POST("/admin/add-menu",admin.MenuController{}.AddMenu )
	r.GET("/admin/edit-menu/:id",admin.MenuController{}.UpdateMenuIndex)
	r.POST("/admin/update-menu/:id",admin.MenuController{}.UpdateMenu )
	r.GET("/admin/delete-menu/:id",admin.MenuController{}.DeleteMenu )

	// Article Operations
	r.GET("/admin/articles",admin.ArticleController{}.Index)
	r.GET("/admin/add-new-article",admin.ArticleController{}.AddNewArticle)
	r.POST("/admin/add-article",admin.ArticleController{}.AddArticle )
	r.GET("/admin/edit-article/:id",admin.ArticleController{}.UpdateArticleIndex)
	r.POST("/admin/update-article/:id",admin.ArticleController{}.UpdateArticle )
	r.GET("/admin/delete-article/:id",admin.ArticleController{}.DeleteArticle )

	// Category Operations
	r.GET("/admin/categories",admin.CategoryController{}.Index)
	r.GET("/admin/add-new-category",admin.CategoryController{}.AddNewCategory )
	r.POST("/admin/add-category",admin.CategoryController{}.AddCategory )
	r.GET("/admin/edit-category/:id",admin.CategoryController{}.UpdateCategoryIndex )
	r.POST("/admin/update-category/:id",admin.CategoryController{}.UpdateCategory )
	r.GET("/admin/delete-category/:id",admin.CategoryController{}.DeleteCategory )

	// Feature Operations
	r.GET("/admin/features",admin.FeatureController{}.Index)
	r.GET("/admin/add-new-feature",admin.FeatureController{}.AddNewFeature)
	r.POST("/admin/add-feature",admin.FeatureController{}.AddFeature )
	r.GET("/admin/edit-feature/:id",admin.FeatureController{}.UpdateFeatureIndex )
	r.POST("/admin/update-feature/:id",admin.FeatureController{}.UpdateFeature )
	r.GET("/admin/delete-feature/:id",admin.FeatureController{}.DeleteFeature )

	// Service Operations
	r.GET("/admin/services",admin.ServiceController{}.Index)
	r.GET("/admin/add-new-service",admin.ServiceController{}.AddNewService)
	r.POST("/admin/add-service",admin.ServiceController{}.AddService)
	r.GET("/admin/edit-service/:id",admin.ServiceController{}.UpdateServiceIndex )
	r.POST("/admin/update-service/:id",admin.ServiceController{}.UpdateService )
	r.GET("/admin/delete-service/:id",admin.ServiceController{}.DeleteService )

	// Login Operations
	r.GET("/admin/login",admin.LoginController{}.Index)
	r.POST("/admin/do_login",admin.LoginController{}.Login)
	r.GET("/admin/logout",admin.LoginController{}.Logout)

	/******************************* <SITE> *******************************/
	r.GET("/",controllers.HomeController{}.Index)
	r.GET("/index",controllers.HomeController{}.Index)
	r.GET("/about",controllers.AboutController{}.Index)
	r.GET("/contact",controllers.ContactController{}.Index)
	r.POST("/contact_us", controllers.ContactController{}.ContactUs)
	r.GET("/service",controllers.ServiceController{}.Index)
	r.GET("/blog", controllers.BlogController{}.Index)
	r.GET("/blog/:Slug",controllers.BlogController{}.Detail)
	r.GET("/category/:CategoryID",controllers.BlogController{}.ArticleByCategory)



	r.ServeFiles("/Site/assets/*filepath",http.Dir("Site/assets"))
	r.ServeFiles("/Admin/assets/*filepath",http.Dir("Admin/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))

	return r
}
