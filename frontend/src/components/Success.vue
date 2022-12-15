<template>
    <div>
        <div>
            <div class="flex flex-row flex-wrap">
                <h1 class="font-bold text-4xl text-green-500 pr-3" v-text="`${first_name} ${last_name},`"></h1>
                <h1 class="text-4xl">you've logged in!</h1>
            </div>
            <p class="py-3">I'm so proud of you.</p>
        </div>

        <router-link to="/update-account" class="font-semibold py-3 text-green-500 hover:underline">
            <p>Update my account information</p>
        </router-link>

        <router-link to="/delete-confirmation" class="font-semibold py-3 text-green-500 hover:underline">
            <p>Delete my account</p>
        </router-link>

        <div>
            <button><a class="font-semibold py-3 text-green-500 hover:underline" @click="logout()">Log
                    out</a></button>
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