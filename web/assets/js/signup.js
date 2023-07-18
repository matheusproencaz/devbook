$('#form-signup').on('submit', createUser);

function createUser(event) {
    event.preventDefault();
    
    if ($('#password').val() != $('#password-confirmation').val()) {
        Swal.fire("Ops...", "As senhas não coincidem!", "error");
        return;
    }

    $.ajax({
        url: '/users',
        method: "POST",
        dataType: "text",
        data: {
            name: $('#name').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            password: $('#password').val()
        }
    }).done(() => {
        Swal.fire("Sucesso!", "Usuário cadastrado com sucesso!", "success")
            .then(() => {
                $.ajax({
                    url: '/login',
                    method: 'POST',
                    dataType: "text",
                    data: {
                        email: $('#email').val(),
                        password: $('#password').val()
                    }
                })
                .done(() => window.location = '/home')
                .fail(() => Swal.fire("Ops...", "Erro ao autenticar usuário!", "error"));
            });
    }).fail((erro) => {
        Swal.fire("Ops...", "Erro ao cadastrar usuário!", "error");
    });
}