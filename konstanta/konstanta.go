package konstanta

import "fmt"

//file  name
var html = "html"
var home = "home"
var newCOA = "new-coa"
var newTransaction = "new-transaction"
var editCOA = "edit-coa"
var editTransaction = "edit-transaction"
var summary = "summary"
var fullJournal = "full-journal"
var manageUser = "manage-user"

//route path
var homeroute = "/"
var newcoaroute = fmt.Sprintf("/%s", newCOA)
var newtransactionroute = fmt.Sprintf("/%s", newTransaction)
var editcoaroute = fmt.Sprintf("/%s", editCOA)
var edittransactionroute = fmt.Sprintf("/%s", editTransaction)
var manageuserroute = fmt.Sprintf("/%s", manageUser)
var summaryroute = fmt.Sprintf("/%s", summary)
var fulljournalroute = fmt.Sprintf("/%s", fullJournal)

type htmlFileName struct {
}

func GetHTMLFileName() *htmlFileName {
	return &htmlFileName{}
}

func (r *htmlFileName) Home() string {
	return fmt.Sprintf("%s.%s", home, html)
}
func (r *htmlFileName) NewCOA() string {
	return fmt.Sprintf("%s.%s", newCOA, html)
}
func (r *htmlFileName) NewTransaction() string {
	return fmt.Sprintf("%s.%s", newTransaction, html)
}
func (r *htmlFileName) EditCOA() string {
	return fmt.Sprintf("%s.%s", editCOA, html)
}
func (r *htmlFileName) EditTransaction() string {
	return fmt.Sprintf("%s.%s", editTransaction, html)
}
func (r *htmlFileName) Summary() string {
	return fmt.Sprintf("%s.%s", summary, html)
}
func (r *htmlFileName) FullJournal() string {
	return fmt.Sprintf("%s.%s", fullJournal, html)
}
func (r *htmlFileName) ManageUser() string {
	return fmt.Sprintf("%s.%s", manageUser, html)
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
func (r *route) FullJournal() string {
	return fulljournalroute
}
func (r *route) ManageUser() string {
	return manageuserroute
}
