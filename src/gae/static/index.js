<script src="https://www.gstatic.com/firebasejs/5.4.1/firebase.js"></script>
    <script>
// FirebaseUI config.
var uiConfig = {
    signInSuccessUrl: '/signin',
    signInOptions: [
        // Leave the lines as is for the providers you want to offer your users.
        firebase.auth.GoogleAuthProvider.PROVIDER_ID,
    ],
    // Terms of service url.
    tosUrl: '<your-tos-url>'
};

firebase.auth().onAuthStateChanged(function (user) {
    if (user) {
        // User is signed in.
        document.getElementById('sign-out').hidden = false;
        document.getElementById('post-form').hidden = false;
        document.getElementById('account-details').textContent =
            'Signed in as ' + user.displayName + ' (' + user.email + ')';
        user.getIdToken().then(function (accessToken) {
            // Add the token to the post form. The user info will be extracted
            // from the token by the server.
            document.getElementById('tokenfirebase').value = accessToken;
        });
    } else {
        // User is signed out.
        // Initialize the FirebaseUI Widget using Firebase.
        var ui = new firebaseui.auth.AuthUI(firebase.auth());
        // Show the Firebase login button.
        ui.start('#firebaseui-auth-container', uiConfig);
        // Update the login state indicators.
        document.getElementById('sign-out').hidden = true;
        document.getElementById('post-form').hidden = true;
        document.getElementById('account-details').textContent = '';
    }
}, function (error) {
    console.log(error);
    alert('Unable to log in: ' + error)
});
</script>
