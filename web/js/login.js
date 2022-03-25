function login() {


    $("#pills-profile").removeClass("show active");
    $("#pills-home").addClass("show active")

    $("#pills-home-tab").addClass('active');
    $('#pills-profile-tab').removeClass('active');

    $("#pills-home").delay(100).fadeIn(100);
    $("#pills-profile").fadeOut(100);




    e.preventDefault();
}

function signup() {


    $("#pills-home").removeClass("show active");
    $("#pills-profile").addClass("show active")

    $("#pills-profile-tab").addClass('active');
    $('#pills-home-tab').removeClass('active');

    $("#pills-profile").delay(100).fadeIn(100);
    $("#pills-home").fadeOut(100);



    e.preventDefault();
}