function DelCacheMember(UID){
    fetch(`/del-cache-member/${UID}`)
    .then(response => response.json())
    .then(data => {
        if (data == "OK"){
            location.reload()
        }
    });
}