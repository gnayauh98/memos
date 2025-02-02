import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import Home from './ui/Home.vue'
import Layout from './ui/Layout.vue'
import Settings from './ui/Settings.vue'
import Explore from './ui/Explore.vue'
import SignIn from './ui/SignIn.vue'
import Resources from './ui/resources/index.vue'
import { verify } from './api/user'

async function auth() {
    const accessToken = localStorage.getItem("access-token")
    if (!accessToken) {
        return { path: '/signin' }
    }

    const isVerify = await verify(accessToken)

    return isVerify
}

const routes: RouteRecordRaw[] = [
    {
        name: 'layout',
        path: '/',
        component: Layout,
        children: [
            {
                name: 'home',
                path: '',
                component: Home,
                beforeEnter: auth
            },
            {
                name: 'resources',
                path: 'resources',
                component: Resources,
                beforeEnter: auth
            },
            {
                name: 'settings',
                path: 'settings',
                component: Settings,
                beforeEnter: auth
            },
            {
                name: 'explore',
                path: 'explore',
                component: Explore,
                // beforeEnter: auth
            },
            {
                name: 'signin',
                path: 'signin',
                component: SignIn
            }
        ]
    }
] as const

export default createRouter({
    history: createWebHistory(),
    routes
})
