function getConnects(){
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "http://localhost:2200/get/connects", true);
    // xhr.setRequestHeader("Content-Type","application/json");
    // xhr.setRequestHeader("Access-Control-Allow-Origin","*");
    // xhr.setRequestHeader("Access-Control-Allow-Headers", "Content-Type")
    xhr.onreadystatechange = function(){
        var XMLHttpReq = xhr;
        /**
            XMLHttpReq.readyState
         0: 请求未初始化
         1: 服务器连接已建立
         2: 请求已接收
         3: 请求处理中
         4: 请求已完成，且响应已就绪
        **/
        if (XMLHttpReq.readyState == 4) {
            if (XMLHttpReq.status == 200) {                
                var data = XMLHttpReq.responseText;
                // alert(data)
                console.log(data);
                var json = JSON.parse(data);
                for(var i in json) {
                    console.log(json[i].name)
                    var dl = document.getElementById("db_list_dl")
                    var dd = document.createElement("dt")
                    // var ddv = document.createElement("p")
                    // ddv.innerText = 0
                    // ddv.setAttribute("hidden",true);

                    dd.type = "db_item"
                    dd.className = "db_list_first_item"
                    dd.innerHTML = json[i].name
                    // dd.onselectstart = function() {
                    //     document
                    // }
                    dd.ondblclick = function(){
                        dd.style.width = 'auto'
                        dd.style.height = '40px'
                        dd.style.textAlign = 'center'
                        dd.style.backgroundColor = 'green'
                    }

                    dl.appendChild(dd)
                }
            }else if(XMLHttpReq.status == 100){
            
            }else if(XMLHttpReq.status == 300){
            
            }else if(XMLHttpReq.status == 400){
            
            }else if(XMLHttpReq.status == 500){
            
            }else if(XMLHttpReq.status == 0){
                /** 0不是http协议的状态,关于XMLHttpReq.status的说明:
                1、If the state is UNSENT or OPENED, return 0.（如果状态是UNSENT或者OPENED，返回0）
                2、If the error flag is set, return 0.（如果错误标签被设置，返回0）
                3、Return the HTTP status code.（返回HTTP状态码）
                第一种情况,例如:url请求的是本地文件,状态会是0
                第二种情况经常出现在跨域请求中,比如url不是本身网站IP或域名,例如请求www.baidu.com时
                第三种,正常请求本站http协议信息时,正常返回http协议状态值
                **/
            }
            
        }
    };
    xhr.send(null);
}

function addConnect(){
    var xhr = new XMLHttpRequest();
    xhr.timeout = 2000;
    xhr.ontimeout = function(event) {
        alert("请求超时");
    }

    var sendData = {name:"database",type:0,server:"./database.db"}

    xhr.open('POST', 'http://localhost:2200/add/connect')
    // xhr.send(json)
    xhr.send(JSON.stringify(sendData));
    xhr.onreadystatechange = function() {
        if (xhr.status == 200) {
            alert("请求成功")
        }
    }
}
function post(){
    var xhr = new XMLHttpRequest();
    xhr.timeout = 3000;
    xhr.ontimeout = function (event) {
        alert("请求超时！");
    }
    var formData = new FormData();
    formData.append('tel', '18217767969');
    formData.append('psw', '111111');
    xhr.open('POST', 'http://www.test.com:8000/login');
    xhr.send(formData);
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4 && xhr.status == 200) {
            alert(xhr.responseText);
        }
        else {
            alert(xhr.statusText);
        }
    }
}

window.onload = function(){
    // alert("onloaded")
    getConnects()
}