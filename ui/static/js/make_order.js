import { getCookie } from "./cookies.js";
import { setDropdown } from "./dropdown.js";

export function makeOrder() {
    fetch('/catalogue', {
        method: 'POST',
        headers: new Headers({
          'Content-Type': 'application/json',
          'Token': getCookie("token")
        })
    }).then(response => {
        let json_data = response.json();
        let authorized = response.headers.get('Authorized');
        console.log(authorized);
        if (authorized == "true") {
          setDropdown();
          window.location.replace("/order");
        } else {
          document.getElementById("login-username").innerHTML=`
          <a href="/auth">  
           <button>Login</button>  
          </a>
          `; 
          document.getElementById("messageBox").innerHTML = "To make orders you need to be authorized!";
        }
    
        return json_data;
    });

}