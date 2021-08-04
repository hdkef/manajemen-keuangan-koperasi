let statementSelect = document.getElementById("statement")
let typeSelect =  document.getElementById("type")
var curVal = "Balance Sheet"
let balancesheetType = ["Assets","Liabilites","Equity"]
let incomestatementType = ["Revenues","Expenses"]

statementSelect.onchange = () => {
    populateSelect(statementSelect,curVal,typeSelect)
}

populateSelect = (fromSelect,currentValue,targetSelect)=>{
    let val = fromSelect.value
    if (val != currentValue && val == "Balance Sheet") {
        deleteChildren(targetSelect)
        for (let i of balancesheetType) {
            var added = document.createElement('option');
            added.value = i;
            added.innerHTML = i;
            targetSelect.append(added);
        }
    }
    if (val != currentValue && val == "Income Statement") {
        deleteChildren(targetSelect)
        for (let i of incomestatementType) {
           var added = document.createElement('option');
           added.value = i;
           added.innerHTML = i;
           targetSelect.append(added);
        }
   }
   curVal = val
}

deleteChildren = (parent)=>{
    while (parent.firstChild) {
        parent.removeChild(parent.firstChild);
    }
}