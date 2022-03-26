// code for navbar
$(function() {
    $("#includedContent").load("../web/assets/navbar.html");
});

// code for creating card

$(function() {
    //hide first div or remove after append using `$(".card:first").remove()`

    $(".column:first").hide()

    $.ajax({
        url: "https://jsonplaceholder.typicode.com/posts",
        success: function(result) {
            $.each(result, function(index, item) {

                var cards = $(".column:first").clone() //clone first divs

                var userId = item.userId;
                var typeId = item.id;
                var imgId = item.img;
                var companyId = item.comapny;
                var modelId = item.model
                var priceId = item.price;

                //add values inside divs
                $(cards).find(".card-header").html("user id: " + userId + " - " + "id: " + typeId);
                $(cards).find(".card-img").attr("src", "../web/images/home1.jpg");
                $(cards).find(".card-company").html("Comapny : " + titleId);
                $(cards).find("card-model").html("Model : " + modelId);
                $(cards).find(".card-price").html("Base Price : " + priceId);
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