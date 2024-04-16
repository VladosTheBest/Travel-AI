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
function loginUser(email, password) {
    fetch('http://localhost:8081/login', {
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
            console.log('Вход выполнен успешно!');
        } else {
            console.error('Ошибка при входе:', response.statusText);
        }
    })
    .catch(error => {
        console.error('Ошибка при отправке запроса:', error);
    });
}
registerUser('user@example.com', 'password123');
loginUser('user@example.com', 'password123');
