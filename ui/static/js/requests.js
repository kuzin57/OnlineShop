function getCookie(cname) {
  let name = cname + "=";
  let decodedCookie = decodeURIComponent(document.cookie);
  let ca = decodedCookie.split(';');
  for(let i = 0; i <ca.length; i++) {
    let c = ca[i];
    while (c.charAt(0) == ' ') {
      c = c.substring(1);
    }
    if (c.indexOf(name) == 0) {
      return c.substring(name.length, c.length);
    }
  }
  return "";
}

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
      document.getElementById("Dropdown").innerHTML=`
      <div class="dropdown">
        <button>` + getCookie("username") + `</button>
        <div class="dropdown-options">
          <a href="#">Orders</a>
          <a href="#">Settings</a>
          <a href="#">Logout</a>
        </div>
      </div>
      `;
    } else {
      document.getElementById("login").innerHTML=`
      <form id="login button" action="/auth">
        <input type="submit" style="float: right;" value="Login"/>
      </form>
      `; 
    }
  })
