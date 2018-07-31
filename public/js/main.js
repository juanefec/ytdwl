let dl;
function download() {
    let id = document.getElementById("idinput").value
    $.ajax({
        url: "http://localhost:8080/api/getURL?id=" + id
    }).then(function (resp) {
            const link = document.createElement('a');
            link.href = "http://localhost:8080/api/video/" + resp;
            link.setAttribute('download', resp);
            document.body.appendChild(link);
            link.click();
        
    })
}