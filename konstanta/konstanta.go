package konstanta

import "fmt"

//name
var html = "html"
var newCOA = "new-coa"
var newTransaction = "new-transaction"
var editCOA = "edit-coa"
var editTransaction = "edit-transaction"
var summary = "summary"
var fullReport = "full-report"
var manageUser = "manage-user"
var admin = "admin"
var errors = "error"
var success = "success"
var login = "login"
var member = "member"
var memrequest = "mem-request"
var meminspect = "mem-inspect"
var depositreq = "deposit-req"
var withdrawreq = "withdraw-req"
var finduser = "find-user"
var meminspectres = "mem-inspect-res"
var register = "register"
var usersetting = "user-setting"
var adminput = "adm-input"
var murobahahreq = "murobahah-req"

func createRoute(route string) string {
	return fmt.Sprintf("/%s", route)
}

//route path
var homeroute = createRoute("")
var newcoaroute = createRoute(newCOA)
var newtransactionroute = createRoute(newTransaction)
var editcoaroute = createRoute(editCOA)
var edittransactionroute = createRoute(editTransaction)
var manageuserroute = createRoute(manageUser)
var summaryroute = createRoute(summary)
var fullReportroute = createRoute(fullReport)
var adminroute = createRoute(admin)
var loginroute = createRoute(login)
var memberroute = createRoute(member)
var memrequestroute = createRoute(memrequest)
var meminspectroute = createRoute(meminspect)
var depositreqroute = createRoute(depositreq)
var withdrawreqroute = createRoute(withdrawreq)
var finduserroute = createRoute(finduser)
var registerroute = createRoute(register)
var usersettingroute = createRoute(usersetting)
var memrequestaccroute = createRoute("mem-request-acc")
var memrequestdecroute = createRoute("mem-request-dec")
var adminputroute = createRoute(adminput)
var murobahahreqroute = createRoute(murobahahreq)
var murobahahaccroute = createRoute("murobahah-acc")
var murobahahdecroute = createRoute("murobahah-dec")
var delcachememberroute = createRoute("del-cache-member/:UID")

var CookiesBearer = "bearer"

var Claims = "Claims"
var RoleADMINInput = "Admin-Input"
var RoleADMINSuper = "Admin-Super"
var RoleMEMBER = "member"

var QueryID = "ID"
var QueryUsername = "Username"
var QueryPass = "Pass"
var QueryRole = "Role"
var QueryMemID = "MemID"
var QueryFilter = "Filter"
var QueryKey = "Key"
var QueryType = "Type"
var QueryAmount = "Amount"
var QueryInfo = "Info"
var QueryDate = "Date"
var QueryUID = "UID"
var QueryIsAgent = "IsAgent"
var QueryDoc = "Doc"
var QueryUIDBuyer = "UIDBuyer"
var QueryUIDAgent = "UIDAgent"
var QueryDueDate = "DueDate"

var TypeIP = "IP"
var TypeIW = "IW"
var TypeSSPos = "SS+"
var TypeSSNeg = "SS-"
var TypeSHU = "SHU"
var TypeBonus = "Bonus"
var TYPEMurobahahPos = "MRBH+"
var TYPEMurobahahNeg = "MRBH-"

var TABLEALLUSER = "alluser"
var TABLEMEMJOURNAL = "member_journal"
var TABLEMEMMUROBAHAH = "member_murobahah"
var TABLEMEMREQMUROBAHAH = "member_req_murobahah"
var TABLEMEMBALANCE = "member_balance"
var TABLEMEMREQ = "member_req"
var TABLEMEMBALANCEHISTORY = "member_balance_history"
var TABLEALLINFO = "allinfo"
var TABLEAGENTHISTORY = "agent_history"

type htmlFileName struct {
}

func GetHTMLFileName() *htmlFileName {
	return &htmlFileName{}
}

func createHTMLFilename(fname string) string {
	return fmt.Sprintf("%s.%s", fname, html)
}

func (r *htmlFileName) NewCOA() string {
	return createHTMLFilename(newCOA)
}
func (r *htmlFileName) NewTransaction() string {
	return createHTMLFilename(newTransaction)
}
func (r *htmlFileName) EditCOA() string {
	return createHTMLFilename(editCOA)
}
func (r *htmlFileName) EditTransaction() string {
	return createHTMLFilename(editTransaction)
}
func (r *htmlFileName) Summary() string {
	return createHTMLFilename(summary)
}
func (r *htmlFileName) FullReport() string {
	return createHTMLFilename(fullReport)
}
func (r *htmlFileName) ManageUser() string {
	return createHTMLFilename(manageUser)
}
func (r *htmlFileName) Admin() string {
	return createHTMLFilename(admin)
}

func (r *htmlFileName) Error() string {
	return createHTMLFilename(errors)
}

func (r *htmlFileName) Success() string {
	return createHTMLFilename(success)
}

func (r *htmlFileName) Login() string {
	return createHTMLFilename(login)
}

func (r *htmlFileName) Member() string {
	return createHTMLFilename(member)
}

func (r *htmlFileName) MemRequest() string {
	return createHTMLFilename(memrequest)
}

func (r *htmlFileName) MemInspect() string {
	return createHTMLFilename(meminspect)
}

func (r *htmlFileName) DepositReq() string {
	return createHTMLFilename(depositreq)
}

func (r *htmlFileName) WithdrawReq() string {
	return createHTMLFilename(withdrawreq)
}

func (r *htmlFileName) FindUser() string {
	return createHTMLFilename(finduser)
}

func (r *htmlFileName) MemInspectRes() string {
	return createHTMLFilename(meminspectres)
}

func (r *htmlFileName) Register() string {
	return createHTMLFilename(register)
}

func (r *htmlFileName) UserSetting() string {
	return createHTMLFilename(usersetting)
}

func (r *htmlFileName) AdmInput() string {
	return createHTMLFilename(adminput)
}

func (r *htmlFileName) MurobahahReq() string {
	return createHTMLFilename(murobahahreq)
}

type route struct {
}

func GetRoute() *route {
	return &route{}
}

func (r *route) Home() string {
	return homeroute
}

func (r *route) EditCOA() string {
	return editcoaroute
}
func (r *route) EditTransaction() string {
	return edittransactionroute
}
func (r *route) NewCOA() string {
	return newcoaroute
}
func (r *route) NewTransaction() string {
	return newtransactionroute
}
func (r *route) Summary() string {
	return summaryroute
}
func (r *route) FullReport() string {
	return fullReportroute
}
func (r *route) ManageUser() string {
	return manageuserroute
}

func (r *route) Admin() string {
	return adminroute
}

func (r *route) Login() string {
	return loginroute
}

func (r *route) Member() string {
	return memberroute
}

func (r *route) MemRequest() string {
	return memrequestroute
}

func (r *route) MemInspect() string {
	return meminspectroute
}

func (r *route) DepositReq() string {
	return depositreqroute
}

func (r *route) WithdrawReq() string {
	return withdrawreqroute
}

func (r *route) FindUser() string {
	return finduserroute
}

func (r *route) Register() string {
	return registerroute
}

func (r *route) UserSetting() string {
	return usersettingroute
}

func (r *route) MemRequestAcc() string {
	return memrequestaccroute
}

func (r *route) MemRequestDec() string {
	return memrequestdecroute
}

func (r *route) AdmInput() string {
	return adminputroute
}

func (r *route) MurobahahReq() string {
	return murobahahreqroute
}

func (r *route) MurobahahAcc() string {
	return murobahahaccroute
}

func (r *route) MurobahahDec() string {
	return murobahahdecroute
}

func (r *route) DelCacheMember() string {
	return delcachememberroute
}
