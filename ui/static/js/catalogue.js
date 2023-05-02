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

let token = getCookie("token")
fetch('/catalogue', {
    method: 'POST',
    headers: new Headers({
      'Content-Type': 'application/json',
      'Token': token
    })
  })
  .then(response => {
    let json_data = response.json();
    let authorized = response.headers.get('Authorized');
    console.log(authorized);
    if (authorized == "true") {
      document.getElementById("Dropdown").innerHTML=`
      <div class="dropdown">
        <button>` + getCookie("username") + `</button>
        <div class="dropdown-options">
          <button>Orders</button>
          <button>Settings</button>
          <button>Logout</button>
        </div>
      </div>
      `;

      // <a href="#">Orders</a>
      //     <a href="#">Settings</a>
      //     <a href="#">Logout</a>
    } else {
      document.cookie = "";
      document.getElementById("login").innerHTML=`
      <form id="login button" action="/auth">
        <input type="submit" style="float: right;" value="Login"/>
      </form>
      `; 
    }

    return json_data;
  }).then(
    (data) => {
      console.log(data.products.length)
      const tbl = document.createElement("table");
      const tblBody = document.createElement("tbody");
      const headRow = document.createElement("tr");
      
      // adding headers to table
      const prodNameColumn = document.createElement("td");
      const prodNameColumnText = document.createTextNode("Название");
      prodNameColumn.appendChild(prodNameColumnText);
      headRow.appendChild(prodNameColumn);

      const brandNameColumn = document.createElement("td");
      const brandNameColumnText = document.createTextNode("Бренд");
      brandNameColumn.appendChild(brandNameColumnText);
      headRow.appendChild(brandNameColumn);

      const priceColumn = document.createElement("td");
      const priceColumnText = document.createTextNode("Цена");
      priceColumn.appendChild(priceColumnText);
      headRow.appendChild(priceColumn);

      const ratingColumn = document.createElement("td");
      const ratingColumnText = document.createTextNode("Рейтинг");
      ratingColumn.appendChild(ratingColumnText);
      headRow.appendChild(ratingColumn);
      tblBody.appendChild(headRow);

      for (let i = 0; i < data.products.length; i++) {
        const row = document.createElement("tr");

        const name = document.createElement("td");
        var a = document.createElement('a');
        var linkText = document.createTextNode(data.products[i].name);
        a.appendChild(linkText);
        a.title = data.products[i].name;
        a.href = "/products/"+data.products[i].name;
        name.appendChild(a);
        row.appendChild(name);

        const brand = document.createElement("td");
        const brandText = document.createTextNode(data.products[i].brand);
        brand.appendChild(brandText);
        row.appendChild(brand);

        const price = document.createElement("td");
        const priceText = document.createTextNode(data.products[i].price + "руб.");
        price.appendChild(priceText);
        row.appendChild(price);

        const rating = document.createElement("td");
        const ratingText = document.createTextNode(data.products[i].rating);
        rating.appendChild(ratingText);
        row.appendChild(rating);

        tblBody.appendChild(row);
      }

      tbl.appendChild(tblBody);
      document.body.appendChild(tbl);
      tbl.setAttribute("border", "2");
    }
  )
