// let btn = document.querySelector('.btn');
// function loginUser(email, password) {
//     fetch('http://localhost:8081/login', {
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
    //         console.log('Вход выполнен успешно!');
    //     } else {
    //         console.error('Ошибка при входе:', response.statusText);
    //     }
    // })
    // .catch(error => {
    //     console.error('Ошибка при отправке запроса:', error);
    // });
// }

let form = document.getElementById('form');
form.addEventListener('submit', function(e){
    e.preventDefault()

    let email = document.getElementById('email').value;
    let password = document.getElementById('password').value;

    fetch("http://localhost:8081/login",{
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
            console.log('Login successful!');
        } else {
            console.error('Sign-in error:', response.statusText);
        }
    })
    // .then(function(data){
    //     console.log(data)
    // })
    .catch(error => {
        console.error('Error sending a request:', error);
    });
})