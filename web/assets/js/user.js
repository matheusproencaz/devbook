$('#stop-follow').on('click', stopFollowing)
$('#follow').on('click', follow)
$('#edit-user').on('submit', edit)
$('#update-password').on('submit', editPassword)
$('#delete-user').on('click', deleteUser)

function stopFollowing() {
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/unfollow`,
        method: 'POST'
    }).done(() => {
        window.location = `/users/${userId}`
    }).fail(() => {
        Swal.fire("Ops...", "Erro ao parar de seguir o usuário!", "error");
        $('#stop-follow').prop('disabled', false);
    });
}

function follow() {
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/follow`,
        method: 'POST'
    }).done(() => {
        window.location = `/users/${userId}`
    }).fail(() => {
        Swal.fire("Ops...", "Erro ao seguir o usuário!", "error");
        $('#follow').prop('disabled', false);
    });
}

function edit(event) {
    event.preventDefault();

    $.ajax({
        url: '/edit-user',
        method: 'PUT',
        dataType: 'text',
        data: {
            name:  $("#name").val(),
            email: $("#email").val(),
            nick: $("#nick").val()
        }
    }).done(() => {
        Swal.fire("Sucesso!", "Usuário atualizado com sucesso!", "success")
            .then(() => {
                window.location = '/profile'
            })
    }).fail(() => {
        Swal.fire("Ops...", "Erro ao atualizar o usuário!", "error")
    })
}

function editPassword(event) {
    event.preventDefault()

    if ($('#new-password').val() != $('#confirm-password').val()) {
        Swal.fire("Ops...", "As senhas não coincidem!", "warning")
        return;
    }

    $.ajax({
        url: '/edit-password',
        method: 'POST',
        dataType: 'text',
        data: {
            new: $('#new-password').val(),
            current: $('#current-password').val()
        }
    }).done(() => {
        Swal.fire("Sucesso!", "A senha foi atualizada com sucesso!", "success")
            .then(() => {
                window.location = '/profile';
            })
    }).fail(() => {
        Swal.fire("Ops...", "Houve um erro ao atualizar a senha!", "error");
    });
}

function deleteUser() {
    Swal.fire({
        title: 'Atenção!',
        text: 'Tem certeza que deseja apagar a sua conta? Essa é uma ação irreversível!',
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then((confirmation) => {
        if (confirmation.value) {
            $.ajax({
                url: '/delete-user',
                method: 'DELETE'
            }).done(() => {
                Swal.fire("Sucesso!", "Seu usuário foi excluido com sucesso!", "success")
                .then(() => {
                    window.location = '/logout';
                })
            }).fail(() => {
                Swal.fire("Ops...", "Houve um erro ao tentar excluir o seu usuário!", "error");
            })
        }
    })
}