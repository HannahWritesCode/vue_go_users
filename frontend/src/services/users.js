import axios from 'axios'

const API_URL = "http://127.0.0.1:50009/api/v1/users"

class UserService {
    attemptSignup(email, first_name, last_name, password, password_confirmation) {
        axios({
            method: "POST",
            url: API_URL,
            data: {
                email: email,
                first_name: first_name,
                last_name: last_name,
                password: password,
                password_confirmation: password_confirmation
            },
            headers: {
                "content-type": "application/json"
            }
        })
            .then(response => {
                console.log(response.data)
                localStorage.setItem("user", response.data)
            })
            .catch(error => {
                //this.error = error.response.status
                console.log(error)
                return error
            })
    }

    attemptUpdate(email, first_name, last_name, password, password_confirmation) {
        let token = localStorage.getItem("user.token")
        axios({
            method: "PUT",
            url: API_URL,
            data: {
                email: email,
                first_name: first_name,
                last_name: last_name,
                password: password,
                password_confirmation: password_confirmation
            },
            headers: {
                "content-type": "application/json",
                "authorization": `Bearer ${token}`
            }
        })
            .then(result => {
                console.log(result)
                localStorage.setItem("user.user", response.data)
            })
            .catch(error => {
                //this.error = error.response.status
                console.log(error)
            })
    }
}

export default new UserService()