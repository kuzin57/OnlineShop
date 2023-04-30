<<<<<<< HEAD
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

=======
>>>>>>> ba7e9b9 (removed useless sql scripts)
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
            response => response.json()
        ).then(
<<<<<<< HEAD
=======
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> 4b8c4fc (removed useless sql scripts)
            (data) => {
                document.getElementById("serverMessageBox").innerHTML=data.description;
                if (data.status != 403) {
                    setCookie("username", data.userName);
                    setCookie("token", data.token);
                    console.log(document.cookie);
                    window.location.replace("/");
                }
            }
        ).catch(
            error => console.error(error)
        )
        
<<<<<<< HEAD
=======
=======
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
>>>>>>> 021580b (fix test + fix communication between front and back)
=======
            (data) => {
                console.log(data)
                if (data.status == 403) {
                    document.getElementById("serverMessageBox").innerHTML="Login failed!"
                } else {
                    document.getElementById("serverMessageBox").innerHTML="Login succeeded!"
                }
            }
        ).catch(
            error => console.error(error)
        )
        
>>>>>>> ba7e9b9 (removed useless sql scripts)
>>>>>>> 4b8c4fc (removed useless sql scripts)
    });
}