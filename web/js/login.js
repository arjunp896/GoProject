function fnSubmitLogin() {

    $('#loginForm')
        .ajaxForm({
            url: '/login',
            dataType: 'json',
            success: function(data) {
                console.log(data)
                if (data.redirect) {
                    window.location.href = data.redirect;
                }
            },
            error: function(data) {
                console.log(data)
                var json = JSON.parse(data.responseText);
                console.log(json)
                alert("Error: " + json.error);
            }
        });
}

function fnSubmitSignUp() {

    $('#signUpForm')
        .ajaxForm({
            url: '/signup',
            dataType: 'json',
            method: "POST",
            success: function(data) {

                alert(data.message);

                console.log("Success: " + data);

            },
            error: function(data) {

                console.log("Error" + data);
                var json = JSON.parse(data.responseText);
                console.log(json);
                alert("Error: " + json.error);
            }
        });




}


function login() {


    $("#pills-profile").removeClass("show active");
    $("#pills-home").addClass("show active")

    $("#pills-home-tab").addClass('active');
    $('#pills-profile-tab').removeClass('active');

    $("#pills-home").delay(100).fadeIn(100);
    $("#pills-profile").fadeOut(100);


}

function signup() {


    $("#pills-home").removeClass("show active");
    $("#pills-profile").addClass("show active")

    $("#pills-profile-tab").addClass('active');
    $('#pills-home-tab').removeClass('active');

    $("#pills-profile").delay(100).fadeIn(100);
    $("#pills-home").fadeOut(100);

}