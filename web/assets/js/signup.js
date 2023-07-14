$('#form-signup').on('submit', createUser);

function createUser(event) {
    event.preventDefault();
    
    if ($('#password').val() != $('#password-confirmation').val()) {
        alert('As senhas não coincidem!')
        return;
    }

    $.ajax({
        url: '/users',
        method: "POST",
        data: {
            name: $('#name').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            password: $('#password').val()
        }
    }).done(() => {
        alert('Usuário cadastrado com sucesso!')
    }).fail((erro) => {
        console.log(erro);
        alert('Falha ao cadastrar usuário!')
    });
}