async function getImg(){
    let requestOptions = {
        method: 'GET',
        redirect: 'follow',
      };
    let section = document.getElementById('mainSection')
    await fetch("http://146.56.199.136:8080/api/picture?page=0&page_size=100", requestOptions)
    .then(response => response.json())
    .then(result => {
        for (let i of result["images"]){
            console.log(i)
            let img = document.createElement("img")
            img.src = "http://" + i['Url']
            section.appendChild(img)
        }
    })
    .catch(error => console.log('error', error));
}

function showFile(){
    
}

async function uploadImg(){
    let imgs = document.getElementById("upload")
    let formdata = new FormData();
    formdata.append("token", "123456");
    formdata.append("tag", "test");
    formdata.append("file", imgs.files[0])
    
    let requestOptions = {
        method: 'POST',
        body: formdata,
        redirect: 'follow'
      };

    fetch("http://146.56.199.136:8080/api/picture", requestOptions)
    .then(response => response.json())
    .catch(error => console.log("error" , error))
    .then(result => console.log(result))
    .catch(error => console.log('error', error));

}
