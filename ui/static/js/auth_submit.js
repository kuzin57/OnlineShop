function auth_submit() {
    // let form = document.forms["auth_form"];
    const url = "/127.0.0.1:7000/auth"

    var inputForm = document.getElementById("auth_form");
    inputForm.addEventListener("submit", async (e)=> {
        e.preventDefault()
        
        const formData = new FormData(inputForm)
        console.log(formData.get("email"))

        fetch("/auth", {
            method: "POST",
            body: JSON.stringify({email: formData.get("email"), password: formData.get("password")}),
        }).then(
            response => response.text()
        ).then(
            (data) => {console.log(data);document.getElementById("serverMessageBox").innerHTML=data}
        ).catch(
            error => console.error(error)
        )
    

        // // (async () => {
        //     const rawResponse = await fetch('/auth', {
        //       method: 'POST',
        //       headers: {
        //         'Accept': 'application/json',
        //         'Content-Type': 'application/json'
        //       },
        //       body: JSON.stringify({a: 1, b: 'Textual content'})
        //     });
        //     const content = await rawResponse.json();
          
        //     console.log(content);
         // })();
    });
}