import { createWebHistory, createRouter } from "vue-router";
import Login from './components/Login.vue'
import Signup from './components/Signup.vue'
import UpdateAccount from './components/UpdateAccount.vue'
import ResetPassword from './components/ResetPassword.vue'
import Success from './components/Success.vue'

const routes = [
    {
        path: "/",
        component: Login,
    },
    {
        path: "/login",
        component: Login,
    },
    {
        path: "/sign-up",
        component: Signup,
    },
    {
        path: "/success",
        component: Success,
    },
    {
        path: "/update-account",
        component: UpdateAccount,
    },
    {
        path: "/reset-password",
        component: ResetPassword,
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
