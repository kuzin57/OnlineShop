import { delete_cookies, getCookie, setCookie } from "./cookies.js";
import { setDropdown } from "./dropdown.js";

fetch('/my_orders', {
    method: 'GET',
    headers: new Headers({
      'Content-Type': 'application/json',
      'Token': getCookie("token"),
      'Email': getCookie("email"),
      'Get-Orders': 'yes'
    })
  })
  .then(response => {
    let data = response.json();
    console.log("data", data);
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

    return data;
  }).then(
    (data) => {
      console.log("orders", data.orders[0]);
      for (var i = 0; i < data.orders.length; i++) {
        var newLine = document.createElement("tr");
  
        var orderID = document.createElement("th");
        orderID.innerHTML = data.orders[i].id;
  
        var date = document.createElement("th");
        date.innerHTML = data.orders[i].date;
  
        var deliveryDate = document.createElement("th");
        deliveryDate.innerHTML = data.orders[i].delivery_date;
  
        var address = document.createElement("th");
        address.innerHTML = data.orders[i].city + ", " + data.orders[i].street + ", " + 
                              data.orders[i].house_number + ", ap. " +
                              data.orders[i].flat_number;
        
        var cost = document.createElement("th");
        cost.innerHTML = data.orders[i].total_sum;
  
        var products = document.createElement("th");
        var listProducts = document.createElement("ul");
        for (var j = 0; j < data.orders[i].chosen_products.length; j++) {
            var point = document.createElement("li");
            point.innerHTML = data.orders[i].chosen_products[j].name + " " + 
                              data.orders[i].chosen_products[j].brand + " " + 
                              data.orders[i].chosen_products[j].amount + "шт.";
            listProducts.appendChild(point);
        }
  
        products.appendChild(listProducts);
  
        newLine.appendChild(orderID);
        newLine.appendChild(date);
        newLine.appendChild(deliveryDate);
        newLine.appendChild(cost);
        newLine.appendChild(address);
        newLine.appendChild(products);

        document.getElementById("table-body").appendChild(newLine);
      }    
    }
  );