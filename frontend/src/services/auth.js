import axios from 'axios'

const API_URL = 'http://127.0.0.1:50009/api/v1/'

class AuthService {
    attemptLogin(email, password) {
        axios({
            method: "POST",
            url: API_URL,
            data: {
                email: email,
                password: password
            },
            headers: {
                "content-type": "application/json"
            }
        })
            .then(response => {
                localStorage.setItem("user", response.data)
                this.$router.push("/success");
            })
            .catch(error => {
                //this.error = error.response.status
                console.log(error)
                return error
            })
    }

    logout() {
        localStorage.removeItem('user');
        this.$router.push("/login");
    }
}

export default new AuthService()