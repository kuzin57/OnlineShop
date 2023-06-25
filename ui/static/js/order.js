import {getCookie} from "./cookies.js";


function createOrder(totalSum) {
    var inputForm = document.getElementById("order_form");
        
    const formData = new FormData(inputForm);
    console.log(getCookie("email"));
    fetch("/order", {
        method: "POST",
        headers: new Headers({
            'Content-Type': 'application/json',
            'Email': getCookie("email")
        }),
        body: JSON.stringify({
            delivery_date: formData.get("delivery_day"),
            city: formData.get("city"),
            street: formData.get("street"),
            house_number: parseInt(formData.get("house_number")),
            flat_number: parseInt(formData.get("flat_number")),
            chosen_products: JSON.parse(sessionStorage.getItem("chosen_products")),
            total_sum: totalSum
        }),
    }).then(
        response => response.json()
    ).then(
        (data) => {
            if (data.status == 200) {
                window.location.replace("/");
            } else {
                document.getElementById("serverMessageBox").innerHTML=data.description
            }
        }
    ).catch(
        error => console.error(error)
    );
}


var jsonData = sessionStorage.getItem("chosen_products");
var chosenProducts = JSON.parse(jsonData);
var total = 0;
for (var i = 0; i < chosenProducts.length; i++) {
    total += chosenProducts[i].price * chosenProducts[i].amount;
    
    var newTr = document.createElement("tr");

    var numberTh = document.createElement("th");
    numberTh.innerHTML = i + 1;

    var nameTh = document.createElement("th");
    nameTh.innerHTML = chosenProducts[i].name;

    var brandTh = document.createElement("th");
    brandTh.innerHTML = chosenProducts[i].brand;

    var amountTh = document.createElement("th");
    amountTh.innerHTML = chosenProducts[i].amount;

    var priceTh = document.createElement("th");
    priceTh.innerHTML = chosenProducts[i].price;

    newTr.appendChild(numberTh);
    newTr.appendChild(nameTh);
    newTr.appendChild(brandTh);
    newTr.appendChild(amountTh);
    newTr.appendChild(priceTh);

    document.getElementById("table-body").appendChild(newTr);
}

var totalSum = document.createElement("mark");
totalSum.setAttribute("style", "margin-left: 15%; font-size: 30px;");
totalSum.innerHTML = "Total: " + total.toString() + "₽";
document.getElementsByClassName("product-list")[0].appendChild(totalSum);
document.getElementById("create-order").addEventListener('click', function() {
    createOrder(total);
});