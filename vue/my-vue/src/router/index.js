import {createRouter, createWebHashHistory} from 'vue-router'

const routes = [
    {
        path:'/',
        component: ()=>import ('../views/Main.vue'),
        children:[
            {
                path:'/home',
                component : ()=>import ('../views/Home.vue')
            },
            {
                path:'/user',
                component : ()=>import ('../views/User.vue')
            },
            {
                path:'/mall',
                component : ()=>import ('../views/Mall.vue')
            },
            {
                path:'/page1',
                component : ()=>import ('../views/Page1.vue')
            },
            {
                path:'/page2',
                component : ()=>import ('../views/Page2.vue')
            }
        ]

    }
]

const router = createRouter({
    routes,
    history: createWebHashHistory()
})

export default router