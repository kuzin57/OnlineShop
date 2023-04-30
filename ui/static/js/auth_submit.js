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
    inputForm.addEventListener("submit", (e)=> {
        e.preventDefault()
        
        const formData = new FormData(inputForm)

        fetch("/auth", {
            method: "POST",
            body: formData,
        }).then(
            response => response.text()
        ).then(
            (data) => {console.log(data); document.getElementById("serverMessageBox").innerHTML = data}
        ).catch(
            error => console.error(error)
        )
    });
}