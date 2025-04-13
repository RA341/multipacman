const changeTextColor = () => {
    // Change the H1 color to blue.
    document.querySelector('h1').style.color = 'blue';
};



// function logout() {
//     fetch('/logout', {
//         method: 'GET',
//     }).then(response => {
//         if (response.redirected) {
//             console.log('Redirecting to ' + response.url)
//             window.location.href = response.url;
//         } else {
//             if (response.statusText) {
//                 alert('Error ' + response.statusText)
//             }
//         }
//     }).catch(error => {
//         console.error('There was a problem sending data to server:', error);
//     });
//
// }