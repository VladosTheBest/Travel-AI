let btn = document.getElementById('.btn');
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
btn.addEventListener("click", loginUser(email, password))