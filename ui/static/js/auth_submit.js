function setCookie(name, value, options = {}) {

    options = {
      path: '/',
      // add other defaults here if necessary
      ...options
    };
  
    if (options.expires instanceof Date) {
      options.expires = options.expires.toUTCString();
    }
  
    let updatedCookie = encodeURIComponent(name) + "=" + encodeURIComponent(value);
  
    for (let optionKey in options) {
      updatedCookie += "; " + optionKey;
      let optionValue = options[optionKey];
      if (optionValue !== true) {
        updatedCookie += "=" + optionValue;
      }
    }
  
    document.cookie = updatedCookie;
  }

function auth_submit() {
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