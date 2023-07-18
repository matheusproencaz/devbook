$('#login').on('submit', Login);

function Login(event) {
    event.preventDefault();

    $.ajax({
        url: '/login',
        method: 'POST',
        dataType: "text",
        data: {
            email: $('#email').val(),
            password: $('#password').val()
        }
    }).done(() => {
        window.location = '/home'
    }).fail((err) => Swal.fire("Ops...", "Usuário ou senha incorretos!", "error"));
}