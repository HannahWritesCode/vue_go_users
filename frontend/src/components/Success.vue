<template>
    <div>
        <div>
            <div class="flex flex-row">
                <p class="font-bold text-4xl text-green-500" v-text="`${first_name} ${last_name}`"></p>
                <p class="text-4xl">, you've logged in!</p>
            </div>
            <p>I'm so proud of you.</p>
            <!-- add methods to determine how many accounts there are total, 
        how many people have logged in that day, 
        how many people in total have logged in? -->
        </div>

        <router-link to="/update-account" class="font-semibold py-3 text-green-500 hover:underline">
            <p>Update your account information</p>
        </router-link>

        <!-- add delete account button -->

        <div>
            <button><a class="font-semibold py-3 text-green-500 hover:underline" @click="logout()">Log out</a></button>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'Success',
    data: function () {
        return {
            email: "",
            first_name: "",
            last_name: "",
            error: ""
        }
    },
    created() {
        let token = localStorage.getItem("token")
        console.log("4")
        if (token == null) {
            console.log("5")
            this.$router.push("/login")
        }
        else {
            this.getCurrentUser(token)
        }
    },
    methods: {
        getCurrentUser: function (token) {
            axios({
                method: "GET",
                url: "http://127.0.0.1:50009/api/v1/users",
                headers: {
                    "content-type": "application/json",
                    "authorization": `Bearer ${token}`
                }
            })
                .then(response => {
                    this.first_name = response.data.first_name
                    this.last_name = response.data.last_name
                })
                .catch(error => {
                    this.error = error.response.status
                    console.log(error)
                    console.log("8")
                    if (this.error == '400') {
                        console.log("9")
                        this.logout()
                    }
                    this.error = ''
                })
        },
        logout: function () {
            console.log("6")
            localStorage.removeItem('token')
            this.$router.push("/login")
            console.log("7")
        }
    }
}
</script>