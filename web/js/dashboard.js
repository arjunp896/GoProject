// code for navbar
$(function() {
    $("#includedContent").load("../web/assets/navbar.html");
});


// fetch data from json
var getCategory = function() {
    $.ajax({
        url: "/car/cars/{category}",
        success: function(data) {
            return result
        }
    })

}

function getMake(make) {

    alert(make);


    // var makes = function() {
    //     $.ajax({
    //         url: "/car/cars/{category}/" + make,
    //         success: function(data) {
    //             return result
    //         }
    //     })

    // }
    var result = ["abc", "def", "xyz"];


    return result;
}

var getYear = function() {
    $.ajax({
        url: "/car/cars/{category}/{make}/{year}",
        success: function(data) {
            return result
        }
    })

}





// load data of category in dropdown
$(function() {

    var result = ["abc", "def", "xyz"];

    // loadDropDownData("category", getCategory)
    loadDropDownData("category", result);

    // result = ["audi", "merc", "bmw"];


    // result = [2015, 2016, 2022];


});

$(document).on('click', '#category .dropdown-content a', function() {

    var btn = $(this).parent().parent().find(".dropbtn");

    btn.html($(this).text());

    $("#make .dropdown-content").not(':first').remove();
    $("#year .dropdown-content").not(':first').remove();

    var category = getMake($("#category .dropbtn").text());

    loadDropDownData("make", category);

});

$(document).on('click', '#make .dropdown-content a', function() {

    var btn = $(this).parent().parent().find(".dropbtn");

    btn.html($(this).text());

    $("#year .dropdown-content").not(':first').remove();

    var data = getYear($("#make .dropbtn").text());

    loadDropDownData("year", category, make);


});

$(document).on('click', '#year .dropdown-content a', function() {

    var btn = $(this).parent().parent().find(".dropbtn");

    btn.html($(this).text());

    var data = getYear($("#make .dropbtn").text());

});

//  load drop down data
function loadDropDownData(id, data) {

    $.each(data, function(index, item) {

        alert(item);

        // var option = $("#category").find(".dropdown-content a:first").clone();
        var option = $("#" + id + " .dropdown-content a:first").clone();

        $(option).text(item);

        $(option).appendTo($("#" + id + " .dropdown-content"));

    })
    $("#" + id).find(".dropdown-content a:first").hide();

}

//code for store data into dropdown from json
// $(function() {
//     let dropdown = $('#dropdown');

//     dropdown.empty();
//     const url = "/car/cars/";

//     $.getJSON(url, function(loadMainItem) {

//         $.each(loadMainItem.data, function(i, f) {
//             var makemenu = "<a href='#' class='dropdown-content dropdown-toggle'  id='" + f.id + "' role='button' data-toggle='dropdown' aria-haspopup='true' aria-expanded='false'>" + f.text + "</a>";

//             $(makemenu).appendTo("dropdown-content");
//         });
//     });

// });

// code for creating card

$(function() {
    //hide first div or remove after append using `$(".card:first").remove()`

    // $(".column:first").hide()

    $.ajax({
        url: "/car/cars/",
        success: function(result) {

            console.log(result);

            $.each(result, function(index, item) {

                var cards = $(".column:first").clone() //clone first divs

                // var userId = item.userId;
                var vehicleid = item.id;
                var imgurl = item.imageurl;
                var company = item.make;
                var model = item.model
                var price = item.price;

                //add values inside divs
                $(cards).find(".card-id").attr("href", "/car/" + vehicleid);

                if (imgurl == "") {
                    $(cards).find(".card-img").attr("src", "../web/images/home1.jpg");
                } else {
                    $(cards).find(".card-img").attr("src", imgurl);
                }

                $(cards).find(".card-company").text("Comapny : " + company);
                $(cards).find(".card-model").text("Model : " + model);
                $(cards).find(".card-price").text("Base Price : " + price);
                $(cards).show() //show cards

                $(cards).appendTo($(".row")) //append to container
            });
        }
    });
});

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