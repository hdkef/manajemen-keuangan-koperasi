var xhttp = new XMLHttpRequest();
xhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {
    document.getElementById("mynav").innerHTML =
    this.responseText;
    }
};
xhttp.open("GET", "public/ajax/nav.html", true);
xhttp.send();