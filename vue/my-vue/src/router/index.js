import {createRouter, createWebHashHistory} from 'vue-router'

const routes = [
    {
        path:'/',
        component: ()=>import ('../views/Main.vue'),
        children:[
            {
                path:'/home',
                component : ()=>import ('../views/Home.vue')
            }
        ]

    }
]

const router = createRouter({
    routes,
    history: createWebHashHistory()
})

export default router