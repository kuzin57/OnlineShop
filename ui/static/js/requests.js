import { delete_cookies, getCookie, setCookie } from "./cookies.js";
import { setDropdown } from "./dropdown.js";

fetch('/', {
    method: 'GET',
    headers: new Headers({
      'Content-Type': 'application/json',
      'Token': getCookie("token")
    })
  })
  .then(response => {
    let authorized = response.headers.get('Authorized');
    let profile = response.headers.get('Profile');
    console.log(authorized);
    if (authorized == "true") {
      setDropdown();
       
      if (document.URL == "http://localhost:7000/") {
        document.getElementById("serverMessageBox").innerHTML = "Hello, " + getCookie("username");  
      }
    } else {
      document.getElementById("login-username").innerHTML=`
      <a href="/auth">  
       <button>Login</button>  
      </a>
      `;
    }
  })
