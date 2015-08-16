var xhr = new XMLHttpRequest();

function ajax(o) {
    xhr.onreadystatechange = function() {
        if (xhr.readyState == XMLHttpRequest.DONE) {
            if(xhr.status == 200){
                o.success();
            } else {
                o.error();
            }
        }
    }

    xhr.open(o.method, o.url);
    xhr.send(o.data);
}