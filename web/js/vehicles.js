// code for navbar
$(function() {
    $("#includedContent").load("../web/assets/navbar.html");
});

// code for printing card
$(function() {
    //hide first div or remove after append using `$(".card:first").remove()`

    // $(".column:first").hide()

    $.ajax({
        url: "",
        success: function(result) {
            $.each(result, function(index, item) {

                var cards = $(".column:first").clone() //clone first divs

                var userId = item.userId;
                var typeId = item.id;
                var imgId = item.img;
                var companyId = item.comapny;
                var modelId = item.model;
                var vehicleTypeId = item.vehicleType;
                var categoryId = item.category;
                var subModelId = item.subModel;
                var modelTypeId = item.modelType;
                var yearId = item.year;
                var colorId = item.color;
                var engineId = item.engine;
                var transmissionId = item.transmission;
                var tyersId = item.tyres;
                var rooftopId = item.rooftop;
                var accessoriesId = item.accessories;

                var priceId = item.price;


                //add values inside divs
                $(cards).find(".card-header").html("user id: " + userId + " - " + "id: " + typeId);
                $(cards).find(".card-img").attr("src", "../web/images/home1.jpg");
                $(cards).find(".company").html(companyId);
                $(cards).find(".model").html(modelId);
                $(cards).find(".vehicleType").html(vehicleTypeId);
                $(cards).find(".category").html(categoryId);
                $(cards).find(".subModel").html(subModelId);
                $(cards).find(".modelType").html(modelTypeId);
                $(cards).find(".year").html(yearId);
                $(cards).find(".color").html(colorId);
                $(cards).find(".engine").html(engineId);
                $(cards).find(".transmission").html(transmissionId);
                $(cards).find(".tyres").html(tyersId);

                if (category == "car") {
                    $(cards).find(".rooftop").html(rooftopId);
                } else {
                    $(cards).find(".rooftop").hide();
                }

                $(cards).find(".accessories").html(accessoriesId);
                $(cards).find(".price").html(priceId);
                $(cards).show() //show cards

                $(cards).appendTo($(".row")) //append to container
            });
        }
    });
});