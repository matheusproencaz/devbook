$('#login').on('submit', Login);

function Login(event) {
    event.preventDefault();

    $.ajax({
        url: '/login',
        method: 'POST',
        data: {
            email: $('#email').val(),
            password: $('#password').val()
        }
    }).done(() => {
        window.location = '/home'
    }).fail((err) => {
        console.log(err)
        alert("Usuário ou senha inválidos!")
    })
}