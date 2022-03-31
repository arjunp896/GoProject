$(function() {
    // $("#includedContent").load("../web/assets/navbar.html");

    var vehicleDetails = loadVehicelDetails(id);

    // var json = vehicleDetails.colors;

    console.log(JSON.stringify(vehicleDetails));

    // console.log("vehicleDetails: " + json[0].color_id);

    // alert("vehicleDetails.tyres: " + vehicleDetails.tyres);

    loadDetails(vehicleDetails.vehicle);

    loadChecks("colors", vehicleDetails.colors);

    loadDropDown("engine", vehicleDetails.engines);
    loadDropDown("tyre", vehicleDetails.tyres);

    loadDropDown("rooftop", vehicleDetails.rooftops);
    loadDropDown("transmission", vehicleDetails.transmission);

    loadChecks("accessories", vehicleDetails.accessories);


});

function loadDetails(vehicle) {

    $("#make").html("Make: " + vehicle.make + " " + vehicle.model);
    $("#submodel").html("Submodel: " + vehicle.submodel);
    $("#modeltype").html("Model Type: " + vehicle.modeltype);
    $("#category").html("Category: " + vehicle.category);


    if (vehicle.imageurl != "") {
        $("#img").attr("src", vehicle.imageurl);

    }


    $("#vehicleid").prop("value", vehicle.id);
    $("#name").prop("value", vehicle.make);
    $("#mfgyear").prop("value", vehicle.yearofmfg);
    $("#model").prop("value", vehicle.model);
    $("#price").prop("value", vehicle.price);

}


function loadDropDown(id, data) {

    if (data == null) {
        $("#" + id).parent().hide();
        return
    } else {
        $("#id").show();
    }


    $.each(data, function(index, item) {


        var value, text;

        switch (id) {
            case "engine":
                value = item.engine_id;
                text = item.engine_type + " - " + item.power + " -   $" + item.price;
                break;

            case "rooftop":
                value = item.engine_id;
                text = item.engine_type + " - " + item.power + " -   $" + item.price;
                break;

            case "transmission":
                value = item.transmission_id;
                text = item.transmission_type + " -   $" + item.price;

                break;


            case "tyre":
                value = item.tyre_id;
                text = "size: " + item.size + " -   $" + item.price;
                break;


            default:
                break;
        }

        var o = new Option(text, value);


        $("#" + id).append(o);

    });

}

function submitForm() {
    $("#custcarform").submit();
}

function fnSubmit() {

    // alert("form");


    $('#custcarform')
        .ajaxForm({
            url: '/createvehicle',
            method: "POST",
            success: function(data) {

                // alert(data.message);

                console.log("Success: " + data);
                window.location.href = "/web/vehicles.html";

            },
            error: function(data) {

                console.log("Error" + data);
                var json = JSON.parse(data.responseText);
                console.log(json);
                // alert("Error: " + json.error);
            }
        });
}


function loadChecks(id, data) {

    if (data == null) {
        $("#" + id).parent().hide();
        return
    } else {
        $("#id").show();
    }


    $.each(data, function(index, item) {

        var value, text;


        var divChild = $("#" + id).find(".form-check:first").clone();

        if (id == "colors") {
            divChild.find(".form-option-color").css("background-color", item.color_name);
            divChild.find(".form-option-label").prop("for", id + index);

            value = item.color_id;

        }

        if (id == "accessories") {
            divChild.find(".form-check-label").prop("for", id + index);

            value = item.accessory_id;
            text = item.accessories_name + " - $" + item.price;

            divChild.find(".form-check-label").html(text);
        }

        divChild.find(".form-check-input").prop("value", value);
        divChild.find(".form-check-input").prop("id", id + index);

        $("#" + id).append(divChild);

    });

    $("#" + id).find(".form-check:first").hide();


}

function loadVehicelDetails(id) {

    var vehicleDetails;

    $.ajax({
        type: "GET",
        url: "/vehicle/" + id,
        async: false,
        success: function(response) {

            vehicleDetails = response
        }
    });

    return vehicleDetails;
}