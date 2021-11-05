var base64Ts

function selectOne(){
    chV = document.getElementById("select").value
    console.log(chV)
    if (chV == "") {
    window.alert("請勿輸入空值")
    }else{
    document.getElementById("fs").style = null;
    const uri = 'http://127.0.0.1:8080/getWho';
  const initDetails = {
      method: 'POST',
      headers: {
          "Content-Type": "application/json; charset=utf-8"
      },
      body:JSON.stringify({"name":chV}),
      mode: "cors"
  }
      fetch( uri, initDetails )
          .then(function (response) {
            return response.json();
          })
          .then(function (data) {
          appendData(data);
         })
          .catch( err =>
          {
              console.log( 'Fetch Error :-S', err );
          } );
  function appendData(data) {
  let mainContainer = document.getElementById("myData");
  for (var i = 0; i < data.length; i++) {
    let tr = document.createElement("tr");
    tr.innerHTML = `<td id="name${data[i].id}" style="width: 20%;">${data[i].name}</td><td><img id="${data[i].id}" style="width: 400px;" src="data:image/jpeg;base64,${data[i].picture}"></td>
    <td><button id="change" onclick="again(this)" value="${data[i].id}">重拍</button>
    <button id="change" onclick="change(this)" value="${data[i].id}">覆蓋</button></td>
    <td><button id="del" onclick="del(this)" value="${data[i].id}">刪除</button></td>`;
    mainContainer.appendChild(tr);
  }
}
}
}

function select(){
    document.getElementById("fs").style = null;
    const uri = 'http://127.0.0.1:8080/get';
  const initDetails = {
      method: 'GET',
      headers: {
          "Content-Type": "application/json; charset=utf-8"
      },
      mode: "cors"
  }
      fetch( uri, initDetails )
          .then(function (response) {
            return response.json();
          })
          .then(function (data) {
          appendData(data);
         })
          .catch( err =>
          {
              console.log( 'Fetch Error :-S', err );
          } );
  function appendData(data) {
  let mainContainer = document.getElementById("myData");
  for (var i = 0; i < data.length; i++) {
    let tr = document.createElement("tr");
    tr.innerHTML = `<td id="name${data[i].id}" style="width: 20%;">${data[i].name}</td><td><img id="${data[i].id}" style="width: 400px;" src="data:image/jpeg;base64,${data[i].picture}"></td>
    <td><button id="change" onclick="again(this)" value="${data[i].id}">重拍</button>
    <button id="change" onclick="change(this)" value="${data[i].id}">覆蓋</button></td>
    <td><button id="del" onclick="del(this)" value="${data[i].id}">刪除</button></td>`;
    mainContainer.appendChild(tr);
  }
}
}

function again(e){
    const uri = 'http://127.0.0.1:8080/take';
    const initDetails = {
        method: 'GET',
        headers: {
            "Content-Type": "application/json; charset=utf-8"
        },
        mode: "cors"
    }
        fetch( uri, initDetails )
            .then(function (response) {
              return response.json();
            })
            .then(function (data) {
              base64Ts = JSON.stringify(data)
              addendData(data)
           })
            .catch( err =>
            {
                console.log( 'Fetch Error :-S', err );
            });
  function addendData(base64Ts){
      var a = "data:image/jpeg;base64,"
      var b = a + base64Ts
      document.getElementById(`${e.value}`).src = b
  }
  }
function take(){
      const uri = 'http://127.0.0.1:8080/take';
      const initDetails = {
          method: 'GET',
          headers: {
              "Content-Type": "application/json; charset=utf-8"
          },
          mode: "cors"
      }
          fetch( uri, initDetails )
              .then(function (response) {
                return response.json();
              })
              .then(function (data) {
                base64Ts = JSON.stringify(data)
                addendData(data)
             })
              .catch( err =>
              {
                  console.log( 'Fetch Error :-S', err );
              });
    function addendData(base64Ts){
        var a = "data:image/jpeg;base64,"
        var b = a + base64Ts
        document.getElementById('img').src = b
    }
    }
function insert(){
    let picByte = JSON.parse(base64Ts)
    var chV = prompt('請輸入檔案名稱');
    if (chV == "") {
    window.alert("請勿輸入空值")
    }else{
      const uri = 'http://127.0.0.1:8080/save';
      const initDetails = {
          method: 'POST',
          headers: {
              "Content-Type": "application/json; charset=utf-8"
          },
          body: JSON.stringify({"name":chV,"picture":picByte}),
          mode: "cors"
      }
          fetch( uri, initDetails )
              .then(function (response) {
                return response.json();
              })
              .then(function (data) {
              window.alert(JSON.stringify(data.name)+' 新增成功')
              // location.reload()
             })
              .catch( err =>
              {
                  console.log( 'Fetch Error :-S', err );
              } );
    }
    }
    function change(e){
      var chV = prompt('請輸入修改值');
      var chId = e.value
      var cid = parseInt(chId,10)
      let picByte = JSON.parse(base64Ts)
      const uri = `http://127.0.0.1:8080/put/${chId}`;
      const initDetails = {
          method: 'PUT',
          headers: {
              "Content-Type": "application/json; charset=utf-8"
          },
          body: JSON.stringify({"id":cid,"name":chV,"picture":picByte}),
          mode: "cors"
      }
          fetch( uri, initDetails )
              .then(function (response) {
                return response.json();
              })
              .then(function (data) {
              window.alert(JSON.stringify(data.name)+'修改成功')
              location.reload()
            // addendData()
             })
              .catch( err =>
              {
                  console.log( 'Fetch Error :-S', err );
              } );
        // function addendData(){
        // document.getElementById(`name${e.value}`)
        // select()
    // }
    }
    
    function del(e){
      var chId = e.value
      var cid = parseInt(chId,10)
      const uri = `http://127.0.0.1:8080/del/${chId}`;
      const initDetails = {
          method: 'DELETE',
          headers: {
              "Content-Type": "application/json; charset=utf-8"
          },
          body: JSON.stringify({"id":cid}),
          mode: "cors"
      }
          fetch( uri, initDetails )
              .then(function (response) {
                return response.json();
              })
              .then(function () {
              // window.alert(JSON.stringify(data))
            // location.reload()
             })
              .catch( err =>
              {
                  console.log( 'Fetch Error :-S', err );
              } );
    }
