let btn = document.querySelector('.btn');
function registerUser(email, password) {
    fetch('http://localhost:8081/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            email: email,
            password: password
        }),
    })
    .then(response => {
        if (response.ok) {
            console.log('Пользователь зарегистрирован успешно!');
        } else {
            console.error('Ошибка при регистрации:', response.statusText);
        }
    })
    .catch(error => {
        console.error('Ошибка при отправке запроса:', error);
    });
}