<template>
    <div class="min-w-full">
        <div id="login-form">
            <h1 class="text-4xl font-semibold">Welcome.</h1>

            <form @submit.stop.prevent="attemptLogin(email, password)">
                <div class="flex flex-col my-2">
                    <label for="login-email">Email</label>
                    <input
                        class="border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 p-0.5"
                        type="email" id="login-email" v-model="email" required />
                </div>

                <div class="flex flex-col my-2">
                    <label for="password">Password</label>
                    <input
                        class="border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 p-0.5"
                        type="password" id="password" v-model="password" required />
                </div>

                <input class="bg-green-500 hover:bg-green-700 text-white font-bold my-2 py-2 px-4 rounded" type="submit"
                    value="Log In" id="submit-button" />
            </form>

            <!-- <hr /> -->

            <div v-if="error">
                <p class="font-bold py-3">Error logging in. <br />Check that the email and/or
                    password you entered are correct.</p>
            </div>
        </div>

        <router-link to="/sign-up" class="font-semibold py-3 text-green-500 hover:underline">
            <p>Create an account</p>
        </router-link>

        <router-link to="/request-password" class="font-semibold py-3 text-green-500 hover:underline">
            <p>Forgot your password?</p>
        </router-link>
    </div>
</template>

<script>
import axios from 'axios'

let token = localStorage.getItem("token")

export default {
    name: 'Login',
    data: function () {
        return {
            email: "",
            password: "",
            first_name: "",
            last_name: "",
            error: ""
        }
    },
    created() {
        console.log("1")
        if (token != null) {
            console.log("2")
            this.$router.push("/success")
        }
    },
    methods: {
        attemptLogin: function () {
            axios({
                method: "POST",
                url: "http://127.0.0.1:50009/api/v1/auth",
                data: {
                    email: this.email,
                    password: this.password
                },
                headers: {
                    "content-type": "application/json"
                }
            })
                .then(response => {
                    console.log("3")
                    localStorage.setItem("token", response.data.token)
                    this.$router.push("/success");
                })
                .catch(error => {
                    this.error = error.response.status
                })
        },
    }
}
</script>