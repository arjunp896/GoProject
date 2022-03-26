// code for navbar
$(function() {
    $("#includedContent").load("../web/assets/navbar.html");
});

// code for creating card

$(function() {
    //hide first div or remove after append using `$(".card:first").remove()`

    $(".column:first").hide()

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