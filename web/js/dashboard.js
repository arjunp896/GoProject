// code for navbar

var type = "all";

$(function() {

    $("#includedContent").load("../web/assets/navbar.html");

    // load data of category in dropdown

    $(".dpcontainer").hide();


    getCards("all");

});


// manage navbars

function initialize() {
    $(".dpcontainer").show();
    $("#category .dropdown-content a").not(':first').remove();

    $("#category .dropbtn").text("Select Category");
    $("#make .dropbtn").text("Select Make");
    $("#year .dropbtn").text("Select Year");

}

function loadCars() {

    initialize();




    var categories = getCategories("cars");
    // alert(categories[0]);

    loadDropDownData("category", categories);

    type = "cars";

    getCards("cars");


}

function loadBikes() {

    initialize();

    type = "bikes";

    var categories = getCategories("bikes");

    loadDropDownData("category", categories);

    getCards("bikes");

}

function getCategories(type) {

    var urlString;

    if (type == "cars") {
        urlString = "/Car/getcategories";

    } else if (type == "bikes") {
        urlString = "/Bike/getcategories"
    }


    var c;
    $.ajax({
        url: urlString,
        async: false,

        success: function(data) {

            // alert(data.categories[0]);

            // console.log("categories  getCategories: " + JSON.parse(data));

            c = data.categories;
        }
    });

    return c;

}

function getCarsByCategory() {


    if (type == "all") {
        urlString = "/vehicles/categories"
    } else if (type == "cars") {
        urlString = "/car/cars/{category}";

    } else if (type == "bikes") {
        urlString = "/bike/bikes/{category}"
    }

    // fetch data from json
    var categories = function() {

        // alert(type);

        $.ajax({
            url: "/car/cars/{category}",
            success: function(data) {
                return result
            }
        })

    }

    return categories;

}

function getMake(category) {

    var urlString;

    if (type == "cars") {
        urlString = "/Car/getmakes/" + category;

    } else if (type == "bikes") {
        urlString = "/Bike/getmakes/" + category;
    }


    var m;

    $.ajax({
        url: urlString,
        async: false,

        success: function(data) {

            // alert(data.makes);

            m = data.makes;
        }
    });

    return m;
}


function getYear(category, make) {

    var urlString;

    if (type == "cars") {
        urlString = "/Car/getyears/" + category + "/" + make;

    } else if (type == "bikes") {
        urlString = "/Bike/getyears/" + category + "/" + make;
    }

    // alert(category);

    var y;

    $.ajax({
        url: urlString,
        async: false,

        success: function(data) {

            // alert(data.years);

            y = data.years;
        }
    });

    return y;

}


// dropdown click change


$(document).on('click', '#category .dropdown-content a', function() {

    var btn = $(this).parent().parent().find(".dropbtn");

    btn.html($(this).text());

    $("#make .dropdown-content a").not(':first').remove();
    $("#year .dropdown-content a").not(':first').remove();

    $("#make .dropbtn").text("Select Make");
    $("#year .dropbtn").text("Select Year");


    var category = $("#category .dropbtn").text();

    var makes = getMake(category);

    // alert(category);

    loadDropDownData("make", makes);

    getCards(type, category);


});

$(document).on('click', '#make .dropdown-content a', function() {

    var btn = $(this).parent().parent().find(".dropbtn");

    btn.html($(this).text());

    $("#year .dropdown-content").not(':first').remove();

    var category = $("#category .dropbtn").text();
    var make = $("#make .dropbtn").text();

    $("#year .dropbtn").text("Select Year");


    var years = getYear(category, make);

    // alert(years);

    loadDropDownData("year", years);

    getCards(type, category, make);


});

$(document).on('click', '#year .dropdown-content a', function() {

    var btn = $(this).parent().parent().find(".dropbtn");

    btn.html($(this).text());

    var category = $("#category .dropbtn").text();
    var make = $("#make .dropbtn").text();;
    var year = $("#year .dropbtn").text();;

    getCards(type, category, make, year);

});


//  load drop down data
function loadDropDownData(id, data) {


    // console.log("loadDropDownData: ", data);

    $.each(data, function(index, item) {

        // alert(item);

        // var option = $("#category").find(".dropdown-content a:first").clone();
        var option = $("#" + id + " .dropdown-content a:first").clone();

        $(option).text(item);

        $(option).appendTo($("#" + id + " .dropdown-content"));
        $(option).show();

    });
    $("#" + id).find(".dropdown-content a:first").hide();

}


// code for creating card

function loadCards(result) {

    console.log(result);
    $(".column:first").hide();

    $.each(result, function(index, item) {

        var cards = $(".column:first").clone() //clone first divs

        // var userId = item.userId;
        var vehicleid = item.id;
        var imgurl = item.imageurl;
        var company = item.make;
        var model = item.model
        var price = item.price;
        var vehicleType = item.vehicletype;

        //add values inside divs
        if (vehicleType.toUpperCase() == "CAR") {
            $(cards).find(".card-id").attr("href", "/car/" + vehicleid);
        }

        if (vehicleType.toUpperCase() == "BIKE") {
            $(cards).find(".card-id").attr("href", "/bike/" + vehicleid);
        }

        if (imgurl == "") {
            $(cards).find(".card-img").attr("src", "../web/images/home1.jpg");
        } else {
            $(cards).find(".card-img").attr("src", imgurl);
        }

        $(cards).find(".card-company").text("Comapny : " + company);
        $(cards).find(".card-model").text("Model : " + model);
        $(cards).find(".card-price").text("Base Price : $" + price);
        $(cards).show() //show cards

        $(cards).appendTo($(".row")) //append to container
    });

}


function getCards(type, category, make, year) {

    // alert(type);

    var urlString;

    if (type == "all") {
        urlString = "/vehicles";
    } else if (type == "cars") {
        urlString = "/car/cars/";
    } else if (type == "bikes") {
        urlString = "/bike/bikes"
    }

    if (category) {
        if (type == "cars") {
            urlString = "/car/cars/" + category;
        } else if (type == "bikes") {
            urlString = "bike/bikes/" + category;
        }
    }

    if (category && make) {
        if (type == "cars") {
            urlString = "/car/cars/" + category + "/" + make;
        } else if (type == "bikes") {
            urlString = "/bike/bikes/" + category + "/" + make;
        }
    }


    if (category && make && year) {
        if (type == "cars") {
            urlString = "/car/cars/" + category + "/" + make + "/" + year;
        } else if (type == "bikes") {
            urlString = "/bike/bikes/" + category + "/" + make + "/" + year;
        }
    }





    $(".column").not(':first').remove();



    $.ajax({
        url: urlString,
        success: function(result) {

            loadCards(result);

        }
    });
}

// code for search filter

/* When the user clicks on the button,
toggle between hiding and showing the dropdown content */
function myFunction() {
    document.getElementById("myDropdown").classList.toggle("show");
}

// Close the dropdown menu if the user clicks outside of it
window.onclick = function(event) {
    if (!event.target.matches('.dropbtn')) {
        var dropdowns = document.getElementsByClassName("dropdown-content");
        var i;
        for (i = 0; i < dropdowns.length; i++) {
            var openDropdown = dropdowns[i];
            if (openDropdown.classList.contains('show')) {
                openDropdown.classList.remove('show');
            }
        }
    }
}