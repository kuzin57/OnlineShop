export class Product {
    constructor(id, name, brand, price) {
        this.id = id;
        this.name = name;
        this.brand = brand;
        this.price = price;
        this.amount = 1;
    }
}

export function addProductToList(product) {
    console.log("product", product);

    var jsonData = sessionStorage.getItem("chosen_products");
    if (jsonData == null) {
        var chosenProducts = new Array(1);
        chosenProducts[0] = product;
        const jsonArray = JSON.stringify(chosenProducts);
        sessionStorage.setItem("chosen_products", jsonArray);
        console.log("array", sessionStorage.getItem("chosen_products"));
        return;
    }

    var chosenProductsArray = JSON.parse(jsonData);
    chosenProductsArray.push(product);
    const jsonArray = JSON.stringify(chosenProductsArray);
    sessionStorage.setItem("chosen_products", jsonArray);
    console.log("array", sessionStorage.getItem("chosen_products"));
}

export function removeProductFromList(productID) {
    var chosenArray = JSON.parse(sessionStorage.getItem("chosen_products"));
    for (var i = 0; i < chosenArray.length; i++) {
        if (chosenArray[i].id == productID) {
            chosenArray.splice(i, 1);
            break;
        }
    }

    const jsonArray = JSON.stringify(chosenArray);
    sessionStorage.setItem("chosen_products", jsonArray);
    console.log("array", sessionStorage.getItem("chosen_products"));
}