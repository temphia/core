package app

import (
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

func (s *Server) buildRoutes(e *gin.Engine) {

	pp.Println("Building routes")

	e.GET("/", s.routes.RootIndex)
	e.GET("/temphia-start", s.routes.StartPage)
	e.GET("/console", s.routes.AdminRoot2)
	e.GET("/auth", s.routes.AuthIndex2)
	s.operatorAPI(e.Group("/operator"))

	e.NoRoute(s.routes.NoRoute)

	{
		apiv1 := e.Group("/api/:tenant_id/v1/", s.CORS)
		s.adminTenantAPI(apiv1)
		s.authAPI(apiv1)
		s.authAPI2(apiv1)
		s.bprintAPI(apiv1)
		s.resourceAPI(apiv1)
		s.userAPI(apiv1)
		s.userSelfAPI(apiv1)
		s.plugAPI(apiv1)
		s.repoAPI(apiv1)
		s.cabinetAPI(apiv1)
		s.dtableAPI(apiv1)
		s.engineAPI(apiv1)

	}

}

func (s *Server) authAPI(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	auth.POST("/login", s.routes.Login)
	auth.POST("/signup", s.routes.SignUp)
	auth.POST("/refresh", s.routes.RefreshServiceToken)
	auth.POST("/refresh_from_pair_token", s.routes.RefreshFromPairToken)
}

func (s *Server) authAPI2(api *gin.RouterGroup) {
	//auth := api.Group("/auth2")
	/*

		/oauth_callback
		/load_methods
		/method_submit -> PreLoggedClaim
		/post_auth -> SessionClaim
		/refresh

	*/

}

func (s *Server) operatorAPI(op *gin.RouterGroup) {

	op.GET("/", s.routes.OperatorIndex)
	op.POST("/login", s.routes.OperatorLogin)
	op.GET("/stats", s.routes.OperatorStats)

	opTen := op.Group("/tenant")

	opTen.GET("/", s.routes.OperatorListTenant)
	opTen.POST("/", s.routes.OperatorAddTenant)
	opTen.PATCH("/", s.routes.OperatorUpdateTenant)
	opTen.DELETE("/", s.routes.OperatorDeleteTenant)
	opTen.POST("/token", s.routes.OperatorTenantToken)

}

// admin api

func (s *Server) adminTenantAPI(adminApi *gin.RouterGroup) {
	r := s.routes

	tenAPI := adminApi.Group("/tenant")
	tenAPI.POST("/", r.Authed(s.routes.UpdateTenant))

	dapi := tenAPI.Group("/domain")

	dapi.GET("/", r.Authed(s.routes.ListTenantDomain))
	dapi.POST("/", r.Authed(s.routes.AddTenantDomain))
	dapi.GET("/:id", r.Authed(s.routes.GetTenantDomain))
	dapi.POST("/:id", r.Authed(s.routes.UpdateTenantDomain))
	dapi.DELETE("/:id", r.Authed(s.routes.RemoveTenantDomain))

	dapi.GET("/:did/widget", r.Authed(s.routes.ListDomainWidget))
	dapi.POST("/:did/widget", r.Authed(s.routes.AddDomainWidget))
	dapi.GET("/:did/widget/:wid", r.Authed(s.routes.GetDomainWidget))
	dapi.POST("/:did/widget/:wid", r.Authed(s.routes.UpdateDomainWidget))
	dapi.DELETE("/:did/widget/:wid", r.Authed(s.routes.RemoveDomainWidget))
}

func (s *Server) bprintAPI(adminApi *gin.RouterGroup) {
	r := s.routes

	bAPI := adminApi.Group("/bprint")
	bAPI.GET("/", r.Authed(r.BprintList))
	bAPI.POST("/", r.Authed(r.BprintCreate))

	bAPI.GET("/:id", r.Authed(r.BprintGet))
	bAPI.POST("/:id", r.Authed(r.BprintUpdate))
	bAPI.DELETE("/:id", r.Authed(r.BprintRemove))
	bAPI.POST("/:id/install", r.Authed(r.BprintInstall))
	bAPI.GET("/:id/file", r.Authed(r.BprintListFiles))
	bAPI.GET("/:id/file/:file_id", r.Authed(r.BprintGetFile))

	bAPI.POST("/:id/file/:file_id", r.Authed(r.BprintNewBlob))
	bAPI.PATCH("/:id/file/:file_id", r.Authed(r.BprintUpdateBlob))

	bAPI.DELETE("/:id/file/:file_id", r.Authed(r.BprintDelFile))
	adminApi.POST("/import_bprint", r.Authed(r.BprintImport))
	bAPI.POST("/:id/push_token", r.Authed(r.BprintPushToken))

}

func (s *Server) resourceAPI(adminApi *gin.RouterGroup) {
	r := s.routes

	bAPI := adminApi.Group("/resource")
	adminApi.POST("/agent_resources", r.Authed(r.ResourceAgentList))
	bAPI.GET("/", r.Authed(r.ResourceList))
	bAPI.POST("/", r.Authed(r.ResourceCreate))
	bAPI.GET("/:slug", r.Authed(r.ResourceGet))
	bAPI.POST("/:slug", r.Authed(r.ResourceUpdate))
	bAPI.DELETE("/:slug", r.Authed(r.ResourceRemove))

}

func (s *Server) userSelfAPI(rg *gin.RouterGroup) {
	r := s.routes
	usAPI := rg.Group("/self")

	usAPI.POST("/message_user", r.Authed(r.SelfMessageUser))
	usAPI.GET("/get_user_info/:user", r.Authed(r.SelfGetUserInfo))
	usAPI.GET("/get_self_info", r.Authed(r.SelfGetInfo))

	usAPI.POST("/change_email", r.Authed(r.SelfChangeEmail))
	usAPI.POST("/change_auth", r.Authed(r.SelfChangeAuth))
	usAPI.POST("/list_messages", r.Authed(r.SelfListMessages))
	usAPI.POST("/modify_messages", r.Authed(r.SelfModifyMessages))
	usAPI.POST("/update_sockd_tags", r.Authed(r.UserSocketUpdate))
	usAPI.GET("/user_ws", r.SelfUserSocket)
	usAPI.POST("/dtable_change", r.Authed(r.SockdDgroupChange))

}

func (s *Server) userAPI(rg *gin.RouterGroup) {
	r := s.routes
	user := rg.Group("/user")

	user.GET("/", r.Authed(r.ListUsers))
	user.POST("/", r.Authed(r.AddUser))
	user.GET("/:user_id", r.Authed(r.GetUserByID))
	user.POST("/:user_id", r.Authed(r.UpdateUser))
	user.DELETE("/:user_id", r.Authed(r.RemoveUser))

	userGroup := rg.Group("/user_group")

	userGroup.GET("/", r.Authed(r.ListUserGroup))
	userGroup.POST("/", r.Authed(r.AddUserGroup))
	userGroup.GET("/:user_group", r.Authed(r.GetUserGroup))
	userGroup.POST("/:user_group", r.Authed(r.UpdateUserGroup))
	userGroup.DELETE("/:user_group", r.Authed(r.RemoveUserGroup))

	auth := rg.Group("/user_auth/:ugroup")
	auth.GET("/", r.Authed(r.ListUserGroupAuth))
	auth.POST("/", r.Authed(r.AddUserGroupAuth))
	auth.GET("/:id", r.Authed(r.GetUserGroupAuth))
	auth.POST("/:id", r.Authed(r.GetUserGroupAuth))
	auth.DELETE("/:id", r.Authed(r.RemoveUserGroupAuth))

	hook := rg.Group("/user_hook/:ugroup")
	hook.GET("/", r.Authed(r.ListUserGroupHook))
	hook.POST("/", r.Authed(r.AddUserGroupHook))
	hook.GET("/:id", r.Authed(r.GetUserGroupHook))
	hook.POST("/:id", r.Authed(r.UpdateUserGroupHook))
	hook.DELETE("/:id", r.Authed(r.RemoveUserGroupHook))

	plug := rg.Group("/user_plug/:ugroup")
	plug.GET("/", r.Authed(r.ListUserGroupPlug))
	plug.POST("/", r.Authed(r.AddUserGroupPlug))
	plug.GET("/:id", r.Authed(r.GetUserGroupPlug))
	plug.POST("/:id", r.Authed(r.UpdateUserGroupPlug))
	plug.DELETE("/:id", r.Authed(r.RemoveUserGroupPlug))

	data := rg.Group("/user_data/:ugroup")
	data.GET("/", r.Authed(r.ListUserGroupData))
	data.POST("/", r.Authed(r.AddUserGroupData))
	data.GET("/:id", r.Authed(r.GetUserGroupData))
	data.POST("/:id", r.Authed(r.UpdateUserGroupData))
	data.DELETE("/:id", r.Authed(r.RemoveUserGroupData))

	perm := rg.Group("/perm")
	perm.GET("/", r.Authed(r.ListAllPerm))
	perm.POST("/", r.Authed(r.AddPerm))
	perm.GET("/:perm", r.Authed(r.GetPerm))
	perm.POST("/:perm", r.Authed(r.UpdatePerm))
	perm.DELETE("/:perm", r.Authed(r.RemovePerm))

	role := rg.Group("/role")
	role.GET("/", r.Authed(r.ListAllRole))
	role.POST("/", r.Authed(r.AddRole))
	role.GET("/:role", r.Authed(r.GetRole))
	role.POST("/:role", r.Authed(r.UpdateRole))
	role.DELETE("/:role", r.Authed(r.RemoveRole))

	userRole := rg.Group("/user_role")
	userRole.GET("/", r.Authed(r.ListAllUserRole))
	userRole.POST("/", r.Authed(r.AddUserRole))
	userRole.DELETE("/", r.Authed(r.RemoveUserRole))

	rg.GET("/user_perm", r.Authed(r.ListUserPerm)) // user query
	rg.GET("/user_profile_image/:user_id", r.UserProfileImage)

}

func (s *Server) plugAPI(rg *gin.RouterGroup) {
	r := s.routes
	plug := rg.Group("/plug")

	plug.GET("/", r.Authed(r.ListPlug))
	plug.POST("/", r.Authed(r.NewPlug))
	plug.GET("/:plug_id", r.Authed(r.GetPlug))
	plug.POST("/:plug_id", r.Authed(r.UpdatePlug))
	plug.DELETE("/:plug_id", r.Authed(r.DelPlug))

	plug.GET("/:plug_id/agent/", r.Authed(r.ListAgent))
	plug.POST("/:plug_id/agent/", r.Authed(r.NewAgent))
	plug.GET("/:plug_id/agent/:agent_id", r.Authed(r.GetAgent))
	plug.POST("/:plug_id/agent/:agent_id", r.Authed(r.UpdateAgent))
	plug.DELETE("/:plug_id/agent/:agent_id", r.Authed(r.DelAgent))
	plug.POST("/:plug_id/agent/:agent_id/pair_token", r.Authed(r.PairAgentToken))

	rg.GET("plug_icon/:plug_id", r.PlugIcon)

}

// admin api end

func (s *Server) repoAPI(apiv1 *gin.RouterGroup) {
	repoApi := apiv1.Group("/repo")
	r := s.routes
	repoApi.GET("/", r.Authed(r.RepoSources))
	repoApi.GET("/:repo", r.Authed(s.routes.RepoList))
	repoApi.GET("/:repo/:group_id/:slug", r.Authed(r.RepoGet))
	repoApi.GET("/:repo/:group_id/:slug/:file", r.Authed(r.RepoGetFile))
}

func (s *Server) devAPI(apiv1 *gin.RouterGroup) {
	devApi := apiv1.Group("/dev")
	r := s.routes

	devApi.POST("/extract_js_methods", r.Authed(nil))
	devApi.POST("/check/lang/yaml", r.Authed(nil))

}

func (s *Server) cabinetAPI(apiv1 *gin.RouterGroup) {
	r := s.routes
	blobapi := apiv1.Group("/cabinet")

	apiv1.GET("/cabinet_sources", r.Authed(r.ListCabinetSources))
	blobapi.GET("/", r.Authed(r.ListRootFolder))
	blobapi.GET("/:folder", r.Authed(r.ListFolder))
	blobapi.POST("/:folder", r.Authed(r.NewFolder))
	blobapi.GET("/:folder/file/:fname", r.Authed(r.GetFile))
	blobapi.POST("/:folder/file/:fname", r.Authed(r.UploadFile))
	blobapi.DELETE("/:folder/file/:fname", r.Authed(r.DeleteFile))

	blobapi.GET("/:folder/preview/:fname", r.Authed(r.GetFilePreview))
	blobapi.POST("/:folder/ticket", r.Authed(r.GetFolderTicket))

	partCab := apiv1.Group("/ticket_cabinet/:ticket")
	partCab.GET("/", r.TicketCabinetList)
	partCab.GET("/:file", r.TicketCabinetFile)
	partCab.GET("/preview/:file", r.TicketCabinetPreviewFile)
	partCab.POST("/:file", r.TicketCabinetUpload)

}

func (s *Server) dtableAPI(apiv1 *gin.RouterGroup) {
	r := s.routes
	dgroup := apiv1.Group("/dgroup")
	dgroup.GET("/", r.Authed(r.ListDtableSources))
	dgroup.GET("/:source", r.Authed(r.ListGroup))
	dgroup.POST("/:source/", r.Authed(r.NewGroup))
	dgroup.PATCH("/:source/:group_id", r.Authed(r.EditGroup))
	dgroup.GET("/:source/:group_id", r.Authed(r.GetGroup))
	dgroup.DELETE("/:source/:group_id", r.Authed(r.DeleteGroup))

	dtable := apiv1.Group("/dtable")

	apiv1.GET("/dgroup_load", r.Authed(r.LoadGroup))

	dtable.GET("/", r.Authed(r.ListTables))
	dtable.POST("/", r.Authed(r.AddTable))
	dtable.GET("/:table_id", r.Authed(r.GetTable))
	dtable.PATCH("/:table_id", r.Authed(r.EditTable))
	dtable.DELETE("/:table_id", r.Authed(r.DeleteTable))
	dtable.GET("/:table_id/column", r.Authed(r.ListColumns))
	dtable.POST("/:table_id/column", r.Authed(r.AddColumn)) // fixme
	dtable.PATCH("/:table_id/column/:column_id", r.Authed(r.EditColumn))
	dtable.GET("/:table_id/column/:column_id", r.Authed(r.GetColumn))
	dtable.DELETE("/:table_id/column/:column_id", r.Authed(r.DeleteColumn))

	dtable.GET("/:table_id/view", r.Authed(r.ListView))
	dtable.POST("/:table_id/view", r.Authed(r.NewView))
	dtable.POST("/:table_id/view/:id", r.Authed(r.ModifyView))
	dtable.GET("/:table_id/view/:id", r.Authed(r.GetView))
	dtable.DELETE("/:table_id/view/:id", r.Authed(r.DelView))

	dtable.GET("/:table_id/hook", r.Authed(r.ListHook))
	dtable.POST("/:table_id/hook", r.Authed(r.NewHook))
	dtable.POST("/:table_id/hook/:id", r.Authed(r.ModifyHook))
	dtable.GET("/:table_id/hook/:id", r.Authed(r.GetHook))
	dtable.DELETE("/:table_id/hook/:id", r.Authed(r.DelHook))

	dops := apiv1.Group("/dtable_ops")

	dops.POST("/:table_id/row", r.Authed(r.NewRow))
	dops.GET("/:table_id/row/:id", r.Authed(r.GetRow))
	dops.POST("/:table_id/row/:id", r.Authed(r.UpdateRow))
	dops.DELETE("/:table_id/row/:id", r.Authed(r.DeleteRow))
	dops.POST("/:table_id/simple_query", r.Authed(r.SimpleQuery))
	dops.POST("/:table_id/fts_query", r.Authed(r.FTSQuery)) // fixme => remove this and consolidate this to simple_query ?
	dops.POST("/:table_id/ref_load", r.Authed(r.RefLoad))
	dops.POST("/:table_id/ref_resolve", r.Authed(r.RefResolve))
	dops.POST("/:table_id/rev_ref_load", r.Authed(r.ReverseRefLoad))

	dops.GET("/:table_id/activity/:row_id", r.Authed(r.ListActivity))
	dops.POST("/:table_id/activity/:row_id", r.Authed(r.CommentRow))

}

func (s *Server) engineAPI(apiv1 *gin.RouterGroup) {
	mux := apiv1.Group("/engine/:plug_id/:agent_id")

	// fixme => redo the APIs

	// mux.Any("/exec_raw/:action", s.routes.EngineExecConsole)

	mux.Any("/exec_con/:action", s.routes.EngineExecConsole)
	mux.GET("/launcher/html", s.routes.EngineLaunchExecHTML)
	mux.POST("/launcher/json", s.routes.EngineLaunchExec)
	mux.GET("/referer_ticket", nil) // fixme => generate referer ticket
	mux.GET("/serve/:file", s.routes.EngineServe)
	mux.GET("/executor/:loader/loader.js", s.routes.EngineExecLoaderScript)
	mux.GET("/executor/:loader/loader.css", s.routes.EngineExecLoaderStyle)
	mux.GET("/exec_ws", s.routes.PlugSocket)
}

const (
	AllowedHeader = "Content-Type, Content-Length, Accept-Encoding, accept, origin, Cache-Control," +
		"X-Requested-With, X-CSRF-Token, Authorization, X-Request-HMAC"
)

func (a *Server) CORS(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", AllowedHeader)
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()

}
