function validateLogin(e) {
    e.preventDefault()
    let loginForm = document.forms["loginForm"]
    let username = loginForm["Username"].value
    let pass = loginForm["Pass"].value
    if (username || pass){
        return true
    }
    alert("Please fill the required field")
    return false
}