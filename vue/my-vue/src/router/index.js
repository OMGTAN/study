import {createRouter, createWebHashHistory} from 'vue-router'

const routes = [
    {
        path:'/',
        component: ()=>import ('../views/Main.vue'),
        redirect: '/home',
        children:[
            {
                path:'/home',
                name:'home',
                component : ()=>import ('../views/Home.vue')
            },
            {
                path:'/user',
                name:'user',
                component : ()=>import ('../views/User.vue')
            },
            {
                path:'/mall',
                name:'mall',
                component : ()=>import ('../views/Mall.vue')
            },
            {
                path:'/page1',
                name:'page1',
                component : ()=>import ('../views/Page1.vue')
            },
            {
                path:'/page2',
                name:'page2',
                component : ()=>import ('../views/Page2.vue')
            }
        ],

    },
    {
        path:'/login',
        component:()=>import('../views/Login.vue')
    }
]

const router = createRouter({
    routes,
    history: createWebHashHistory()
})

export default router