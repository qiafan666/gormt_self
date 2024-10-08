package gtools

import (
	"fmt"
	"os/exec"

	"github.com/xxjwxc/public/mylog"

	"github.com/xxjwxc/gormt/data/dlg"
	"github.com/xxjwxc/gormt/data/view/model"

	"github.com/xxjwxc/gormt/data/config"

	"github.com/xxjwxc/gormt/data/view/model/genmssql"
	"github.com/xxjwxc/gormt/data/view/model/genmysql"
	"github.com/xxjwxc/gormt/data/view/model/gensqlite"
	"github.com/xxjwxc/public/tools"
)

// Execute exe the cmd
func Execute() {
	if config.GetIsGUI() {
		dlg.WinMain()
	} else {
		showCmd()
	}
}

func showCmd() {
	// var tt oauth_db.UserInfoTbl
	// tt.Nickname = "ticket_001"
	// orm.Where("nickname = ?", "ticket_001").Find(&tt)
	// fmt.Println(tt)
	var modeldb model.IModel
	switch config.GetDbInfo().Type {
	case 0: // mysql
		modeldb = genmysql.GetModel()
	case 1: // sqllite
		modeldb = gensqlite.GetModel()
	case 2: //
		modeldb = genmssql.GetModel()
	}
	if modeldb == nil {
		mylog.Error(fmt.Errorf("modeldb not fund : please check db_info.type (0:mysql , 1:sqlite , 2:mssql) "))
		return
	}

	pkg := modeldb.GenModel()
	// gencnf.GenOutPut(&pkg)
	// just for test
	// out, _ := json.Marshal(pkg)
	// tools.WriteFile("test.txt", []string{string(out)}, true)

	list, _ := model.Generate(pkg)

	for _, v := range list {
		path := config.GetOutDir() + "/" + v.FileName
		tools.WriteFile(path, []string{v.FileCtx}, true)

		mylog.Info("fix structure fields for memory alignment")
		cmd, _ := exec.Command("fieldalignment", "-fix", path).Output()
		mylog.Info(string(cmd))

		mylog.Info("formatting differs from goimport's:")
		cmd, _ = exec.Command("goimports", "-l", "-w", path).Output()
		mylog.Info(string(cmd))

		mylog.Info("formatting differs from gofmt's:")
		cmd, _ = exec.Command("gofmt", "-l", "-w", path).Output()
		mylog.Info(string(cmd))
	}

	out, _ := model.GenerateReq(pkg)
	path := config.GetOutDir() + "/../gen/gen_pojo/gen_request/gen_req.go"
	tools.WriteFile(path, []string{out.FileCtx}, true)
	mylog.Info("fix structure fields for memory alignment")
	cmd, _ := exec.Command("fieldalignment", "-fix", path).Output()
	mylog.Info(string(cmd))

	mylog.Info("formatting differs from goimport's:")
	cmd, _ = exec.Command("goimports", "-l", "-w", path).Output()
	mylog.Info(string(cmd))

	mylog.Info("formatting differs from gofmt's:")
	cmd, _ = exec.Command("gofmt", "-l", "-w", path).Output()
	mylog.Info(string(cmd))

	out2, _ := model.GenerateResp(pkg)
	path2 := config.GetOutDir() + "/../gen/gen_pojo/gen_response/gen_resp.go"
	tools.WriteFile(path2, []string{out2.FileCtx}, true)
	mylog.Info("fix structure fields for memory alignment")
	cmd2, _ := exec.Command("fieldalignment", "-fix", path2).Output()
	mylog.Info(string(cmd2))

	mylog.Info("formatting differs from goimport's:")
	cmd2, _ = exec.Command("goimports", "-l", "-w", path2).Output()
	mylog.Info(string(cmd2))

	mylog.Info("formatting differs from gofmt's:")
	cmd2, _ = exec.Command("gofmt", "-l", "-w", path2).Output()
	mylog.Info(string(cmd2))

	out3, _ := model.GenerateServices(pkg)
	path3 := config.GetOutDir() + "/../gen/gen_services/gen_service.go"
	tools.WriteFile(path3, []string{out3.FileCtx}, true)
	cmd3, _ := exec.Command("fieldalignment", "-fix", path3).Output()
	mylog.Info(string(cmd3))

	cmd3, _ = exec.Command("goimports", "-l", "-w", path3).Output()
	mylog.Info(string(cmd3))

	cmd3, _ = exec.Command("gofmt", "-l", "-w", path3).Output()
	mylog.Info(string(cmd3))

	out4, _ := model.GenerateControllers(pkg)
	path4 := config.GetOutDir() + "/../gen/gen_controllers/gen_controller.go"
	tools.WriteFile(path4, []string{out4.FileCtx}, true)
	cmd4, _ := exec.Command("fieldalignment", "-fix", path4).Output()
	mylog.Info(string(cmd4))

	cmd4, _ = exec.Command("goimports", "-l", "-w", path4).Output()
	mylog.Info(string(cmd4))

	cmd4, _ = exec.Command("gofmt", "-l", "-w", path4).Output()
	mylog.Info(string(cmd4))

	out5, _ := model.GenerateReqCommon(pkg)
	path5 := config.GetOutDir() + "/../gen/gen_pojo/gen_request/gen_common.go"
	tools.WriteFile(path5, []string{out5.FileCtx}, true)
	cmd5, _ := exec.Command("fieldalignment", "-fix", path5).Output()
	mylog.Info(string(cmd5))

	cmd5, _ = exec.Command("goimports", "-l", "-w", path5).Output()
	mylog.Info(string(cmd5))

	cmd5, _ = exec.Command("gofmt", "-l", "-w", path5).Output()
	mylog.Info(string(cmd5))

	out6, _ := model.GenerateRespCommon(pkg)
	path6 := config.GetOutDir() + "/../gen/gen_pojo/gen_response/gen_common.go"
	tools.WriteFile(path6, []string{out6.FileCtx}, true)
	cmd6, _ := exec.Command("fieldalignment", "-fix", path6).Output()
	mylog.Info(string(cmd6))

	cmd6, _ = exec.Command("goimports", "-l", "-w", path6).Output()
	mylog.Info(string(cmd6))

	cmd6, _ = exec.Command("gofmt", "-l", "-w", path6).Output()
	mylog.Info(string(cmd6))
}
