$('#new-post').on('submit', createPost);

function createPost(event) {
    event.preventDefault()

    $.ajax({
        url: '/posts',
        method: "POST",
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        }
    }).done(() => {
        window.location = '/home';
    }).fail(() => {
        alert("Houve um erro ao criar a publicação!");
    })

}