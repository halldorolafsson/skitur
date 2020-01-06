function sendJSON(){
    let dato = document.querySelector('#dato'); 
    let antallkilometer = document.querySelector('#antallkilometer'); 
    let antallminutter = document.querySelector('#antallminutter'); 
    let sted = document.querySelector('#sted'); 
  
    // // Debuggery 
    console.log("Started");
    
    var data = JSON.stringify(
    {
        "Dato": dato.value, 
        "Antallkilometer": antallkilometer.value , 
        "Antallminutter": antallminutter.value, 
        "Sted": sted.value 
    }); 
    
    // // Debuggery 
    console.log("Innsendt data: "+ data);

    //Hard coded link. Need som sort og config here
    let url_1 = 'http://localhost:8080/skitur';
    // // Debuggery 
    console.log("Url fra config: "+ url_1);

    fetch(url_1, {
        method: "POST", 
        body: data,
        headers:{
          'Content-Type': 'application/json'
        }
    })
    .then((response) => response.json())
    .then((data) => {
        console.log('Success:', data)
        print('Lagret følgende ' + JSON.stringify(data));
    })
    .catch((error) => {
        console.error('Error:', error);
    });
    //Debuggery 
    console.log(JSON.stringify(data));
}

var r = document.getElementById('result');
function print(s){
    r.innerHTML += s + '<br>';
}

function goBack() {
    window.history.back();
}



