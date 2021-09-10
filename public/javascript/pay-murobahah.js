var id = null

function createDialog(id_){
    var divFrm = document.getElementById("frm")
    if (!id){
        id = id_
        divFrm.innerHTML = `
       <div>
       <p>Cicil Murobahah ID : ${id}</p>
       <form method="POST" action="pay-murobahah">
       <input type="hidden" value="${id_}" name="ID">
       <input type="text" placeholder="jumlah (Rp)" name="Amount" required>
       <textarea name="Info" placeholder="Info"></textarea>
       <button type="submit">Submit</button>
       </form>
       </div>
        `
    }else{
        divFrm.innerHTML = null
        id = null
    }
}