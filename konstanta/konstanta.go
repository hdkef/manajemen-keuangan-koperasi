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
var loanreq = "loan-req"
var meminspect = "mem-inspect"

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
var loanreqroute = createRoute(loanreq)
var meminspectroute = createRoute(meminspect)

var CookiesBearer = "bearer"

var Claims = "Claims"
var RoleADMINInput = "Admin-Input"
var RoleADMINSuper = "Admin-Super"
var RoleMEMBER = "member"

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

func (r *htmlFileName) LoanReq() string {
	return createHTMLFilename(loanreq)
}

func (r *htmlFileName) MemInspect() string {
	return createHTMLFilename(meminspect)
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

func (r *route) LoanReq() string {
	return loanreqroute
}

func (r *route) MemInspect() string {
	return meminspectroute
}
