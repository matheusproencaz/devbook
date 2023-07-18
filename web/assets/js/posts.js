$('#new-post').on('submit', createPost);
$(document).on('click', '.like-post', likePost);
$(document).on('click', '.dislike-post', dislikePost);

$('#edit-post').on('click', updatePost)

$('.delete-post').on('click', deletePost)

function createPost(event) {
    event.preventDefault()

    $.ajax({
        url: '/posts',
        method: 'POST',
        dataType: 'text',
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        }
    }).done(() => {
        window.location = '/home';
    }).fail(() => {
        Swal.fire("Ops...", "Houve um erro ao criar a publicação!", "error");
    })
}

function likePost(event) {
    event.preventDefault();
    const elementClicked = $(event.target);
    const postId = elementClicked.closest('div').data('post-id');

    elementClicked.prop('disabled', true);
    $.ajax({
        url: `/posts/${postId}/like`,
        method: 'POST'
    }).done(() => {
        const spanQntLikes = elementClicked.next('span');
        const qntLikes = parseInt(spanQntLikes.text());
        spanQntLikes.text(qntLikes + 1);

        elementClicked.addClass('dislike-post');
        elementClicked.addClass('text-danger');
        elementClicked.removeClass('like-post');
    }).fail(() => {
        Swal.fire("Ops...", "Houve um erro ao curtir a publicação!", "error");
    }).always(() => {
        elementClicked.prop('disabled', false);
    })
}

function dislikePost(event) {
    event.preventDefault();
    const elementClicked = $(event.target);
    const postId = elementClicked.closest('div').data('post-id');

    elementClicked.prop('disabled', true);
    $.ajax({
        url: `/posts/${postId}/dislike`,
        method: 'POST'
    }).done(() => {
        const spanQntLikes = elementClicked.next('span');
        const qntLikes = parseInt(spanQntLikes.text());
        spanQntLikes.text(qntLikes - 1);

        elementClicked.addClass('like-post');
        elementClicked.removeClass('dislike-post');
        elementClicked.removeClass('text-danger');
    }).fail(() => {
        Swal.fire("Ops...", "Houve um erro ao descurtir a publicação!", "error");
    }).always(() => {
        elementClicked.prop('disabled', false);
    })
}

function updatePost(event) {
    $(this).prop('disabled', true);

    const postId = $(this).data('post-id');
    
    $.ajax({
        url: `/posts/${postId}`,
        method: 'PUT',
        dataType: 'text',
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        }
    }).done(() => {
        Swal.fire('Sucesso!', 'Publicação criado com sucesso', 'success')
            .then(() => {
                window.location = '/home'
            })
    }).fail(() => {
        Swal.fire("Ops...", "Houve um erro ao editar a publicação!", "error");
    }).always(() => {
        $('#edit-post').prop('disabled', false)
    })
}

function deletePost(event) {
    event.preventDefault();
    
    Swal.fire({
        title: 'Atenção!',
        text: 'Tem certeza que deseja excluir essa publicação? Essa ação é irreversível!',
        showCancelButton: true,
        cancelButtonText: 'Cancelar',
        icon: 'warning'
    }).then((confirmation) => {
        if(!confirmation.value) return;

        const elementClicked = $(event.target);
        const postJumbotron = elementClicked.closest('div');
        const postId = elementClicked.closest('div').data('post-id');
        elementClicked.prop('disabled', true);

        $.ajax({
            url: `/posts/${postId}`,
            method: 'DELETE',
        }).done(() => {
            postJumbotron.fadeOut('slow', () => {
                $(this).remove();
            });
        }).fail(() => {
            Swal.fire("Ops...", "Houve um erro ao excluir a publicação!", "error");
        }).always(() => {
            elementClicked.prop('disabled', false);
        })
    });
}
