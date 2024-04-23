// let btn = document.querySelector('.btn');
// function registerUser(email, password) {
//     fetch('http://localhost:8081/register', {
//         method: 'POST',
//         headers: {
//             'Content-Type': 'application/json',
//         },
//         body: JSON.stringify({
//             email: email,
//             password: password
//         }),
//     })
    // .then(response => {
    //     if (response.ok) {
    //         console.log('Пользователь зарегистрирован успешно!');
    //     } else {
    //         console.error('Ошибка при регистрации:', response.statusText);
    //     }
    // })
//     .catch(error => {
//         console.error('Ошибка при отправке запроса:', error);
//     });
// }

let form = document.getElementById('form');
form.addEventListener('submit', function(e){
    e.preventDefault()

    let email = document.getElementById('email').value;
    let password = document.getElementById('password').value;

    fetch("http://localhost:8081/register",{
        method: 'POST',
        body:JSON.stringify({
            email:email,
            password:password
        }),
        headers:{
            "Content-Type":"application/json; charset=UTF-8"
        }
    })
    .then(response => {
        if (response.ok) {
            console.log('Registration successful!');
        } else {
            console.error('Sign-up error:', response.statusText);
        }
    })
    // .then(function(data){
    //     console.log(data)
    // })
    .catch(error => {
        console.error('Error sending a request:', error);
    });
})